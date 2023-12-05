import { UserList } from "@quarkloop/components";

interface PageProps {
    params: {
        orgId: string;
        workspaceId: string;
    };
    searchParams: any;
}

const Page = async (props: PageProps) => {
    const {
        params: { orgId, workspaceId },
        searchParams,
    } = props;

    return <UserList />;
};

export default Page;
