import { ServiceType } from "@bufbuild/protobuf";
import {
  Interceptor,
  PromiseClient,
  createPromiseClient,
} from "@connectrpc/connect";
import { createGrpcWebTransport } from "@connectrpc/connect-web";
import { accessTokenSetter, logger } from "./defaultInterceptors";

interface GRPCClientConfig<TService extends ServiceType> {
  baseUrl: string;
  customInterceptors?: Interceptor[];
  service: TService;
}

class GRPCClient<TService extends ServiceType> {
  private instance: PromiseClient<TService>;

  constructor({
    baseUrl,
    service,
    customInterceptors,
  }: GRPCClientConfig<TService>) {
    const transport = createGrpcWebTransport({
      baseUrl: baseUrl,
      interceptors: [
        accessTokenSetter,
        logger(true),
        ...(customInterceptors || []),
      ],
    });
    this.instance = createPromiseClient(service, transport);
  }

  getInstance(): PromiseClient<TService> {
    return this.instance;
  }
}

export default GRPCClient;
