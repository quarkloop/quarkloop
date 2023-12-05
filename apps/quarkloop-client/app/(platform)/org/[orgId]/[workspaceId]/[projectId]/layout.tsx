import { AppLayout } from "@quarkloop/components";

interface LayoutProps {
    children: React.ReactNode;
    params: {
        orgId: string;
        workspaceId: string;
        projectId: string;
    };
}

const Layout = async (props: LayoutProps) => {
    const {
        children,
        params: { orgId, workspaceId, projectId },
    } = props;

    return (
        <AppLayout params={{ orgId, workspaceId, projectId }}>
            {children}
        </AppLayout>
    );
};

export default Layout;
