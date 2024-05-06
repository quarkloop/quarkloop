"use client";

import { useEffect, useState } from "react";
import { useGetOrgByIdQuery, useGetOrgUsersQuery } from "./Org.endpoint";
import { StatusCode } from "@quarkloop/lib";
import { Org } from "./Org.schema";

//import { GetOrgUsersApiResponse } from "@quarkloop/types";

interface useGetOrgsProps {
    id: string;
}

interface useGetOrgsResult {
    data: Org | null;
    status: {
        isLoading: boolean;
        isError: boolean;
        isSuccess: boolean;
    };
}

export const useGetOrgById = (props: useGetOrgsProps): useGetOrgsResult => {
    const { id } = props;

    const [org, setOrg] = useState<Org | null>(null);

    const {
        data: osQuery,
        isLoading,
        isError,
        isSuccess,
    } = useGetOrgByIdQuery({ orgSid: id });

    //   useEffect(() => {
    //     if (isLoading) {
    //     }
    //     if (isSuccess) {
    //     }
    //     if (isError) {
    //     }
    //   }, [isLoading, isSuccess, isError]);

    useEffect(() => {
        if (osQuery == null) {
            return;
        }

        // const status = osQuery.status;
        // if (status !== StatusCode.OK) {
        //     // TODO: make showErrorAlert to accept error messages
        //     //showErrorNotification({ message: status.details?.message });
        //     //showErrorAlert(true);
        //     return;
        // }

        const _os = osQuery.data as Org | null;
        if (_os) {
            setOrg(_os);
        }
    }, [osQuery]);

    return {
        data: org,
        status: {
            isLoading,
            isError,
            isSuccess,
        },
    };
};

interface useGetOrgUsersProps {
    orgSid: string;
}

interface useGetOrgUsersResult {
    //data: GetOrgUsersApiResponse | undefined;
    data: any | undefined;
    status: {
        isLoading: boolean;
        isError: boolean;
        isSuccess: boolean;
    };
}

export const useGetOrgUsers = (
    props: useGetOrgUsersProps
): useGetOrgUsersResult => {
    const { orgSid } = props;

    const {
        data: osUsersQuery,
        isLoading,
        isError,
        isSuccess,
    } = useGetOrgUsersQuery({ orgSid });

    //   useEffect(() => {
    //     if (isLoading) {
    //     }
    //     if (isSuccess) {
    //     }
    //     if (isError) {
    //     }
    //   }, [isLoading, isSuccess, isError]);

    useEffect(() => {
        if (osUsersQuery == null) {
            return;
        }

        const status = osUsersQuery.status;
        if (status !== StatusCode.OK) {
            // TODO: make showErrorAlert to accept error messages
            //showErrorNotification({ message: status.details?.message });
            //showErrorAlert(true);
            return;
        }
    }, [osUsersQuery]);

    return {
        data: osUsersQuery,
        status: {
            isLoading,
            isError,
            isSuccess,
        },
    };
};
