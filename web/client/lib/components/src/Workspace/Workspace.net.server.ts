"use server";

import { cookies } from "next/headers";
import { ZodSchema } from "zod";

import {
    GetWorkspaceByIdApiArgs,
    GetWorkspaceByIdApiResponse,
    getWorkspaceByIdSchema,
} from "./Workspace.net.schema";

interface Link {
    query: (params: any) => string;
    queryFn: (link: Link, params: any) => Promise<unknown | null>;
    schema: {
        arg?: ZodSchema;
        response: ZodSchema;
    };
}

interface ServerLink {
    [key: string]: Link;
}

const serverLinks: ServerLink = {
    getWorkspaceById: {
        query: ({ workspaceSid }: GetWorkspaceByIdApiArgs) =>
            `http://localhost:8000/api/v1/manage/${workspaceSid}`,
        queryFn: getFn,
        schema: {
            response: getWorkspaceByIdSchema,
        },
    },
};

const getWorkspaceById = async (
    params: GetWorkspaceByIdApiArgs
): Promise<GetWorkspaceByIdApiResponse | null> => {
    const link = serverLinks.getWorkspaceById;
    return (await link.queryFn(link, params)) as GetWorkspaceByIdApiResponse;
};

export async function getFn(link: Link, params: any): Promise<unknown | null> {
    const res = await fetch(link.query(params), {
        credentials: "include",
        headers: { Cookie: cookies().toString() },
    });
    if (res.status !== 200) {
        return null;
    }

    const jsonBody = await res.json();
    return await link.schema.response.parseAsync(jsonBody);
}

export { getWorkspaceById };
