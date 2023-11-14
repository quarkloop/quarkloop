import { z } from "zod";
import {
    apiResponseV2Schema,
    operatingSystemSchema,
    OperatingSystem,
} from "@quarkloop/types";

/// GetOperatingSystems
export const getOperatingSystemsSchema = apiResponseV2Schema.merge(
    z.object({
        data: operatingSystemSchema,
    })
);
export type GetOperatingSystemsApiResponse = z.infer<
    typeof getOperatingSystemsSchema
>;
export type GetOperatingSystemsApiArgs = void;

// TODO
/// GetOperatingSystemUsers
export const getOperatingSystemUsersSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(operatingSystemSchema),
    })
);
export type GetOperatingSystemUsersApiResponse = z.infer<
    typeof getOperatingSystemUsersSchema
>;
export type GetOperatingSystemUsersApiArgs = {
    id: string;
    workspaceId?: string;
};

/// GetOperatingSystemsByUserId
export const getOperatingSystemsByUserIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(operatingSystemSchema),
    })
);
export type GetOperatingSystemsByUserIdApiResponse = z.infer<
    typeof getOperatingSystemsByUserIdSchema
>;
export type GetOperatingSystemsByUserIdApiArgs = void;

/// GetOperatingSystemById
export const getOperatingSystemByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: operatingSystemSchema,
    })
);
export type GetOperatingSystemByIdApiResponse = z.infer<
    typeof getOperatingSystemByIdSchema
>;
export type GetOperatingSystemByIdApiArgs = {
    id: string;
};

/// CreateOperatingSystem
export const createOperatingSystemSchema = apiResponseV2Schema.merge(
    z.object({
        data: operatingSystemSchema,
    })
);
export type CreateOperatingSystemApiResponse = z.infer<
    typeof createOperatingSystemSchema
>;
export type CreateOperatingSystemApiArgs = {
    name: string;
    description: string;
};

/// UpdateOperatingSystem
export const updateOperatingSystemSchema = apiResponseV2Schema.merge(
    z.object({
        data: operatingSystemSchema,
    })
);
export type UpdateOperatingSystemApiResponse = z.infer<
    typeof updateOperatingSystemSchema
>;
export type UpdateOperatingSystemApiArgs = OperatingSystem;

/// DeleteOperatingSystem
export const deleteOperatingSystemSchema = apiResponseV2Schema;
export type DeleteOperatingSystemApiResponse = z.infer<
    typeof deleteOperatingSystemSchema
>;
export type DeleteOperatingSystemApiArgs = {
    id: string;
};
