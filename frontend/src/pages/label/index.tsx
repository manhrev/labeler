import DefaultLayout from "@/layouts/default";
import { Button, Image, RadioGroup } from "@nextui-org/react";
import { CustomRadio } from "./component/CustomRadio";
import { Controller, useForm } from "react-hook-form";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { authClient } from "@/service/auth";
import { toast } from "react-hot-toast";
import {
  BackgroundType,
  Image as ImagePb,
} from "@/lib/api/auth/model/image_pb";
import { UpdateImageAfterLabeledRequest } from "@/lib/api/auth/rpc/update_image_after_labeled_pb";
import { PartialMessage } from "@bufbuild/protobuf";
import { RollbackLabeledImageRequest } from "@/lib/api/auth/rpc/rollback_labeled_image_pb";
import { ConnectError } from "@connectrpc/connect";
import { useNavigate } from "react-router-dom";
import React from "react";

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
  const [hasPrevious, setHasPrevious] = React.useState(false);
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const { mutateAsync: updateImageMutate, isPending: mutateLoading } =
    useMutation({
      mutationFn: async (
        request: PartialMessage<UpdateImageAfterLabeledRequest>
      ) => {
        await authClient.updateImageAfterLabeled(request);
      },
    });

  const { data: image, isFetching: queryLoading } = useQuery({
    queryKey: ["getImageToLabel"],
    queryFn: async () => {
      try {
        const { image } = await authClient.getImageToLabel({});
        return image;
      } catch (error) {
        const err = ConnectError.from(error);
        console.log(err.code);
        if (err.code === 5) {
          setTimeout(() => {
            navigate("/");
          }, 1000);
          toast.success("No image to label");
        } else toast.error("Get image failed, please try again");
      }
    },
  });

  const { mutateAsync: rollbackImageMutate, isPending: rollbackLoading } =
    useMutation({
      mutationFn: async (
        request: PartialMessage<RollbackLabeledImageRequest>
      ) => {
        await authClient.rollbackLabeledImage(request);
      },
    });

  const loading = queryLoading || mutateLoading;
  const { category, displayName, id, url1, url2, url3 } =
    image ?? new ImagePb();

  const handleNext = async (data: LabelFields) => {
    try {
      await updateImageMutate({
        id: id.toString(),
        backgroundType: parseInt(data.backgroundType),
        category: category,
        urlSelected: data.imageNumber,
      });
      toast.success("Update image success");
      queryClient.invalidateQueries({ queryKey: ["getImageToLabel"] });
      reset();
      setHasPrevious(true);
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
      toast.success("Image discarded");
      queryClient.invalidateQueries({ queryKey: ["getImageToLabel"] });
      reset();
    } catch {
      toast.error("Discard image failed");
    }
  };

  const handlePrevious = async () => {
    try {
      await rollbackImageMutate({
        category: category,
        id: id.toString(),
      });
      setHasPrevious(false);
      toast.success("Rollback previous image success");
    } catch {
      toast.error("Get previous image failed");
    }
  };

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
                <div>
                  <RadioGroup
                    label={
                      <div className="text-center box w-full">Chọn hình</div>
                    }
                    orientation="horizontal"
                    defaultValue={getValues("imageNumber")}
                    isInvalid={!!errors.imageNumber}
                    errorMessage={errors.imageNumber?.message}
                  >
                    <div className="grid grid-cols-3 gap-4 ">
                      <Controller
                        rules={{ required: "This is required" }}
                        control={control}
                        name="imageNumber"
                        render={({ field: { onChange } }) => (
                          <CustomRadio
                            value="1"
                            description="Hình 1"
                            onChange={onChange}
                          >
                            <Image
                              alt="NextUI hero Image"
                              src={url1}
                              fallbackSrc="https://placehold.co/400"
                            />
                          </CustomRadio>
                        )}
                      />
                      <Controller
                        control={control}
                        name="imageNumber"
                        render={({ field: { onChange } }) => (
                          <CustomRadio
                            value="2"
                            description="Hình 2"
                            onChange={onChange}
                          >
                            <Image
                              alt="NextUI hero Image"
                              src={url2}
                              fallbackSrc="https://placehold.co/400"
                            />
                          </CustomRadio>
                        )}
                      />
                      <Controller
                        control={control}
                        name="imageNumber"
                        render={({ field: { onChange } }) => (
                          <CustomRadio
                            value="3"
                            description="Hình 3"
                            onChange={onChange}
                          >
                            <Image
                              alt="NextUI hero Image"
                              src={url3}
                              fallbackSrc="https://placehold.co/400"
                            />
                          </CustomRadio>
                        )}
                      />
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
              isDisabled={mutateLoading}
            >
              Bỏ
            </Button>
            <Button
              size="lg"
              onClick={handlePrevious}
              isLoading={rollbackLoading}
              isDisabled={!hasPrevious}
            >
              Huỷ ảnh trước
            </Button>
            <Button
              size="lg"
              color="primary"
              onClick={handleSubmit(handleNext)}
              isDisabled={!isDirty}
              //   disabled={!isDirty}
            >
              Tiếp
            </Button>
          </div>
        </div>
      </form>
    </DefaultLayout>
  );
}
