import { z } from "zod";
import { workspaceSchema } from "./Workspace.schema";

export const workspaceRowSchema = workspaceSchema
    // .pick({
    //     id: true,
    //     sid: true,
    //     name: true,
    //     description: true,
    //     visibility: true,
    //     path: true,
    //     createdAt: true,
    //     createdBy: true,
    //     updatedAt: true,
    //     updatedBy: true,
    // })
    .merge(
        z.object({
            orgSid: z.string(),
        })
    );

export type WorkspaceRow = z.infer<typeof workspaceRowSchema> & {
    orgSid: string;
};
