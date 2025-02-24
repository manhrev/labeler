// @generated by protoc-gen-es v1.2.1 with parameter "target=js+dts"
// @generated from file auth/rpc/get_my_labeled_image.proto (package rpc, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Category, Image } from "../model/image_pb.js";

/**
 * @generated from message rpc.GetMyLabeledImageRequest
 */
export declare class GetMyLabeledImageRequest extends Message<GetMyLabeledImageRequest> {
  /**
   * @generated from field: model.Category category = 1;
   */
  category: Category;

  /**
   * @generated from field: string id = 2;
   */
  id: string;

  constructor(data?: PartialMessage<GetMyLabeledImageRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "rpc.GetMyLabeledImageRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMyLabeledImageRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMyLabeledImageRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMyLabeledImageRequest;

  static equals(a: GetMyLabeledImageRequest | PlainMessage<GetMyLabeledImageRequest> | undefined, b: GetMyLabeledImageRequest | PlainMessage<GetMyLabeledImageRequest> | undefined): boolean;
}

/**
 * @generated from message rpc.GetMyLabeledImageResponse
 */
export declare class GetMyLabeledImageResponse extends Message<GetMyLabeledImageResponse> {
  /**
   * @generated from field: model.Image image = 1;
   */
  image?: Image;

  constructor(data?: PartialMessage<GetMyLabeledImageResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "rpc.GetMyLabeledImageResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMyLabeledImageResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMyLabeledImageResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMyLabeledImageResponse;

  static equals(a: GetMyLabeledImageResponse | PlainMessage<GetMyLabeledImageResponse> | undefined, b: GetMyLabeledImageResponse | PlainMessage<GetMyLabeledImageResponse> | undefined): boolean;
}

