import { Role } from "@/components/Utils";
import { getUserAccess } from "@/components/Auth/server";

import { Header } from "./Layout.header";
import { Sidebar } from "./Layout.sidebar";
import { getLinks } from "./Layout.links";

interface LayoutProps {
    children: React.ReactNode;
    userId?: string;
    orgSid?: string;
    workspaceSid?: string;
}

const Layout = async (props: LayoutProps) => {
    const { children, userId, orgSid, workspaceSid } = props;

    let role: Role | null = null;
    if (orgSid) {
        role = await getUserAccess({ orgSid, workspaceSid });
    }
    const links = getLinks(role, userId, orgSid, workspaceSid);

    return (
        <main className="flex-1 flex">
            <Sidebar
                links={links}
                role={role}
            />
            <div className="flex-1 flex flex-col dark:bg-[#343434]">
                <nav className="basis-14 max-h-14 flex flex-col border-b border-b-neutral-200">
                    <Header />
                </nav>
                <div className="flex-1 flex">{children}</div>
            </div>
        </main>
    );
};

export { Layout };
