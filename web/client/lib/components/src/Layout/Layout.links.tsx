import {
    IconAntenna,
    IconBox,
    IconCirclesRelation,
    IconPlaceholder,
    IconSettings,
    IconSquare,
    IconUser,
    IconUserCircle,
    IconUsers,
} from "@tabler/icons-react";
import { Role } from "@/components/Utils";

type SideBarLinkType = "user" | "org" | "workspace";

export interface SideBarLink {
    type?: SideBarLinkType[];
    label: string;
    priority: number;
    href?: string;
    icon?: JSX.Element;
    sublinks: SideBarLink[];
}

export function getLinks(
    role: Role | null,
    userId?: string,
    orgSid?: string,
    workspaceSid?: string
): SideBarLink[] {
    if (workspaceSid) {
        const sidebar: LinkBase = new WorkspaceSideBarLink(
            role,
            orgSid!,
            workspaceSid
        );
        return sidebar.links();
    }
    if (orgSid) {
        const sidebar = new OrgSideBarLink(role, orgSid);
        return sidebar.links();
    }
    if (userId) {
        const sidebar: LinkBase = new UserSettingsSideBarLink(userId);
        return sidebar.links();
    }

    const sidebar: LinkBase = new UserSideBarLink();
    return sidebar.links();
}

export interface LinkBase {
    links(): SideBarLink[];
}

const org: SideBarLink = {
    label: "Orgs",
    priority: 1,
    href: "/manage/orgs",
    icon: (
        <IconBox
            size="1.3rem"
            stroke={1.5}
        />
    ),
    sublinks: [],
};

const workspaces: SideBarLink = {
    label: "Workspaces",
    priority: 1,
    href: "/manage/workspaces",
    icon: (
        <IconPlaceholder
            size="1.3rem"
            stroke={1.5}
        />
    ),
    sublinks: [],
};

const manage: SideBarLink = {
    label: "Manage",
    priority: 9,
    icon: (
        <IconUsers
            size="1.3rem"
            stroke={1.5}
        />
    ),
    sublinks: [
        {
            label: "Activity",
            priority: 1,
            href: "activity",
            sublinks: [],
        },
        {
            label: "Members",
            priority: 1,
            href: "members",
            sublinks: [],
        },
    ],
};

const settings: SideBarLink = {
    label: "Settings",
    priority: 10,
    icon: (
        <IconSettings
            size="1.3rem"
            stroke={1.5}
        />
    ),
    sublinks: [
        {
            label: "General",
            priority: 1,
            href: "settings/general",
            sublinks: [],
        },
        {
            label: "Integrations",
            priority: 1,
            href: "settings/integratons",
            sublinks: [],
        },
        {
            label: "Access tokens",
            priority: 1,
            href: "settings/access_tokens",
            sublinks: [],
        },
        {
            label: "Webhooks",
            priority: 1,
            href: "settings/webhooks",
            sublinks: [],
        },
        {
            label: "Danger zone",
            priority: 1,
            href: "settings/danger-zone",
            sublinks: [],
        },
    ],
};

// const tables: SideBarLink = {
//     label: "Tables",
//     priority: 1,
//     icon: (
//         <IconDatabase
//             size="1.3rem"
//             stroke={1.5}
//         />
//     ),
//     sublinks: [
//         {
//             label: "Personal Info",
//             priority: 1,
//             href: "personal_info",
//             sublinks: [],
//         },
//         {
//             label: "Documents",
//             priority: 1,
//             href: "documents",
//             sublinks: [],
//         },
//         {
//             label: "Payments",
//             priority: 1,
//             href: "payments",
//             sublinks: [],
//         },
//         {
//             label: "Reservations",
//             priority: 1,
//             href: "reservations",
//             sublinks: [],
//         },
//         {
//             label: "Location",
//             priority: 1,
//             href: "location",
//             sublinks: [],
//         },
//     ],
// };

// const services: SideBarLink = {
//     label: "Services",
//     priority: 1,
//     icon: (
//         <IconCode
//             size="1.3rem"
//             stroke={1.5}
//         />
//     ),
//     sublinks: [
//         {
//             label: "Services",
//             priority: 1,
//             href: "services",
//             sublinks: [],
//         },
//         {
//             label: "Submission requests",
//             priority: 1,
//             href: "submission_requests",
//             sublinks: [],
//         },
//     ],
// };

