import { AppFormSettings as PrismaAppFormSettings } from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface AppFormSettings extends Partial<PrismaAppFormSettings> {}
export interface AppFormField {
  id: string;
  name: string;
  type: string;
}

export type AppFormSettingsPluginArgs =
  | GetAppFormSettingsByIdPluginArgs
  | GetAppFormsSettingsByAppIdPluginArgs
  | CreateAppFormSettingsPluginArgs
  | UpdateAppFormSettingsPluginArgs
  | DeleteAppFormSettingsPluginArgs;

/// GetAppFormSettingsById
export interface GetAppFormSettingsById {}
export interface GetAppFormSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppFormSettingsByIdApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAppFormSettingsByIdPluginArgs
  extends GetAppFormSettingsByIdApiArgs {}

/// GetAppFormsSettingsByAppId
export interface GetAppFormsSettingsByAppId {}
export interface GetAppFormsSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppFormsSettingsByAppIdApiArgs {
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAppFormsSettingsByAppIdPluginArgs
  extends GetAppFormsSettingsByAppIdApiArgs {}

/// CreateAppFormSettings
export interface CreateAppFormSettings {}
export interface CreateAppFormSettingsApiResponse extends ApiResponse {}
export interface CreateAppFormSettingsApiArgs extends Partial<AppFormSettings> {
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface CreateAppFormSettingsPluginArgs
  extends CreateAppFormSettingsApiArgs {}

/// UpdateAppFormSettings
export interface UpdateAppFormSettings {}
export interface UpdateAppFormSettingsApiResponse extends ApiResponse {}
export interface UpdateAppFormSettingsApiArgs extends Partial<AppFormSettings> {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
  formFieldCreate?: boolean;
  formFieldUpdate?: boolean;
  formFieldDelete?: boolean;
}
export interface UpdateAppFormSettingsPluginArgs
  extends UpdateAppFormSettingsApiArgs {}

/// DeleteAppFormSettings
export interface DeleteAppFormSettings {}
export interface DeleteAppFormSettingsApiResponse extends ApiResponse {}
export interface DeleteAppFormSettingsApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface DeleteAppFormSettingsPluginArgs
  extends DeleteAppFormSettingsApiArgs {}
