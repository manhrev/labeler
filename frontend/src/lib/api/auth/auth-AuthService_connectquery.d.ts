// @generated by protoc-gen-connect-query v0.2.3 with parameter "target=js+dts"
// @generated from file auth/auth.proto (package auth, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { UnaryHooks } from "@bufbuild/connect-query";
import { SignUpRequest, SignUpResponse } from "./rpc/sign_up_pb.js";
import { LoginRequest, LoginResponse } from "./rpc/login_pb.js";
import { MeRequest, MeResponse } from "./rpc/me_pb.js";
import { GetImageToLabelRequest, GetImageToLabelResponse } from "./rpc/get_image_to_label_pb.js";
import { UpdateImageAfterLabeledRequest, UpdateImageAfterLabeledResponse } from "./rpc/update_image_after_labeled_pb.js";
import { RollbackLabeledImageRequest, RollbackLabeledImageResponse } from "./rpc/rollback_labeled_image_pb.js";
import { GetMyLabeledImageRequest, GetMyLabeledImageResponse } from "./rpc/get_my_labeled_image_pb.js";

export const signUp: UnaryHooks<SignUpRequest, SignUpResponse>;
export const login: UnaryHooks<LoginRequest, LoginResponse>;
export const me: UnaryHooks<MeRequest, MeResponse>;
export const getImageToLabel: UnaryHooks<GetImageToLabelRequest, GetImageToLabelResponse>;
export const updateImageAfterLabeled: UnaryHooks<UpdateImageAfterLabeledRequest, UpdateImageAfterLabeledResponse>;
export const rollbackLabeledImage: UnaryHooks<RollbackLabeledImageRequest, RollbackLabeledImageResponse>;
export const getMyLabeledImage: UnaryHooks<GetMyLabeledImageRequest, GetMyLabeledImageResponse>;