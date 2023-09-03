import { AppThread as PrismaAppThread } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppThread extends Partial<PrismaAppThread> {}

export type AppThreadPluginArgs =
  | GetAppThreadByIdPluginArgs
  | GetAppThreadByAppInstanceIdPluginArgs
  | CreateAppThreadPluginArgs
  | UpdateAppThreadPluginArgs
  | DeleteAppThreadPluginArgs;

/// GetAppThreadById
export interface GetAppThreadById {}
export interface GetAppThreadByIdApiResponse extends ApiResponse {}
export interface GetAppThreadByIdApiArgs {
  id: number;
  appInstanceId: string;
}
export interface GetAppThreadByIdPluginArgs extends GetAppThreadByIdApiArgs {}

/// GetAppThreadByAppInstanceId
export interface GetAppThreadByAppInstanceId {}
export interface GetAppThreadByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppThreadByAppInstanceIdApiArgs {
  appInstanceId: string;
}
export interface GetAppThreadByAppInstanceIdPluginArgs
  extends GetAppThreadByAppInstanceIdApiArgs {}

/// CreateAppThread
export interface CreateAppThread {}
export interface CreateAppThreadApiResponse extends ApiResponse {}
export interface CreateAppThreadApiArgs extends Partial<AppThread> {
  appInstanceId: string;
}
export interface CreateAppThreadPluginArgs extends CreateAppThreadApiArgs {}

/// UpdateAppThread
export interface UpdateAppThread {}
export interface UpdateAppThreadApiResponse extends ApiResponse {}
export interface UpdateAppThreadApiArgs extends Partial<AppThread> {
  id: number;
  appInstanceId: string;
}
export interface UpdateAppThreadPluginArgs extends UpdateAppThreadApiArgs {}

/// DeleteAppThread
export interface DeleteAppThread {}
export interface DeleteAppThreadApiResponse extends ApiResponse {}
export interface DeleteAppThreadApiArgs {
  id: number;
  appInstanceId: string;
}
export interface DeleteAppThreadPluginArgs extends DeleteAppThreadApiArgs {}
