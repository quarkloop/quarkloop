import { NextResponse } from "next/server";
import { getServerSession } from "next-auth";

import { authOptions } from "@/auth";

export async function GET(request: Request, params: { params: any }) {
    const session = await getServerSession(authOptions);
    const body = session?.user ?? {};

    return NextResponse.json(body, {
        status: session ? 200 : 401,
    });
}
