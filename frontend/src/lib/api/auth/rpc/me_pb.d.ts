// @generated by protoc-gen-es v1.2.1 with parameter "target=js+dts"
// @generated from file auth/rpc/me.proto (package rpc, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { UserInfo } from "../model/user_info_pb.js";

/**
 * @generated from message rpc.MeRequest
 */
export declare class MeRequest extends Message<MeRequest> {
  constructor(data?: PartialMessage<MeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "rpc.MeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeRequest;

  static equals(a: MeRequest | PlainMessage<MeRequest> | undefined, b: MeRequest | PlainMessage<MeRequest> | undefined): boolean;
}

/**
 * @generated from message rpc.MeResponse
 */
export declare class MeResponse extends Message<MeResponse> {
  /**
   * @generated from field: model.UserInfo info = 1;
   */
  info?: UserInfo;

  constructor(data?: PartialMessage<MeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "rpc.MeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeResponse;

  static equals(a: MeResponse | PlainMessage<MeResponse> | undefined, b: MeResponse | PlainMessage<MeResponse> | undefined): boolean;
}

