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
import { WorkspaceRow } from "./Workspace.list.schema";
import { useWorkspaceData } from "./Workspace.util";

interface CellProps extends HTMLAttributes<HTMLDivElement> {
    row: Row<WorkspaceRow>;
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
    const rowData = useWorkspaceData(row.original);

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
                    href={`/manage/${row.original.orgSid}/${rowData.sid}`}
                    className="flex items-center text-base font-semibold hover:text-neutral-700">
                    {rowData.name}
                </Link>
                <div className="px-2 flex items-center bg-neutral-200 text-neutral-500 text-sm rounded-lg">
                    {rowData.visibility}
                </div>
            </div>
            <div className="w-[300px] md:w-[600px] inline-block truncate">
                {rowData.description}
            </div>
            <div className="flex items-center gap-1">{history}</div>
        </div>
    );
};

export const columns: ColumnDef<WorkspaceRow>[] = [
    {
        accessorKey: "workspace",
        header: ({ column }) => (
            <DataTableV3ColumnHeader
                column={column}
                title="Workspace"
            />
        ),
        cell: ({ row }) => <Cell row={row} />,
        enableSorting: false,
        enableHiding: false,
    },
];
