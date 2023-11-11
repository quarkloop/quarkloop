import { z } from "zod";

import { ApiResponse } from "./api-response";

export const workspaceSchema = z.object({
    relId: z.number(),
    id: z.string(),
    osId: z.string(),
    name: z.string(),
    path: z.string(),
    description: z.string().optional(),
    accessType: z.string().optional(),
    imageUrl: z.string().optional(),
    createdAt: z.coerce.date(),
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
//     osId: string;
// }

export enum WorkspaceAccessType {
    Public = 1,
    Private = 2,
}
