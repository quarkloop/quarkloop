"use client";

import { HTMLAttributes, useCallback, useMemo } from "react";

import {
    useGetWorkspaceByIdQuery,
    useUpdateWorkspaceByIdMutation,
} from "./Workspace.net.client";
import {
    WorkspaceGeneralSettingsForm,
    WorkspaceGeneralSettingsFormData,
} from "./Workspace.settings.general.form";

const useWorkspaceGeneralSettings = (orgSid: string, workspaceSid: string) => {
    const { data: workspace } = useGetWorkspaceByIdQuery({
        orgSid,
        workspaceSid,
    });
    const [updateWorkspace] = useUpdateWorkspaceByIdMutation();

    const onFormSubmit = useCallback(
        async (data: WorkspaceGeneralSettingsFormData) => {
            try {
                await updateWorkspace({
                    orgSid,
                    workspaceSid,
                    payload: data,
                }).unwrap();
            } catch (error) {
                console.log(error);
            }
        },
        []
    );

    const initialValues = useMemo(
        () =>
            ({
                id: workspace?.data.id ?? 0,
                sid: workspace?.data.sid ?? "",
                name: workspace?.data.name ?? "",
                description: workspace?.data.description ?? "",
                visibility: workspace?.data.visibility ?? "private",
                path: workspace?.data.path ?? "",
                createdBy: workspace?.data.createdBy ?? "",
                updatedBy: workspace?.data.updatedBy ?? "",
            } as WorkspaceGeneralSettingsFormData),
        [workspace?.data]
    );

    return {
        initialValues,
        onFormSubmit,
    };
};

interface WorkspaceGeneralSettingsProps extends HTMLAttributes<HTMLDivElement> {
    orgSid: string;
    workspaceSid: string;
}

const WorkspaceGeneralSettings = (props: WorkspaceGeneralSettingsProps) => {
    const { orgSid, workspaceSid } = props;
    const { initialValues, onFormSubmit } = useWorkspaceGeneralSettings(
        orgSid,
        workspaceSid
    );

    return (
        <div className="py-14 flex-1 flex flex-col items-center gap-5">
            <div className="flex items-center gap-3 text-3xl font-medium">
                General Settings
            </div>
            <WorkspaceGeneralSettingsForm
                readOnly={false}
                initialValues={initialValues}
                onFormSubmit={onFormSubmit}
            />
        </div>
    );
};

export { WorkspaceGeneralSettings };
