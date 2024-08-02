import DefaultLayout from "@/layouts/default";
import { Button, Image, RadioGroup } from "@nextui-org/react";
import { CustomRadio } from "./component/CustomRadio";
import { Controller, useForm } from "react-hook-form";
import { useMutation } from "@tanstack/react-query";
import { authClient } from "@/service/auth";
import { toast } from "react-hot-toast";
import {
  BackgroundType,
  Image as ImagePb,
} from "@/lib/api/auth/model/image_pb";
import { UpdateImageAfterLabeledRequest } from "@/lib/api/auth/rpc/update_image_after_labeled_pb";
import { PartialMessage } from "@bufbuild/protobuf";
import { ConnectError } from "@connectrpc/connect";
import { useNavigate } from "react-router-dom";
import React from "react";
import { GetMyLabeledImageRequest } from "@/lib/api/auth/rpc/get_my_labeled_image_pb";

type LabelFields = {
  imageNumber: string;
  backgroundType: string;
};

export default function LabelPage() {
  const {
    formState: { errors, isDirty },
    control,
    getValues,
    handleSubmit,
    reset,
  } = useForm<LabelFields>({
    defaultValues: { imageNumber: "", backgroundType: "" },
    mode: "onChange",
  });
  const navigate = useNavigate();
  const [images, setImages] = React.useState<ImagePb[]>([]);
  const [canPrevious, setCanPrevious] = React.useState(false);

  React.useEffect(() => {
    console.debug(images);
  }, [images]);

  const { mutateAsync: getMyImageMutate, isPending: getMyImageLoading } =
    useMutation({
      mutationFn: async (request: PartialMessage<GetMyLabeledImageRequest>) => {
        const { image } = await authClient.getMyLabeledImage(request);
        return image;
      },
    });

  const { mutateAsync: updateImageMutate, isPending: updateLoading } =
    useMutation({
      mutationFn: async (
        request: PartialMessage<UpdateImageAfterLabeledRequest>
      ) => {
        await authClient.updateImageAfterLabeled(request);
      },
    });

  const { mutateAsync: getImageToLabelMutate, isPending: getImageLoading } =
    useMutation({
      mutationFn: async () => {
        try {
          const { image } = await authClient.getImageToLabel({});
          return image;
        } catch (error) {
          const err = ConnectError.from(error);
          console.log(err.code);
          if (err.code === 5) {
            toast.success("No image to label");
            setTimeout(() => {
              navigate("/");
            }, 1000);
          } else throw error;
        }
      },
    });

  //   const { mutateAsync: getMyLabeledImageMutate, isPending: rollbackLoading } =
  //     useMutation({
  //       mutationFn: async (request: PartialMessage<GetMyLabeledImageRequest>) => {
  //         await authClient.getMyLabeledImage(request);
  //       },
  //     });

  const loading = getMyImageLoading || getImageLoading || updateLoading;
  const { category, displayName, id, url1, url2, url3 } =
    images?.[images.length - 1] ?? new ImagePb();

  const handleNext = async (data: LabelFields) => {
    try {
      await updateImageMutate({
        id: id.toString(),
        backgroundType: parseInt(data.backgroundType),
        category: category,
        urlSelected: data.imageNumber,
      });

      fetchImage();

      reset();
      setCanPrevious(true);

      toast.success("Update image success");
    } catch {
      toast.error("Update image failed");
    }
  };

  const handleDiscard = async () => {
    try {
      await updateImageMutate({
        id: id.toString(),
        backgroundType: BackgroundType.BT_OUTSIDE,
        category: category,
        urlSelected: "0",
      });

      fetchImage();
      setCanPrevious(true);
      toast.success("Image discarded");

      reset();
    } catch {
      toast.error("Discard image failed");
    }
  };

  const handlePrevious = async () => {
    try {
      const image = await getMyImageMutate({
        category: images[images.length - 2].category,
        id: images[images.length - 2].id.toString(),
      });
      setImages((prev) => [...prev.slice(0, -2), image ?? new ImagePb()]);
      reset();
    } catch {
    } finally {
      setCanPrevious(false);
    }
  };

  const fetchImage = async () => {
    const image = await getImageToLabelMutate();
    setImages((prev) => [...prev, image ?? new ImagePb()]);
  };

  React.useEffect(() => {
    fetchImage();
  }, []);

  return (
    <DefaultLayout loading={loading}>
      <form>
        <div className="flex items-center flex-col gap-8">
          <h1 className="text-xl font-semibold">
            Hình {displayName}-{id.toString()}
          </h1>

          <div className="md:h-[60vh] h-[50vh] w-full flex flex-col gap-8 items-center">
            {!loading && (
              <>
                <div className="w-full">
                  <RadioGroup
                    label={
                      <div className="text-center box w-full">Chọn hình</div>
                    }
                    orientation="horizontal"
                    defaultValue={getValues("imageNumber")}
                    isInvalid={!!errors.imageNumber}
                    errorMessage={errors.imageNumber?.message}
                  >
                    <div className="flex justify-center w-full">
                      <div className="grid grid-cols-3 gap-4 w-full justify-center max-w-[800px] justify-self-center">
                        {[url1, url2, url3].map((url, index) => (
                          <div
                            className="flex justify-center w-full"
                            key={index}
                          >
                            <Controller
                              rules={{ required: "This is required" }}
                              control={control}
                              name="imageNumber"
                              render={({ field: { onChange } }) => (
                                <CustomRadio
                                  value={(index + 1).toString()}
                                  description={"Hình " + (index + 1)}
                                  onChange={onChange}
                                >
                                  <Image
                                    className="min-w-[110px] sm:min-w-[170px] md:min-w-[230px]"
                                    alt="NextUI hero Image"
                                    src={url}
                                    fallbackSrc="https://placehold.co/400"
                                  />
                                </CustomRadio>
                              )}
                            />
                          </div>
                        ))}
                      </div>
                    </div>
                  </RadioGroup>
                </div>
                <RadioGroup
                  label={
                    <div className="text-center box w-full">Chọn loại nền</div>
                  }
                  orientation="horizontal"
                  isInvalid={!!errors.backgroundType}
                  errorMessage={errors.backgroundType?.message}
                >
                  <div className="grid grid-cols-2 gap-4 max-w-[600px]">
                    <Controller
                      rules={{ required: "This is required" }}
                      control={control}
                      name="backgroundType"
                      render={({ field: { onChange } }) => (
                        <CustomRadio
                          description="Nền bên trong"
                          value="1"
                          onChange={onChange}
                        ></CustomRadio>
                      )}
                    />
                    <Controller
                      rules={{ required: "This is required" }}
                      control={control}
                      name="backgroundType"
                      render={({ field: { onChange } }) => (
                        <CustomRadio
                          description="Nền toàn bộ"
                          value="2"
                          onChange={onChange}
                        ></CustomRadio>
                      )}
                    />
                  </div>
                </RadioGroup>
              </>
            )}
          </div>
          <div className="mt-0 flex gap-3">
            <Button
              color="danger"
              size="lg"
              onClick={handleDiscard}
              isDisabled={updateLoading}
            >
              Bỏ, lấy hình mới
            </Button>
            <Button
              size="lg"
              onClick={handlePrevious}
              //   isLoading={rollbackLoading}
              isDisabled={images.length < 2 || !canPrevious}
            >
              Trước
            </Button>
            <Button
              size="lg"
              color="primary"
              onClick={handleSubmit(handleNext)}
              isDisabled={!isDirty}
              //   disabled={!isDirty}
            >
              Ok, Lấy hình mới
            </Button>
          </div>
        </div>
      </form>
    </DefaultLayout>
  );
}