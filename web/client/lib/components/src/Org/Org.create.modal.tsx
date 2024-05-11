"use client";

import { useCallback, useMemo } from "react";
import { Drawer } from "@mantine/core";

import { OrgCreateForm, OrgCreateFormData } from "./Org.create.form";

interface OrgCreateModalProps {
    initialValue?: OrgCreateFormData;
    defaultOpened: boolean;
    onFormSubmit: (data: OrgCreateFormData) => Promise<void>;
    readonly open: () => void;
    readonly close: () => void;
}

const OrgCreateModal = (props: OrgCreateModalProps) => {
    const { initialValue, defaultOpened, onFormSubmit, open, close } = props;

    const initialValues = useMemo(
        () =>
            ({
                sid: initialValue?.sid ?? "",
                name: initialValue?.name ?? "",
                description: initialValue?.description ?? "",
                visibility: initialValue?.visibility ?? "private",
            } as OrgCreateFormData),
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
            title="New organization"
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
            <OrgCreateForm
                readOnly={false}
                initialValues={initialValues}
                onFormSubmit={onFormSubmit}
            />
        </Drawer>
    );
};

export { OrgCreateModal };
