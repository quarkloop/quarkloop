import { NextAuthOptions } from "next-auth";

//import CredentialsProvider from "next-auth/providers/credentials";
import GoogleProvider from "next-auth/providers/google";

import { PrismaAdapter } from "@next-auth/prisma-adapter";
import { prisma } from "@quarkloop/prisma/client";

// const credentialProvider = CredentialsProvider({
//     // The name to display on the sign in form (e.g. "Sign in with...")
//     name: "Credentials",
//     // `credentials` is used to generate a form on the sign in page.
//     // You can specify which fields should be submitted, by adding keys to the `credentials` object.
//     // e.g. domain, username, password, 2FA token, etc.
//     // You can pass any HTML attribute to the <input> tag through the object.
//     credentials: {
//         username: { label: "Username", type: "text", placeholder: "jsmith" },
//         password: { label: "Password", type: "password" }
//     },

//     async authorize(credentials, req) {
//         console.log("authorize", credentials, req);

//         // Add logic here to look up the user from the credentials supplied
//         const user = { id: "1", name: "J Smith", email: "jsmith@example.com" }

//         if (user) {
//             // Any object returned will be saved in `user` property of the JWT
//             return user
//         } else {
//             // If you return null then an error will be displayed advising the user to check their details.
//             return null

//             // You can also Reject this callback with an Error thus the user will be sent to the error page with the error message as a query parameter
//         }
//     }
// });

const googleProvider = GoogleProvider({
    clientId: process.env.GOOGLE_ID!,
    clientSecret: process.env.GOOGLE_SECRET!,
});

export const authOptions: NextAuthOptions = {
    debug: false,
    adapter: PrismaAdapter(prisma),
    secret: process.env.NEXTAUTH_SECRET,
    providers: [
        //credentialProvider,
        googleProvider,
    ],
    callbacks: {
        // async signIn({ user, account, profile, email, credentials }) {
        //   console.log("signIn", { user, account, profile, email, credentials });
        //   return true;
        // },

        async redirect({ url, baseUrl }) {
            //console.log("redirectEventt", { url, baseUrl });
            return Promise.resolve(process.env.QUARKLOOP_WEB_SERVER_URL);
        },

        async session({ session, user, token }) {
            // TODO: this check remains here until the problem with
            // type augmentation in next-auth.d.s is resolved
            if (typeof user.id !== "bigint") {
                throw new Error(`user.id has unknown type: ${typeof user.id}`);
            }

            if (session.user) {
                const userId: bigint = BigInt(user.id);

                const { id, password, passwordSalt, emailVerified, ...u } =
                    user;
                session.user = {
                    id: userId,
                    emailVerified: emailVerified ?? undefined,
                    ...u,
                };
            }

            //console.log("session", { session, user, token });
            return session;
        },

        // async jwt({ token, user, account, profile }) {
        //     console.log("jwt", { token, user, account, profile });
        //     token.access_token = account?.access_token;
        //     return token;
        // },
    },
    session: {
        strategy: "database",
        maxAge: 30 * 24 * 60 * 60, // 30 days
        updateAge: 24 * 60 * 60, // 24 hours
    },
};
