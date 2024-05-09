"use client";

import { useMemo, useState } from "react";
import { z } from "zod";

import {
    Form,
    FormButton,
    FormInput,
    FormSegmentedControl,
    FormTextArea,
} from "@/ui/form";
import { visibilityData } from "@/components/Utils";

import { workspaceSchema } from "./Workspace.schema";
import { workspaceRowSchema } from "./Workspace.list.schema";

export type WorkspaceCreateFormData = z.infer<typeof workspaceRowSchema>;

interface FormProps {
    readOnly: boolean;
    initialValues: WorkspaceCreateFormData;
    onFormSubmit: (data: WorkspaceCreateFormData) => Promise<void>;
}

const WorkspaceCreateForm = (props: FormProps) => {
    const { readOnly, initialValues, onFormSubmit } = props;

    const [visibility, setVisibility] = useState<string>(
        initialValues.visibility ?? "private"
    );
    const visData = useMemo(visibilityData, []);
    const FormStatus = useMemo(() => {
        const vis = <span className="font-semibold">{visibility}</span>;
        return (
            <p>You are creating a {vis} workspace in your personal account.</p>
        );
    }, [visibility]);

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
                    <p>An workspace contains all services.</p>
                </div>
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
                <div className="flex flex-col gap-1">
                    <div className="font-medium text-sm">Visibility</div>
                    <div className="flex gap-4">
                        <FormSegmentedControl
                            readOnly={readOnly}
                            data={visData}
                            name="visibility"
                            orientation="vertical"
                            transitionDuration={250}
                            transitionTimingFunction="linear"
                            onChange={setVisibility}
                        />
                        <div className="py-2 flex flex-col gap-3 text-sm text-neutral-500">
                            <p className="flex-1">
                                Anyone on the internet can see this workspace.
                            </p>
                            <p className="flex-1">
                                You choose who can see and work in this
                                workspace.
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
                <FormButton className="w-[150px]">Create</FormButton>
            </div>
        </Form>
    );
};

export { WorkspaceCreateForm };
