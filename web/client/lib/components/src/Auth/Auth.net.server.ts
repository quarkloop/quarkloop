"use server";

import { cookies } from "next/headers";

import {
    GetUserAccessApiArgs,
    getUserAccessApiResponseSchema,
} from "./Auth.net.schema";
import { Role } from "@/components/Utils";

const url = {
    userAccess: ({ orgSid, workspaceSid }: GetUserAccessApiArgs) => {
        if (orgSid !== "" && workspaceSid && workspaceSid !== "") {
            return `http://localhost:8000/api/v1/auth/role?orgSid=${orgSid}&workspaceSid=${workspaceSid}`;
        }
        return `http://localhost:8000/api/v1/auth/role?orgSid=${orgSid}`;
    },
};

export async function getUserAccess(
    params: GetUserAccessApiArgs
): Promise<Role | null> {
    const res = await fetch(url.userAccess(params), {
        credentials: "include",
        headers: { Cookie: cookies().toString() },
    });
    if (res.status !== 200 && res.status !== 404) {
        throw new Error("Failed to fetch user access data");
    }

    const jsonBody = await res.json();
    const role = await getUserAccessApiResponseSchema.parseAsync(jsonBody);
    return role as Role | null;
}
