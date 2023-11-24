import { z } from "zod";
import { apiResponseV2Schema, orgSchema, Organization } from "@quarkloop/types";

/// GetOrganizations
export const getOrganizationsSchema = apiResponseV2Schema.merge(
    z.object({
        data: orgSchema,
    })
);
export type GetOrganizationsApiResponse = z.infer<
    typeof getOrganizationsSchema
>;
export type GetOrganizationsApiArgs = void;

// TODO
/// GetOrganizationUsers
export const getOrganizationUsersSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(orgSchema),
    })
);
export type GetOrganizationUsersApiResponse = z.infer<
    typeof getOrganizationUsersSchema
>;
export type GetOrganizationUsersApiArgs = {
    id: string;
    workspaceId?: string;
};

/// GetOrganizationsByUserId
export const getOrganizationsByUserIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(orgSchema),
    })
);
export type GetOrganizationsByUserIdApiResponse = z.infer<
    typeof getOrganizationsByUserIdSchema
>;
export type GetOrganizationsByUserIdApiArgs = void;

/// GetOrganizationById
export const getOrganizationByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: orgSchema,
    })
);
export type GetOrganizationByIdApiResponse = z.infer<
    typeof getOrganizationByIdSchema
>;
export type GetOrganizationByIdApiArgs = {
    id: string;
};

/// CreateOrganization
export const createOrganizationSchema = apiResponseV2Schema.merge(
    z.object({
        data: orgSchema,
    })
);
export type CreateOrganizationApiResponse = z.infer<
    typeof createOrganizationSchema
>;
export type CreateOrganizationApiArgs = {
    name: string;
    description: string;
    accessType: number;
};

/// UpdateOrganization
export const updateOrganizationSchema = apiResponseV2Schema.merge(
    z.object({
        data: orgSchema,
    })
);
export type UpdateOrganizationApiResponse = z.infer<
    typeof updateOrganizationSchema
>;
export type UpdateOrganizationApiArgs = Partial<Organization>;

/// DeleteOrganization
export const deleteOrganizationSchema = apiResponseV2Schema;
export type DeleteOrganizationApiResponse = void;
export type DeleteOrganizationApiArgs = {
    id: string;
};
