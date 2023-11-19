//import { AppForm as PrismaAppForm } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppForm {
    id: string;
    name: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    appInstanceId: string | null;
}

// export type AppFormPluginArgs =
//     | GetAppFormByAppInstanceIdPluginArgs
//     | GetAppFormByIdPluginArgs
//     | CreateAppFormPluginArgs
//     | UpdateAppFormPluginArgs
//     | DeleteAppFormPluginArgs;

/// GetAppFormByAppInstanceId
export interface GetAppFormByAppInstanceId {}
export interface GetAppFormByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppFormByAppInstanceIdApiArgs {
    projectId: string;
    instanceId: string;
}
export interface GetAppFormByAppInstanceIdPluginArgs
    extends GetAppFormByAppInstanceIdApiArgs {}

/// GetAppFormById
export interface GetAppFormById {}
export interface GetAppFormByIdApiResponse extends ApiResponse {}
export interface GetAppFormByIdApiArgs {
    projectId: string;
    instanceId: string;
    formId: string;
}
export interface GetAppFormByIdPluginArgs extends GetAppFormByIdApiArgs {}

/// CreateAppForm
export interface CreateAppForm {}
export interface CreateAppFormApiResponse extends ApiResponse {}
export interface CreateAppFormApiArgs {
    projectId: string;
    instanceId: string;
    form: Partial<AppForm>;
}
export interface CreateAppFormPluginArgs extends CreateAppFormApiArgs {}

/// UpdateAppForm
export interface UpdateAppForm {}
export interface UpdateAppFormApiResponse extends ApiResponse {}
export interface UpdateAppFormApiArgs {
    projectId: string;
    instanceId: string;
    form: Partial<AppForm>;
}
export interface UpdateAppFormPluginArgs extends UpdateAppFormApiArgs {}

/// DeleteAppForm
export interface DeleteAppForm {}
export interface DeleteAppFormApiResponse extends ApiResponse {}
export interface DeleteAppFormApiArgs {
    projectId: string;
    instanceId: string;
    formId: string;
}
export interface DeleteAppFormPluginArgs extends DeleteAppFormApiArgs {}
