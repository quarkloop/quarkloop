"use client";

import { useMemo } from "react";
import { z } from "zod";

import {
    Form,
    FormButton,
    FormInput,
    FormSegmentedControl,
    FormTextArea,
} from "@/ui/form";

import { orgSchema } from "./Org.schema";
import { orgRowSchema } from "./Org.list.schema";
import { orgVisibilityData } from "./Org.util";

export type OrgGeneralSettingsFormData = z.infer<typeof orgRowSchema>;

interface FormProps {
    readOnly: boolean;
    initialValues: OrgGeneralSettingsFormData;
    onFormSubmit: (data: OrgGeneralSettingsFormData) => Promise<void>;
}

const OrgGeneralSettingsForm = (props: FormProps) => {
    const { readOnly, initialValues, onFormSubmit } = props;

    const visibilityData = useMemo(orgVisibilityData, []);
    const FormStatus = useMemo(() => {
        const vis = (
            <span className="font-semibold">{initialValues.visibility}</span>
        );
        return (
            <p>
                You are updating a {vis} organization in your personal account.
            </p>
        );
    }, [initialValues.visibility]);

    return (
        <Form
            initialValues={initialValues}
            schema={orgSchema}
            onFormSubmit={onFormSubmit}
            className="py-3 flex-1 flex flex-col gap-4">
            <div className="border-t border-t-neutral-200" />
            <div className="flex flex-col gap-2 text-neutral-500 text-sm">
                <div className="flex items gap-3">
                    <p>&#9432;</p>
                    <p className="italic">
                        Required fields are marked with an asterisk (
                        <span className="text-red-600">*</span>).
                    </p>
                </div>
            </div>
            <div className="border-t border-t-neutral-200" />
            <div className="flex flex-col gap-5">
                <div className="flex items-start gap-2">
                    <FormInput
                        readOnly
                        name="owner"
                        label="Owner"
                        value="Reza Ebrahimi"
                        styles={{
                            input: {
                                backgroundColor: "var(--mantine-color-gray-1)",
                            },
                        }}
                    />
                    <p className="pt-7 flex items-center text-xl">/</p>
                    <FormInput
                        readOnly={readOnly}
                        required
                        name="sid"
                        label="Org id"
                        // styles={{
                        //     error: {
                        //         color: "green",
                        //     },
                        // }}
                        //error="ssdf is available"
                    />
                </div>
                <div className="flex items-start gap-2">
                    <FormInput
                        readOnly={readOnly}
                        required
                        name="name"
                        label="Org name"
                    />
                </div>
                <div className="flex flex-col gap-1">
                    <div className="font-medium text-sm">Visibility</div>
                    <div className="flex gap-4">
                        <FormSegmentedControl
                            readOnly={readOnly}
                            data={visibilityData}
                            name="visibility"
                            orientation="vertical"
                            transitionDuration={250}
                            transitionTimingFunction="linear"
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
                <FormTextArea
                    autosize
                    required
                    readOnly={readOnly}
                    name="description"
                    label="Description"
                    minRows={4}
                    maxRows={4}
                />
            </div>
            <div className="border-t border-t-neutral-200" />
            <div className="flex items gap-3 text-neutral-500">
                <p>&#9432;</p>
                {FormStatus}
            </div>
            <div className="border-t border-t-neutral-200" />
            <div className="flex items-center justify-end">
                <FormButton className="w-[150px]">Save changes</FormButton>
            </div>
        </Form>
    );
};

export { OrgGeneralSettingsForm };