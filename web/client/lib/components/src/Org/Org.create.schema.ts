import { z } from "zod";

import { UseHookReturnType } from "@quarkloop/lib";
import { orgSchema } from "./Org.schema";

export const orgCreateFormSchema = orgSchema.pick({
    sid: true,
    name: true,
    description: true,
    visibility: true,
});

export type OrgCreateForm = z.infer<typeof orgCreateFormSchema>;

export type OrgCreateHookReturnType = UseHookReturnType<
    unknown,
    {
        onCreateOrg: (values: OrgCreateForm) => Promise<void>;
    }
>;

// export type OrgCreateProps = {
//     onCreateOrg: (values: OrgCreateForm) => Promise<void>;
// };
