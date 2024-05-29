"use client";

import { enpointApi } from "@/api/net";
import {
    GetUserAccessApiResponse,
    GetUserAccessApiArgs,
    getUserAccessApiArgsSchema,
    getUserAccessApiResponseSchema,
} from "./Auth.net.schema";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["role"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getUserAccess endpoint
            getUserAccess: builder.query<
                GetUserAccessApiResponse,
                GetUserAccessApiArgs
            >({
                query: (queryArg) => ({
                    url: `/auth/role`,
                    method: "GET",
                    params: {
                        orgSid: queryArg.orgSid,
                        workspaceSid: queryArg.workspaceSid,
                    },
                }),
                providesTags: ["role"],
                extraOptions: {
                    argSchema: getUserAccessApiArgsSchema,
                    responseSchema: getUserAccessApiResponseSchema,
                },
            }),

            // // grantUserAccess endpoint
            // grantUserAccess: builder.mutation<GrantUserAccessApiResponse, GrantUserAccessApiArgs>(
            //     {
            //         query: (queryArg) => ({
            //             url: `/manage/orgs`,
            //             method: "POST",
            //             body: queryArg,
            //         }),
            //         invalidatesTags: ["org"],
            //         extraOptions: {
            //             argSchema: grantUserAccessApiArgsSchema,
            //             responseSchema: grantUserAccessApiResponseSchema,
            //         },
            //     }
            // ),

            // // revokeUserAccess endpoint
            // revokeUserAccess: builder.mutation<
            //     DeleteUserAccessApiResponse,
            //     DeleteUserAccessApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/manage/${queryArg.orgSid}`,
            //         method: "DELETE",
            //     }),
            //     invalidatesTags: ["org"],
            //     extraOptions: {
            //         argSchema: revokeUserAccessApiArgsSchema,
            //     },
            // }),
        }),
    });

export const {
    useGetUserAccessQuery,
    // useGrantUserAccessMutation,
    // useRevokeUserAccessMutation,
} = endpoint;
