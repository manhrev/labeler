import { useForm } from "react-hook-form";
import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Divider,
  Button,
  Input,
} from "@nextui-org/react";
import UnauthenticatedLayout from "@/layouts/unauth";
import { AccessTokenKey } from "@/const/token";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { authClient } from "@/service/auth";

type LoginFields = {
  username: string;
  password: string;
};

export default function LoginPage() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFields>();
  const { mutateAsync, isPending: loading } = useMutation({
    mutationFn: async (data: LoginFields) => {
      return await authClient.login(data);
    },
  });

  const navigate = useNavigate();
  const handleLogin = async (data: LoginFields) => {
    try {
      const res = await mutateAsync(data);

      localStorage.setItem(AccessTokenKey, res.token);
      navigate("/");
    } catch {
      toast.error("Login failed");
    }
  };

  return (
    <UnauthenticatedLayout>
      <form>
        <div className="flex justify-center w-full">
          <Card className="max-w-[500px] min-w-[380px]">
            <CardHeader className="flex gap-3 justify-center">
              <h1 className="text-3xl font-semibold">Login</h1>
            </CardHeader>
            <Divider />
            <CardBody className="flex flex-col gap-4 p-4 py-6">
              <div>
                <Input
                  type="text"
                  label="Username"
                  {...register("username", { required: true })}
                  isInvalid={!!errors.username}
                  errorMessage="Username is required"
                  labelPlacement="outside"
                  placeholder="Username"
                />
              </div>
              <div>
                <Input
                  type="password"
                  label="Password"
                  {...register("password", { required: true })}
                  isInvalid={!!errors.password}
                  errorMessage="Password is required"
                  labelPlacement="outside"
                  placeholder="Password"
                />
              </div>
            </CardBody>
            <Divider />
            <CardFooter className="flex justify-center">
              <Button
                color="primary"
                size="md"
                className="text-md"
                onClick={handleSubmit(handleLogin)}
                isLoading={loading}
              >
                Login
              </Button>
            </CardFooter>
          </Card>
        </div>
      </form>
    </UnauthenticatedLayout>
  );
}
