import { FileUploader } from "@quarkloop/components";

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

    return <FileUploader />;
};

export default Page;
