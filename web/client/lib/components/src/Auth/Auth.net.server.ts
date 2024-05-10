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
    try {
        const res = await fetch(url.userAccess(params), {
            credentials: "include",
            headers: { Cookie: cookies().toString() },
            cache: "no-store",
        });
        if (res.status !== 200 && res.status !== 404) {
            return null;
        }

        const jsonBody = await res.json();
        const role = await getUserAccessApiResponseSchema.parseAsync(jsonBody);
        return role as Role | null;
    } catch (error) {
        console.error(error);
        return null;
    }
}
