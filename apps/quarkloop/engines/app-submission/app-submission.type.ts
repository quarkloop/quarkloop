import { AppSubmission as PrismaAppSubmission } from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface AppSubmission extends Partial<PrismaAppSubmission> {}

export type AppSubmissionPluginArgs =
  | GetAdminAppSubmissionByIdPluginArgs
  | GetAppSubmissionByIdPluginArgs
  | GetAppSubmissionsByAppIdPluginArgs
  | CreateAppSubmissionPluginArgs
  | UpdateAppSubmissionPluginArgs
  | DeleteAppSubmissionPluginArgs;

/// GetAdminAppSubmissionById
export interface GetAdminAppSubmissionById {}
export interface GetAdminAppSubmissionByIdApiResponse extends ApiResponse {}
export interface GetAdminAppSubmissionByIdApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAdminAppSubmissionByIdPluginArgs
  extends GetAdminAppSubmissionByIdApiArgs {}

/// GetAppSubmissionById
export interface GetAppSubmissionById {}
export interface GetAppSubmissionByIdApiResponse extends ApiResponse {}
export interface GetAppSubmissionByIdApiArgs {
  id: string;
}
export interface GetAppSubmissionByIdPluginArgs
  extends GetAppSubmissionByIdApiArgs {}

/// GetAppSubmissionsByAppId
export interface GetAppSubmissionsByAppId {}
export interface GetAppSubmissionsByAppIdApiResponse extends ApiResponse {}
export interface GetAppSubmissionsByAppIdApiArgs {
  osId?: string;
  workspaceId?: string;
  appId?: string;
}
export interface GetAppSubmissionsByAppIdPluginArgs
  extends GetAppSubmissionsByAppIdApiArgs {}

/// CreateAppSubmission
export interface CreateAppSubmission {}
export interface CreateAppSubmissionApiResponse extends ApiResponse {}
export interface CreateAppSubmissionApiArgs extends Partial<AppSubmission> {}
export interface CreateAppSubmissionPluginArgs
  extends CreateAppSubmissionApiArgs {}

/// UpdateAppSubmission
export interface UpdateAppSubmission {}
export interface UpdateAppSubmissionApiResponse extends ApiResponse {}
export interface UpdateAppSubmissionApiArgs extends Partial<AppSubmission> {}
export interface UpdateAppSubmissionPluginArgs
  extends UpdateAppSubmissionApiArgs {}

/// DeleteAppSubmission
export interface DeleteAppSubmission {}
export interface DeleteAppSubmissionApiResponse extends ApiResponse {}
export interface DeleteAppSubmissionApiArgs {
  id: string;
}
export interface DeleteAppSubmissionPluginArgs
  extends DeleteAppSubmissionApiArgs {}
