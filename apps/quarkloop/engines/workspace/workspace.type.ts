import { Workspace as PrismaWorkspace } from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface Workspace extends Partial<PrismaWorkspace> {}
export enum WorkspaceAccessType {
  Public = 1,
  Private = 2,
}
export type WorkspacePluginArgs =
  | GetWorkspaceByIdPluginArgs
  | GetWorkspacesByOsIdPluginArgs
  | CreateWorkspacePluginArgs
  | UpdateWorkspacePluginArgs
  | DeleteWorkspacePluginArgs;

/// GetWorkspaceById
export interface GetWorkspaceById {}
export interface GetWorkspaceByIdApiResponse extends ApiResponse {}
export interface GetWorkspaceByIdApiArgs {
  id: string;
  osId: string;
}
export interface GetWorkspaceByIdPluginArgs extends GetWorkspaceByIdApiArgs {}

/// GetWorkspacesByOsId
export interface GetWorkspacesByOsId {}
export interface GetWorkspacesByOsIdApiResponse extends ApiResponse {}
export interface GetWorkspacesByOsIdApiArgs {
  osId: string;
}
export interface GetWorkspacesByOsIdPluginArgs
  extends GetWorkspacesByOsIdApiArgs {}

/// CreateWorkspace
export interface CreateWorkspace {}
export interface CreateWorkspaceApiResponse extends ApiResponse {}
export interface CreateWorkspaceApiArgs extends Partial<Workspace> {
  osId: string;
}
export interface CreateWorkspacePluginArgs extends CreateWorkspaceApiArgs {}

/// UpdateWorkspace
export interface UpdateWorkspace {}
export interface UpdateWorkspaceApiResponse extends ApiResponse {}
export interface UpdateWorkspaceApiArgs extends Partial<Workspace> {
  osId: string;
}
export interface UpdateWorkspacePluginArgs extends UpdateWorkspaceApiArgs {}

/// DeleteWorkspace
export interface DeleteWorkspace {}
export interface DeleteWorkspaceApiResponse extends ApiResponse {}
export interface DeleteWorkspaceApiArgs {
  id: string;
  osId: string;
}
export interface DeleteWorkspacePluginArgs extends DeleteWorkspaceApiArgs {}
