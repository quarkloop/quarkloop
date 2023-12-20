import { ProjectSettings } from "@quarkloop/components";

interface PageProps {
    params: {
        orgId: string;
        workspaceId: string;
        projectId: string;
    };
    searchParams: any;
}

const Page = async (props: PageProps) => {
    const {
        params: { orgId, workspaceId, projectId },
        searchParams,
    } = props;

    return <ProjectSettings />;
};

export default Page;
