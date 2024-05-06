import { Role } from "@/components/Auth";
import { SideBarLink } from "./Layout.links";

export interface SidebarProps {
    role: Role | null;
    links: SideBarLink[];
}
