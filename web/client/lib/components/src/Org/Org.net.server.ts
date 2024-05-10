"use server";

import { cookies } from "next/headers";
import { ZodSchema } from "zod";

import {
    GetOrgByIdApiArgs,
    GetOrgByIdApiResponse,
    getOrgByIdSchema,
} from "./Org.net.schema";

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
    getOrgById: {
        query: ({ orgSid }: GetOrgByIdApiArgs) =>
            `http://localhost:8000/api/v1/manage/${orgSid}`,
        queryFn: getFn,
        schema: {
            response: getOrgByIdSchema,
        },
    },
};

const getOrgById = async (
    params: GetOrgByIdApiArgs
): Promise<GetOrgByIdApiResponse | null> => {
    const link = serverLinks.getOrgById;
    return (await link.queryFn(link, params)) as GetOrgByIdApiResponse;
};

export async function getFn(link: Link, params: any): Promise<unknown | null> {
    try {
        const res = await fetch(link.query(params), {
            credentials: "include",
            headers: { Cookie: cookies().toString() },
        });
        if (res.status !== 200) {
            return null;
        }

        const jsonBody = await res.json();
        return await link.schema.response.parseAsync(jsonBody);
    } catch (error) {
        console.error(error);
        return null;
    }
}

export { getOrgById };
