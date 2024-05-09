"use server";

import { notFound } from "next/navigation";

import { getOrgById } from "./Org.net.server";

interface OrgGaurdProps {
    children: React.ReactNode;
    orgSid: string;
}

const OrgGaurd = async (props: OrgGaurdProps) => {
    const { children, orgSid } = props;

    const org = await getOrgById({ orgSid });
    if (org == null) {
        notFound();
    }

    return <>{children}</>;
};

export { OrgGaurd };
