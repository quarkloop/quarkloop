import { z } from "zod";
import { workspaceSchema } from "./Workspace.schema";

export const workspaceRowSchema = workspaceSchema
    .pick({
        id: true,
        sid: true,
        name: true,
        description: true,
        visibility: true,
        path: true,
        createdAt: true,
        createdBy: true,
        updatedAt: true,
        updatedBy: true,
    })
    .merge(
        z.object({
            onViewClick: z.function().args(z.any()).optional(),
            onUpdateClick: z.function().args(z.any()).optional(),
            onDeleteClick: z.function().args(z.any()).optional(),
        })
    );

export type WorkspaceRow = z.infer<typeof workspaceRowSchema>;
