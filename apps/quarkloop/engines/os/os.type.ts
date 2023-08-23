import {
  OperatingSystem as PrismaOperatingSystem,
  PermissionType,
  PermissionRole,
} from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface OperatingSystem extends Partial<PrismaOperatingSystem> {
  isOwner: boolean;
}

export interface OperatingSystemUser {
  osId: string | null;
  type: PermissionType;
  role: PermissionRole;
  createdAt: Date;
  user: {
    id: string;
    name: string | null;
    email: string;
    image: string | null;
  };
}

export type OperatingSystemPluginArgs =
  | GetOperatingSystemByIdPluginArgs
  | GetOperatingSystemUsersPluginArgs
  | CreateOperatingSystemPluginArgs
  | UpdateOperatingSystemPluginArgs
  | DeleteOperatingSystemPluginArgs;

export type OperatingSystemApiResponse = GetOperatingSystemUsersApiResponse;

/// GetOperatingSystems
export interface GetOperatingSystems {}
export interface GetOperatingSystemsApiResponse extends ApiResponse {}
export type GetOperatingSystemsApiArgs = void;
export type GetOperatingSystemsPluginArgs = void;

/// GetOperatingSystemUsers
export interface GetOperatingSystemUsers {}
export interface GetOperatingSystemUsersApiResponse extends ApiResponse {
  mydata?: OperatingSystemUser[];
}
export interface GetOperatingSystemUsersApiArgs {
  id: string;
  workspaceId?: string;
}
export interface GetOperatingSystemUsersPluginArgs
  extends GetOperatingSystemUsersApiArgs {}

/// GetOperatingSystemById
export interface GetOperatingSystemById {}
export interface GetOperatingSystemByIdApiResponse extends ApiResponse {}
export interface GetOperatingSystemByIdApiArgs {
  id: string;
}
export interface GetOperatingSystemByIdPluginArgs
  extends GetOperatingSystemByIdApiArgs {}

/// GetOperatingSystemsByUserId
export interface GetOperatingSystemsByUserId {}
export interface GetOperatingSystemsByUserIdApiResponse extends ApiResponse {}
export type GetOperatingSystemsByUserIdApiArgs = void;
export type GetOperatingSystemsByUserIdPluginArgs = void;

/// CreateOperatingSystem
export interface CreateOperatingSystem {}
export interface CreateOperatingSystemApiResponse extends ApiResponse {}
export interface CreateOperatingSystemApiArgs
  extends Partial<OperatingSystem> {}
export interface CreateOperatingSystemPluginArgs
  extends CreateOperatingSystemApiArgs {}

/// UpdateOperatingSystem
export interface UpdateOperatingSystem {}
export interface UpdateOperatingSystemApiResponse extends ApiResponse {}
export interface UpdateOperatingSystemApiArgs
  extends Partial<OperatingSystem> {}
export interface UpdateOperatingSystemPluginArgs
  extends UpdateOperatingSystemApiArgs {}

/// DeleteOperatingSystem
export interface DeleteOperatingSystem {}
export interface DeleteOperatingSystemApiResponse extends ApiResponse {}
export interface DeleteOperatingSystemApiArgs {
  id: string;
}
export interface DeleteOperatingSystemPluginArgs
  extends DeleteOperatingSystemApiArgs {}
