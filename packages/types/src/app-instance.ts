import { z } from "zod";

export const appInstanceSchema = z.object({
    id: z.string(),
    name: z.string(),
    path: z.string(),
    description: z.string(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type AppInstance = z.infer<typeof appInstanceSchema>;
