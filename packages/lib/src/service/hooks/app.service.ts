"use client";

import { enpointApi } from "../api";
import {
    GetAppListApiResponse,
    GetAppListApiArgs,
    GetAppByIdApiResponse,
    GetAppByIdApiArgs,
    CreateAppApiResponse,
    CreateAppApiArgs,
    UpdateAppApiResponse,
    UpdateAppApiArgs,
    DeleteAppApiResponse,
    DeleteAppApiArgs,
} from "./app.type";

import {
    getAppListSchema,
    getAppByIdSchema,
    createAppSchema,
    updateAppSchema,
    deleteAppSchema,
} from "./app.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["app"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getAppById endpoint
            getAppById: builder.query<GetAppByIdApiResponse, GetAppByIdApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/apps/${queryArg.id}`,
                        method: "GET",
                    }),
                    providesTags: ["app"],
                    extraOptions: {
                        dataSchema: getAppByIdSchema,
                    },
                }
            ),

            // getAppList endpoint
            getAppList: builder.query<GetAppListApiResponse, GetAppListApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/apps`,
                        method: "GET",
                        params: {
                            osId: queryArg.osId,
                            workspaceId: queryArg.workspaceId,
                        },
                    }),
                    providesTags: ["app"],
                    extraOptions: {
                        dataSchema: getAppListSchema,
                    },
                }
            ),

            // createApp endpoint
            createApp: builder.mutation<CreateAppApiResponse, CreateAppApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/apps`,
                        method: "POST",
                        body: queryArg,
                    }),
                    invalidatesTags: ["app"],
                    extraOptions: {
                        dataSchema: createAppSchema,
                    },
                }
            ),

            // updateApp endpoint
            updateApp: builder.mutation<UpdateAppApiResponse, UpdateAppApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/apps/${queryArg.id}`,
                        method: "PUT",
                        body: queryArg,
                    }),
                    invalidatesTags: ["app"],
                    extraOptions: {
                        dataSchema: updateAppSchema,
                    },
                }
            ),

            // deleteApp endpoint
            deleteApp: builder.mutation<DeleteAppApiResponse, DeleteAppApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/apps/${queryArg.id}`,
                        method: "DELETE",
                    }),
                    invalidatesTags: ["app"],
                    extraOptions: {
                        dataSchema: deleteAppSchema,
                    },
                }
            ),
        }),
    });

export const {
    useGetAppByIdQuery,
    useLazyGetAppByIdQuery,
    useGetAppListQuery,
    useLazyGetAppListQuery,

    useCreateAppMutation,
    useUpdateAppMutation,
    useDeleteAppMutation,
} = endpoint;
