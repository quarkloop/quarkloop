"use client";

import { useCallback, useMemo } from "react";
import { Drawer } from "@mantine/core";

import {
    WorkspaceCreateForm,
    WorkspaceCreateFormData,
} from "./Workspace.create.form";

interface WorkspaceCreateModalProps {
    initialValue?: WorkspaceCreateFormData;
    defaultOpened: boolean;
    onFormSubmit: (data: WorkspaceCreateFormData) => Promise<void>;
    readonly open: () => void;
    readonly close: () => void;
}

const WorkspaceCreateModal = (props: WorkspaceCreateModalProps) => {
    const { initialValue, defaultOpened, onFormSubmit, open, close } = props;

    const initialValues = useMemo(
        () =>
            ({
                id: initialValue?.id ?? 0,
                sid: initialValue?.sid ?? "",
                name: initialValue?.name ?? "",
                description: initialValue?.description ?? "",
                visibility: initialValue?.visibility ?? "private",
                path: initialValue?.path ?? "",
                createdBy: initialValue?.createdBy ?? "",
                updatedBy: initialValue?.updatedBy ?? "",
            } as WorkspaceCreateFormData),
        [initialValue]
    );

    const onCloseDrawer = useCallback(() => {
        close();
    }, []);

    return (
        <Drawer
            opened={defaultOpened}
            onClose={onCloseDrawer}
            position="right"
            title="New workspace"
            size="lg"
            styles={{
                content: {
                    display: "flex",
                    flexDirection: "column",
                },
                body: {
                    flex: 1,
                    display: "flex",
                    flexDirection: "column",
                },
            }}>
            <WorkspaceCreateForm
                readOnly={false}
                initialValues={initialValues}
                onFormSubmit={onFormSubmit}
            />
        </Drawer>
    );
};

export { WorkspaceCreateModal };
