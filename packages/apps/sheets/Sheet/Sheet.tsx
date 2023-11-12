"use client";

import { SheetView } from "./Sheet.view";
import { useSheet } from "./Sheet.hook";
import { build } from "./Sheet.func";

const Sheet = () => {
    const workspaceList = useSheet();

    // if (!workspaceList?.status) {
    //     return null;
    // }

    const tableInfo = build();

    console.log("json", tableInfo);

    return (
        <SheetView
            columns={tableInfo.columns as any}
            data={tableInfo.data}
        />
    );
};

export { Sheet };
