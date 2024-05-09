import { z } from "zod";

import { userSchema } from "@quarkloop/types";
import { memberRolesSchema } from "@/components/Utils";

export const workspaceMemberSchema = z.object({
    user: userSchema,
    role: memberRolesSchema,
});

export const workspaceMemberRowSchema = workspaceMemberSchema.merge(
    z.object({
        currentLoggedInUser: userSchema,
    })
);

export type WorkspaceMemberRow = z.infer<typeof workspaceMemberRowSchema>;
