import { AppFile as PrismaAppFile } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppFile extends Partial<PrismaAppFile> {}

export type AppFilePluginArgs =
  | GetAppFileByIdPluginArgs
  | GetAppFileByAppInstanceIdPluginArgs
  | CreateAppFilePluginArgs
  | UpdateAppFilePluginArgs
  | DeleteAppFilePluginArgs;

/// GetAppFileById
export interface GetAppFileById {}
export interface GetAppFileByIdApiResponse extends ApiResponse {}
export interface GetAppFileByIdApiArgs {
  id: string;
  appInstanceId: string;
}
export interface GetAppFileByIdPluginArgs extends GetAppFileByIdApiArgs {}

/// GetAppFileByAppInstanceId
export interface GetAppFileByAppInstanceId {}
export interface GetAppFileByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppFileByAppInstanceIdApiArgs {
  appInstanceId: string;
}
export interface GetAppFileByAppInstanceIdPluginArgs
  extends GetAppFileByAppInstanceIdApiArgs {}

/// CreateAppFile
export interface CreateAppFile {}
export interface CreateAppFileApiResponse extends ApiResponse {}
export interface CreateAppFileApiArgs extends Partial<AppFile> {
  appInstanceId: string;
}
export interface CreateAppFilePluginArgs extends CreateAppFileApiArgs {}

/// UpdateAppFile
export interface UpdateAppFile {}
export interface UpdateAppFileApiResponse extends ApiResponse {}
export interface UpdateAppFileApiArgs extends Partial<AppFile> {
  id: string;
  appInstanceId: string;
}
export interface UpdateAppFilePluginArgs extends UpdateAppFileApiArgs {}

/// DeleteAppFile
export interface DeleteAppFile {}
export interface DeleteAppFileApiResponse extends ApiResponse {}
export interface DeleteAppFileApiArgs {
  id: string;
  appInstanceId: string;
}
export interface DeleteAppFilePluginArgs extends DeleteAppFileApiArgs {}
