// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=js+dts"
// @generated from file auth/auth.proto (package auth, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SignUpRequest, SignUpResponse } from "./rpc/sign_up_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { LoginRequest, LoginResponse } from "./rpc/login_pb.js";
import { MeRequest, MeResponse } from "./rpc/me_pb.js";
import { GetImageToLabelRequest, GetImageToLabelResponse } from "./rpc/get_image_to_label_pb.js";
import { UpdateImageAfterLabeledRequest, UpdateImageAfterLabeledResponse } from "./rpc/update_image_after_labeled_pb.js";
import { RollbackLabeledImageRequest, RollbackLabeledImageResponse } from "./rpc/rollback_labeled_image_pb.js";

/**
 * @generated from service auth.AuthService
 */
export const AuthService = {
  typeName: "auth.AuthService",
  methods: {
    /**
     * @generated from rpc auth.AuthService.SignUp
     */
    signUp: {
      name: "SignUp",
      I: SignUpRequest,
      O: SignUpResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.AuthService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.AuthService.Me
     */
    me: {
      name: "Me",
      I: MeRequest,
      O: MeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.AuthService.GetImageToLabel
     */
    getImageToLabel: {
      name: "GetImageToLabel",
      I: GetImageToLabelRequest,
      O: GetImageToLabelResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.AuthService.UpdateImageAfterLabeled
     */
    updateImageAfterLabeled: {
      name: "UpdateImageAfterLabeled",
      I: UpdateImageAfterLabeledRequest,
      O: UpdateImageAfterLabeledResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc auth.AuthService.RollbackLabeledImage
     */
    rollbackLabeledImage: {
      name: "RollbackLabeledImage",
      I: RollbackLabeledImageRequest,
      O: RollbackLabeledImageResponse,
      kind: MethodKind.Unary,
    },
  }
};

