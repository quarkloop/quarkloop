import { z } from "zod";
import {
    apiResponseSchema,
    userSchema,
    userSessionSchema,
} from "@quarkloop/types";

/// GetUser
export const getUserApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type GetUserApiResponse = z.infer<typeof getUserApiResponseSchema>;
export type GetUserApiArgs = void;

/// GetUserPermissions
export const getUserPermissionsApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type GetUserPermissionsApiResponse = z.infer<
    typeof getUserPermissionsApiResponseSchema
>;
export const getUserPermissionsApiArgsSchema = z.object({
    userId: z.string(),
});
export type GetUserPermissionsApiArgs = z.infer<
    typeof getUserPermissionsApiArgsSchema
>;

/// DeleteUser
export const deleteUserApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type DeleteUserApiResponse = z.infer<typeof deleteUserApiResponseSchema>;
export const deleteUserApiArgsSchema = z.object({
    email: z.string(),
});
export type DeleteUserApiArgs = z.infer<typeof deleteUserApiArgsSchema>;

/// UpdateUser
export const updateUserApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type UpdateUserApiResponse = z.infer<typeof updateUserApiResponseSchema>;
export const updateUserApiArgsSchema = userSchema
    .merge(
        z.object({
            email: z.string(),
        })
    )
    .extend(apiResponseSchema.shape);
export type UpdateUserApiArgs = z.infer<typeof updateUserApiArgsSchema>;

/// GetUserLinkedAccounts
export const getUserLinkedAccountsApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type GetUserLinkedAccountsApiResponse = z.infer<
    typeof getUserLinkedAccountsApiResponseSchema
>;
export type GetUserLinkedAccountsApiArgs = void;

/// GetServerSession
export const getServerSessionApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type GetServerSessionApiResponse = z.infer<
    typeof getServerSessionApiResponseSchema
>;
export type GetServerSessionApiArgs = void;

/// GetUserSessions
export const getUserSessionsApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type GetUserSessionsApiResponse = z.infer<
    typeof getUserSessionsApiResponseSchema
>;
export type GetUserSessionsApiArgs = void;

/// TerminateUserSession
export const terminateUserSessionApiResponseSchema = userSchema.extend(
    apiResponseSchema.shape
);
export type TerminateUserSessionApiResponse = z.infer<
    typeof terminateUserSessionApiResponseSchema
>;
const terminateUserSessionApiArgsSchema = z.object({
    sessionToken: z.string(),
});
export type TerminateUserSessionApiArgs = z.infer<
    typeof terminateUserSessionApiArgsSchema
>;
