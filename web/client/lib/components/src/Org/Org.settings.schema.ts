import { z } from "zod";

import { orgSchema } from "./Org.schema";

export const orgSettingsFormSchema = orgSchema.pick({
    name: true,
    description: true,
    visibility: true,
});

export type OrgSettingsForm = z.infer<typeof orgSettingsFormSchema>;

// export type OrgSettingsHookReturnType = UseHookReturnType<
//     {
//         org: OrgSettingsForm;
//         workspaceList: any[];
//     },
//     {
//         onUpdateOrg: (values: OrgSettingsForm) => Promise<void>;
//         onDeleteOrg: () => Promise<void>;
//     }
// >;
