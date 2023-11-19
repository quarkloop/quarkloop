//import { AppPageSettings as PrismaAppPageSettings } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppPageSettings {
    id: string;
    name: string;
    entryPoint: boolean;
    content: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    projectId: string;
}

// export type AppPageSettingsPluginArgs =
//   | GetAppPageSettingsByIdPluginArgs
//   | GetAppPagesSettingsByAppIdPluginArgs
//   | CreateAppPageSettingsPluginArgs
//   | UpdateAppPageSettingsPluginArgs
//   | DeleteAppPageSettingsPluginArgs;

/// GetAppPageSettingsById
export interface GetAppPageSettingsById {}
export interface GetAppPageSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppPageSettingsByIdApiArgs {
    id: string;
    projectId: string;
}
export interface GetAppPageSettingsByIdPluginArgs
    extends GetAppPageSettingsByIdApiArgs {}

/// GetAppPagesSettingsByAppId
export interface GetAppPagesSettingsByAppId {}
export interface GetAppPagesSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppPagesSettingsByAppIdApiArgs {
    projectId: string;
}
export interface GetAppPagesSettingsByAppIdPluginArgs
    extends GetAppPagesSettingsByAppIdApiArgs {}

/// CreateAppPageSettings
export interface CreateAppPageSettings {}
export interface CreateAppPageSettingsApiResponse extends ApiResponse {}
export interface CreateAppPageSettingsApiArgs extends Partial<AppPageSettings> {
    projectId: string;
}
export interface CreateAppPageSettingsPluginArgs
    extends CreateAppPageSettingsApiArgs {}

/// UpdateAppPageSettings
export interface UpdateAppPageSettings {}
export interface UpdateAppPageSettingsApiResponse extends ApiResponse {}
export interface UpdateAppPageSettingsApiArgs extends Partial<AppPageSettings> {
    id: string;
    projectId: string;
}
export interface UpdateAppPageSettingsPluginArgs
    extends UpdateAppPageSettingsApiArgs {}

/// DeleteAppPageSettings
export interface DeleteAppPageSettings {}
export interface DeleteAppPageSettingsApiResponse extends ApiResponse {}
export interface DeleteAppPageSettingsApiArgs {
    id: string;
    projectId: string;
}
export interface DeleteAppPageSettingsPluginArgs
    extends DeleteAppPageSettingsApiArgs {}