// const service: SideBarLink = {
//     label: "Service",
//     priority: 1,
//     icon: (
//         <IconCode
//             size="1.3rem"
//             stroke={1.5}
//         />
//     ),
//     sublinks: [
//         {
//             label: "General",
//             priority: 1,
//             href: "",
//             sublinks: [],
//         },
//         {
//             label: "Intro Page",
//             priority: 1,
//             href: "intro",
//             sublinks: [],
//         },
//         {
//             label: "Components",
//             priority: 1,
//             href: "components",
//             sublinks: [],
//         },
//         {
//             label: "Integrations",
//             priority: 1,
//             href: "itegrations",
//             sublinks: [],
//         },
//         {
//             label: "Manage",
//             priority: 1,
//             href: "manage",
//             sublinks: [],
//         },
//     ],
// };

export class UserSideBarLink implements LinkBase {
    constructor() {}

    public links(): SideBarLink[] {
        return [org, workspaces];
    }
}

export class UserSettingsSideBarLink implements LinkBase {
    private href: string;

    constructor(userId: string) {
        this.href = `/users/${userId}/settings`;
    }

    public links(): SideBarLink[] {
        return [
            {
                label: "Public profile",
                priority: 1,
                href: `${this.href}/profile`,
                icon: (
                    <IconUser
                        size="1.3rem"
                        stroke={1.5}
                    />
                ),
                sublinks: [],
            },
            {
                label: "Account",
                priority: 1,
                href: `${this.href}/account`,
                icon: (
                    <IconUserCircle
                        size="1.3rem"
                        stroke={1.5}
                    />
                ),
                sublinks: [],
            },
            {
                label: "Linked accounts",
                priority: 1,
                href: `${this.href}/linked-accounts`,
                icon: (
                    <IconCirclesRelation
                        size="1.3rem"
                        stroke={1.5}
                    />
                ),
                sublinks: [],
            },
            {
                label: "Sessions",
                priority: 1,
                href: `${this.href}/sessions`,
                icon: (
                    <IconAntenna
                        size="1.3rem"
                        stroke={1.5}
                    />
                ),
                sublinks: [],
            },
            {
                label: "Plan",
                priority: 1,
                href: `${this.href}/plan`,
                icon: (
                    <IconSquare
                        size="1.3rem"
                        stroke={1.5}
                    />
                ),
                sublinks: [],
            },
        ];
    }
}

export class OrgSideBarLink implements LinkBase {
    private role: Role | null;
    private orgSid: string | string[];
    private href: string;

    constructor(role: Role | null, orgSid: string | string[]) {
        this.role = role;
        this.orgSid = orgSid;
        this.href = `/manage/${this.orgSid}`;
    }

    public links(): SideBarLink[] {
        let links = [
            {
                ...workspaces,
                href: `${this.href}/workspaces`,
            },
            {
                ...manage,
                sublinks: manage.sublinks.map((link) => ({
                    ...link,
                    href: `${this.href}/${link.href}`,
                })),
            },
        ];

        if (this.role && this.role === "owner") {
            links = [
                ...links,
                {
                    ...settings,
                    sublinks: settings.sublinks.map((link) => ({
                        ...link,
                        href: `${this.href}/${link.href}`,
                    })),
                },
            ];
        }

        return links;
    }
}

export class WorkspaceSideBarLink implements LinkBase {
    private role: Role | null;
    private orgSid: string | string[];
    private workspaceSid: string | string[];
    private href: string;

    constructor(
        role: Role | null,
        orgSid: string | string[],
        workspaceSid: string | string[]
    ) {
        this.role = role;
        this.orgSid = orgSid;
        this.workspaceSid = workspaceSid;
        this.href = `/manage/${this.orgSid}/${this.workspaceSid}`;
    }

    public links(): SideBarLink[] {
        let links = [manage];

        if (this.role && this.role === "owner") {
            links = [
                ...links,
                {
                    ...settings,
                    sublinks: settings.sublinks.map((link) => ({
                        ...link,
                        href: `${this.href}/${link.href}`,
                    })),
                },
            ];
        }

        return links;
    }
}
