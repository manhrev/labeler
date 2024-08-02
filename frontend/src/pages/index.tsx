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

  //   const handleGetInfo = () => {
  //     alert(info?.info?.email);
  //   };

  return (
    <DefaultLayout loading={loading}>
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
            {/* <Link
            isExternal
            className={buttonStyles({ variant: "bordered", radius: "full" })}
            href={siteConfig.links.github}
          >
            <GithubIcon size={20} />
            GitHub
          </Link> */}
          </div>

          <div className="mt-8">
            <Snippet hideCopyButton hideSymbol variant="bordered">
              <span>
                Bấm "bắt đầu" để chọn lựa ảnh{" "}
                <Code color="primary">.........</Code>
              </span>
            </Snippet>
          </div>
        </section>
      )}
    </DefaultLayout>
  );
}
