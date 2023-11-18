import { z } from "zod";
import {
    apiResponseV2Schema,
    appInstanceSchema,
    AppInstance,
} from "@quarkloop/types";

/// GetAppInstanceList
export const getAppInstanceListSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(appInstanceSchema),
    })
);
export type GetAppInstanceListApiResponse = z.infer<
    typeof getAppInstanceListSchema
>;
export type GetAppInstanceListApiArgs = {
    projectId?: string;
};

/// GetAppInstanceById
export const getAppInstanceByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: appInstanceSchema,
    })
);
export type GetAppInstanceByIdApiResponse = z.infer<
    typeof getAppInstanceByIdSchema
>;
export type GetAppInstanceByIdApiArgs = {
    projectId: string;
    instanceId: string;
};

/// CreateAppInstance
export const createAppInstanceSchema = apiResponseV2Schema.merge(
    z.object({
        data: appInstanceSchema,
    })
);
export type CreateAppInstanceApiResponse = z.infer<
    typeof getAppInstanceByIdSchema
>;
export type CreateAppInstanceApiArgs = {
    projectId: string;
    project: AppInstance;
};

/// UpdateAppInstance
export const updateAppInstanceSchema = apiResponseV2Schema.merge(
    z.object({
        data: appInstanceSchema,
    })
);
export type UpdateAppInstanceApiResponse = z.infer<
    typeof updateAppInstanceSchema
>;
export type UpdateAppInstanceApiArgs = {
    projectId: string;
    instanceId: string;
    project: AppInstance;
};

/// DeleteAppInstance
export const deleteAppInstanceSchema = apiResponseV2Schema;
export type DeleteAppInstanceApiResponse = z.infer<
    typeof deleteAppInstanceSchema
>;
export type DeleteAppInstanceApiArgs = {
    projectId: string;
    instanceId: string;
};
