import { AppThread as PrismaAppThread } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppThread extends Partial<PrismaAppThread> {}

export type AppThreadPluginArgs =
    | GetAppThreadByAppInstanceIdPluginArgs
    | GetAppThreadByIdPluginArgs
    | CreateAppThreadPluginArgs
    | UpdateAppThreadPluginArgs
    | DeleteAppThreadPluginArgs;

/// GetAppThreadByAppInstanceId
export interface GetAppThreadByAppInstanceId {}
export interface GetAppThreadByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppThreadByAppInstanceIdApiArgs {
    appId: string;
    instanceId: string;
}
export interface GetAppThreadByAppInstanceIdPluginArgs
    extends GetAppThreadByAppInstanceIdApiArgs {}

/// GetAppThreadById
export interface GetAppThreadById {}
export interface GetAppThreadByIdApiResponse extends ApiResponse {}
export interface GetAppThreadByIdApiArgs {
    appId: string;
    instanceId: string;
    threadId: string;
}
export interface GetAppThreadByIdPluginArgs extends GetAppThreadByIdApiArgs {}

/// CreateAppThread
export interface CreateAppThread {}
export interface CreateAppThreadApiResponse extends ApiResponse {}
export interface CreateAppThreadApiArgs {
    appId: string;
    instanceId: string;
    thread: Partial<AppThread>;
}
export interface CreateAppThreadPluginArgs extends CreateAppThreadApiArgs {}

/// UpdateAppThread
export interface UpdateAppThread {}
export interface UpdateAppThreadApiResponse extends ApiResponse {}
export interface UpdateAppThreadApiArgs {
    appId: string;
    instanceId: string;
    thread: Partial<AppThread>;
}
export interface UpdateAppThreadPluginArgs extends UpdateAppThreadApiArgs {}

/// DeleteAppThread
export interface DeleteAppThread {}
export interface DeleteAppThreadApiResponse extends ApiResponse {}
export interface DeleteAppThreadApiArgs {
    appId: string;
    instanceId: string;
    threadId: string;
}
export interface DeleteAppThreadPluginArgs extends DeleteAppThreadApiArgs {}
