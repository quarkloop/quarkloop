"use client";

import React from "react";
import Link from "next/link";
import { Skeleton } from "@mantine/core";

import { useGetOrgByIdQuery } from "@/components/Org";
import { useGetWorkspaceByIdQuery } from "@/components/Workspace";

const Name = ({ href }: { href: string }) => {
    return (
        <Link
            href={href}
            className="flex-1 px-2 flex items-center rounded-lg hover:bg-neutral-100">
            Overview
        </Link>
    );
};

const Org = ({ orgSid }: { orgSid: string }) => {
    const { data: orgData } = useGetOrgByIdQuery({ orgSid: orgSid });

    if (orgData?.data == null) {
        return <NameSkeleton />;
    }
    return <Name href={orgData.data.path} />;
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
    return <Name href={wsData.data.path} />;
};

const NameSkeleton = () => (
    <div className="flex-1 p-2">
        <Skeleton
            height={"100%"}
            radius="lg"
        />
    </div>
);

export { Org, Workspace };
