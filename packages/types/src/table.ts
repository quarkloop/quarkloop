import { z } from "zod";

export const tableSchema = z.object({
    id: z.string(),
    name: z.string(),
    type: z.number(),
    description: z.string(),
    metadata: z.any(),
    data: z.any(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Table = z.infer<typeof tableSchema>;
