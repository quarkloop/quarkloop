"use client";

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

interface RowActionsProps<TData> {
    row: Row<TData>;
}

export function RowActions<TData>(props: RowActionsProps<TData>) {
    const { row } = props;
    const actionRow = orgMemberSchema.parse(row.original);

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    variant="ghost"
                    className="flex h-8 w-8 p-0 data-[state=open]:bg-muted">
                    <IconDots className="h-4 w-4" />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent
                align="end"
                className="w-[160px]">
                <DropdownMenuItem
                    onClick={() => actionRow.onViewClick?.(actionRow)}>
                    View
                </DropdownMenuItem>
                <DropdownMenuItem
                    onClick={() => actionRow.onUpdateClick?.(actionRow)}>
                    Update
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem
                    onClick={() => actionRow.onDeleteClick?.(actionRow)}>
                    Delete
                    <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}

export const columns: ColumnDef<OrgMemberRow>[] = [
    {
        accessorKey: "sid",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Scope ID"
            />
        ),
        cell: ({ row }) => <div>{row.getValue("sid")}</div>,
        enableSorting: false,
        enableHiding: false,
    },
    {
        accessorKey: "name",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Name"
            />
        ),
        cell: ({ row }) => {
            return (
                <div>
                    <Link
                        href={row.original.path}
                        className="max-w-[500px] truncate font-medium">
                        {row.getValue("name")}
                    </Link>
                </div>
            );
        },
    },
    {
        accessorKey: "visibility",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Visibility"
            />
        ),
        cell: ({ row }) => {
            return (
                <div className="capitalize">{row.getValue("visibility")}</div>
            );
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id));
        },
    },
    {
        accessorKey: "description",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Description"
            />
        ),
        cell: ({ row }) => {
            return <div>{row.getValue("description")}</div>;
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id));
        },
    },
    {
        accessorKey: "updatedAt",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Updated At"
            />
        ),
        cell: ({ row }) => {
            const date: Date = row.getValue("updatedAt");
            return (
                <div>
                    <span>{moment(date).fromNow()}</span>
                </div>
            );
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id));
        },
    },
    {
        id: "actions",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Action"
            />
        ),
        cell: ({ row }) => <RowActions row={row} />,
    },
];
