//import { AppInstance as PrismaAppInstance } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppInstance {
    id: string;
    name: string | null;
    createdAt: Date | null;
    updatedAt: Date | null;
    appId: string;
    stage: string;
}

// export type AppInstancePluginArgs =
//     | GetAppInstanceByIdPluginArgs
//     | GetAppInstanceByIdPluginArgs
//     | GetAppInstancesByAppIdPluginArgs
//     | CreateAppInstancePluginArgs
//     | UpdateAppInstancePluginArgs
//     | DeleteAppInstancePluginArgs;

/// GetAppInstanceById
export interface GetAppInstanceById {}
export interface GetAppInstanceByIdApiResponse extends ApiResponse {}
export interface GetAppInstanceByIdApiArgs {
    id: string;
    appId: string;
}
export interface GetAppInstanceByIdPluginArgs
    extends GetAppInstanceByIdApiArgs {}

/// GetAppInstancesByAppId
export interface GetAppInstancesByAppId {}
export interface GetAppInstancesByAppIdApiResponse extends ApiResponse {}
export interface GetAppInstancesByAppIdApiArgs {
    appId?: string;
}
export interface GetAppInstancesByAppIdPluginArgs
    extends GetAppInstancesByAppIdApiArgs {}

/// CreateAppInstance
export interface CreateAppInstance {}
export interface CreateAppInstanceApiResponse extends ApiResponse {}
export interface CreateAppInstanceApiArgs extends Partial<AppInstance> {}
export interface CreateAppInstancePluginArgs extends CreateAppInstanceApiArgs {}

/// UpdateAppInstance
export interface UpdateAppInstance {}
export interface UpdateAppInstanceApiResponse extends ApiResponse {}
export interface UpdateAppInstanceApiArgs extends Partial<AppInstance> {}
export interface UpdateAppInstancePluginArgs extends UpdateAppInstanceApiArgs {}

/// DeleteAppInstance
export interface DeleteAppInstance {}
export interface DeleteAppInstanceApiResponse extends ApiResponse {}
export interface DeleteAppInstanceApiArgs {
    id: string;
    appId: string;
}
export interface DeleteAppInstancePluginArgs extends DeleteAppInstanceApiArgs {}
