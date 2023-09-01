import { AppThreadData as PrismaAppThreadData } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

/// Types
export interface AppThreadData extends Partial<PrismaAppThreadData> {}

export type AppThreadDataPluginArgs =
  | GetAppThreadDataByIdPluginArgs
  | GetAppThreadDataByAppSubmissionIdPluginArgs
  | CreateAppThreadDataPluginArgs
  | UpdateAppThreadDataPluginArgs
  | DeleteAppThreadDataPluginArgs;

/// GetAppThreadDataById
export interface GetAppThreadDataById {}
export interface GetAppThreadDataByIdApiResponse extends ApiResponse {}
export interface GetAppThreadDataByIdApiArgs {
  id: number;
  appSubmissionId: string;
}
export interface GetAppThreadDataByIdPluginArgs
  extends GetAppThreadDataByIdApiArgs {}

/// GetAppThreadDataByAppSubmissionId
export interface GetAppThreadDataByAppSubmissionId {}
export interface GetAppThreadDataByAppSubmissionIdApiResponse
  extends ApiResponse {}
export interface GetAppThreadDataByAppSubmissionIdApiArgs {
  appSubmissionId: string;
}
export interface GetAppThreadDataByAppSubmissionIdPluginArgs
  extends GetAppThreadDataByAppSubmissionIdApiArgs {}

/// CreateAppThreadData
export interface CreateAppThreadData {}
export interface CreateAppThreadDataApiResponse extends ApiResponse {}
export interface CreateAppThreadDataApiArgs extends Partial<AppThreadData> {
  appSubmissionId: string;
}
export interface CreateAppThreadDataPluginArgs
  extends CreateAppThreadDataApiArgs {}

/// UpdateAppThreadData
export interface UpdateAppThreadData {}
export interface UpdateAppThreadDataApiResponse extends ApiResponse {}
export interface UpdateAppThreadDataApiArgs extends Partial<AppThreadData> {
  id: number;
  appSubmissionId: string;
}
export interface UpdateAppThreadDataPluginArgs
  extends UpdateAppThreadDataApiArgs {}

/// DeleteAppThreadData
export interface DeleteAppThreadData {}
export interface DeleteAppThreadDataApiResponse extends ApiResponse {}
export interface DeleteAppThreadDataApiArgs {
  id: number;
  appSubmissionId: string;
}
export interface DeleteAppThreadDataPluginArgs
  extends DeleteAppThreadDataApiArgs {}
