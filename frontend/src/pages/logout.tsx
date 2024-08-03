import UnauthenticatedLayout from "@/layouts/unauth";
import { AccessTokenKey } from "@/const/token";
import React from "react";
import { toast } from "react-hot-toast";

export default function LogoutPage() {
  React.useMemo(() => {
    localStorage.removeItem(AccessTokenKey);
    toast.success("Logout successfully");
    window.location.replace("/login");
  }, []);

  return (
    <UnauthenticatedLayout>
      <></>
    </UnauthenticatedLayout>
  );
}
