import { z } from "zod";

export const workspaceSchema = z.object({
    id: z.string(),
    name: z.string(),
    accessType: z.number(),
    path: z.string(),
    description: z.string(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Workspace = z.infer<typeof workspaceSchema>;

// export interface Workspace {
//     relId: number;
//     id: string;
//     name: string;
//     path: string;
//     description: string | null;
//     accessType: number | null;
//     imageUrl: string | null;
//     createdAt: Date;
//     orgId: string;
// }

export enum WorkspaceAccessType {
    Public = 1,
    Private = 2,
}
