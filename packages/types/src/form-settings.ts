//import { AppFormSettings as PrismaAppFormSettings } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppFormSettings {
    id: string;
    name: string;
    fields: any[];
    fieldCount: number | null;
    createdAt: Date | null;
    updatedAt: Date | null;
    appId: string;
}
export interface AppFormField {
    id: string;
    name: string;
    type: string;
}

// export type AppFormSettingsPluginArgs =
//   | GetAppFormSettingsByIdPluginArgs
//   | GetAppFormsSettingsByAppIdPluginArgs
//   | CreateAppFormSettingsPluginArgs
//   | UpdateAppFormSettingsPluginArgs
//   | DeleteAppFormSettingsPluginArgs;

/// GetAppFormSettingsById
export interface GetAppFormSettingsById {}
export interface GetAppFormSettingsByIdApiResponse extends ApiResponse {}
export interface GetAppFormSettingsByIdApiArgs {
    id: string;
    appId: string;
}
export interface GetAppFormSettingsByIdPluginArgs
    extends GetAppFormSettingsByIdApiArgs {}

/// GetAppFormsSettingsByAppId
export interface GetAppFormsSettingsByAppId {}
export interface GetAppFormsSettingsByAppIdApiResponse extends ApiResponse {}
export interface GetAppFormsSettingsByAppIdApiArgs {
    appId: string;
}
export interface GetAppFormsSettingsByAppIdPluginArgs
    extends GetAppFormsSettingsByAppIdApiArgs {}

/// CreateAppFormSettings
export interface CreateAppFormSettings {}
export interface CreateAppFormSettingsApiResponse extends ApiResponse {}
export interface CreateAppFormSettingsApiArgs extends Partial<AppFormSettings> {
    appId: string;
}
export interface CreateAppFormSettingsPluginArgs
    extends CreateAppFormSettingsApiArgs {}

/// UpdateAppFormSettings
export interface UpdateAppFormSettings {}
export interface UpdateAppFormSettingsApiResponse extends ApiResponse {}
export interface UpdateAppFormSettingsApiArgs extends Partial<AppFormSettings> {
    id: string;
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
    appId: string;
}
export interface DeleteAppFormSettingsPluginArgs
    extends DeleteAppFormSettingsApiArgs {}
