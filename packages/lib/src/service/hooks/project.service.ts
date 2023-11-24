"use client";

import { enpointApi } from "../api";
import {
    GetProjectListApiResponse,
    GetProjectListApiArgs,
    GetProjectByIdApiResponse,
    GetProjectByIdApiArgs,
    CreateProjectApiResponse,
    CreateProjectApiArgs,
    UpdateProjectApiResponse,
    UpdateProjectApiArgs,
    DeleteProjectApiResponse,
    DeleteProjectApiArgs,
} from "./project.type";

import {
    getProjectListSchema,
    getProjectByIdSchema,
    createProjectSchema,
    updateProjectSchema,
    deleteProjectSchema,
} from "./project.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["project"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getProjectById endpoint
            getProjectById: builder.query<
                GetProjectByIdApiResponse,
                GetProjectByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.id}`,
                    method: "GET",
                }),
                providesTags: ["project"],
                extraOptions: {
                    dataSchema: getProjectByIdSchema,
                },
            }),

            // getProjectList endpoint
            getProjectList: builder.query<
                GetProjectListApiResponse,
                GetProjectListApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects`,
                    method: "GET",
                    params: {
                        orgId: queryArg.orgId,
                        workspaceId: queryArg.workspaceId,
                    },
                }),
                providesTags: ["project"],
                extraOptions: {
                    dataSchema: getProjectListSchema,
                },
            }),

            // createProject endpoint
            createProject: builder.mutation<
                CreateProjectApiResponse,
                CreateProjectApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["project"],
                extraOptions: {
                    dataSchema: createProjectSchema,
                },
            }),

            // updateProject endpoint
            updateProject: builder.mutation<
                UpdateProjectApiResponse,
                UpdateProjectApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.id}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["project"],
                extraOptions: {
                    dataSchema: updateProjectSchema,
                },
            }),

            // deleteProject endpoint
            deleteProject: builder.mutation<
                DeleteProjectApiResponse,
                DeleteProjectApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.id}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["project"],
                extraOptions: {
                    dataSchema: deleteProjectSchema,
                },
            }),
        }),
    });

export const {
    useGetProjectByIdQuery,
    useLazyGetProjectByIdQuery,
    useGetProjectListQuery,
    useLazyGetProjectListQuery,

    useCreateProjectMutation,
    useUpdateProjectMutation,
    useDeleteProjectMutation,
} = endpoint;
