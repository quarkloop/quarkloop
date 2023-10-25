//import { App as PrismaApp } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface App {
    id: string;
    type: number;
    name: string | null;
    icon: string | null;
    metadata: any | null;
    status: AppStatus | null;
    createdAt: Date | null;
    updatedAt: Date | null;
}

export type AppStatus = "On" | "Off" | "Archived";

// export type AppPluginArgs =
//   | GetAppByIdPluginArgs
//   | CreateAppPluginArgs
//   | UpdateAppPluginArgs
//   | DeleteAppPluginArgs;

/// GetAppById
export interface GetAppById {}
export interface GetAppByIdApiResponse extends ApiResponse {}
export interface GetAppByIdApiArgs {
    id: string;
}
export interface GetAppByIdPluginArgs extends GetAppByIdApiArgs {}

/// CreateApp
export interface CreateApp {}
export interface CreateAppApiResponse extends ApiResponse {}
export interface CreateAppApiArgs extends App {}
export interface CreateAppPluginArgs extends CreateAppApiArgs {}

/// UpdateApp
export interface UpdateApp {}
export interface UpdateAppApiResponse extends ApiResponse {}
export interface UpdateAppApiArgs extends App {
    id: string;
}
export interface UpdateAppPluginArgs extends UpdateAppApiArgs {}

/// DeleteApp
export interface DeleteApp {}
export interface DeleteAppApiResponse extends ApiResponse {}
export interface DeleteAppApiArgs {
    id: string;
}
export interface DeleteAppPluginArgs extends DeleteAppApiArgs {}
