import { ConnectError } from "@connectrpc/connect";
import { headerGrpcMessage } from "@connectrpc/connect/protocol-grpc";

export const getConnectError = (err: any) => {
  return ConnectError.from(err);
};

//get grpc-status header value from ConnectError object
export const getGrpcStatusCode = (err: any) => {
  const connectErr = getConnectError(err);
  return Number(connectErr?.code);
};

export const getGrpcMessage = (err: any) => {
  const connectErr = ConnectError.from(err);
  return connectErr.metadata.get(headerGrpcMessage);
};

export const getGrpcErrorDetail = (err: any) => {
  const connectErr = ConnectError.from(err);
  const detailString = connectErr.metadata.get("details") ?? "";
  try {
    const details: {
      metadata: {
        code: number;
        msg: string;
      };
    }[] = JSON.parse(atob(detailString));

    return details;
  } catch {
    return [];
  }
};
