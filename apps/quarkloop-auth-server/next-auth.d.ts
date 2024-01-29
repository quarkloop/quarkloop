import NextAuth, { DefaultUser, DefaultSession } from "next-auth";
import type { AdapterUser as DefaultAdapterUser } from "next-auth/adapters";

import { User as PrismaUser } from "@quarkloop/types";

type NewUser = Omit<PrismaUser, "password" | "passwordSalt">;

declare module "next-auth" {
    interface Session {
        user: Omit<DefaultUser, "id"> & NewUser;
    }

    // With using AdapterUser below, this doesn't have any effect
    interface User extends Omit<DefaultUser, "id">, NewUser {}
}

// TODO: Type augmentation is not working with 'id' field
// So the problem is resolved with a check in session callback
declare module "next-auth/adapters" {
    interface AdapterUser extends Omit<DefaultAdapterUser, "id">, PrismaUser {
        id: bigint; // id is still string type, augmentation not works
        name: string; // augmentation works, from DefaultUser
        email: string; // augmentation works, from DefaultUser
        image: string; // augmentation works, from DefaultUser
    }
}
