"use client";

import React from "react";
import Link from "next/link";
import { IconPlus } from "@tabler/icons-react";

import { useGetPlanSubscriptionByUserSession } from "@quarkloop/lib";

interface WorkspaceCreateButtonProps {
    label: string;
}

const WorkspaceCreateButton = (props: WorkspaceCreateButtonProps) => {
    const { label } = props;
    const { subscription, metrics } = useGetPlanSubscriptionByUserSession();
    const disable =
        (metrics?.Workspace.used || 0) <
        (subscription?.plan.features.maxWorkspaces || 0);

    return (
        <div className="px-5 py-7 flex items-center justify-center">
            {disable ? (
                <Link
                    href="/new/setup-workspace"
                    prefetch={false}
                    className="px-5 py-2 flex-1 flex items-center justify-center gap-2 bg-blue-600 rounded-lg">
                    <IconPlus
                        size="1.7rem"
                        stroke={2.0}
                    />
                    <p className="flex items-center text-lg text-white font-medium">
                        {label}
                    </p>
                </Link>
            ) : (
                <div className="px-5 py-2 flex-1 flex items-center justify-center gap-2 bg-blue-600 rounded-lg">
                    <IconPlus
                        size="1.7rem"
                        stroke={2.0}
                    />
                    <p className="flex items-center text-lg text-white font-medium">
                        {label}
                    </p>
                </div>
            )}
        </div>
    );
};

export { WorkspaceCreateButton };
