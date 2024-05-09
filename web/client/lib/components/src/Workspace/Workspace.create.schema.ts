import { z } from "zod";

import { UseHookReturnType } from "@quarkloop/lib";
import { workspaceSchema } from "./Workspace.schema";

export const workspaceCreateFormSchema = workspaceSchema.pick({
    sid: true,
    name: true,
    description: true,
    visibility: true,
});

export type WorkspaceCreateForm = z.infer<typeof workspaceCreateFormSchema>;

export type WorkspaceCreateHookReturnType = UseHookReturnType<
    unknown,
    {
        onCreateWorkspace: (values: WorkspaceCreateForm) => Promise<void>;
    }
>;

// export type WorkspaceCreateProps = {
//     onCreateWorkspace: (values: WorkspaceCreateForm) => Promise<void>;
// };
