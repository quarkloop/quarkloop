import { z } from "zod";
import { PermissionRole, PermissionType } from "@quarkloop/types";

export const history = z.object({
    createdAt: z.any(),
    createdBy: z.string(),
    updatedAt: z.any().optional(),
    updatedBy: z.string().optional(),
});

export const orgVisibilitySchema = z.enum(["public", "private"]);

export const orgSchema = history.merge(
    z.object({
        id: z.number(),
        sid: z.string(),
        name: z.string(),
        visibility: orgVisibilitySchema,
        description: z.string(),
        path: z.string(),
        createdAt: z.coerce.date().optional(),
        updatedAt: z.coerce.date().optional(),
    })
);

export type Org = z.infer<typeof orgSchema>;
export type OrgVisibility = z.infer<typeof orgVisibilitySchema>;

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
