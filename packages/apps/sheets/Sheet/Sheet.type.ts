import { z } from "zod";

import { workspaceSchema } from "@quarkloop/types";
import { UseHookReturnType } from "@quarkloop/lib";
import { ColumnDef } from "@tanstack/react-table";

export const workspaceRowSchema = workspaceSchema.pick({
    id: true,
    name: true,
    description: true,
    path: true,
    updatedAt: true,
});

export type SheetRow = z.infer<typeof workspaceRowSchema>;

export type SheetHookReturnType = UseHookReturnType<SheetRow[], null> | null;

export type SheetProps = {
    columns: ColumnDef<unknown>[];
    data: unknown[];
};
