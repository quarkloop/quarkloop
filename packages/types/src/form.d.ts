import { AppForm as PrismaAppForm } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppForm extends Partial<PrismaAppForm> {}

export type AppFormPluginArgs =
  | GetAppFormByIdPluginArgs
  | GetAppFormByAppInstanceIdPluginArgs
  | CreateAppFormPluginArgs
  | UpdateAppFormPluginArgs
  | DeleteAppFormPluginArgs;

/// GetAppFormById
export interface GetAppFormById {}
export interface GetAppFormByIdApiResponse extends ApiResponse {}
export interface GetAppFormByIdApiArgs {
  id: number;
  appInstanceId: string;
}
export interface GetAppFormByIdPluginArgs extends GetAppFormByIdApiArgs {}

/// GetAppFormByAppInstanceId
export interface GetAppFormByAppInstanceId {}
export interface GetAppFormByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppFormByAppInstanceIdApiArgs {
  appInstanceId: string;
}
export interface GetAppFormByAppInstanceIdPluginArgs
  extends GetAppFormByAppInstanceIdApiArgs {}

/// CreateAppForm
export interface CreateAppForm {}
export interface CreateAppFormApiResponse extends ApiResponse {}
export interface CreateAppFormApiArgs extends Partial<AppForm> {
  appInstanceId: string;
}
export interface CreateAppFormPluginArgs extends CreateAppFormApiArgs {}

/// UpdateAppForm
export interface UpdateAppForm {}
export interface UpdateAppFormApiResponse extends ApiResponse {}
export interface UpdateAppFormApiArgs extends Partial<AppForm> {
  id: number;
  appInstanceId: string;
}
export interface UpdateAppFormPluginArgs extends UpdateAppFormApiArgs {}

/// DeleteAppForm
export interface DeleteAppForm {}
export interface DeleteAppFormApiResponse extends ApiResponse {}
export interface DeleteAppFormApiArgs {
  id: number;
  appInstanceId: string;
}
export interface DeleteAppFormPluginArgs extends DeleteAppFormApiArgs {}
