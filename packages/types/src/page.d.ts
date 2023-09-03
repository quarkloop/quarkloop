import { AppPage as PrismaAppPage } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppPage extends Partial<PrismaAppPage> {}

export type AppPagePluginArgs =
  | GetAppPageByIdPluginArgs
  | GetAppPageByAppInstanceIdPluginArgs
  | CreateAppPagePluginArgs
  | UpdateAppPagePluginArgs
  | DeleteAppPagePluginArgs;

/// GetAppPageById
export interface GetAppPageById {}
export interface GetAppPageByIdApiResponse extends ApiResponse {}
export interface GetAppPageByIdApiArgs {
  id: string;
  appInstanceId: string;
}
export interface GetAppPageByIdPluginArgs extends GetAppPageByIdApiArgs {}

/// GetAppPageByAppInstanceId
export interface GetAppPageByAppInstanceId {}
export interface GetAppPageByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppPageByAppInstanceIdApiArgs {
  appInstanceId: string;
}
export interface GetAppPageByAppInstanceIdPluginArgs
  extends GetAppPageByAppInstanceIdApiArgs {}

/// CreateAppPage
export interface CreateAppPage {}
export interface CreateAppPageApiResponse extends ApiResponse {}
export interface CreateAppPageApiArgs extends Partial<AppPage> {
  appInstanceId: string;
}
export interface CreateAppPagePluginArgs extends CreateAppPageApiArgs {}

/// UpdateAppPage
export interface UpdateAppPage {}
export interface UpdateAppPageApiResponse extends ApiResponse {}
export interface UpdateAppPageApiArgs extends Partial<AppPage> {
  id: string;
  appInstanceId: string;
}
export interface UpdateAppPagePluginArgs extends UpdateAppPageApiArgs {}

/// DeleteAppPage
export interface DeleteAppPage {}
export interface DeleteAppPageApiResponse extends ApiResponse {}
export interface DeleteAppPageApiArgs {
  id: string;
  appInstanceId: string;
}
export interface DeleteAppPagePluginArgs extends DeleteAppPageApiArgs {}
