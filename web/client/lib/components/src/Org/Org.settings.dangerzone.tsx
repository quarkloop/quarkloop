"use client";

import { useCallback, useMemo } from "react";
import { useRouter } from "next/navigation";
import { z } from "zod";

import { Button } from "@/ui/primitives";
import { Form, FormButton, FormSegmentedControl } from "@/ui/form";

import {
    useGetOrgByIdQuery,
    useDeleteOrgByIdMutation,
    useGetOrgWorkspaceListQuery,
    useChangeOrgVisibilityMutation,
} from "./Org.endpoint";
import { orgVisibilityData } from "./Org.util";
import { OrgVisibility, orgVisibilitySchema } from "./Org.schema";
import { OrgSettingsForm } from "./Org.settings.schema";

const useOrgDangerZoneSettings = (orgSid: string) => {
    const router = useRouter();

    const { data: org } = useGetOrgByIdQuery({ orgSid });
    const { data: wsList } = useGetOrgWorkspaceListQuery({ orgSid });

    const [deleteOrg] = useDeleteOrgByIdMutation();
    const [changeOrgVisibility] = useChangeOrgVisibilityMutation();

    const onDeleteOrg = useCallback(async () => {
        try {
            // await deleteOrg({
            //     id: orgSid as string,
            // }).unwrap();
            // router.push("/");
        } catch (error) {
            console.error("[onDeleteOrg] error", error);
        }
    }, [router, orgSid]);

    const onChangeOrgVisibility = useCallback(
        async (data: { visibility: OrgVisibility }) => {
            try {
                const _ = await changeOrgVisibility({
                    orgSid,
                    visibility: data.visibility,
                }).unwrap();
            } catch (error) {
                console.error("[onChangeOrgVisibility] error", error);
            }
        },
        [orgSid]
    );

    return {
        data: {
            org: {
                name: org?.data.name || "",
                description: org?.data.description || "",
                visibility: org?.data.visibility ?? "private",
            },
            workspaceList: wsList?.data || [],
        },
        triggers: {
            onDeleteOrg,
            onChangeOrgVisibility,
        },
    };
};

interface SettingsProps {
    data: {
        org: OrgSettingsForm;
        workspaceList: any[];
    };
    triggers: {
        onDeleteOrg: () => Promise<void>;
        onChangeOrgVisibility: (values: {
            visibility: OrgVisibility;
        }) => Promise<void>;
    };
}

const TransferOrgOwnershipSettings = (props: SettingsProps) => {
    const {
        triggers: { onDeleteOrg },
        data: { org, workspaceList },
    } = props;

    return (
        <div className="w-1/2 flex flex-col rounded-lg border">
            <p className="p-3 flex items-center font-medium">
                Transfer ownership
            </p>
            <hr />
            <div className="p-3">
                Transfer this organization to another user.
            </div>
            <hr />
            <div className="p-3 flex justify-end">
                <Button
                    onClick={onDeleteOrg}
                    className="flex items-center bg-red-600 hover:bg-red-500">
                    Transfer
                </Button>
            </div>
        </div>
    );
};

const ChangeOrgVisibilitySettings = (props: SettingsProps) => {
    const {
        triggers: { onChangeOrgVisibility },
        data: { org, workspaceList },
    } = props;

    const visibilityData = useMemo(orgVisibilityData, []);

    return (
        <div className="w-1/2 flex flex-col rounded-lg border">
            <p className="p-3 flex items-center font-medium">
                Change organization visibility
            </p>
            <hr />
            <div className="p-3 flex items-center gap-1">
                <p>This organization is currently</p>
                <span>
                    <span className="font-medium">{org.visibility}</span>.
                </span>
            </div>
            <hr />
            <Form
                initialValues={{ visibility: org.visibility }}
                schema={z.object({ visibility: orgVisibilitySchema })}
                onFormSubmit={onChangeOrgVisibility}
                className="p-3 flex justify-end gap-5">
                <FormSegmentedControl
                    data={visibilityData}
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

const DeleteOrgSettings = (props: SettingsProps) => {
    const {
        triggers: { onDeleteOrg },
        data: { org, workspaceList },
    } = props;

    return (
        <div className="w-1/2 flex flex-col rounded-lg border">
            <p className="p-3 flex items-center font-medium">
                Delete this organization
            </p>
            <hr />
            <div className="p-3">
                <div className="self-start flex items-center gap-1">
                    <p>Organization:</p>
                    <p className="font-medium">{org.name}</p>
                </div>
                <p>
                    Please note that by deleting organization, all your data
                    will be permanently deleted.
                </p>
                {workspaceList.length > 0 && (
                    <>
                        <p>Unable to delete this organization.</p>
                        <p>Please first delete following workspaces:</p>
                        {workspaceList.map((workspace: any, idx) => (
                            <p key={idx}>{workspace.name}</p>
                        ))}
                    </>
                )}
            </div>
            <hr />
            <div className="p-3 flex justify-end">
                <Button
                    disabled={workspaceList.length > 0}
                    onClick={onDeleteOrg}
                    className="flex items-center gap-[6px] bg-red-600 hover:bg-red-500">
                    Delete
                    <p className="px-2 bg-red-400 rounded-lg">{org.name}</p>
                    organization
                </Button>
            </div>
        </div>
    );
};

interface OrgDangerZoneSettingsProps {
    orgSid: string;
}

const OrgDangerZoneSettings = ({ orgSid }: OrgDangerZoneSettingsProps) => {
    const settings = useOrgDangerZoneSettings(orgSid);

    return (
        <div className="py-14 flex-1 flex flex-col items-center gap-4">
            <p className="w-1/2 flex items-center text-3xl">Danger zone</p>
            <ChangeOrgVisibilitySettings {...settings} />
            <TransferOrgOwnershipSettings {...settings} />
            <DeleteOrgSettings {...settings} />
        </div>
    );
};

export { OrgDangerZoneSettings };
