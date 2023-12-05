import { OrganizationSettings } from "@quarkloop/components";

interface PageProps {
    params: { orgId: string };
    searchParams: any;
}

const Page = (props: PageProps) => {
    const {
        params: { orgId },
        searchParams,
    } = props;

    return <OrganizationSettings />;
};

export default Page;
