import { AppConversationData as PrismaAppConversationData } from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface AppConversationData
  extends Partial<PrismaAppConversationData> {}

export type AppConversationDataPluginArgs =
  | GetAppConversationDataByIdPluginArgs
  | GetAppConversationDataByAppSubmissionIdPluginArgs
  | CreateAppConversationDataPluginArgs
  | UpdateAppConversationDataPluginArgs
  | DeleteAppConversationDataPluginArgs;

/// GetAppConversationDataById
export interface GetAppConversationDataById {}
export interface GetAppConversationDataByIdApiResponse extends ApiResponse {}
export interface GetAppConversationDataByIdApiArgs {
  id: number;
  appSubmissionId: string;
}
export interface GetAppConversationDataByIdPluginArgs
  extends GetAppConversationDataByIdApiArgs {}

/// GetAppConversationDataByAppSubmissionId
export interface GetAppConversationDataByAppSubmissionId {}
export interface GetAppConversationDataByAppSubmissionIdApiResponse
  extends ApiResponse {}
export interface GetAppConversationDataByAppSubmissionIdApiArgs {
  appSubmissionId: string;
}
export interface GetAppConversationDataByAppSubmissionIdPluginArgs
  extends GetAppConversationDataByAppSubmissionIdApiArgs {}

/// CreateAppConversationData
export interface CreateAppConversationData {}
export interface CreateAppConversationDataApiResponse extends ApiResponse {}
export interface CreateAppConversationDataApiArgs
  extends Partial<AppConversationData> {
  appSubmissionId: string;
}
export interface CreateAppConversationDataPluginArgs
  extends CreateAppConversationDataApiArgs {}

/// UpdateAppConversationData
export interface UpdateAppConversationData {}
export interface UpdateAppConversationDataApiResponse extends ApiResponse {}
export interface UpdateAppConversationDataApiArgs
  extends Partial<AppConversationData> {
  id: number;
  appSubmissionId: string;
}
export interface UpdateAppConversationDataPluginArgs
  extends UpdateAppConversationDataApiArgs {}

/// DeleteAppConversationData
export interface DeleteAppConversationData {}
export interface DeleteAppConversationDataApiResponse extends ApiResponse {}
export interface DeleteAppConversationDataApiArgs {
  id: number;
  appSubmissionId: string;
}
export interface DeleteAppConversationDataPluginArgs
  extends DeleteAppConversationDataApiArgs {}
