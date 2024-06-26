import { z } from "zod";

//import { PermissionRole, PermissionType } from "@quarkloop/types";
import { historySchema, visibilitySchema } from "@/components/Utils";

export const workspaceSchema = historySchema.merge(
    z.object({
        id: z.number(),
        sid: z.string(),
        name: z.string(),
        visibility: visibilitySchema,
        description: z.string(),
        path: z.string(),
    })
);

export type Workspace = z.infer<typeof workspaceSchema>;
export type WorkspaceVisibility = z.infer<typeof visibilitySchema>;

export const mutationFormSchema = workspaceSchema.omit({
    id: true,
    path: true,
    createdAt: true,
    createdBy: true,
    updatedAt: true,
    updatedBy: true,
});

// export interface WorkspaceUser {
//     workspaceSid: string | null;
//     type: PermissionType;
//     role: PermissionRole;
//     createdAt: Date;
//     user: {
//         id: string;
//         name: string | null;
//         email: string;
//         image: string | null;
//     };
// }
