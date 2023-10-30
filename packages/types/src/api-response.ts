import { z } from "zod";

export const apiResponseSchema = z.object({
    status: z.number(),
    statusString: z.string(),
    error: z.string().optional(),
    errorDetails: z.record(z.string(), z.any()),
    database: z.any(),
});

export const apiResponseV2Schema = z.object({
    status: z.number(),
    statusString: z.string(),
    error: z.string().optional(),
    errorDetails: z.record(z.string(), z.any()).optional(),
});

export type ApiResponse = z.infer<typeof apiResponseSchema>;
export type ApiResponseV2 = z.infer<typeof apiResponseV2Schema>;

// export type ApiResponse = {
//     status: any; // StatusState
//     error?: string;
//     errorDetails?: Record<string, any>;
//     database?: any; // DatabaseState
// };
