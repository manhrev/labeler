// @generated by protoc-gen-es v1.2.1 with parameter "target=js+dts"
// @generated from file auth/rpc/me.proto (package rpc, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { UserInfo } from "../model/user_info_pb.js";

/**
 * @generated from message rpc.MeRequest
 */
export const MeRequest = proto3.makeMessageType(
  "rpc.MeRequest",
  [],
);

/**
 * @generated from message rpc.MeResponse
 */
export const MeResponse = proto3.makeMessageType(
  "rpc.MeResponse",
  () => [
    { no: 1, name: "info", kind: "message", T: UserInfo },
  ],
);
