"use client";

import React from "react";
import Link from "next/link";

import { Button } from "@quarkloop/radix/components/ui/button";
import { DataTable } from "@quarkloop/components";

import { SheetProps } from "./Sheet.type";

const SheetView = (props: SheetProps) => {
    const { columns, data } = props;

    return (
        <div className="hidden p-3 h-full flex-1 flex-col space-y-8 md:flex">
            <div className="p-2 flex items-center justify-between">
                <Link href="/os/new">
                    <Button>New record</Button>
                </Link>
            </div>
            <DataTable
                data={data}
                columns={columns}
            />
        </div>
    );
};

export { SheetView };
