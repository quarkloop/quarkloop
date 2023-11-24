"use client";

import { enpointApi } from "../api";
import {
    GetServiceListApiResponse,
    GetServiceListApiArgs,
    GetServiceByIdApiResponse,
    GetServiceByIdApiArgs,
    CreateServiceApiResponse,
    CreateServiceApiArgs,
    UpdateServiceApiResponse,
    UpdateServiceApiArgs,
    DeleteServiceApiResponse,
    DeleteServiceApiArgs,
} from "./service.type";

import {
    getServiceListSchema,
    getServiceByIdSchema,
    createServiceSchema,
    updateServiceSchema,
    deleteServiceSchema,
} from "./service.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["service"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getServiceById endpoint
            getServiceById: builder.query<
                GetServiceByIdApiResponse,
                GetServiceByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/services/${queryArg.serviceId}`,
                    method: "GET",
                }),
                providesTags: ["service"],
                extraOptions: {
                    dataSchema: getServiceByIdSchema,
                },
            }),

            // getServiceList endpoint
            getServiceList: builder.query<
                GetServiceListApiResponse,
                GetServiceListApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/services`,
                    method: "GET",
                }),
                providesTags: ["service"],
                extraOptions: {
                    dataSchema: getServiceListSchema,
                },
            }),

            // createService endpoint
            createService: builder.mutation<
                CreateServiceApiResponse,
                CreateServiceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/services`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["service"],
                extraOptions: {
                    dataSchema: createServiceSchema,
                },
            }),

            // updateService endpoint
            updateService: builder.mutation<
                UpdateServiceApiResponse,
                UpdateServiceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/services/${queryArg.serviceId}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["service"],
                extraOptions: {
                    dataSchema: updateServiceSchema,
                },
            }),

            // deleteService endpoint
            deleteService: builder.mutation<
                DeleteServiceApiResponse,
                DeleteServiceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/services/${queryArg.serviceId}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["service"],
                extraOptions: {
                    dataSchema: deleteServiceSchema,
                },
            }),
        }),
    });

export const {
    useGetServiceByIdQuery,
    useLazyGetServiceByIdQuery,
    useGetServiceListQuery,
    useLazyGetServiceListQuery,

    useCreateServiceMutation,
    useUpdateServiceMutation,
    useDeleteServiceMutation,
} = endpoint;
