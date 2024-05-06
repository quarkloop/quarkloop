import { useMemo } from "react";
import { Box, Center } from "@mantine/core";
import { IconLock, IconWorld } from "@tabler/icons-react";

import moment from "moment";

import { Org } from "./Org.schema";

export const useOrgData = (data?: Org) => {
    const orgData = useMemo(() => {
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

    return orgData;
};

export const orgVisibilityData = () => [
    {
        value: "public",
        label: (
            <Center>
                <IconWorld size={16} />
                <Box ml={10}>Public</Box>
            </Center>
        ),
    },
    {
        value: "private",
        label: (
            <Center>
                <IconLock size={16} />
                <Box ml={10}>Private</Box>
            </Center>
        ),
    },
];
