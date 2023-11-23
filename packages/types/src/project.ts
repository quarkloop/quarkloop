import { z } from "zod";

export const projectSchema = z.object({
    id: z.string(),
    name: z.string(),
    accessType: z.number(),
    path: z.string(),
    description: z.string(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Project = z.infer<typeof projectSchema>;

export enum ProjectAccessType {
    public = 1,
    private,
}
