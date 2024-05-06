import { z } from "zod";

/// GetUserAccess
export const getUserAccessApiArgsSchema = z.object({
    orgSid: z.string(),
    workspaceSid: z.string().optional(),
});
export type GetUserAccessApiArgs = z.infer<typeof getUserAccessApiArgsSchema>;

export const getUserAccessApiResponseSchema = z.string();
export type GetUserAccessApiResponse = z.infer<
    typeof getUserAccessApiResponseSchema
>;

/// GrantUserAccess
export const grantUserAccessApiArgsSchema = z.object({});
export type GrantUserAccessApiArgs = z.infer<
    typeof grantUserAccessApiArgsSchema
>;

export const grantUserAccessApiResponseSchema = z.object({});
export type GrantUserAccessApiResponse = z.infer<
    typeof grantUserAccessApiResponseSchema
>;

/// RevokeUserAccess
export const revokeUserAccessByIdApiArgsSchema = z.object({});
export type RevokeUserAccessByIdApiArgs = z.infer<
    typeof revokeUserAccessByIdApiArgsSchema
>;

export type RevokeUserAccessByIdApiResponse = void;
