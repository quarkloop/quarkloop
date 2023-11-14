import { z } from "zod";
import { apiResponseV2Schema, appSchema, App } from "@quarkloop/types";

/// GetAppList
export const getAppListSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(appSchema),
    })
);
export type GetAppListApiResponse = z.infer<typeof getAppListSchema>;
export type GetAppListApiArgs = {
    osId?: string;
    workspaceId?: string;
};

/// GetAppById
export const getAppByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: appSchema,
    })
);
export type GetAppByIdApiResponse = z.infer<typeof getAppByIdSchema>;
export type GetAppByIdApiArgs = {
    id: string;
};

/// CreateApp
export const createAppSchema = apiResponseV2Schema.merge(
    z.object({
        data: appSchema,
    })
);
export type CreateAppApiResponse = z.infer<typeof getAppByIdSchema>;
export type CreateAppApiArgs = {
    osId: string;
    workspaceId: string;
    app: Partial<App>;
};

/// UpdateApp
export const updateAppSchema = apiResponseV2Schema.merge(
    z.object({
        data: appSchema,
    })
);
export type UpdateAppApiResponse = z.infer<typeof updateAppSchema>;
export type UpdateAppApiArgs = {
    id: string;
    app: Partial<App>;
};

/// DeleteApp
export const deleteAppSchema = apiResponseV2Schema;
export type DeleteAppApiResponse = z.infer<typeof deleteAppSchema>;
export type DeleteAppApiArgs = {
    id: string;
};
