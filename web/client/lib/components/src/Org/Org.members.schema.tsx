import { z } from "zod";

import { userSchema } from "@quarkloop/types";
import { memberRolesSchema } from "@/components/Utils";

export const orgMemberSchema = z.object({
    user: userSchema,
    role: memberRolesSchema,
});

export const orgMemberRowSchema = orgMemberSchema.merge(
    z.object({
        currentLoggedInUser: userSchema,
    })
);

export type OrgMemberRow = z.infer<typeof orgMemberRowSchema>;
