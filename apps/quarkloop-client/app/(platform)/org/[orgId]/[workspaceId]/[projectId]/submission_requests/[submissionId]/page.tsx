import { redirect } from "next/navigation";

interface PageProps {
    params: {
        orgId: string;
        workspaceId: string;
        projectId: string;
        submissionId: string;
    };
    searchParams: any;
}

const Page = async (props: PageProps) => {
    const {
        params: { orgId, workspaceId, projectId, submissionId },
        searchParams,
    } = props;

    redirect(`${submissionId}/discussions`);
};

export default Page;
