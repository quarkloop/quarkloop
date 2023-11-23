import { z } from "zod";

export const serviceSchema = z.object({
    id: z.string(),
    name: z.string(),
    type: z.number(),
    description: z.string(),
    metadata: z.any(),
    data: z.any(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Service = z.infer<typeof serviceSchema>;
