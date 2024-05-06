"use client";

import React, { HTMLAttributes, useMemo } from "react";
import Link from "next/link";
import moment from "moment";
import { ColumnDef, Row } from "@tanstack/react-table";
import { IconDots } from "@tabler/icons-react";

import { cn } from "@/ui/lib";
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
import { OrgRow, orgRowSchema } from "./Org.list.schema";
import { useOrgData } from "./Org.util";

interface CellProps extends HTMLAttributes<HTMLDivElement> {
    row: Row<OrgRow>;
}

export const Cell = (props: CellProps) => {
    const { row, className, ...rest } = props;
    // const rowData = useMemo(() => {
    //     const id = row.original.id;
    //     const sid = row.original.sid;
    //     const name = row.original.name;
    //     const description = row.original.description;
    //     const visibility = row.original.visibility;
    //     const updatedAt = moment(row.original.updatedAt).fromNow();
    //     const updatedBy = row.original.updatedBy;
    //     const createdAt = moment(row.original.createdAt).fromNow();
    //     const createdBy = row.original.createdBy;
    //     const path = row.original.path;

    //     return {
    //         id,
    //         sid,
    //         name,
    //         description,
    //         visibility,
    //         createdBy,
    //         createdAt,
    //         updatedAt,
    //         updatedBy,
    //         path,
    //     };
    // }, [row.original]);
    const rowData = useOrgData(row.original);

    const history = useMemo((): JSX.Element => {
        return (
            <>
                <Link
                    href={rowData.path ?? ""}
                    className="hover:text-neutral-700"
                    rel="noopener noreferrer"
                    target="_blank">
                    @{rowData.sid}
                </Link>
                {rowData.updatedBy ? (
                    <>
                        <p>updated {rowData.updatedAt} by</p>
                        <Link
                            href={`/users/${rowData.updatedBy}`}
                            className="hover:text-neutral-700">
                            {rowData.updatedBy}
                        </Link>
                    </>
                ) : (
                    <>
                        <p>created {rowData.createdAt} by</p>
                        <Link
                            href={`/users/${rowData.createdBy}`}
                            className="hover:text-neutral-700">
                            {rowData.createdBy}
                        </Link>
                    </>
                )}
            </>
        );
    }, [rowData]);

    return (
        <div
            className={cn("p-4 flex flex-col gap-1", className)}
            {...rest}>
            <div className="flex items-center gap-4">
                <Link
                    href={`/manage/${rowData.sid}`}
                    className="flex items-center text-base font-semibold hover:text-neutral-700">
                    {rowData.name}
                </Link>
                <div className="px-2 flex items-center bg-neutral-200 text-neutral-500 text-sm rounded-lg">
                    {rowData.visibility}
                </div>
                <Link
                    href={rowData.path ?? ""}
                    rel="noopener noreferrer"
                    target="_blank"
                    className="flex items-center underline hover:text-neutral-700">
                    view profile
                </Link>
            </div>
            <div className="w-[300px] md:w-[600px] inline-block truncate">
                {rowData.description}
            </div>
            <div className="flex items-center gap-1">{history}</div>
        </div>
    );
};

export const columns: ColumnDef<OrgRow>[] = [
    {
        accessorKey: "orgRow",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Organization"
            />
        ),
        cell: ({ row }) => <Cell row={row} />,
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id));
        },
        enableSorting: false,
        enableHiding: false,
    },
];

// interface RowActionsProps<TData> {
//     row: Row<TData>;
// }

// export function RowActions<TData>(props: RowActionsProps<TData>) {
//     const { row } = props;
//     const actionRow = orgRowSchema.parse(row.original);

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

// export const columns: ColumnDef<OrgRow>[] = [
//     {
//         accessorKey: "sid",
//         header: ({ column }) => (
//             <DataTableV3ColumnHeader
//                 column={column}
//                 title="Scope ID"
//             />
//         ),
//         cell: ({ row }) => <Cell>{row.getValue("sid")}</Cell>,
//         enableSorting: false,
//         enableHiding: false,
//     },
//     {
//         accessorKey: "name",
//         header: ({ column }) => (
//             <DataTableV3ColumnHeader
//                 column={column}
//                 title="Name"
//             />
//         ),
//         cell: ({ row }) => {
//             return (
//                 <Cell>
//                     <Link
//                         href={row.original.path}
//                         className="max-w-[500px] truncate font-medium">
//                         {row.getValue("name")}
//                     </Link>
//                 </Cell>
//             );
//         },
//     },
//     {
//         accessorKey: "visibility",
//         header: ({ column }) => (
//             <DataTableV3ColumnHeader
//                 column={column}
//                 title="Visibility"
//             />
//         ),
//         cell: ({ row }) => {
//             return (
//                 <Cell className="capitalize">{row.getValue("visibility")}</Cell>
//             );
//         },
//         filterFn: (row, id, value) => {
//             return value.includes(row.getValue(id));
//         },
//     },
//     {
//         accessorKey: "description",
//         header: ({ column }) => (
//             <DataTableV3ColumnHeader
//                 column={column}
//                 title="Description"
//             />
//         ),
//         cell: ({ row }) => {
//             return <Cell>{row.getValue("description")}</Cell>;
//         },
//         filterFn: (row, id, value) => {
//             return value.includes(row.getValue(id));
//         },
//     },
//     {
//         accessorKey: "updatedAt",
//         header: ({ column }) => (
//             <DataTableV3ColumnHeader
//                 column={column}
//                 title="Updated At"
//             />
//         ),
//         cell: ({ row }) => {
//             const date: Date = row.getValue("updatedAt");
//             return (
//                 <Cell>
//                     <span>{moment(date).fromNow()}</span>
//                 </Cell>
//             );
//         },
//         filterFn: (row, id, value) => {
//             return value.includes(row.getValue(id));
//         },
//     },
//     {
//         id: "actions",
//         header: ({ column }) => (
//             <DataTableV3ColumnHeader
//                 column={column}
//                 title="Action"
//             />
//         ),
//         cell: ({ row }) => <RowActions row={row} />,
//     },
// ];
