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
    GetUserWorkspacesApiResponse,
    GetUserWorkspacesApiArgs,
    getUserWorkspacesSchema,
    GetOrgWorkspacesApiResponse,
    GetOrgWorkspacesApiArgs,
    getOrgWorkspacesApiResponseSchema,
    getOrgWorkspacesApiArgsSchema,
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
                    url: `/manage/${queryArg.orgSid}/workspaces/${queryArg.workspaceSid}`,
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    responseSchema: getWorkspaceByIdSchema,
                },
            }),

            // getUserWorkspaces endpoint
            getUserWorkspaces: builder.query<
                GetUserWorkspacesApiResponse,
                GetUserWorkspacesApiArgs
            >({
                query: (queryArg) => ({
                    url: "/manage/workspaces",
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    responseSchema: getUserWorkspacesSchema,
                },
            }),

            // getOrgWorkspaces endpoint
            getOrgWorkspaces: builder.query<
                GetOrgWorkspacesApiResponse,
                GetOrgWorkspacesApiArgs
            >({
                query: (queryArg) => ({
                    url: "/manage/workspaces",
                    method: "GET",
                }),
                providesTags: ["workspace"],
                extraOptions: {
                    argSchema: getOrgWorkspacesApiArgsSchema,
                    responseSchema: getOrgWorkspacesApiResponseSchema,
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
    useGetUserWorkspacesQuery,
    useLazyGetUserWorkspacesQuery,
    useGetOrgWorkspacesQuery,
    useLazyGetOrgWorkspacesQuery,
    useCreateWorkspaceMutation,
    useUpdateWorkspaceByIdMutation,
    useDeleteWorkspaceByIdMutation,
    useChangeWorkspaceVisibilityMutation,
    useGetWorkspaceMembersQuery,
} = endpoint;
