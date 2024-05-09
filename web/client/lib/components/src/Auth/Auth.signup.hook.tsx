import { useCallback, useEffect } from "react";
import { useRouter } from "next/navigation";
import { signIn, useSession } from "next-auth/react";
import { Session } from "next-auth";
import { useForm } from "@mantine/form";

import { UseHookReturnType } from "@quarkloop/lib";

interface SignUpForm {
    email: string;
    password: string;
}

export const useSignUp = (): any => {
    // UseHookReturnType<SignUpForm, Session>
    const router = useRouter();
    const form = useForm<SignUpForm>({
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

    const signInWithCredentials = useCallback(async (values: SignUpForm) => {
        const response = await signIn("credentials", {
            redirect: false,
            username: values.email,
            password: values.password,
        });
        console.log("signIn response", response, values.password);
    }, []);

    const signInWithGoogle = useCallback(async () => {
        const response = await signIn("google", { redirect: false });
        console.log("signIn response", response);
    }, []);

    const submitForm = useCallback(
        async (values: SignUpForm) => {
            return signInWithCredentials(values);
        },
        [signInWithCredentials]
    );

    // useEffect(() => {
    //     return () => {
    //         redirectTo("/");
    //     }
    // }, []);

    useEffect(() => {
        if (session) {
            redirectTo("/");
        }
    }, [session, redirectTo]);

    return {
        auth: session,
        triggers: {
            signInWithCredentials,
            signInWithGoogle,
            submitForm,
        },
        ui: { form },
    };
};
