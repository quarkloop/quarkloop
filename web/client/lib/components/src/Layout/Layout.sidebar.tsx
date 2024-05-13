"use client";

import { useCallback, useMemo } from "react";
import { useParams, usePathname } from "next/navigation";
import { NavLink } from "@mantine/core";
import Link from "next/link";

import { SideBarLink } from "@/components/Layout";

import { SidebarProps } from "./Layout.sidebar.type";
import { SidebarUser } from "./Layout.sidebar.user";
import { QuarkloopLogo } from "./Layout.sidebar.logo";
import { Org, Workspace } from "./Layout.sidebar.component";

const Sidebar = (props: SidebarProps) => {
    const { links } = props;

    const { orgSid, workspaceSid }: { orgSid: string; workspaceSid: string } =
        useParams();
    const pathname = usePathname();

    const NameComponent = useMemo(() => {
        if (workspaceSid) {
            return (
                <Workspace
                    orgSid={orgSid}
                    workspaceSid={workspaceSid}
                />
            );
        } else if (orgSid) {
            return <Org orgSid={orgSid} />;
        }
        return null;
    }, [orgSid, workspaceSid]);

    const linkProps = useCallback(
        (link: SideBarLink) => {
            if (link.sublinks.length === 0) {
                return {
                    defaultOpened: false,
                    component: Link,
                    label: link.label,
                    href: link.href || "",
                    className:
                        pathname === link.href
                            ? "bg-violet-100 text-violet-700 font-medium"
                            : "",
                };
            }

            return {
                defaultOpened: true,
                component: undefined,
                label: <p className="font-medium">{link.label}</p>,
                href: "",
                childrenOffset: 35,
            };
        },
        [pathname]
    );

    return (
        <nav className="min-h-screen h-screen sticky top-0 w-80 flex flex-col border-r border-r-neutral-200 text-sm">
            <div className="p-3 basis-14 h-14 flex items-center bg-neutral-100 border-b border-b-neutral-200">
                <QuarkloopLogo />
            </div>

            {NameComponent && (
                <div className="basis-12 max-h-12 flex flex-col border-b border-b-neutral-200">
                    {NameComponent}
                </div>
            )}

            <div className="flex-1 flex flex-col overflow-y-auto">
                {links.map((link, idx) => (
                    <NavLink
                        key={idx}
                        leftSection={link.icon}
                        {...linkProps(link)}>
                        {link.sublinks.length === 0
                            ? null
                            : link.sublinks.map((sublink, subIdx: number) => (
                                  <NavLink
                                      key={subIdx}
                                      {...linkProps(sublink)}
                                      //className="p-2 mx-2 flex items-center gap-3 rounded-lg hover:bg-neutral-200"
                                  />
                              ))}
                    </NavLink>
                ))}
            </div>
            <SidebarUser />
        </nav>
    );
};

export { Sidebar };
