import { Snippet } from "@nextui-org/snippet";
import { Code } from "@nextui-org/code";

import { title, subtitle } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import { Button } from "@nextui-org/button";
import { useQuery } from "@tanstack/react-query";
import { authClient } from "@/service/auth";
import { useNavigate } from "react-router-dom";
import { ConnectError } from "@connectrpc/connect";

export default function IndexPage() {
  const navigate = useNavigate();
  const { data: info, isLoading: loading } = useQuery({
    queryKey: ["getInfo"],
    queryFn: async () => {
      try {
        return await authClient.me({});
      } catch (error) {
        const err = ConnectError.from(error);
        if (err.code === 16) {
          navigate("/login");
        }
      }
    },
  });

  const { data: imagesCount, isLoading: countLoading } = useQuery({
    queryKey: ["countMyLabeledImages"],
    queryFn: async () => {
      return (await authClient.countMyLabeledImages({})).count;
    },
  });

  return (
    <DefaultLayout loading={loading || countLoading}>
      {!loading && (
        <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
          <div className="inline-block max-w-lg text-center justify-center">
            <h1 className={title()}>Bắt đầu label</h1>
            <h4 className={subtitle({ class: "mt-4" })}>
              Hello there! {info?.info?.displayName}
            </h4>
          </div>

          <div className="flex gap-3">
            <Button color="primary" onClick={() => navigate("/label")}>
              Bắt đầu
            </Button>
          </div>

          <div className="mt-8">
            <Snippet hideCopyButton hideSymbol variant="bordered">
              <span>
                Tổng số ảnh đã label của bạn là{" "}
                <Code color="primary">{(imagesCount ?? 0).toString()}</Code>
              </span>
            </Snippet>
          </div>
        </section>
      )}
    </DefaultLayout>
  );
}
