"use client";

import { useMemo, useState } from "react";
import { Input, InputBase, Combobox, useCombobox } from "@mantine/core";

import Link from "next/link";
import moment from "moment";
import { ColumnDef, Row } from "@tanstack/react-table";
import { IconDots } from "@tabler/icons-react";

import {
    Button,
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuShortcut,
    DropdownMenuTrigger,
} from "@/ui/primitives";
import { DataTableV3ColumnHeader } from "@/components/DataTable";
import { OrgMemberRow, orgMemberSchema } from "./Org.members.schema";
import { Role, memberRolesSchema } from "@/components/Utils";

const RoleCombobox = ({ member }: { member: OrgMemberRow }) => {
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

interface RowActionsProps<TData> {
    row: Row<TData>;
}

// export function RowActions<TData>(props: RowActionsProps<TData>) {
//     const { row } = props;
//     const actionRow = orgMemberSchema.parse(row.original);

//     return (
//         <DropdownMenu>
//             <DropdownMenuTrigger asChild>
//                 <Button
//                     variant="ghost"
//                     className="flex h-8 w-8 p-0 data-[state=open]:bg-muted">
//                     <IconDots className="h-4 w-4" />
//                 </Button>
//             </DropdownMenuTrigger>
//             <DropdownMenuContent
//                 align="end"
//                 className="w-[160px]">
//                 <DropdownMenuItem
//                     onClick={() => actionRow.onViewClick?.(actionRow)}>
//                     View
//                 </DropdownMenuItem>
//                 <DropdownMenuItem
//                     onClick={() => actionRow.onUpdateClick?.(actionRow)}>
//                     Update
//                 </DropdownMenuItem>
//                 <DropdownMenuSeparator />
//                 <DropdownMenuItem
//                     onClick={() => actionRow.onDeleteClick?.(actionRow)}>
//                     Delete
//                     <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
//                 </DropdownMenuItem>
//             </DropdownMenuContent>
//         </DropdownMenu>
//     );
// }

export const columns: ColumnDef<OrgMemberRow>[] = [
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
    // {
    //     accessorKey: "updatedAt",
    //     header: ({ column }) => (
    //         <DataTableV3ColumnHeader
    //             column={column}
    //             title="Updated At"
    //         />
    //     ),
    //     cell: ({ row }) => {
    //         const date: Date = row.getValue("updatedAt");
    //         return (
    //             <div>
    //                 <span>{moment(date).fromNow()}</span>
    //             </div>
    //         );
    //     },
    //     filterFn: (row, id, value) => {
    //         return value.includes(row.getValue(id));
    //     },
    // },
    // {
    //     id: "actions",
    //     header: ({ column }) => (
    //         <DataTableV3ColumnHeader
    //             column={column}
    //             title="Action"
    //         />
    //     ),
    //     cell: ({ row }) => <RowActions row={row} />,
    // },
];
