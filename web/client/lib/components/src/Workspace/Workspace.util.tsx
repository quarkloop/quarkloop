import { useMemo } from "react";
import moment from "moment";

import { Workspace } from "./Workspace.schema";

export const useWorkspaceData = (data?: Workspace) => {
    const workspaceData = useMemo(() => {
        if (data == null) {
            return {};
        }

        const id = data.id;
        const sid = data.sid;
        const name = data.name;
        const description = data.description;
        const visibility = data.visibility;
        const updatedAt = moment(data.updatedAt).fromNow();
        const updatedBy = data.updatedBy;
        const createdAt = moment(data.createdAt).fromNow();
        const createdBy = data.createdBy;
        const path = data.path;

        return {
            id,
            sid,
            name,
            description,
            visibility,
            createdBy,
            createdAt,
            updatedAt,
            updatedBy,
            path,
        };
    }, [data]);

    return workspaceData;
};
