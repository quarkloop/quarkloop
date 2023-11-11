import { z } from "zod";

import { PermissionType, PermissionRole } from "./planSubscription";

export const operatingSystemSchema = z.object({
    id: z.string(),
    name: z.string(),
    path: z.string(),
    description: z.string(),
    overview: z.string().optional(),
    imageUrl: z.string().optional(),
    websiteUrl: z.string().optional(),
    url1: z.string().optional(),
    url2: z.string().optional(),
    url3: z.string().optional(),
    url4: z.string().optional(),
    isVerified: z.boolean().optional(),
    isOwner: z.boolean().optional(),
    createdAt: z.coerce.date().optional(),
    updatedAt: z.coerce.date().optional(),
});

export type OperatingSystem = z.infer<typeof operatingSystemSchema>;

export interface OperatingSystemUser {
    osId: string | null;
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
