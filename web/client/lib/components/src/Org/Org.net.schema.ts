import { z } from "zod";

import { apiResponseV2Schema } from "@quarkloop/types";
import { visibilitySchema } from "@/components/Utils";

import { Org, orgSchema } from "./Org.schema";
import { orgMemberSchema } from "./Org.members.schema";

/// GetOrgById
export const getOrgByIdSchema = z.object({
    data: orgSchema,
});
export type GetOrgByIdApiArgs = {
    orgSid: string;
};
export type GetOrgByIdApiResponse = z.infer<typeof getOrgByIdSchema>;

/// GetOrgWorkspaceList
export const getOrgWorkspaceListSchema = z.object({
    data: z.array(z.any()), // TODO: change with workspace array
});
export type GetOrgWorkspaceListApiArgs = {
    orgSid: string;
};
export type GetOrgWorkspaceListApiResponse = z.infer<
    typeof getOrgWorkspaceListSchema
>;

/// CreateOrg
export const createOrgApiArgsSchema = z.object({
    payload: orgSchema.pick({
        sid: true,
        name: true,
        description: true,
        visibility: true,
    }),
});
export const createOrgApiResponseSchema = z.object({
    data: orgSchema,
});
export type CreateOrgApiArgs = z.infer<typeof createOrgApiArgsSchema>;
export type CreateOrgApiResponse = z.infer<typeof createOrgApiResponseSchema>;

/// UpdateOrg
export const updateOrgByIdApiArgsSchema = z.object({
    orgSid: z.string(),
    payload: orgSchema.pick({
        sid: true,
        name: true,
        description: true,
        visibility: true,
    }),
});
export type UpdateOrgByIdApiArgs = z.infer<typeof updateOrgByIdApiArgsSchema>;
export type UpdateOrgByIdApiResponse = void;

/// DeleteOrg
export const deleteOrgByIdApiArgsSchema = z.object({
    orgSid: z.string(),
});
export type DeleteOrgByIdApiArgs = z.infer<typeof deleteOrgByIdApiArgsSchema>;
export type DeleteOrgByIdApiResponse = void;

/// ChangeOrgVisibility
export const changeOrgVisibilityApiArgsSchema = z.object({
    orgSid: z.string(),
    visibility: visibilitySchema,
});
export type ChangeOrgVisibilityApiArgs = z.infer<
    typeof changeOrgVisibilityApiArgsSchema
>;
export type ChangeOrgVisibilityApiResponse = void;

/// GetOrgMembers
export const getOrgMembersApiArgsSchema = z.object({
    orgSid: z.string(),
});
export const getOrgMembersApiResponseSchema = z.object({
    data: z.array(orgMemberSchema),
});

export type GetOrgMembersApiArgs = z.infer<typeof getOrgMembersApiArgsSchema>;
export type GetOrgMembersApiResponse = z.infer<
    typeof getOrgMembersApiResponseSchema
>;

/// GetOrgs
export const getOrgsSchema = z.object({
    data: z.array(orgSchema),
});
export type GetOrgsApiArgs = void;
export type GetOrgsApiResponse = z.infer<typeof getOrgsSchema>;

/// GetOrgsByUserId
export const getOrgsByUserIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(orgSchema),
    })
);
export type GetOrgsByUserIdApiArgs = void;
export type GetOrgsByUserIdApiResponse = z.infer<typeof getOrgsByUserIdSchema>;
