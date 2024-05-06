import { z } from "zod";

export const orgMemberSchema = z.object({
    id: z.bigint(),
});

export type OrgMemberRow = z.infer<typeof orgMemberSchema>;
