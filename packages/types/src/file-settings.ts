//import { AppFileSettings as PrismaAppFileSettings } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppFileSettings {
    id: string;
    enable: boolean;
    createdAt: Date | null;
    updatedAt: Date | null;
    projectId: string;
}

// export type AppFileSettingsPluginArgs =
//   | GetAppFileSettingsByIdPluginArgs
//   | GetAppFileSettingsByAppIdPluginArgs
//   | CreateAppFileSettingsPluginArgs
//   | UpdateAppFileSettingsPluginArgs
//   | DeleteAppFileSettingsPluginArgs;

/// GetAppFileSettingsById
export interface GetAppFileSettingsById {}
export interface GetAppFileSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppFileSettingsByIdApiArgs {
    id: string;
}
export interface GetAppFileSettingsByIdPluginArgs
    extends GetAppFileSettingsByIdApiArgs {}

/// GetAppFileSettingsByAppId
export interface GetAppFileSettingsByAppId {}
export interface GetAppFileSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppFileSettingsByAppIdApiArgs {
    projectId: string;
}
export interface GetAppFileSettingsByAppIdPluginArgs
    extends GetAppFileSettingsByAppIdApiArgs {}

/// CreateAppFileSettings
export interface CreateAppFileSettings {}
export interface CreateAppFileSettingsApiResponse extends ApiResponse {}
export interface CreateAppFileSettingsApiArgs extends Partial<AppFileSettings> {
    projectId: string;
}
export interface CreateAppFileSettingsPluginArgs
    extends CreateAppFileSettingsApiArgs {}

/// UpdateAppFileSettings
export interface UpdateAppFileSettings {}
export interface UpdateAppFileSettingsApiResponse extends ApiResponse {}
export interface UpdateAppFileSettingsApiArgs extends Partial<AppFileSettings> {
    id: string;
    projectId: string;
}
export interface UpdateAppFileSettingsPluginArgs
    extends UpdateAppFileSettingsApiArgs {}

/// DeleteAppFileSettings
export interface DeleteAppFileSettings {}
export interface DeleteAppFileSettingsApiResponse extends ApiResponse {}
export interface DeleteAppFileSettingsApiArgs {
    id: string;
    projectId: string;
}
export interface DeleteAppFileSettingsPluginArgs
    extends DeleteAppFileSettingsApiArgs {}
