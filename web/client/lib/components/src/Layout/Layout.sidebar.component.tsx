"use client";

import React from "react";
import Link from "next/link";
import { Skeleton } from "@mantine/core";

import { useGetWorkspaceByIdQuery } from "@quarkloop/lib";
import { useGetOrgByIdQuery } from "@/components/Org";

const Name = ({ href }: any) => {
    return (
        <Link
            href={href}
            className="flex-1 px-2 flex items-center rounded-lg hover:bg-neutral-100">
            Overview
        </Link>
    );
};

const Org = ({ orgSid }: any) => {
    const { data: orgData } = useGetOrgByIdQuery({ orgSid: orgSid });

    if (orgData?.data == null) {
        return <NameSkeleton />;
    }
    return <Name href={orgData.data.path} />;
};

const Workspace = ({ workspaceSid }: any) => {
    const { data: wsData } = useGetWorkspaceByIdQuery({ id: workspaceSid });

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
