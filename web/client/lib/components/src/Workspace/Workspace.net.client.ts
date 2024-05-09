"use client";

import { enpointApi } from "@quarkloop/lib";
import {
    GetWorkspaceByIdApiResponse,
    GetWorkspaceByIdApiArgs,
    GetWorkspaceMembersApiResponse,
    GetWorkspaceMembersApiArgs,
    CreateWorkspaceApiResponse,
    CreateWorkspaceApiArgs,
    UpdateWorkspaceByIdApiResponse,
    UpdateWorkspaceByIdApiArgs,
    DeleteWorkspaceByIdApiResponse,
    DeleteWorkspaceByIdApiArgs,
    createWorkspaceApiArgsSchema,
    createWorkspaceApiResponseSchema,
    GetWorkspacesApiResponse,
    GetWorkspacesApiArgs,
    getWorkspacesSchema,
    getWorkspaceByIdSchema,
    updateWorkspaceByIdApiArgsSchema,
    deleteWorkspaceByIdApiArgsSchema,
    ChangeWorkspaceVisibilityApiResponse,
    ChangeWorkspaceVisibilityApiArgs,
    changeWorkspaceVisibilityApiArgsSchema,
    getWorkspaceMembersApiArgsSchema,
    getWorkspaceMembersApiResponseSchema,
} from "./Workspace.net.schema";

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
                    url: `/manage/${queryArg.orgSid}/${queryArg.workspaceSid}`,
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    responseSchema: getWorkspaceByIdSchema,
                },
            }),

            // getWorkspaces endpoint
            getWorkspaces: builder.query<
                GetWorkspacesApiResponse,
                GetWorkspacesApiArgs
            >({
                query: (queryArg) => ({
                    url: "/manage/workspaces",
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    responseSchema: getWorkspacesSchema,
                },
            }),

            // createWorkspace endpoint
            createWorkspace: builder.mutation<
                CreateWorkspaceApiResponse,
                CreateWorkspaceApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/workspaces`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    argSchema: createWorkspaceApiArgsSchema,
                    responseSchema: createWorkspaceApiResponseSchema,
                },
            }),

            // updateWorkspaceById endpoint
            updateWorkspaceById: builder.mutation<
                UpdateWorkspaceByIdApiResponse,
                UpdateWorkspaceByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/workspaces/${queryArg.workspaceSid}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    argSchema: updateWorkspaceByIdApiArgsSchema,
                },
            }),

            // deleteWorkspaceById endpoint
            deleteWorkspaceById: builder.mutation<
                DeleteWorkspaceByIdApiResponse,
                DeleteWorkspaceByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/workspaces/${queryArg.workspaceSid}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    argSchema: deleteWorkspaceByIdApiArgsSchema,
                },
            }),

            // changeWorkspaceVisibility endpoint
            changeWorkspaceVisibility: builder.mutation<
                ChangeWorkspaceVisibilityApiResponse,
                ChangeWorkspaceVisibilityApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/workspaces/${queryArg.workspaceSid}/visibility`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["workspace"],
                extraOptions: {
                    argSchema: changeWorkspaceVisibilityApiArgsSchema,
                },
            }),

            // getWorkspaceMembers endpoint
            getWorkspaceMembers: builder.query<
                GetWorkspaceMembersApiResponse,
                GetWorkspaceMembersApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/workspaces/${queryArg.workspaceSid}/members`,
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    argSchema: getWorkspaceMembersApiArgsSchema,
                    responseSchema: getWorkspaceMembersApiResponseSchema,
                },
            }),
        }),
    });

export const {
    useGetWorkspaceByIdQuery,
    useLazyGetWorkspaceByIdQuery,
    useGetWorkspacesQuery,
    useLazyGetWorkspacesQuery,
    useCreateWorkspaceMutation,
    useUpdateWorkspaceByIdMutation,
    useDeleteWorkspaceByIdMutation,
    useChangeWorkspaceVisibilityMutation,
    useGetWorkspaceMembersQuery,
} = endpoint;
