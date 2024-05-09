"use server";

import { HTMLAttributes } from "react";
import { notFound } from "next/navigation";

import { getUserAccess } from "./server";
import { Role } from "@/components/Utils";

interface RoleProviderProps extends HTMLAttributes<HTMLDivElement> {
    orgSid: string;
    workspaceSid?: string;
    allowedRoles: Role[];
}

const RoleProvider = async (props: RoleProviderProps) => {
    const { orgSid, workspaceSid, allowedRoles, children } = props;

    const role = await getUserAccess({ orgSid, workspaceSid });

    if (role == null || !allowedRoles.includes(role)) {
        notFound();
    }
    return <>{children}</>;
};

export { RoleProvider };
