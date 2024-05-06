"use client";

import { useCallback, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useForm } from "@mantine/form";

import { UseHookReturnType } from "@quarkloop/lib";
import { SigninProps } from "./Signin.type";

interface SigninForm {
    email: string;
    password: string;
}

export const useSignin = (props: SigninProps): any => {
    // : UseHookReturnType<SigninForm, Session>
    const { variant } = props;
    const router = useRouter();

    const form = useForm<SigninForm>({
        initialValues: {
            email: "",
            password: "",
        },
    });
    const { data: session } = useSession();

    const redirectTo = useCallback(
        (path: string) => {
            router.push(path);
        },
        [router]
    );

    const signInWithCredentials = useCallback(async (values: SigninForm) => {
        // const response = await signIn("credentials", {
        //     redirect: false,
        //     username: values.email,
        //     password: values.password,
        // });
    }, []);

    const submitForm = useCallback(
        async (values: SigninForm) => {
            return signInWithCredentials(values);
        },
        [signInWithCredentials]
    );

    useEffect(() => {
        console.log(
            "getSession",
            session,
            process.env.NEXTAUTH_URL_INTERNAL,
            process.env.NEXTAUTH_SECRET
        );

        if (session) {
            redirectTo("/");
        }
    }, [session, redirectTo]);

    return {
        auth: session,
        triggers: {
            signInWithCredentials,
            submitForm,
            redirectTo,
        },
        ui: { form },
    };
};
