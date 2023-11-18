"use client";

import { enpointApi } from "../api";
import {
    GetWorkspaceByIdApiResponse,
    GetWorkspaceByIdApiArgs,
    // GetWorkspaceByNameApiResponse,
    // GetWorkspaceByNameApiArgs,
    GetWorkspacesByOrgIdApiResponse,
    GetWorkspacesByOrgIdApiArgs,
    CreateWorkspaceApiResponse,
    CreateWorkspaceApiArgs,
    UpdateWorkspaceApiResponse,
    UpdateWorkspaceApiArgs,
    DeleteWorkspaceApiResponse,
    DeleteWorkspaceApiArgs,
} from "./workspace.type";

import {
    getWorkspaceByIdSchema,
    getWorkspacesByOrgIdSchema,
    createWorkspaceSchema,
    updateWorkspaceSchema,
    deleteWorkspaceSchema,
} from "./workspace.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["workspace"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getWorkspaceById endpoint
            getWorkspaceById: builder.query<
                GetWorkspaceByIdApiResponse,
                GetWorkspaceByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/workspaces/${queryArg.id}`,
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    dataSchema: getWorkspaceByIdSchema,
                },
            }),

            // // getWorkspaceByName endpoint
            // getWorkspaceByName: builder.query<
            //     GetWorkspaceByNameApiResponse,
            //     GetWorkspaceByNameApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/workspace`,
            //         method: "GET",
            //         params: { name: queryArg.name },
            //     }),
            //     providesTags: ['workspace'],
            // }),

            // getWorkspacesByOrgId endpoint
            getWorkspacesByOrgId: builder.query<
                GetWorkspacesByOrgIdApiResponse,
                GetWorkspacesByOrgIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/workspaces`,
                    method: "GET",
                    params: {
                        orgId: queryArg.orgId,
                    },
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    dataSchema: getWorkspacesByOrgIdSchema,
                },
            }),

            // createWorkspace endpoint
            createWorkspace: builder.mutation<
                CreateWorkspaceApiResponse,
                CreateWorkspaceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/workspaces`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    dataSchema: createWorkspaceSchema,
                },
            }),

            // updateWorkspace endpoint
            updateWorkspace: builder.mutation<
                UpdateWorkspaceApiResponse,
                UpdateWorkspaceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/workspaces/${queryArg.workspace!.id}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    dataSchema: updateWorkspaceSchema,
                },
            }),

            // deleteWorkspace endpoint
            deleteWorkspace: builder.mutation<
                DeleteWorkspaceApiResponse,
                DeleteWorkspaceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/workspaces/${queryArg.id}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    dataSchema: deleteWorkspaceSchema,
                },
            }),
        }),
    });

export const {
    useGetWorkspaceByIdQuery,
    useLazyGetWorkspaceByIdQuery,
    useGetWorkspacesByOrgIdQuery,
    useCreateWorkspaceMutation,
    useUpdateWorkspaceMutation,
    useDeleteWorkspaceMutation,
} = endpoint;
