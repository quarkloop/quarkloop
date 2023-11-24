import { z } from "zod";
import {
    apiResponseV2Schema,
    submissionSchema,
    Submission,
} from "@quarkloop/types";

/// GetSubmissionList
export const getSubmissionListSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(submissionSchema),
    })
);
export type GetSubmissionListApiResponse = z.infer<
    typeof getSubmissionListSchema
>;
export type GetSubmissionListApiArgs = {
    projectId: string;
};

/// GetSubmissionById
export const getSubmissionByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: submissionSchema,
    })
);
export type GetSubmissionByIdApiResponse = z.infer<
    typeof getSubmissionByIdSchema
>;
export type GetSubmissionByIdApiArgs = {
    projectId: string;
    submissionId: string;
};

/// CreateSubmission
export const createSubmissionSchema = apiResponseV2Schema.merge(
    z.object({
        data: submissionSchema,
    })
);
export type CreateSubmissionApiResponse = z.infer<
    typeof createSubmissionSchema
>;
export type CreateSubmissionApiArgs = {
    projectId: string;
    submission: Partial<Submission>;
};

/// UpdateSubmission
export const updateSubmissionSchema = apiResponseV2Schema.merge(
    z.object({
        data: submissionSchema,
    })
);
export type UpdateSubmissionApiResponse = z.infer<
    typeof updateSubmissionSchema
>;
export type UpdateSubmissionApiArgs = {
    projectId: string;
    submissionId: string;
    submission: Partial<Submission>;
};

/// DeleteSubmission
export const deleteSubmissionSchema = apiResponseV2Schema;
export type DeleteSubmissionApiResponse = z.infer<
    typeof deleteSubmissionSchema
>;
export type DeleteSubmissionApiArgs = {
    projectId: string;
    submissionId: string;
};
