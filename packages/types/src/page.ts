//import { AppPage as PrismaAppPage } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface AppPage {
    id: string;
    name: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    appInstanceId: string | null;
}

// export type AppPagePluginArgs =
//     | GetAppPageByAppInstanceIdPluginArgs
//     | GetAppPageByIdPluginArgs
//     | CreateAppPagePluginArgs
//     | UpdateAppPagePluginArgs
//     | DeleteAppPagePluginArgs;

/// GetAppPageByAppInstanceId
export interface GetAppPageByAppInstanceId {}
export interface GetAppPageByAppInstanceIdApiResponse extends ApiResponse {}
export interface GetAppPageByAppInstanceIdApiArgs {
    projectId: string;
    instanceId: string;
}
export interface GetAppPageByAppInstanceIdPluginArgs
    extends GetAppPageByAppInstanceIdApiArgs {}

/// GetAppPageById
export interface GetAppPageById {}
export interface GetAppPageByIdApiResponse extends ApiResponse {}
export interface GetAppPageByIdApiArgs {
    projectId: string;
    instanceId: string;
    pageId: string;
}
export interface GetAppPageByIdPluginArgs extends GetAppPageByIdApiArgs {}

/// CreateAppPage
export interface CreateAppPage {}
export interface CreateAppPageApiResponse extends ApiResponse {}
export interface CreateAppPageApiArgs {
    projectId: string;
    instanceId: string;
    page: Partial<AppPage>;
}
export interface CreateAppPagePluginArgs extends CreateAppPageApiArgs {}

/// UpdateAppPage
export interface UpdateAppPage {}
export interface UpdateAppPageApiResponse extends ApiResponse {}
export interface UpdateAppPageApiArgs {
    projectId: string;
    instanceId: string;
    page: Partial<AppPage>;
}
export interface UpdateAppPagePluginArgs extends UpdateAppPageApiArgs {}

/// DeleteAppPage
export interface DeleteAppPage {}
export interface DeleteAppPageApiResponse extends ApiResponse {}
export interface DeleteAppPageApiArgs {
    projectId: string;
    instanceId: string;
    pageId: string;
}
export interface DeleteAppPagePluginArgs extends DeleteAppPageApiArgs {}
