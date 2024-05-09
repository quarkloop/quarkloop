import { z } from "zod";

import { apiResponseV2Schema } from "@quarkloop/types";
import { visibilitySchema } from "@/components/Utils";

import { Workspace, workspaceSchema } from "./Workspace.schema";
import { workspaceMemberSchema } from "./Workspace.members.schema";

/// GetWorkspaceById
export const getWorkspaceByIdSchema = z.object({
    data: workspaceSchema,
});
export type GetWorkspaceByIdApiArgs = {
    orgSid: string;
    workspaceSid: string;
};
export type GetWorkspaceByIdApiResponse = z.infer<
    typeof getWorkspaceByIdSchema
>;

/// CreateWorkspace
export const createWorkspaceApiArgsSchema = z.object({
    orgSid: z.string(),
    payload: workspaceSchema.pick({
        sid: true,
        name: true,
        description: true,
        visibility: true,
    }),
});
export const createWorkspaceApiResponseSchema = z.object({
    data: workspaceSchema,
});
export type CreateWorkspaceApiArgs = z.infer<
    typeof createWorkspaceApiArgsSchema
>;
export type CreateWorkspaceApiResponse = z.infer<
    typeof createWorkspaceApiResponseSchema
>;

/// UpdateWorkspace
export const updateWorkspaceByIdApiArgsSchema = z.object({
    orgSid: z.string(),
    workspaceSid: z.string(),
    payload: workspaceSchema.pick({
        sid: true,
        name: true,
        description: true,
        visibility: true,
    }),
});
export type UpdateWorkspaceByIdApiArgs = z.infer<
    typeof updateWorkspaceByIdApiArgsSchema
>;
export type UpdateWorkspaceByIdApiResponse = void;

/// DeleteWorkspace
export const deleteWorkspaceByIdApiArgsSchema = z.object({
    orgSid: z.string(),
    workspaceSid: z.string(),
});
export type DeleteWorkspaceByIdApiArgs = z.infer<
    typeof deleteWorkspaceByIdApiArgsSchema
>;
export type DeleteWorkspaceByIdApiResponse = void;

/// ChangeWorkspaceVisibility
export const changeWorkspaceVisibilityApiArgsSchema = z.object({
    orgSid: z.string(),
    workspaceSid: z.string(),
    visibility: visibilitySchema,
});
export type ChangeWorkspaceVisibilityApiArgs = z.infer<
    typeof changeWorkspaceVisibilityApiArgsSchema
>;
export type ChangeWorkspaceVisibilityApiResponse = void;

/// GetWorkspaceMembers
export const getWorkspaceMembersApiArgsSchema = z.object({
    orgSid: z.string(),
    workspaceSid: z.string(),
});
export const getWorkspaceMembersApiResponseSchema = z.object({
    data: z.array(workspaceMemberSchema),
});

export type GetWorkspaceMembersApiArgs = z.infer<
    typeof getWorkspaceMembersApiArgsSchema
>;
export type GetWorkspaceMembersApiResponse = z.infer<
    typeof getWorkspaceMembersApiResponseSchema
>;

/// GetWorkspaces
export const getWorkspacesSchema = z.object({
    data: z.array(workspaceSchema),
});
export type GetWorkspacesApiArgs = void;
export type GetWorkspacesApiResponse = z.infer<typeof getWorkspacesSchema>;

/// GetWorkspacesByUserId
export const getWorkspacesByUserIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(workspaceSchema),
    })
);
export type GetWorkspacesByUserIdApiArgs = void;
export type GetWorkspacesByUserIdApiResponse = z.infer<
    typeof getWorkspacesByUserIdSchema
>;
