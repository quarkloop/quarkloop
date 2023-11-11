import { z } from "zod";

export const appSchema = z.object({
    id: z.string(),
    name: z.string(),
    accessType: z.number(),
    path: z.string(),
    description: z.string(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type App = z.infer<typeof appSchema>;

export enum AccessType {
    public = 1,
    private,
}

export type AppStatus = "On" | "Off" | "Archived";
