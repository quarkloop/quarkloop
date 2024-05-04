import { NextResponse } from "next/server";
import { getServerSession } from "next-auth";

import { authOptions } from "@/lib";

export async function GET(request: Request, params: { params: any }) {
    const session = await getServerSession(authOptions);
    const body = session?.user ?? {};
    // console.log(request.headers.get("cookie"));

    console.log(body);

    // 401: Unauthenticated
    // 200: OK
    return NextResponse.json(body, {
        status: session ? 200 : 401,
    });
}
