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

import { workspaceSchema } from "./Workspace.schema";
import { workspaceRowSchema } from "./Workspace.list.schema";

export type WorkspaceGeneralSettingsFormData = z.infer<
    typeof workspaceRowSchema
>;

interface FormProps {
    readOnly: boolean;
    initialValues: WorkspaceGeneralSettingsFormData;
    onFormSubmit: (data: WorkspaceGeneralSettingsFormData) => Promise<void>;
}

const WorkspaceGeneralSettingsForm = (props: FormProps) => {
    const { readOnly, initialValues, onFormSubmit } = props;

    const FormStatus = useMemo(() => {
        const vis = (
            <span className="font-semibold">{initialValues.visibility}</span>
        );
        return (
            <p>You are updating a {vis} workspace in your personal account.</p>
        );
    }, [initialValues.visibility]);

    return (
        <Form
            initialValues={initialValues}
            schema={workspaceSchema}
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
                        label="Workspace id"
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
                        label="Workspace name"
                    />
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

export { WorkspaceGeneralSettingsForm };
