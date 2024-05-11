"use client";

import { forwardRef, useCallback, useMemo } from "react";
import { useRouter } from "next/navigation";
import { SegmentedControl, Textarea, TextInput } from "@mantine/core";
import { useForm, zodResolver } from "@mantine/form";

import { Button } from "@/ui/primitives";
import { useCreateOrgMutation } from "./Org.net.client";
import { orgVisibilityData } from "./Org.util";
import { OrgCreateFormData } from "./Org.create.form";
import { mutationOrgFormSchema } from "./Org.schema";

export const useOrgCreate = () => {
    const [createOrg] = useCreateOrgMutation();
    const onCreateOrg = useCallback(
        async (data: OrgCreateFormData) => {
            try {
                const org = await createOrg({ payload: data }).unwrap();
            } catch (error) {
                console.error("[onCreateOrg] error:", error);
            }
        },
        [createOrg]
    );

    return {
        onCreateOrg,
    };
};

const OrgCreate = forwardRef<HTMLFormElement>((_, ref) => {
    const { onCreateOrg } = useOrgCreate();

    const form = useForm<OrgCreateFormData>({
        validate: zodResolver(mutationOrgFormSchema),
        transformValues: (value) => mutationOrgFormSchema.parse(value),
        initialValues: {
            sid: "",
            name: "",
            description: "",
            visibility: "private",
        },
    });
    const visibilityData = useMemo(orgVisibilityData, []);

    return (
        <form
            ref={ref}
            className="pt-14 pb-8 flex-1 flex flex-col items-center gap-4"
            onSubmit={form.onSubmit(onCreateOrg)}>
            <p className="w-[45%] text-3xl font-medium">
                Create a new organization
            </p>
            <div className="w-[45%] border-t border-t-neutral-200" />
            <div className="w-[45%] flex flex-col gap-2 text-neutral-500 text-sm">
                <div className="flex items gap-3">
                    <p>&#9432;</p>
                    <p>
                        An organization contains all workspaces, including the
                        services.
                    </p>
                </div>
                <div className="flex items gap-3">
                    <p>&#9432;</p>
                    <p className="italic">
                        Required fields are marked with an asterisk (
                        <span className="text-red-600">*</span>).
                    </p>
                </div>
            </div>
            <div className="w-[45%] border-t border-t-neutral-200" />
            <div className="w-[45%] flex flex-col gap-5">
                <div className="flex items-start gap-2">
                    <TextInput
                        readOnly
                        label="Owner"
                        value="Reza Ebrahimi"
                        styles={{
                            input: {
                                backgroundColor: "var(--mantine-color-gray-1)",
                            },
                        }}
                    />
                    <p className="pt-7 flex items-center text-xl">/</p>
                    <TextInput
                        required
                        label="Org id"
                        // styles={{
                        //     error: {
                        //         color: "green",
                        //     },
                        // }}
                        {...form.getInputProps("sid")}
                        //error="ssdf is available"
                    />
                </div>
                <div className="flex items-start gap-2">
                    <TextInput
                        required
                        label="Org name"
                        {...form.getInputProps("name")}
                    />
                </div>
                <div className="flex flex-col gap-1">
                    <div className="font-medium text-sm">Visibility</div>
                    <div className="flex gap-4">
                        <SegmentedControl
                            data={visibilityData}
                            orientation="vertical"
                            transitionDuration={250}
                            transitionTimingFunction="linear"
                            styles={{
                                root: {
                                    width: "120px",
                                    border: "1px solid var(--mantine-color-gray-3)",
                                },
                            }}
                            {...form.getInputProps("visibility")}
                        />
                        <div className="py-2 flex flex-col gap-3 text-sm text-neutral-500">
                            <p className="flex-1">
                                Anyone on the internet can see this
                                organization.
                            </p>
                            <p className="flex-1">
                                You choose who can see and work in this
                                organization.
                            </p>
                        </div>
                    </div>
                </div>
                <Textarea
                    required
                    label="Description"
                    autosize
                    minRows={4}
                    maxRows={4}
                    {...form.getInputProps("description")}
                />
            </div>
            <div className="w-[45%] border-t border-t-neutral-200" />
            <div className="w-[45%] flex items gap-3 text-neutral-500">
                <p>&#9432;</p>
                <p>
                    You are creating a{" "}
                    <span className="font-semibold">
                        {form.values.visibility}
                    </span>{" "}
                    organization in your personal account.
                </p>
            </div>
            <div className="w-[45%] border-t border-t-neutral-200" />
            <div className="w-[45%] flex items-center justify-end">
                <Button type="submit">Create organization</Button>
            </div>
        </form>
    );
});

OrgCreate.displayName = "OrgCreate";

export { OrgCreate };
