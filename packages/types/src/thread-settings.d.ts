import { AppThreadSettings as PrismaAppThreadSettings } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppThreadSettings extends Partial<PrismaAppThreadSettings> {}

export type AppThreadSettingsPluginArgs =
  | GetAppThreadSettingsByIdPluginArgs
  | GetAppThreadSettingsByAppIdPluginArgs
  | CreateAppThreadSettingsPluginArgs
  | UpdateAppThreadSettingsPluginArgs
  | DeleteAppThreadSettingsPluginArgs;

/// GetAppThreadSettingsById
export interface GetAppThreadSettingsById {}
export interface GetAppThreadSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppThreadSettingsByIdApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAppThreadSettingsByIdPluginArgs
  extends GetAppThreadSettingsByIdApiArgs {}

/// GetAppThreadSettingsByAppId
export interface GetAppThreadSettingsByAppId {}
export interface GetAppThreadSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppThreadSettingsByAppIdApiArgs {
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAppThreadSettingsByAppIdPluginArgs
  extends GetAppThreadSettingsByAppIdApiArgs {}

/// CreateAppThreadSettings
export interface CreateAppThreadSettings {}
export interface CreateAppThreadSettingsApiResponse extends ApiResponse {}
export interface CreateAppThreadSettingsApiArgs
  extends Partial<AppThreadSettings> {
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface CreateAppThreadSettingsPluginArgs
  extends CreateAppThreadSettingsApiArgs {}

/// UpdateAppThreadSettings
export interface UpdateAppThreadSettings {}
export interface UpdateAppThreadSettingsApiResponse extends ApiResponse {}
export interface UpdateAppThreadSettingsApiArgs
  extends Partial<AppThreadSettings> {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface UpdateAppThreadSettingsPluginArgs
  extends UpdateAppThreadSettingsApiArgs {}

/// DeleteAppThreadSettings
export interface DeleteAppThreadSettings {}
export interface DeleteAppThreadSettingsApiResponse extends ApiResponse {}
export interface DeleteAppThreadSettingsApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface DeleteAppThreadSettingsPluginArgs
  extends DeleteAppThreadSettingsApiArgs {}
