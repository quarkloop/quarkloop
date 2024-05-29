"use client";

import { useMemo } from "react";

import { useGetUserQuery } from "@/components/User";
import { DataTableV3 } from "@/components/DataTable";

import { useGetOrgMembersQuery } from "./Org.net.client";
import { columns } from "./Org.members.columns";
import { OrgMemberRow } from "./Org.members.schema";

interface OrgMemberListProps {
    orgSid: string;
}

const OrgMemberList = ({ orgSid }: OrgMemberListProps) => {
    const { data: currentLoggedInUser } = useGetUserQuery();
    const { data: memberList } = useGetOrgMembersQuery({ orgSid });

    const data: OrgMemberRow[] = useMemo(() => {
        if (currentLoggedInUser) {
            const members = memberList?.data.map((member, idx) => ({
                ...member,
                currentLoggedInUser,
            }));
            return members || [];
        }
        return [];
    }, [memberList, currentLoggedInUser]);

    //const [opened, { open, close }] = useDisclosure(false);

    return (
        <div className="px-20 py-10 flex-1 flex-col gap-3 md:flex">
            <div className="py-2 flex items-center justify-between">
                <div className="flex flex-col gap-3">
                    <h2 className="text-3xl font-semibold tracking-tight">
                        Organization members
                    </h2>
                    <p className="text-muted-foreground">All members</p>
                </div>
                {/* <Button onClick={open}>New member</Button> */}
            </div>
            <DataTableV3
                enableHeader
                enablePagination
                toolbar={{
                    filterColumnName: "account",
                    filterColumns: [
                        {
                            columnName: "account",
                            columnTitle: "Account",
                            options: [
                                {
                                    label: "Account",
                                    value: "account",
                                },
                            ],
                        },
                    ],
                }}
                data={data}
                columns={columns}
            />
        </div>
    );
};

export { OrgMemberList };
