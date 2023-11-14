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
    osId: string;
};

/// GetWorkspacesByOsId
export const getWorkspacesByOsIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(workspaceSchema),
    })
);
export type GetWorkspacesByOsIdApiResponse = z.infer<
    typeof getWorkspacesByOsIdSchema
>;
export type GetWorkspacesByOsIdApiArgs = {
    osId: string;
};

/// CreateWorkspace
export const createWorkspaceSchema = apiResponseV2Schema.merge(
    z.object({
        data: workspaceSchema,
    })
);
export type CreateWorkspaceApiResponse = z.infer<typeof createWorkspaceSchema>;
export type CreateWorkspaceApiArgs = {
    osId: string;
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
    osId: string;
    workspace: Partial<z.infer<typeof workspaceSchema>>;
};

/// DeleteWorkspace
export const deleteWorkspaceSchema = apiResponseV2Schema.merge(
    z.object({
        data: workspaceSchema,
    })
);
export type DeleteWorkspaceApiResponse = z.infer<typeof deleteWorkspaceSchema>;
export type DeleteWorkspaceApiArgs = {
    id: string;
    osId: string;
};
