// import {
//     OperatingSystem as PrismaOperatingSystem,
//     PermissionType,
//     PermissionRole,
// } from "@quarkloop/prisma/types";
import { z } from "zod";

import { PermissionType, PermissionRole } from "./planSubscription";
import { ApiResponse } from "./api-response";

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

// export interface OperatingSystem {
//     id: string;
//     name: string;
//     path: string;
//     description: string;
//     overview: string | null;
//     imageUrl: string | null;
//     websiteUrl: string | null;
//     url1: string | null;
//     url2: string | null;
//     url3: string | null;
//     url4: string | null;
//     isVerified: boolean | null;
//     createdAt: Date | null;
//     isOwner: boolean | null;
// }

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

// export type OperatingSystemPluginArgs =
//     | GetOperatingSystemByIdPluginArgs
//     | GetOperatingSystemUsersPluginArgs
//     | CreateOperatingSystemPluginArgs
//     | UpdateOperatingSystemPluginArgs
//     | DeleteOperatingSystemPluginArgs;

// export type OperatingSystemApiResponse = GetOperatingSystemUsersApiResponse;
