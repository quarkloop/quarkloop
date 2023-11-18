import { z } from "zod";
import { apiResponseV2Schema, workspaceSchema } from "@quarkloop/types";

/// GetWorkspaceById
export const getWorkspaceByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: workspaceSchema,
    })
);
export type GetWorkspaceByIdApiResponse = z.infer<
    typeof getWorkspaceByIdSchema
>;
export type GetWorkspaceByIdApiArgs = {
    id: string;
    orgId: string;
};

/// GetWorkspacesByOrgId
export const getWorkspacesByOrgIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(workspaceSchema),
    })
);
export type GetWorkspacesByOrgIdApiResponse = z.infer<
    typeof getWorkspacesByOrgIdSchema
>;
export type GetWorkspacesByOrgIdApiArgs = {
    orgId: string[];
};

/// CreateWorkspace
export const createWorkspaceSchema = apiResponseV2Schema.merge(
    z.object({
        data: workspaceSchema,
    })
);
export type CreateWorkspaceApiResponse = z.infer<typeof createWorkspaceSchema>;
export type CreateWorkspaceApiArgs = {
    orgId: string;
    workspace: Partial<z.infer<typeof workspaceSchema>>;
};

/// UpdateWorkspace
export const updateWorkspaceSchema = apiResponseV2Schema.merge(
    z.object({
        data: workspaceSchema,
    })
);
export type UpdateWorkspaceApiResponse = z.infer<typeof updateWorkspaceSchema>;
export type UpdateWorkspaceApiArgs = {
    orgId: string;
    workspace: Partial<z.infer<typeof workspaceSchema>>;
};

/// DeleteWorkspace
export const deleteWorkspaceSchema = apiResponseV2Schema.merge(
    z.object({
        data: workspaceSchema,
    })
);
export type DeleteWorkspaceApiResponse = void;
export type DeleteWorkspaceApiArgs = {
    id: string;
    orgId: string;
};
