import { AppConversationSettings as PrismaAppConversationSettings } from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface AppConversationSettings
  extends Partial<PrismaAppConversationSettings> {}

export type AppConversationSettingsPluginArgs =
  | GetAppConversationSettingsByIdPluginArgs
  | GetAppConversationSettingsByAppIdPluginArgs
  | CreateAppConversationSettingsPluginArgs
  | UpdateAppConversationSettingsPluginArgs
  | DeleteAppConversationSettingsPluginArgs;

/// GetAppConversationSettingsById
export interface GetAppConversationSettingsById {}
export interface GetAppConversationSettingsByIdApiResponse
  extends ApiResponse {}
export interface GetAppConversationSettingsByIdApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAppConversationSettingsByIdPluginArgs
  extends GetAppConversationSettingsByIdApiArgs {}

/// GetAppConversationSettingsByAppId
export interface GetAppConversationSettingsByAppId {}
export interface GetAppConversationSettingsByAppIdApiResponse
  extends ApiResponse {}
export interface GetAppConversationSettingsByAppIdApiArgs {
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface GetAppConversationSettingsByAppIdPluginArgs
  extends GetAppConversationSettingsByAppIdApiArgs {}

/// CreateAppConversationSettings
export interface CreateAppConversationSettings {}
export interface CreateAppConversationSettingsApiResponse extends ApiResponse {}
export interface CreateAppConversationSettingsApiArgs
  extends Partial<AppConversationSettings> {
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface CreateAppConversationSettingsPluginArgs
  extends CreateAppConversationSettingsApiArgs {}

/// UpdateAppConversationSettings
export interface UpdateAppConversationSettings {}
export interface UpdateAppConversationSettingsApiResponse extends ApiResponse {}
export interface UpdateAppConversationSettingsApiArgs
  extends Partial<AppConversationSettings> {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface UpdateAppConversationSettingsPluginArgs
  extends UpdateAppConversationSettingsApiArgs {}

/// DeleteAppConversationSettings
export interface DeleteAppConversationSettings {}
export interface DeleteAppConversationSettingsApiResponse extends ApiResponse {}
export interface DeleteAppConversationSettingsApiArgs {
  id: string;
  osId: string;
  workspaceId: string;
  appId: string;
}
export interface DeleteAppConversationSettingsPluginArgs
  extends DeleteAppConversationSettingsApiArgs {}
