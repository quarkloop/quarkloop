"use server";

import { notFound } from "next/navigation";

import { getWorkspaceById } from "./Workspace.net.server";

interface WorkspaceGaurdProps {
    children: React.ReactNode;
    orgSid: string;
    workspaceSid: string;
}

const WorkspaceGaurd = async (props: WorkspaceGaurdProps) => {
    const { children, orgSid, workspaceSid } = props;

    const workspace = await getWorkspaceById({ orgSid, workspaceSid });
    if (workspace == null) {
        notFound();
    }

    return <>{children}</>;
};

export { WorkspaceGaurd };
