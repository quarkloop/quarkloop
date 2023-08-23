import {
  User as PrismaUser,
  Account as PrismaUserAccount,
  Session as PrismaUserSession,
} from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface User extends Partial<PrismaUser> {}
export interface UserAccount extends Partial<PrismaUserAccount> {}
export interface UserSession extends Partial<PrismaUserSession> {
  isCurrent: boolean;
}

export interface UserPermissions {
  os: {
    canRead: boolean;
    canCreate: boolean;
    canUpdate: boolean;
    canDelete: boolean;
  };
  workspace: {
    canRead: boolean;
    canCreate: boolean;
    canUpdate: boolean;
    canDelete: boolean;
  };
  app: {
    canRead: boolean;
    canCreate: boolean;
    canUpdate: boolean;
    canDelete: boolean;
  };
}

export type UserPluginArgs =
  | GetUserPluginArgs
  | GetUserPermissionsPluginArgs
  | DeleteUserPluginArgs
  | UpdateUserPluginArgs
  | GetUserLinkedAccountsPluginArgs
  | GetServerSessionPluginArgs
  | GetUserSessionsPluginArgs
  | TerminateUserSessionPluginArgs;

export interface UserAccountPluginArgs
  extends GetUserLinkedAccountsPluginArgs {}

export interface UserSessionPluginArgs
  extends GetServerSessionPluginArgs,
    GetUserSessionsPluginArgs,
    TerminateUserSessionPluginArgs {}

export enum UserAccountStatus {
  Inactive = 0,
  Active = 1,
  Suspended = 2,
  Deleted = 3,
}

/// GetUser
export interface GetUser {}
export interface GetUserApiResponse extends ApiResponse {}
export type GetUserApiArgs = void;
export interface GetUserPluginArgs {}

/// GetUserPermissions
export interface GetUserPermissions {}
export interface GetUserPermissionsApiResponse extends ApiResponse {}
export interface GetUserPermissionsApiArgs {
  osId?: string;
  workspaceId?: string;
  appId?: string;
}
export interface GetUserPermissionsPluginArgs
  extends GetUserPermissionsApiArgs {}

/// DeleteUser
export interface DeleteUser extends PrismaUser {}
export interface DeleteUserApiResponse extends ApiResponse {}
export interface DeleteUserApiArgs extends Partial<User> {}
export interface DeleteUserPluginArgs extends DeleteUserApiArgs {}

/// UpdateUser
export interface UpdateUser {}
export interface UpdateUserApiResponse extends ApiResponse {}
export interface UpdateUserApiArgs extends Partial<User> {}
export interface UpdateUserPluginArgs extends UpdateUserApiArgs {}

/// GetUserLinkedAccounts
export interface GetUserLinkedAccounts {}
export interface GetUserLinkedAccountsApiResponse extends ApiResponse {}
export type GetUserLinkedAccountsApiArgs = void;
export interface GetUserLinkedAccountsPluginArgs {}

/// GetServerSession
export interface GetServerSession {}
export interface GetServerSessionApiResponse extends ApiResponse {}
export type GetServerSessionApiArgs = void;
export interface GetServerSessionPluginArgs {}

/// GetUserSessions
export interface GetUserSessions {}
export interface GetUserSessionsApiResponse extends ApiResponse {}
export type GetUserSessionsApiArgs = void;
export interface GetUserSessionsPluginArgs {}

/// TerminateUserSession
export interface TerminateUserSession {}
export interface TerminateUserSessionApiResponse extends ApiResponse {}
export interface TerminateUserSessionApiArgs extends Partial<UserSession> {}
export interface TerminateUserSessionPluginArgs
  extends TerminateUserSessionApiArgs {}
