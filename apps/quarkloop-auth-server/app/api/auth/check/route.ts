import { NextResponse } from "next/server";
import { getServerSession } from "next-auth";

import { authOptions } from "@quarkloop/lib";

export async function GET(request: Request, params: { params: any }) {
    const session = await getServerSession(authOptions);
    return NextResponse.json(
        {},
        {
            status: session ? 200 : 401,
            statusText: session ? "OK" : "Unauthorized",
        }
    );
}
