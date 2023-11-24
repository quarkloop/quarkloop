"use client";

import { enpointApi } from "../api";
import {
    GetSubmissionListApiResponse,
    GetSubmissionListApiArgs,
    GetSubmissionByIdApiResponse,
    GetSubmissionByIdApiArgs,
    CreateSubmissionApiResponse,
    CreateSubmissionApiArgs,
    UpdateSubmissionApiResponse,
    UpdateSubmissionApiArgs,
    DeleteSubmissionApiResponse,
    DeleteSubmissionApiArgs,
} from "./submission.type";

import {
    getSubmissionListSchema,
    getSubmissionByIdSchema,
    createSubmissionSchema,
    updateSubmissionSchema,
    deleteSubmissionSchema,
} from "./submission.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["submission"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getSubmissionById endpoint
            getSubmissionById: builder.query<
                GetSubmissionByIdApiResponse,
                GetSubmissionByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/submissions/${queryArg.submissionId}`,
                    method: "GET",
                }),
                providesTags: ["submission"],
                extraOptions: {
                    dataSchema: getSubmissionByIdSchema,
                },
            }),

            // getSubmissionList endpoint
            getSubmissionList: builder.query<
                GetSubmissionListApiResponse,
                GetSubmissionListApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/submissions`,
                    method: "GET",
                }),
                providesTags: ["submission"],
                extraOptions: {
                    dataSchema: getSubmissionListSchema,
                },
            }),

            // createSubmission endpoint
            createSubmission: builder.mutation<
                CreateSubmissionApiResponse,
                CreateSubmissionApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/submissions`,
                    method: "POST",
                    body: queryArg.submission,
                }),
                invalidatesTags: ["submission"],
                extraOptions: {
                    dataSchema: createSubmissionSchema,
                },
            }),

            // updateSubmission endpoint
            updateSubmission: builder.mutation<
                UpdateSubmissionApiResponse,
                UpdateSubmissionApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/submissions/${queryArg.submissionId}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["submission"],
                extraOptions: {
                    dataSchema: updateSubmissionSchema,
                },
            }),

            // deleteSubmission endpoint
            deleteSubmission: builder.mutation<
                DeleteSubmissionApiResponse,
                DeleteSubmissionApiArgs
            >({
                query: (queryArg) => ({
                    url: `/projects/${queryArg.projectId}/submissions/${queryArg.submissionId}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["submission"],
                extraOptions: {
                    dataSchema: deleteSubmissionSchema,
                },
            }),
        }),
    });

export const {
    useGetSubmissionByIdQuery,
    useLazyGetSubmissionByIdQuery,
    useGetSubmissionListQuery,
    useLazyGetSubmissionListQuery,

    useCreateSubmissionMutation,
    useUpdateSubmissionMutation,
    useDeleteSubmissionMutation,
} = endpoint;
