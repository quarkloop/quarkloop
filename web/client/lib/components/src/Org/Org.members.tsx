"use client";

import { useDisclosure } from "@mantine/hooks";

import { Button } from "@/ui/primitives";
import { DataTableV3 } from "@/components/DataTable";

import { columns } from "./Org.members.columns";

const OrgMemberList = () => {
    const [opened, { open, close }] = useDisclosure(false);

    return (
        <div className="px-20 py-10 flex-1 flex-col gap-3 md:flex">
            <div className="py-2 flex items-center justify-between">
                <div className="flex flex-col gap-3">
                    <h2 className="text-3xl font-semibold tracking-tight">
                        Organization members
                    </h2>
                    <p className="text-muted-foreground">All members</p>
                </div>
                <Button onClick={open}>New member</Button>
            </div>
            <DataTableV3
                enableHeader
                enablePagination
                toolbar={{
                    filterColumnName: "name",
                    filterColumns: [
                        {
                            columnName: "name",
                            columnTitle: "Name",
                            options: [
                                {
                                    label: "Name",
                                    value: "name",
                                },
                            ],
                        },
                    ],
                }}
                data={orgList || []}
                columns={columns}
            />
        </div>
    );
};

export { OrgMemberList };
