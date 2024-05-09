"use client";

import { Suspense } from "react";
import Link from "next/link";
import { IconCircleCheck } from "@tabler/icons-react";

import { useSignin } from "./Auth.signin.hook";
import { SigninProps } from "./Auth.signin.schema";
import { GoogleButton } from "./Auth.signin.google";

const Signin = (props: SigninProps) => {
    const { variant } = props;

    const {
        auth: session,
        triggers: { signInWithCredentials, submitForm, redirectTo },
        ui: { form },
    } = useSignin(props);

    return (
        <div className="px-3 py-10 flex-1 flex flex-col gap-16 md:px-10 md:py-32 md:flex-row md:justify-center">
            <div className="py-5 flex flex-col gap-4 md:gap-9 md:basis-1/3">
                <p className="text-2xl md:text-4xl font-semibold">
                    {variant === "login"
                        ? "Ready to scale your business?"
                        : "Join thousands worldwide who automate their work using Quarkloop."}
                </p>
                <div className="flex flex-col gap-2">
                    {variant === "login" ? (
                        <p>
                            Easily collaborate with your team with shared and
                            project connections, a centralized login, and more
                        </p>
                    ) : (
                        <>
                            <div className="flex items-center gap-2">
                                <IconCircleCheck stroke={1.6} />
                                <p>Easy setup, no coding required</p>
                            </div>
                            <div className="flex items-center gap-2">
                                <IconCircleCheck stroke={1.6} />
                                <p>Free forever for core features</p>
                            </div>
                        </>
                    )}
                </div>
            </div>

            <form
                onSubmit={form.onSubmit(submitForm)}
                className="flex flex-col gap-10 md:basis-80">
                <div className="flex flex-col items-center justify-center gap-4 md:flex-row text-4xl font-medium">
                    <p className="">Quarkloop</p>
                </div>

                <div className="p-6 flex flex-col gap-6 rounded-lg border">
                    {variant === "login" && (
                        <p className="flex items-center justify-center font-medium text-xl md:text-2xl">
                            Log in to your account
                        </p>
                    )}
                    <Suspense>
                        <GoogleButton variant={variant} />
                    </Suspense>
                    <div className="flex flex-col items-center gap-2 md:flex-row">
                        <p>
                            {variant === "login"
                                ? "Don't have an account yet?"
                                : "Already have an account?"}
                        </p>
                        <Link
                            href={variant === "login" ? "/signup" : "/login"}
                            className="underline">
                            {variant === "login" ? "Sign up" : "Log in"}
                        </Link>
                    </div>
                </div>
            </form>
        </div>
    );
};

export { Signin };

{
    /* <div className="px-5 py-7 flex flex-col gap-5 rounded-lg border bg-slate-50">
      <TextInput
          required={true}
          label="Email address"
          style={{ label: { fontWeight: 400 } }}
          {...form.getInputProps("email")} />
      <PasswordInput
          required={true}
          label="Password"
          style={{ label: { fontWeight: 400 } }}
          {...form.getInputProps("password")} />
      <Button
          fullWidth
          type="submit"
          fw={400}>
          Sign in
      </Button>
  </div>

  <div className="flex items-center justify-stretch gap-4">
      <div className="flex-1 border-b"></div>
      <div>or</div>
      <div className="flex-1 border-b"></div>
  </div> */
}
