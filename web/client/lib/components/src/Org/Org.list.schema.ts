import { z } from "zod";
import { orgSchema } from "./Org.schema";

export const orgRowSchema = orgSchema
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

export type OrgRow = z.infer<typeof orgRowSchema>;
