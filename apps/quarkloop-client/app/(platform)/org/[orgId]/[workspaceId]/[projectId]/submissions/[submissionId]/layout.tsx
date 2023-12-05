import { SubmissionLayout } from "@quarkloop/components";
import { notFound } from "next/navigation";

interface LayoutProps {
    children: React.ReactNode;
    params: {
        orgId: string;
        workspaceId: string;
        projectId: string;
        submissionId: string;
    };
}

const Layout = async (props: LayoutProps) => {
    const {
        children,
        params: { orgId, workspaceId, projectId, submissionId },
    } = props;

    const subId = Number(submissionId);
    if (isNaN(subId)) {
        notFound();
    }

    return (
        <SubmissionLayout
            params={{ orgId, workspaceId, projectId, submissionId: subId }}>
            {children}
        </SubmissionLayout>
    );
};

export default Layout;
