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

    return <div>Page</div>;
};

export default Page;
