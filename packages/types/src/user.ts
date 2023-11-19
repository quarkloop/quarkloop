// import {
//     User as PrismaUser,
//     Account as PrismaUserAccount,
//     Session as PrismaUserSession,
// } from "@quarkloop/prisma/types";
import { z } from "zod";
import { apiResponseSchema } from "./api-response";

export const userSchema = z.object({
    id: z.string(),
    name: z.string().min(2).max(40),
    email: z.string().email(),
    emailVerified: z.date().nullable(),
    password: z.string().min(6).max(20),
    passwordSalt: z.string().nullable(),
    birthdate: z.date().nullable(),
    country: z.string().nullable(),
    image: z.string().nullable(),
    createdAt: z.date(),
    updatedAt: z.date().nullable(),
    status: z.number(),
});

export const userAccountSchema = z.object({
    id: z.string(),
    type: z.string(),
    provider: z.string(),
    providerAccountId: z.string(),
    refresh_token: z.string().nullable(),
    access_token: z.string().nullable(),
    expires_at: z.number().nullable(),
    token_type: z.string().nullable(),
    scope: z.string().nullable(),
    id_token: z.string().nullable(),
    session_state: z.string().nullable(),
    userId: z.string(),
});

export const userSessionSchema = z.object({
    id: z.string(),
    userId: z.string(),
    sessionToken: z.string(),
    expires: z.date().nullable(),
    expiresAsString: z.string(),
    isCurrent: z.boolean(),
});

// export interface User {
//     id: string | null;
//     name: string | null;
//     email: string | null;
//     emailVerified: Date | null;
//     password: string | null;
//     passwordSalt: string | null;
//     birthdate: Date | null;
//     country: string | null;
//     image: string | null;
//     createdAt: Date | null;
//     updatedAt: Date | null;
//     status: number;
// }

// export interface UserAccount {
//     id: string;
//     type: string;
//     provider: string;
//     providerAccountId: string;
//     refresh_token: string | null;
//     access_token: string | null;
//     expires_at: number | null;
//     token_type: string | null;
//     scope: string | null;
//     id_token: string | null;
//     session_state: string | null;
//     userId: string;
// }

export interface UserPermissions {
    org: {
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
    project: {
        canRead: boolean;
        canCreate: boolean;
        canUpdate: boolean;
        canDelete: boolean;
    };
}

export enum UserAccountStatus {
    Inactive = 0,
    Active = 1,
    Suspended = 2,
    Deleted = 3,
}

export type User = z.infer<typeof userSchema>;
export type UserSession = z.infer<typeof userSessionSchema>;
export type UserAccount = z.infer<typeof userAccountSchema>;

// /// GetUser
// export const getUserApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type GetUserApiResponse = z.infer<typeof getUserApiResponseSchema>;
// export type GetUserApiArgs = void;

// /// GetUserPermissions
// export const getUserPermissionsApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type GetUserPermissionsApiResponse = z.infer<
//     typeof getUserPermissionsApiResponseSchema
// >;

// export const getUserPermissionsApiArgsSchema = userSchema
//     .merge(
//         z.object({
//             orgId: z.string().optional(),
//             workspaceId: z.string().optional(),
//             projectId: z.string().optional(),
//         })
//     )
//     .extend(apiResponseSchema.shape);
// export type GetUserPermissionsApiArgs = z.infer<
//     typeof getUserPermissionsApiArgsSchema
// >;

// /// DeleteUser
// export const deleteUserApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type DeleteUserApiResponse = z.infer<typeof deleteUserApiResponseSchema>;
// export const deleteUserApiArgsSchema = userSchema
//     .merge(
//         z.object({
//             orgId: z.string().optional(),
//             workspaceId: z.string().optional(),
//             projectId: z.string().optional(),
//         })
//     )
//     .extend(apiResponseSchema.shape);
// export type DeleteUserApiArgs = z.infer<typeof deleteUserApiArgsSchema>;

// /// UpdateUser
// export const updateUserApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type UpdateUserApiResponse = z.infer<typeof updateUserApiResponseSchema>;
// export const updateUserApiArgsSchema = userSchema
//     .merge(
//         z.object({
//             orgId: z.string().optional(),
//             workspaceId: z.string().optional(),
//             projectId: z.string().optional(),
//         })
//     )
//     .extend(apiResponseSchema.shape);
// export type UpdateUserApiArgs = z.infer<typeof updateUserApiArgsSchema>;

// /// GetUserLinkedAccounts
// export const GetUserLinkedAccountsApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type GetUserLinkedAccountsApiResponse = z.infer<
//     typeof GetUserLinkedAccountsApiResponseSchema
// >;
// export type GetUserLinkedAccountsApiArgs = void;

// /// GetServerSession
// export const getServerSessionApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type GetServerSessionApiResponse = z.infer<
//     typeof getServerSessionApiResponseSchema
// >;
// export type GetServerSessionApiArgs = void;

// /// GetUserSessions
// export const getUserSessionsApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type GetUserSessionsApiResponse = z.infer<
//     typeof getUserSessionsApiResponseSchema
// >;
// export type GetUserSessionsApiArgs = void;

// /// TerminateUserSession
// export const terminateUserSessionApiResponseSchema = userSchema.extend(
//     apiResponseSchema.shape
// );
// export type TerminateUserSessionApiResponse = z.infer<
//     typeof terminateUserSessionApiResponseSchema
// >;
// export type TerminateUserSessionApiArgs = z.infer<typeof userSessionSchema>;
