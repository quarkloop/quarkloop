"use client";

import { enpointApi } from "../api";
import {
    GetAppInstanceListApiResponse,
    GetAppInstanceListApiArgs,
    GetAppInstanceByIdApiResponse,
    GetAppInstanceByIdApiArgs,
    CreateAppInstanceApiResponse,
    CreateAppInstanceApiArgs,
    UpdateAppInstanceApiResponse,
    UpdateAppInstanceApiArgs,
    DeleteAppInstanceApiResponse,
    DeleteAppInstanceApiArgs,
} from "./app_instance.type";

import {
    getAppInstanceListSchema,
    getAppInstanceByIdSchema,
    createAppInstanceSchema,
    updateAppInstanceSchema,
    deleteAppInstanceSchema,
} from "./app_instance.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["appInstance"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getAppInstanceById endpoint
            getAppInstanceById: builder.query<
                GetAppInstanceByIdApiResponse,
                GetAppInstanceByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/apps/${queryArg.projectId}/instances/${queryArg.instanceId}`,
                    method: "GET",
                }),
                providesTags: ["appInstance"],
                extraOptions: {
                    dataSchema: getAppInstanceByIdSchema,
                },
            }),

            // getAppInstanceList endpoint
            getAppInstanceList: builder.query<
                GetAppInstanceListApiResponse,
                GetAppInstanceListApiArgs
            >({
                query: (queryArg) => ({
                    url: `/apps/${queryArg.projectId}/instances`,
                    method: "GET",
                }),
                providesTags: ["appInstance"],
                extraOptions: {
                    dataSchema: getAppInstanceListSchema,
                },
            }),

            // createAppInstance endpoint
            createAppInstance: builder.mutation<
                CreateAppInstanceApiResponse,
                CreateAppInstanceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/apps/${queryArg.projectId}/instances`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["appInstance"],
                extraOptions: {
                    dataSchema: createAppInstanceSchema,
                },
            }),

            // updateAppInstance endpoint
            updateAppInstance: builder.mutation<
                UpdateAppInstanceApiResponse,
                UpdateAppInstanceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/apps/${queryArg.projectId}/instances/${queryArg.instanceId}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["appInstance"],
                extraOptions: {
                    dataSchema: updateAppInstanceSchema,
                },
            }),

            // deleteAppInstance endpoint
            deleteAppInstance: builder.mutation<
                DeleteAppInstanceApiResponse,
                DeleteAppInstanceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/apps/${queryArg.projectId}/instances/${queryArg.instanceId}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["appInstance"],
                extraOptions: {
                    dataSchema: deleteAppInstanceSchema,
                },
            }),
        }),
    });

export const {
    useGetAppInstanceByIdQuery,
    useLazyGetAppInstanceByIdQuery,
    useGetAppInstanceListQuery,
    useLazyGetAppInstanceListQuery,

    useCreateAppInstanceMutation,
    useUpdateAppInstanceMutation,
    useDeleteAppInstanceMutation,
} = endpoint;
