import Papa from "papaparse";

import { DataTableColumnHeader } from "@quarkloop/components";
import { Badge } from "@mantine/core";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";

import { columns } from "./Sheet.columns";
import { data } from "./Sheet.data";

export function parseCsv() {}

export function buildData(): any {
    return data;
}

export function buildColumns(columnsData: unknown[]): any {
    return columnsData.map<ColumnDef<any>>((col) => ({
        accessorKey: col as string,
        header: ({ column }) => (
            <DataTableColumnHeader
                column={column}
                title={col as string}
            />
        ),
        cell: ({ row }) => {
            return (
                <div className="flex space-x-2">
                    {row.getValue(col as string)}
                </div>
            );
        },
        enableSorting: true,
        enableMultiSort: true,
        enableResizing: true,
        enableHiding: true,
    }));
}

const csv = `
Column 1,Column 2,Column 3,Column 4, Column 5, Column 6
1-1,1-2,1-3,1-4,1-5
2-1,2-2,2-3,2-4,2-5
3-1,3-2,3-3,3-4,3-5
4,5,6,7
`;

type TableInfo = {
    data: unknown[];
    columns: unknown[];
};

export function build(): TableInfo {
    const json = Papa.parse(csv.trim(), {
        header: true,
        skipEmptyLines: true,
    });

    console.log("RRRRR", json);

    if (json.data.length > 0) {
        const columns = buildColumns(json.meta.fields as any);
        const rows = json.data;

        return {
            columns: columns,
            data: rows,
        };
    }

    return {
        columns: [],
        data: [],
    };
}
