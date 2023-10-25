//import { AppThreadSettings as PrismaAppThreadSettings } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppThreadSettings {
    id: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    appId: string;
}

// export type AppThreadSettingsPluginArgs =
//   | GetAppThreadSettingsByIdPluginArgs
//   | GetAppThreadSettingsByAppIdPluginArgs
//   | CreateAppThreadSettingsPluginArgs
//   | UpdateAppThreadSettingsPluginArgs
//   | DeleteAppThreadSettingsPluginArgs;

/// GetAppThreadSettingsById
export interface GetAppThreadSettingsById {}
export interface GetAppThreadSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppThreadSettingsByIdApiArgs {
    id: string;
    appId: string;
}
export interface GetAppThreadSettingsByIdPluginArgs
    extends GetAppThreadSettingsByIdApiArgs {}

/// GetAppThreadSettingsByAppId
export interface GetAppThreadSettingsByAppId {}
export interface GetAppThreadSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppThreadSettingsByAppIdApiArgs {
    appId: string;
}
export interface GetAppThreadSettingsByAppIdPluginArgs
    extends GetAppThreadSettingsByAppIdApiArgs {}

/// CreateAppThreadSettings
export interface CreateAppThreadSettings {}
export interface CreateAppThreadSettingsApiResponse extends ApiResponse {}
export interface CreateAppThreadSettingsApiArgs
    extends Partial<AppThreadSettings> {
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
    appId: string;
}
export interface UpdateAppThreadSettingsPluginArgs
    extends UpdateAppThreadSettingsApiArgs {}

/// DeleteAppThreadSettings
export interface DeleteAppThreadSettings {}
export interface DeleteAppThreadSettingsApiResponse extends ApiResponse {}
export interface DeleteAppThreadSettingsApiArgs {
    id: string;
    appId: string;
}
export interface DeleteAppThreadSettingsPluginArgs
    extends DeleteAppThreadSettingsApiArgs {}
