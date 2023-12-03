"use client";

import { ReactElement, ReactNode, useMemo } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

interface TabLinkProps {
    href: string;
    icon?: ReactElement;
    borderBottomColor: string;
    children: ReactNode;
}

const TabLink = (props: TabLinkProps) => {
    const { href, icon, borderBottomColor, children } = props;
    const pathname = usePathname();
    const Icon = useMemo(() => () => icon, [icon]);

    return (
        <Link
            prefetch={false}
            href={href}>
            <div
                className={`${
                    pathname === href ? `border-b-2 ${borderBottomColor}` : ""
                } px-5 py-2 basis-36 flex items-center justify-center gap-2 font-medium rounded-t-lg hover:text-neutral-700 hover:bg-red-50`}>
                {icon && <Icon />}
                <div>{children}</div>
            </div>
        </Link>
    );
};

export { TabLink };
