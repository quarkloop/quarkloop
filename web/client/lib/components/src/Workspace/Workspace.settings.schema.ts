import { z } from "zod";

import { workspaceSchema } from "./Workspace.schema";

export const workspaceSettingsFormSchema = workspaceSchema.pick({
    name: true,
    description: true,
    visibility: true,
});

export type WorkspaceSettingsForm = z.infer<typeof workspaceSettingsFormSchema>;

// export type WorkspaceSettingsHookReturnType = UseHookReturnType<
//     {
//         workspace: WorkspaceSettingsForm;
//         workspaceList: any[];
//     },
//     {
//         onUpdateWorkspace: (values: WorkspaceSettingsForm) => Promise<void>;
//         onDeleteWorkspace: () => Promise<void>;
//     }
// >;
