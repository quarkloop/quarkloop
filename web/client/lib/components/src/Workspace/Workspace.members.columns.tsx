"use client";

import { useMemo, useState } from "react";
import { Input, InputBase, Combobox, useCombobox } from "@mantine/core";
import { ColumnDef } from "@tanstack/react-table";

import { DataTableV3ColumnHeader } from "@/components/DataTable";
import { WorkspaceMemberRow } from "./Workspace.members.schema";
import { Role, memberRolesSchema } from "@/components/Utils";

const RoleCombobox = ({ member }: { member: WorkspaceMemberRow }) => {
    const combobox = useCombobox({
        onDropdownClose: () => combobox.resetSelectedOption(),
    });

    const [value, setValue] = useState(member.role);

    const options = useMemo(() => {
        const options = Object.values(memberRolesSchema.Enum).map((item) => (
            <Combobox.Option
                value={item}
                key={item}>
                <span className="capitalize">{item}</span>
            </Combobox.Option>
        ));
        return options;
    }, []);
    const InputValue = useMemo(() => {
        if (value) {
            return <span className="capitalize">{value}</span>;
        }
        return <Input.Placeholder>Select role</Input.Placeholder>;
    }, [value]);

    const isCurrentUser = useMemo(
        () => member.currentLoggedInUser?.id === member.user.id,
        [member]
    );

    return (
        <Combobox
            store={combobox}
            onOptionSubmit={(val) => {
                setValue(val as Role);
                combobox.closeDropdown();
            }}>
            <Combobox.Target>
                {isCurrentUser ? (
                    <span className="capitalize font-medium">{value}</span>
                ) : (
                    <InputBase
                        component="button"
                        type="button"
                        pointer
                        rightSection={<Combobox.Chevron />}
                        rightSectionPointerEvents="none"
                        onClick={() => combobox.toggleDropdown()}>
                        {InputValue}
                    </InputBase>
                )}
            </Combobox.Target>
            <Combobox.Dropdown>
                <Combobox.Options>{options}</Combobox.Options>
            </Combobox.Dropdown>
        </Combobox>
    );
};

export const columns: ColumnDef<WorkspaceMemberRow>[] = [
    {
        accessorKey: "account",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Account"
            />
        ),
        cell: ({ row }) => <div>{row.original.user.name}</div>,
        enableSorting: false,
        enableHiding: false,
    },
    {
        accessorKey: "role",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Role"
            />
        ),
        cell: ({ row }) => {
            return <RoleCombobox member={row.original} />;
        },
    },
];
