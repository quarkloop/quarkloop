"use client";

import React, { useCallback, useMemo } from "react";
import { useDisclosure } from "@mantine/hooks";

import { Button } from "@/ui/primitives";
import { DataTableV3 } from "@/components/DataTable";

import { columns } from "./Workspace.list.columns";
import {
    useGetOrgWorkspacesQuery,
    useCreateWorkspaceMutation,
} from "./Workspace.net.client";
import { WorkspaceCreateFormData } from "./Workspace.create.form";
import { WorkspaceCreateModal } from "./Workspace.create.modal";

interface WorkspaceListProps {
    orgSid: string;
}

const WorkspaceList = ({ orgSid }: WorkspaceListProps) => {
    const [opened, { open, close }] = useDisclosure(false);

    const { data: workspaceData } = useGetOrgWorkspacesQuery({ orgSid });
    const [createWorkspace] = useCreateWorkspaceMutation();

    const onCreateClick = useCallback(async (data: WorkspaceCreateFormData) => {
        try {
            const result = await createWorkspace({
                orgSid,
                payload: data,
            }).unwrap();
            if (result) {
                close();
            }
        } catch (error) {
            console.error("[onCreateClick] error:", error);
        }
    }, []);
    const workspaceList = useMemo(() => workspaceData?.data, [workspaceData]);

    return (
        <div className="px-20 py-10 flex-1 flex-col gap-3 md:flex">
            <div className="py-2 flex items-center justify-between">
                <div className="flex flex-col gap-3">
                    <h2 className="text-3xl font-semibold tracking-tight">
                        Workspaces
                    </h2>
                    <p className="text-muted-foreground">All workspaces</p>
                </div>
                <Button onClick={open}>New workspace</Button>
                <WorkspaceCreateModal
                    defaultOpened={opened}
                    open={open}
                    close={close}
                    onFormSubmit={onCreateClick}
                />
            </div>
            <DataTableV3
                enableHeader
                enablePagination
                toolbar={{
                    filterColumnName: "name",
                    filterColumns: [
                        {
                            columnName: "name",
                            columnTitle: "Name",
                            options: [
                                {
                                    label: "Name",
                                    value: "name",
                                },
                            ],
                        },
                    ],
                }}
                data={workspaceList || []}
                columns={columns}
            />
        </div>
    );
};

export { WorkspaceList };
