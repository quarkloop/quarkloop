//import { AppFile as PrismaAppFile } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppFile {
    id: string;
    enable: boolean;
    createdAt: Date | null;
    updatedAt: Date | null;
    appInstanceId: string | null;
}

// export type AppFilePluginArgs =
//     | GetAppFileByAppInstanceIdPluginArgs
//     | GetAppFileByIdPluginArgs
//     | CreateAppFilePluginArgs
//     | UpdateAppFilePluginArgs
//     | DeleteAppFilePluginArgs;

/// GetAppFileByAppInstanceId
export interface GetAppFileByAppInstanceId {}
export interface GetAppFileByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppFileByAppInstanceIdApiArgs {
    appId: string;
    instanceId: string;
}
export interface GetAppFileByAppInstanceIdPluginArgs
    extends GetAppFileByAppInstanceIdApiArgs {}

/// GetAppFileById
export interface GetAppFileById {}
export interface GetAppFileByIdApiResponse extends ApiResponse {}
export interface GetAppFileByIdApiArgs {
    appId: string;
    instanceId: string;
    fileId: string;
}
export interface GetAppFileByIdPluginArgs extends GetAppFileByIdApiArgs {}

/// CreateAppFile
export interface CreateAppFile {}
export interface CreateAppFileApiResponse extends ApiResponse {}
export interface CreateAppFileApiArgs {
    appId: string;
    instanceId: string;
    file: Partial<AppFile>;
}
export interface CreateAppFilePluginArgs extends CreateAppFileApiArgs {}

/// UpdateAppFile
export interface UpdateAppFile {}
export interface UpdateAppFileApiResponse extends ApiResponse {}
export interface UpdateAppFileApiArgs {
    appId: string;
    instanceId: string;
    file: Partial<AppFile>;
}
export interface UpdateAppFilePluginArgs extends UpdateAppFileApiArgs {}

/// DeleteAppFile
export interface DeleteAppFile {}
export interface DeleteAppFileApiResponse extends ApiResponse {}
export interface DeleteAppFileApiArgs {
    appId: string;
    instanceId: string;
    fileId: string;
}
export interface DeleteAppFilePluginArgs extends DeleteAppFileApiArgs {}
