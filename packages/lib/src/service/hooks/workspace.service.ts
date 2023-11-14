"use client";

import { enpointApi } from "../api";
import {
    GetWorkspaceByIdApiResponse,
    GetWorkspaceByIdApiArgs,
    // GetWorkspaceByNameApiResponse,
    // GetWorkspaceByNameApiArgs,
    GetWorkspacesByOsIdApiResponse,
    GetWorkspacesByOsIdApiArgs,
    CreateWorkspaceApiResponse,
    CreateWorkspaceApiArgs,
    UpdateWorkspaceApiResponse,
    UpdateWorkspaceApiArgs,
    DeleteWorkspaceApiResponse,
    DeleteWorkspaceApiArgs,
} from "./workspace.type";

import {
    getWorkspaceByIdSchema,
    getWorkspacesByOsIdSchema,
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
                    url: `/os/${queryArg.osId}/workspaces/${queryArg.id}`,
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
            //         url: `/os/${queryArg.osId}/workspace`,
            //         method: "GET",
            //         params: { name: queryArg.name },
            //     }),
            //     providesTags: ['workspace'],
            // }),

            // getWorkspacesByOsId endpoint
            getWorkspacesByOsId: builder.query<
                GetWorkspacesByOsIdApiResponse,
                GetWorkspacesByOsIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os/${queryArg.osId}/workspaces`,
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    dataSchema: getWorkspacesByOsIdSchema,
                },
            }),

            // createWorkspace endpoint
            createWorkspace: builder.mutation<
                CreateWorkspaceApiResponse,
                CreateWorkspaceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os/${queryArg.osId}/workspaces`,
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
                    url: `/os/${queryArg.osId}/workspaces/${
                        queryArg.workspace!.id
                    }`,
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
                    url: `/os/${queryArg.osId}/workspaces/${queryArg.id}`,
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
    useGetWorkspacesByOsIdQuery,
    useCreateWorkspaceMutation,
    useUpdateWorkspaceMutation,
    useDeleteWorkspaceMutation,
} = endpoint;
