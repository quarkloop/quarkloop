import { WorkspaceList } from "@quarkloop/components";

interface PageProps {
    params: { orgId: string };
    searchParams: any;
}

const Page = (props: PageProps) => {
    const {
        params: { orgId },
        searchParams,
    } = props;

    return <WorkspaceList />;
};

export default Page;
