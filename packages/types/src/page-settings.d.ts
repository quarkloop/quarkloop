import { AppPageSettings as PrismaAppPageSettings } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppPageSettings extends Partial<PrismaAppPageSettings> {}

export type AppPageSettingsPluginArgs =
  | GetAppPageSettingsByIdPluginArgs
  | GetAppPagesSettingsByAppIdPluginArgs
  | CreateAppPageSettingsPluginArgs
  | UpdateAppPageSettingsPluginArgs
  | DeleteAppPageSettingsPluginArgs;

/// GetAppPageSettingsById
export interface GetAppPageSettingsById {}
export interface GetAppPageSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppPageSettingsByIdApiArgs {
  id: string;
  appId: string;
}
export interface GetAppPageSettingsByIdPluginArgs
  extends GetAppPageSettingsByIdApiArgs {}

/// GetAppPagesSettingsByAppId
export interface GetAppPagesSettingsByAppId {}
export interface GetAppPagesSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppPagesSettingsByAppIdApiArgs {
  appId: string;
}
export interface GetAppPagesSettingsByAppIdPluginArgs
  extends GetAppPagesSettingsByAppIdApiArgs {}

/// CreateAppPageSettings
export interface CreateAppPageSettings {}
export interface CreateAppPageSettingsApiResponse extends ApiResponse {}
export interface CreateAppPageSettingsApiArgs extends Partial<AppPageSettings> {
  appId: string;
}
export interface CreateAppPageSettingsPluginArgs
  extends CreateAppPageSettingsApiArgs {}

/// UpdateAppPageSettings
export interface UpdateAppPageSettings {}
export interface UpdateAppPageSettingsApiResponse extends ApiResponse {}
export interface UpdateAppPageSettingsApiArgs extends Partial<AppPageSettings> {
  id: string;
  appId: string;
}
export interface UpdateAppPageSettingsPluginArgs
  extends UpdateAppPageSettingsApiArgs {}

/// DeleteAppPageSettings
export interface DeleteAppPageSettings {}
export interface DeleteAppPageSettingsApiResponse extends ApiResponse {}
export interface DeleteAppPageSettingsApiArgs {
  id: string;
  appId: string;
}
export interface DeleteAppPageSettingsPluginArgs
  extends DeleteAppPageSettingsApiArgs {}
