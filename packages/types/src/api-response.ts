import { z } from "zod";

export const apiResponseSchema = z.object({
    status: z.any(),
    error: z.string(),
    errorDetails: z.record(z.string(), z.any()),
    database: z.any(),
});

export type ApiResponse = z.infer<typeof apiResponseSchema>;

// export type ApiResponse = {
//     status: any; // StatusState
//     error?: string;
//     errorDetails?: Record<string, any>;
//     database?: any; // DatabaseState
// };
