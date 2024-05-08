import { z } from "zod";

import { PermissionRole, PermissionType } from "@quarkloop/types";
import { historySchema, visibilitySchema } from "@/components/Utils";

export const orgSchema = historySchema.merge(
    z.object({
        id: z.number(),
        sid: z.string(),
        name: z.string(),
        visibility: visibilitySchema,
        description: z.string(),
        path: z.string(),
        createdAt: z.coerce.date().optional(),
        updatedAt: z.coerce.date().optional(),
    })
);

export type Org = z.infer<typeof orgSchema>;
export type OrgVisibility = z.infer<typeof visibilitySchema>;

export interface OrgUser {
    orgSid: string | null;
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
