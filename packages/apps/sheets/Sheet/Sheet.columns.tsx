"use client";

import { Badge } from "@mantine/core";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import moment from "moment";

import { DataTableColumnHeader } from "@quarkloop/components";

import { labels } from "./Sheet.data";
import { SheetRow } from "./Sheet.type";

export const columns: ColumnDef<SheetRow>[] = [
    // {
    //     id: "select",
    //     header: ({ table }) => (
    //         <Checkbox
    //             checked={table.getIsAllPageRowsSelected()}
    //             onChange={(value) => table.toggleAllPageRowsSelected(!!value)}
    //             aria-label="Select all"
    //             className="translate-y-[2px]"
    //         />
    //     ),
    //     cell: ({ row }) => (
    //         <Checkbox
    //             checked={row.getIsSelected()}
    //             onChange={(value) => row.toggleSelected(!!value)}
    //             aria-label="Select row"
    //             className="translate-y-[2px]"
    //         />
    //     ),
    //     enableSorting: false,
    //     enableHiding: false,
    // },
    {
        accessorKey: "id",
        header: ({ column }) => (
            <DataTableColumnHeader
                column={column}
                title="ID"
            />
        ),
        cell: ({ row }) => <div className="w-[80px]">{row.getValue("id")}</div>,
        enableSorting: false,
        enableHiding: false,
    },
    {
        accessorKey: "name",
        header: ({ column }) => (
            <DataTableColumnHeader
                column={column}
                title="Name"
            />
        ),
        cell: ({ row }) => {
            const label = labels.find(
                (label) => label.value === row.original.name
            );

            return (
                <div className="flex space-x-2">
                    {label && <Badge variant="outline">{label.label}</Badge>}
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
        accessorKey: "description",
        header: ({ column }) => (
            <DataTableColumnHeader
                column={column}
                title="Description"
            />
        ),
        cell: ({ row }) => {
            return (
                <div className="flex w-[100px] items-center">
                    {row.getValue("description")}
                </div>
            );
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id));
        },
    },
    {
        accessorKey: "updatedAt",
        header: ({ column }) => (
            <DataTableColumnHeader
                column={column}
                title="Updated At"
            />
        ),
        cell: ({ row }) => {
            const date: Date = row.getValue("updatedAt");
            return (
                <div className="flex items-center">
                    <span>{moment(date).fromNow()}</span>
                </div>
            );
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id));
        },
    },
];
