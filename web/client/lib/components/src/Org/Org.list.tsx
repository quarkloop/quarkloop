"use client";

import React, { useCallback, useMemo } from "react";
import { useDisclosure } from "@mantine/hooks";

import { Button } from "@/ui/primitives";
import { DataTableV3 } from "@/components/DataTable";

import { columns } from "./Org.list.columns";
import { useGetOrgsQuery, useCreateOrgMutation } from "./Org.endpoint";
import { OrgCreateFormData } from "./Org.create.form";
import { OrgCreateModal } from "./Org.create.modal";

const OrgList = () => {
    const [opened, { open, close }] = useDisclosure(false);

    const { data: orgData } = useGetOrgsQuery();
    const [createOrg] = useCreateOrgMutation();

    const onCreateClick = useCallback(async (data: OrgCreateFormData) => {
        try {
            const result = await createOrg({ payload: data }).unwrap();
            if (result) {
                close();
            }
        } catch (error) {
            console.error("[onCreateClick] error:", error);
        }
    }, []);
    const orgList = useMemo(() => orgData?.data, [orgData]);

    return (
        <div className="px-20 py-10 flex-1 flex-col gap-3 md:flex">
            <div className="py-2 flex items-center justify-between">
                <div className="flex flex-col gap-3">
                    <h2 className="text-3xl font-semibold tracking-tight">
                        Orgs
                    </h2>
                    <p className="text-muted-foreground">All organizations</p>
                </div>
                <Button onClick={open}>New organization</Button>
                <OrgCreateModal
                    defaultOpened={opened}
                    open={open}
                    close={close}
                    onFormSubmit={onCreateClick}
                />
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

export { OrgList };
