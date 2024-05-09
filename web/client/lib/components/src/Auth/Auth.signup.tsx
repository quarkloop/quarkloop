"use client";

import { Button, TextInput, PasswordInput } from "@mantine/core";
import { IconBrandGoogle } from "@tabler/icons-react";

import { useSignUp } from "./Auth.signup.hook";

export interface SignUpProps {}

const SignUp = (props: SignUpProps) => {
    const {
        auth: session,
        triggers: { signInWithCredentials, signInWithGoogle, submitForm },
        ui: { form },
    } = useSignUp();

    return (
        <div className="py-20 flex-1 flex justify-center">
            <form onSubmit={form.onSubmit(submitForm)}>
                <div className="w-[22rem] flex flex-col gap-6">
                    <p className="text-2xl font-medium">
                        Sign in to organization
                    </p>

                    <div className="px-5 py-7 flex flex-col gap-5 rounded-lg border bg-slate-50">
                        <TextInput
                            required={true}
                            label="Email address"
                            style={{ label: { fontWeight: 400 } }}
                            {...form.getInputProps("email")}
                        />
                        <PasswordInput
                            required={true}
                            label="Password"
                            style={{ label: { fontWeight: 400 } }}
                            {...form.getInputProps("password")}
                        />
                        <Button
                            fullWidth
                            type="submit"
                            fw={400}>
                            Register
                        </Button>
                    </div>

                    <div className="flex items-center justify-stretch gap-4">
                        <div className="flex-1 border-b"></div>
                        <div>or</div>
                        <div className="flex-1 border-b"></div>
                    </div>

                    <Button
                        fullWidth
                        fw={400}
                        leftSection={
                            <IconBrandGoogle
                                size="1.4rem"
                                stroke={2.0}
                            />
                        }
                        onClick={signInWithGoogle}>
                        Register with Google
                    </Button>
                </div>
            </form>
        </div>
    );
};

export { SignUp };
