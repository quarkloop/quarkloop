import { Role } from "@/components/Utils";
import { SideBarLink } from "./Layout.links";

export interface SidebarProps {
    role: Role | null;
    links: SideBarLink[];
}
