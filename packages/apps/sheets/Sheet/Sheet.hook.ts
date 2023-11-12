import { useCallback, useEffect } from "react";

import {
    StatusCode,
    useGetWorkspacesByOsIdQuery,
    useGetPlanSubscriptionByUserSession,
} from "@quarkloop/lib";
import { SheetHookReturnType } from "./Sheet.type";
import { useParams } from "next/navigation";

export const useSheet = (): SheetHookReturnType => {
    const { osId } = useParams();

    const { data: workspaceList } = useGetWorkspacesByOsIdQuery({
        osId: osId as string,
    });

    // const [createApp] = useCreateAppMutation();

    // const { subscription, metrics, refetch } =
    //     useGetPlanSubscriptionByUserSession();

    // const disableCreateApp = useMemo(
    //     () =>
    //         (metrics?.App.used || 0) <
    //         (subscription?.plan.features.maxworkspaceList || 0),
    //     [metrics?.App.used, subscription?.plan.features.maxworkspaceList]
    // );

    // const [
    //     isCreateAppModalOpened,
    //     { open: openCreateAppModal, clworkspacee: clworkspaceeCreateAppModal },
    // ] = useDisclworkspaceure(false);

    const onCreateNewApp = useCallback(
        async () => {
            // const metadata: AppMetadata = {
            //     stage: "draft",
            //     title: "Setup application form",
            //     type: "bluecard",
            //     typeLabel: "Blue Card residency permit",
            // };
            // const createdApp = await createApp({
            //     workspaceId: workspaceId as string,
            //     workspaceId: form.values.workspaceId,
            //     name: form.values.name,
            //     status: "Off",
            //     visibility: "PrivateAnyOneWithLink",
            //     // TODO: make updatedAt optional because the createdAt is older than updatedAt at time of creation
            //     metadata: { ...metadata },
            //     // TODO: AppType
            //     type: 0,
            // }).unwrap();
            // const status = createdApp.status;
            // if (status.statusCode !== StatusCode.CREATED) {
            //     // TODO: make showErrorAlert to accept error messages
            //     //showErrorNotification({ message: status.details?.message });
            //     //showErrorAlert(true);
            //     clworkspaceeCreateAppModal();
            //     return;
            // }
            // clworkspaceeCreateAppModal();
        },
        [
            // createApp,
            // clworkspaceeCreateAppModal,
            // workspaceId,
            // form.values.name,
            // form.values.workspaceId,
        ]
    );

    if (workspaceList?.status !== StatusCode.OK) {
        return null;
    }

    return {
        status: true,
        data: workspaceList.data.map((workspace) => ({
            id: workspace.id,
            name: workspace.name,
            description: workspace.description,
            path: workspace.path,
            updatedAt: workspace.createdAt,
        })),
        triggers: null,
        //triggers: { onCreateNewApp, openCreateAppModal, clworkspaceeCreateAppModal },
    };
};
