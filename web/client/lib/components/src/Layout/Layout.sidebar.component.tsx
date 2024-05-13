"use client";

import React from "react";
import Link from "next/link";
import { Skeleton } from "@mantine/core";

import { useGetOrgByIdQuery } from "@/components/Org";
import { useGetWorkspaceByIdQuery } from "@/components/Workspace";

interface NameProps {
    type: "Organization" | "Workspace";
    label: string;
    href: string;
}

const Name = (props: NameProps) => {
    const { type, label, href } = props;
    return (
        <Link
            href={href}
            className="px-3 flex-1 flex items-center gap-2 hover:bg-neutral-100">
            <div className="w-8 h-8 flex items-center justify-center font-medium rounded-full bg-sky-200 text-sky-700">
                {label.charAt(0).toUpperCase()}
            </div>
            <div className="font-semibold">{label}</div>
            <div className="text-neutral-500">|</div>
            <div className="text-xs text-neutral-500">{type}</div>
        </Link>
    );
};

const Org = ({ orgSid }: { orgSid: string }) => {
    const { data: orgData } = useGetOrgByIdQuery({ orgSid: orgSid });

    if (orgData?.data == null) {
        return <NameSkeleton />;
    }
    return (
        <Name
            type="Organization"
            label={orgData.data.name}
            href={orgData.data.path}
        />
    );
};

const Workspace = ({
    orgSid,
    workspaceSid,
}: {
    orgSid: string;
    workspaceSid: string;
}) => {
    const { data: wsData } = useGetWorkspaceByIdQuery({
        orgSid,
        workspaceSid,
    });

    if (wsData?.data == null) {
        return <NameSkeleton />;
    }
    return (
        <Name
            type="Workspace"
            label={wsData.data.name}
            href={wsData.data.path}
        />
    );
};

const NameSkeleton = () => (
    <Skeleton
        height="100%"
        radius={0}
        style={{ flex: 1 }}
    />
);

export { Org, Workspace };
