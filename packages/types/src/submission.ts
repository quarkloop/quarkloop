import { z } from "zod";

export const submissionSchema = z.object({
    id: z.string(),
    title: z.string(),
    metadata: z.any(),
    data: z.any(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Submission = z.infer<typeof submissionSchema>;
