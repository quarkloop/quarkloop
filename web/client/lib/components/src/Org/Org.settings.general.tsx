"use client";

import { HTMLAttributes, useCallback, useMemo } from "react";

import { useGetOrgByIdQuery, useUpdateOrgByIdMutation } from "./Org.endpoint";
import {
    OrgGeneralSettingsForm,
    OrgGeneralSettingsFormData,
} from "./Org.settings.general.form";

const useOrgGeneralSettings = ({ orgSid }: { orgSid: string }) => {
    const { data: org } = useGetOrgByIdQuery({ orgSid });
    const [updateOrg] = useUpdateOrgByIdMutation();

    const onFormSubmit = useCallback(
        async (data: OrgGeneralSettingsFormData) => {
            try {
                await updateOrg({
                    orgSid: data.sid,
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
                id: org?.data.id ?? 0,
                sid: org?.data.sid ?? "",
                name: org?.data.name ?? "",
                description: org?.data.description ?? "",
                path: org?.data.path ?? "",
                createdBy: org?.data.createdBy ?? "",
                updatedBy: org?.data.updatedBy ?? "",
            } as OrgGeneralSettingsFormData),
        [org?.data]
    );

    return {
        initialValues,
        onFormSubmit,
    };
};

interface OrgGeneralSettingsProps extends HTMLAttributes<HTMLDivElement> {
    orgSid: string;
}

const OrgGeneralSettings = ({ orgSid }: OrgGeneralSettingsProps) => {
    const { initialValues, onFormSubmit } = useOrgGeneralSettings({ orgSid });

    return (
        <div className="py-14 flex-1 flex flex-col items-center gap-5">
            <div className="flex items-center gap-3 text-3xl font-medium">
                General Settings
            </div>
            <OrgGeneralSettingsForm
                readOnly={false}
                initialValues={initialValues}
                onFormSubmit={onFormSubmit}
            />
        </div>
    );
};

export { OrgGeneralSettings };
