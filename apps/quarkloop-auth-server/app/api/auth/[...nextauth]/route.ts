import NextAuth from "next-auth";

import { authOptions } from "@/auth";

const handler = NextAuth(authOptions);

export async function GET(request: Request, params: { params: any }) {
    return handler(request, params);
}

export async function POST(request: Request, params: { params: any }) {
    return handler(request, params);
}

// resources
// https://stackoverflow.com/a/73510917/1789729
// https://github.com/nextauthjs/next-auth/discussions/4378
// https://github.com/wpcodevo/nextauth-nextjs13-prisma/tree/main
// https://stackoverflow.com/questions/73507135/nextauth-handle-refresh-token-rotation-reddit-with-database
// https://stackoverflow.com/questions/71363829/next-auth-js-i-cant-get-token-with-gettokenreq
// https://stackoverflow.com/questions/74984195/how-to-decrypt-session-token-next-auth-jwt-token
// https://stackoverflow.com/questions/69068495/how-to-get-the-provider-access-token-in-next-auth
