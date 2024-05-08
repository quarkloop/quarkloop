"use client";

import Link from "next/link";
import { Skeleton } from "@mantine/core";
import { IconSettings } from "@tabler/icons-react";

import { useGetUserQuery } from "@quarkloop/lib";
import { UserImage } from "@/components/User";

const SidebarUser = () => {
    const { data: user } = useGetUserQuery();

    if (user == null) {
        return <UserSkeleton />;
    }

    return (
        <div className="px-3 py-2 flex items-center gap-3 bg-neutral-100 border-t border-t-neutral-200">
            <Link
                href={`/users/${user.id}`}
                className="basis-10 w-10 h-10 relative rounded-full bg-[#d5d5d5]">
                <UserImage
                    value={user.image}
                    onChange={() => {}}
                    imageAlt={user.name}
                />
            </Link>
            <div className="flex-1 flex flex-col justify-center">
                <p>{user.name}</p>
                <p className="text-neutral-500">@{user.username}</p>
            </div>
            <Link
                href={`/users/${user.id}/settings`}
                className="rounded-full hover:bg-neutral-100">
                <IconSettings
                    size="1.5rem"
                    stroke={1.1}
                />
            </Link>
        </div>
    );
};

import React from "react";

export const UserSkeleton = () => (
    <div className="p-3 flex items-center gap-3 border-t border-t-[#eeeeee]">
        <Skeleton
            circle
            height={50}
        />
        <div className="flex-1 flex flex-col gap-2">
            <Skeleton
                height={9}
                width="40%"
            />
            <Skeleton
                height={9}
                width="70%"
            />
            <Skeleton
                height={9}
                width="90%"
            />
        </div>
    </div>
);

export { SidebarUser };
