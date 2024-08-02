import { Code, ConnectError, Interceptor } from "@connectrpc/connect";
import { isTokenExpired, isValidToken } from "@/util/jwt";
import { AccessTokenKey } from "@/const/token";

export const accessTokenSetter: Interceptor = (next) => async (req) => {
  const tokenFromHeader = req.header.get("Authorization");
  let accessToken =
    tokenFromHeader || localStorage.getItem(AccessTokenKey) || "";

  if (!isValidToken(accessToken) || isTokenExpired(accessToken)) {
    localStorage.removeItem(AccessTokenKey);
  }

  req.header.set("Authorization", accessToken);
  return await next(req);
};

export const logger: (enableLog?: boolean) => Interceptor =
  (enableLog = true) =>
  (next) =>
  async (req) => {
    if (enableLog)
      console.log(
        `%c gRPCClientRequest -> [${req.service.typeName} > ${req.method.name}] -> REQUEST:`,
        "background-color: #deeb34; color: #000; font-size: 14px",
        req.message
      );
    try {
      const res = await next(req);
      if (enableLog)
        console.log(
          `%c>>>>> gRPCClientResponse -> [${req.service.typeName} > ${req.method.name}] -> SUCCESS:`,
          "background-color: #23d947; color: #000; font-size: 14px",
          res.message
        );

      return res;
    } catch (err) {
      const connectErr = ConnectError.from(err);
      switch (connectErr?.code) {
        case Code.Unauthenticated:
          window.location.replace("/login");
          if (enableLog)
            console.log(
              `%c>>>>> gRPCClientResponse -> [${req.service.typeName} > ${req.method.name}] -> ERROR -> UNAUTHENTICATED: `,
              "background-color: #c0392b; color: #000; font-size: 14px",
              err
            );
          break;
        default:
          if (enableLog)
            console.log(
              `%c>>>>> gRPCClientResponse  -> [${req.service.typeName}] > ${req.method.name}] -> ERROR: `,
              "background-color: #c0392b; color: #000; font-size: 14px",
              err
            );
          break;
      }

      return Promise.reject(err);
    }
  };
