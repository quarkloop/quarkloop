"use client";

import { useMemo } from "react";
import Link from "next/link";
import { useParams } from "next/navigation";

import { CommandMenu } from "@/components/OrgProfile";

const Header = () => {
    const { orgSid, workspaceSid } = useParams();

    const link = useMemo(() => {
        if (orgSid && workspaceSid) {
            return {
                orgHref: `/manage/${orgSid}`,
                workspaceHref: `/manage/${orgSid}/${workspaceSid}`,
            };
        } else if (orgSid) {
            return {
                orgHref: `/manage/${orgSid}`,
            };
        }

        return {
            orgHref: "",
        };
    }, [orgSid, workspaceSid]);

    return (
        <div className="px-4 flex-1 flex items-center gap-3 bg-neutral-100">
            <Link
                href="/"
                className="w-10 h-10 rounded-full bg-neutral-300"
            />
            <div className="flex items-center gap-1">
                {orgSid && (
                    <Link
                        href={link.orgHref}
                        className="px-2 py-1 hover:bg-neutral-200 hover:rounded-lg">
                        {orgSid}
                    </Link>
                )}
                {workspaceSid && (
                    <>
                        <p>/</p>
                        <Link
                            href={link.workspaceHref!}
                            className="px-2 py-1 font-semibold hover:bg-neutral-200 hover:rounded-lg">
                            {workspaceSid}
                        </Link>
                    </>
                )}
            </div>
            <div className="flex-1" />
            <CommandMenu />
        </div>
    );
};

export { Header };
