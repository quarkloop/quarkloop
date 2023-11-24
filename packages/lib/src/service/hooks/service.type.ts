import { z } from "zod";
import { apiResponseV2Schema, serviceSchema, Service } from "@quarkloop/types";

/// GetServiceList
export const getServiceListSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(serviceSchema),
    })
);
export type GetServiceListApiResponse = z.infer<typeof getServiceListSchema>;
export type GetServiceListApiArgs = {
    projectId: string;
};

/// GetServiceById
export const getServiceByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: serviceSchema,
    })
);
export type GetServiceByIdApiResponse = z.infer<typeof getServiceByIdSchema>;
export type GetServiceByIdApiArgs = {
    projectId: string;
    serviceId: string;
};

/// CreateService
export const createServiceSchema = apiResponseV2Schema.merge(
    z.object({
        data: serviceSchema,
    })
);
export type CreateServiceApiResponse = z.infer<typeof createServiceSchema>;
export type CreateServiceApiArgs = {
    projectId: string;
    project: Partial<Service>;
};

/// UpdateService
export const updateServiceSchema = apiResponseV2Schema.merge(
    z.object({
        data: serviceSchema,
    })
);
export type UpdateServiceApiResponse = z.infer<typeof updateServiceSchema>;
export type UpdateServiceApiArgs = {
    projectId: string;
    serviceId: string;
    project: Partial<Service>;
};

/// DeleteService
export const deleteServiceSchema = apiResponseV2Schema;
export type DeleteServiceApiResponse = z.infer<typeof deleteServiceSchema>;
export type DeleteServiceApiArgs = {
    projectId: string;
    serviceId: string;
};
