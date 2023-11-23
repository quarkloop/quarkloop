import { z } from "zod";

import { PermissionType, PermissionRole } from "./planSubscription";

export const orgSchema = z.object({
    id: z.string(),
    name: z.string(),
    accessType: z.number(),
    path: z.string(),
    description: z.string(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type Organization = z.infer<typeof orgSchema>;

export interface OrganizationUser {
    orgId: string | null;
    type: PermissionType;
    role: PermissionRole;
    createdAt: Date;
    user: {
        id: string;
        name: string | null;
        email: string;
        image: string | null;
    };
}
