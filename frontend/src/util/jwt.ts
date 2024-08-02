import { JwtPayload, jwtDecode } from "jwt-decode";

export function isTokenExpired(token: string): boolean {
  try {
    jwtDecode(token, { header: true });
    const d = jwtDecode(token);
    if (!d) throw new Error("Can not decode token");
    const { exp } = d as JwtPayload;
    if (!exp) throw new Error("Token expired");
    return Number(exp) * 1000 < Date.now();
  } catch (error) {
    return true;
  }
}

export function isValidToken(token?: string) {
  try {
    if (!token) {
      return false;
    }
    if (isTokenExpired(token)) {
      return false;
    }
    return true;
  } catch (error) {
    return false;
  }
}
