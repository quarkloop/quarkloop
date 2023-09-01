import {
  App as PrismaApp,
  UserRoleAssignmentMap as PrismaUserRoleAssignmentMap,
} from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface App extends Partial<PrismaApp> {}

export type AppPluginArgs =
  | GetAppByIdPluginArgs
  | GetAppByPathPluginArgs
  | GetAppsByOsIdPluginArgs
  //| GetAppsByWorkspaceIdPluginArgs
  | CreateAppPluginArgs
  | UpdateAppPluginArgs
  | DeleteAppPluginArgs;

/// GetAppById
export interface GetAppById {}
export interface GetAppByIdApiResponse extends ApiResponse {}
export interface GetAppByIdApiArgs {
  id: string;
  osId: string;
  workspaceId?: string;
}
export interface GetAppByIdPluginArgs extends GetAppByIdApiArgs {}

/// GetAppByPath
export interface GetAppByPath {}
export interface GetAppByPathApiResponse extends ApiResponse {}
export interface GetAppByPathApiArgs {
  osId: string;
  workspaceId: string;
  workspacePath?: string;
  profilePath?: string;
}
export interface GetAppByPathPluginArgs extends GetAppByPathApiArgs {}

/// GetAppsByOsId
export interface GetAppsByOsId {}
export interface GetAppsByOsIdApiResponse extends ApiResponse {}
export interface GetAppsByOsIdApiArgs {
  osId: string;
  workspaceId?: string;
}
export interface GetAppsByOsIdPluginArgs extends GetAppsByOsIdApiArgs {}

// /// GetAppsByWorkspaceId
// export interface GetAppsByWorkspaceId {}
// export interface GetAppsByWorkspaceIdApiResponse extends ApiResponse {}
// export interface GetAppsByWorkspaceIdApiArgs {
//   osId: string;
//   workspaceId: string;
// }
// export interface GetAppsByWorkspaceIdPluginArgs
//   extends GetAppsByWorkspaceIdApiArgs {}

/// CreateApp
export interface CreateApp {}
export interface CreateAppApiResponse extends ApiResponse {}
export interface CreateAppApiArgs extends App {
  osId: string;
  workspaceId: string;
}
export interface CreateAppPluginArgs extends CreateAppApiArgs {}

/// UpdateApp
export interface UpdateApp {}
export interface UpdateAppApiResponse extends ApiResponse {}
export interface UpdateAppApiArgs extends App {
  id: string;
  osId: string;
  workspaceId: string;
}
export interface UpdateAppPluginArgs extends UpdateAppApiArgs {}

/// DeleteApp
export interface DeleteApp {}
export interface DeleteAppApiResponse extends ApiResponse {}
export interface DeleteAppApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
}
export interface DeleteAppPluginArgs extends DeleteAppApiArgs {}
