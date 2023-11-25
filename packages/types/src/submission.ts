import { z } from "zod";

export const submissionSchema = z.object({
    id: z.number(),
    title: z.string(),
    status: z.number(),
    labels: z.array(z.string()),
    dueDate: z.coerce.date().optional(),
    metadata: z.any(),
    data: z.any(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Submission = z.infer<typeof submissionSchema>;
