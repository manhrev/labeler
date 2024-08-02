import { Config } from "@/config/env";
import GRPCClient from "./utils/grpcClient";
import { AuthService as Auth } from "@/lib/api/auth/auth_connect";

export const authClient = new GRPCClient({
  baseUrl: Config.API_BASE_URL,
  service: Auth,
}).getInstance();
