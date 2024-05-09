"use client";

import { useCallback } from "react";
import { Button } from "@mantine/core";
import { useSearchParams } from "next/navigation";
import { signIn } from "next-auth/react";
import { IconBrandGoogle } from "@tabler/icons-react";

import { AuthVariant } from "./Auth.signin.schema";

const planPath = {
    free: "/user/account-setup?plan=free&next=setup",
    standard: "/user/account-setup?plan=standard&next=setup",
    pro: "/user/account-setup?plan=pro&next=setup",
    enterprise: "/user/account-setup?plan=enterprise&next=setup",
};

const GoogleButton = ({ variant }: { variant: AuthVariant }) => {
    const searchParams = useSearchParams();

    const signInWithGoogle = useCallback(async () => {
        let callbackUrl: string | undefined = "http://localhost:3000";
        if (variant === "signup") {
            const next = searchParams.get("next");
            if (next) {
                if (next === "setup-standard-account") {
                    callbackUrl = planPath.standard;
                } else if (next === "setup-pro-account") {
                    callbackUrl = planPath.pro;
                } else if (next === "setup-enterprise-account") {
                    callbackUrl = planPath.enterprise;
                }
            } else {
                callbackUrl = planPath.free;
            }
        }

        const _ = await signIn("google", {
            redirect: true,
            callbackUrl,
        });
    }, [searchParams, variant]);

    return (
        <Button
            fullWidth
            fw={700}
            leftSection={<IconBrandGoogle size="1.5rem" />}
            onClick={signInWithGoogle}>
            {variant === "login" ? "Log in with Google" : "Sign up with Google"}
        </Button>
    );
};

export { GoogleButton };
