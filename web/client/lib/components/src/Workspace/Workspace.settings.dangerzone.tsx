"use client";

import React, { useCallback, useMemo } from "react";
import { useRouter } from "next/navigation";
import { z } from "zod";

import { Button } from "@/ui/primitives";
import { Form, FormButton, FormSegmentedControl } from "@/ui/form";
import { visibilitySchema, visibilityData } from "@/components/Utils";

import {
    useGetWorkspaceByIdQuery,
    useDeleteWorkspaceByIdMutation,
    useChangeWorkspaceVisibilityMutation,
} from "./Workspace.net.client";
import { WorkspaceVisibility } from "./Workspace.schema";
import { WorkspaceSettingsForm } from "./Workspace.settings.schema";

const useWorkspaceDangerZoneSettings = (
    orgSid: string,
    workspaceSid: string
) => {
    const router = useRouter();

    const { data: workspace } = useGetWorkspaceByIdQuery({
        orgSid,
        workspaceSid,
    });
    const [deleteWorkspace] = useDeleteWorkspaceByIdMutation();
    const [changeWorkspaceVisibility] = useChangeWorkspaceVisibilityMutation();

    const onDeleteWorkspace = useCallback(async () => {
        try {
            // await deleteWorkspace({
            //     id: workspaceSid as string,
            // }).unwrap();
            // router.push("/");
        } catch (error) {
            console.error("[onDeleteWorkspace] error", error);
        }
    }, [router, workspaceSid]);

    const onChangeWorkspaceVisibility = useCallback(
        async (data: { visibility: WorkspaceVisibility }) => {
            try {
                const _ = await changeWorkspaceVisibility({
                    orgSid,
                    workspaceSid,
                    visibility: data.visibility,
                }).unwrap();
            } catch (error) {
                console.error("[onChangeWorkspaceVisibility] error", error);
            }
        },
        [workspaceSid]
    );

    return {
        data: {
            workspace: {
                name: workspace?.data.name || "",
                description: workspace?.data.description || "",
                visibility: workspace?.data.visibility ?? "private",
            },
        },
        triggers: {
            onDeleteWorkspace,
            onChangeWorkspaceVisibility,
        },
    };
};

interface SettingsProps {
    data: {
        workspace: WorkspaceSettingsForm;
    };
    triggers: {
        onDeleteWorkspace: () => Promise<void>;
        onChangeWorkspaceVisibility: (values: {
            visibility: WorkspaceVisibility;
        }) => Promise<void>;
    };
}

const TransferWorkspaceOwnershipSettings = (props: SettingsProps) => {
    const {
        triggers: { onDeleteWorkspace },
        data: { workspace },
    } = props;

    return (
        <div className="w-1/2 flex flex-col rounded-lg border">
            <p className="p-3 flex items-center font-medium">
                Transfer ownership
            </p>
            <hr />
            <div className="p-3">Transfer this workspace to another user.</div>
            <hr />
            <div className="p-3 flex justify-end">
                <Button
                    onClick={onDeleteWorkspace}
                    className="flex items-center bg-red-600 hover:bg-red-500">
                    Transfer
                </Button>
            </div>
        </div>
    );
};

const ChangeWorkspaceVisibilitySettings = (props: SettingsProps) => {
    const {
        triggers: { onChangeWorkspaceVisibility },
        data: { workspace },
    } = props;

    const visData = useMemo(visibilityData, []);

    return (
        <div className="w-1/2 flex flex-col rounded-lg border">
            <p className="p-3 flex items-center font-medium">
                Change workspace visibility
            </p>
            <hr />
            <div className="p-3 flex items-center gap-1">
                <p>This workspace is currently</p>
                <span>
                    <span className="font-medium">{workspace.visibility}</span>.
                </span>
            </div>
            <hr />
            <Form
                initialValues={{ visibility: workspace.visibility }}
                schema={z.object({ visibility: visibilitySchema })}
                onFormSubmit={onChangeWorkspaceVisibility}
                className="p-3 flex justify-end gap-5">
                <FormSegmentedControl
                    data={visData}
                    name="visibility"
                    orientation="horizontal"
                    transitionDuration={250}
                    transitionTimingFunction="linear"
                />
                <FormButton className="w-[150px] flex items-center bg-red-600 hover:bg-red-500">
                    Change visibility
                </FormButton>
            </Form>
        </div>
    );
};

const DeleteWorkspaceSettings = (props: SettingsProps) => {
    const {
        triggers: { onDeleteWorkspace },
        data: { workspace },
    } = props;

    return (
        <div className="w-1/2 flex flex-col rounded-lg border">
            <p className="p-3 flex items-center font-medium">
                Delete this workspace
            </p>
            <hr />
            <div className="p-3">
                <div className="self-start flex items-center gap-1">
                    <p>Workspace:</p>
                    <p className="font-medium">{workspace.name}</p>
                </div>
                <p>
                    Please note that by deleting workspace, all your data will
                    be permanently deleted.
                </p>
            </div>
            <hr />
            <div className="p-3 flex justify-end">
                <Button
                    onClick={onDeleteWorkspace}
                    className="flex items-center gap-[6px] bg-red-600 hover:bg-red-500">
                    Delete
                    <p className="px-2 bg-red-400 rounded-lg">
                        {workspace.name}
                    </p>
                    workspace
                </Button>
            </div>
        </div>
    );
};

interface WorkspaceDangerZoneSettingsProps {
    orgSid: string;
    workspaceSid: string;
}

const WorkspaceDangerZoneSettings = (
    props: WorkspaceDangerZoneSettingsProps
) => {
    const { orgSid, workspaceSid } = props;
    const settings = useWorkspaceDangerZoneSettings(orgSid, workspaceSid);

    return (
        <div className="py-14 flex-1 flex flex-col items-center gap-4">
            <p className="w-1/2 flex items-center text-3xl">Danger zone</p>
            <ChangeWorkspaceVisibilitySettings {...settings} />
            <TransferWorkspaceOwnershipSettings {...settings} />
            <DeleteWorkspaceSettings {...settings} />
        </div>
    );
};

export { WorkspaceDangerZoneSettings };
