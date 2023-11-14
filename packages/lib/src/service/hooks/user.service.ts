"use client";

import { enpointApi } from "../api";

import {
    GetUserApiResponse,
    GetUserApiArgs,
    UpdateUserApiResponse,
    UpdateUserApiArgs,
    DeleteUserApiResponse,
    DeleteUserApiArgs,
    GetUserLinkedAccountsApiResponse,
    GetUserLinkedAccountsApiArgs,
    GetUserSessionsApiResponse,
    GetUserSessionsApiArgs,
    TerminateUserSessionApiResponse,
    TerminateUserSessionApiArgs,
    terminateUserSessionApiResponseSchema,
} from "./user.type";

import {
    getUserApiResponseSchema,
    getUserPermissionsApiResponseSchema,
    getUserSessionsApiResponseSchema,
    getServerSessionApiResponseSchema,
    getUserLinkedAccountsApiResponseSchema,
    updateUserApiResponseSchema,
    deleteUserApiResponseSchema,
} from "./user.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["user"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            /// getUser endpoint
            getUser: builder.query<GetUserApiResponse, GetUserApiArgs>({
                query: (queryArg) => ({
                    url: `/user`,
                    method: "GET",
                }),
                providesTags: ["user"],
                extraOptions: {
                    dataSchema: getUserApiResponseSchema,
                },
            }),

            /// updateUser endpoint
            updateUser: builder.mutation<
                UpdateUserApiResponse,
                UpdateUserApiArgs
            >({
                query: (queryArg) => ({
                    url: `/user`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["user"],
                extraOptions: {
                    dataSchema: updateUserApiResponseSchema,
                },
            }),

            /// deleteUser endpoint
            deleteUser: builder.mutation<
                DeleteUserApiResponse,
                DeleteUserApiArgs
            >({
                query: (queryArg) => ({
                    url: `/user`,
                    method: "DELETE",
                    body: queryArg,
                }),
                invalidatesTags: ["user"],
                extraOptions: {
                    dataSchema: deleteUserApiResponseSchema,
                },
            }),

            /// getUserLinkedAccounts endpoint
            getUserLinkedAccounts: builder.query<
                GetUserLinkedAccountsApiResponse,
                GetUserLinkedAccountsApiArgs
            >({
                query: (queryArg) => ({
                    url: `/user/linked-accounts`,
                    method: "GET",
                }),
                extraOptions: {
                    dataSchema: getUserLinkedAccountsApiResponseSchema,
                },
            }),

            /// getUserSessions endpoint
            getUserSessions: builder.query<
                GetUserSessionsApiResponse,
                GetUserSessionsApiArgs
            >({
                query: (queryArg) => ({
                    url: `/user/sessions`,
                    method: "GET",
                }),
                extraOptions: {
                    dataSchema: getUserSessionsApiResponseSchema,
                },
            }),

            /// terminateUserSession endpoint
            terminateUserSession: builder.mutation<
                TerminateUserSessionApiResponse,
                TerminateUserSessionApiArgs
            >({
                query: (queryArg) => ({
                    url: `/user/sessions`,
                    method: "DELETE",
                    body: queryArg,
                }),
                extraOptions: {
                    dataSchema: terminateUserSessionApiResponseSchema,
                },
            }),

            // /// registerUser endpoint
            // registerUser: builder.mutation<
            //     RegisterUserApiResponse,
            //     RegisterUserApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/registerUser`,
            //         method: "POST",
            //     }),
            // }),

            // /// authenticateUser endpoint
            // authenticateUser: builder.mutation<
            //     AuthenticateUserApiResponse,
            //     AuthenticateUserApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/authenticateUser`,
            //         method: "POST",
            //     }),
            // }),

            // /// changePassword endpoint
            // changePassword: builder.mutation<
            //     ChangePasswordApiResponse,
            //     ChangePasswordApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/changePassword`,
            //         method: "PUT",
            //     }),
            // }),

            // /// forgotPassword endpoint
            // forgotPassword: builder.mutation<
            //     ForgotPasswordApiResponse,
            //     ForgotPasswordApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/forgotPassword`,
            //         method: "PUT",
            //     }),
            // }),

            // /// createUserGroup endpoint
            // createUserGroup: builder.mutation<
            //     CreateUserGroupApiResponse,
            //     CreateUserGroupApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/createUserGroup`,
            //         method: "POST",
            //     }),
            // }),

            // /// getUserGroup endpoint
            // getUserGroup: builder.query<
            //     GetUserGroupApiResponse,
            //     GetUserGroupApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/getUserGroup`,
            //         method: "GET",
            //     }),
            // }),

            // /// joinUserGroup endpoint
            // joinUserGroup: builder.mutation<
            //     JoinUserGroupApiResponse,
            //     JoinUserGroupApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/joinUserGroup`,
            //         method: "PUT",
            //     }),
            // }),

            // /// leaveUserGroup endpoint
            // leaveUserGroup: builder.mutation<
            //     LeaveUserGroupApiResponse,
            //     LeaveUserGroupApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/user/leaveUserGroup`,
            //         method: "PUT",
            //     }),
            // }),
        }),
    });

export const {
    useGetUserQuery,
    useLazyGetUserQuery,
    useDeleteUserMutation,
    useUpdateUserMutation,
    useGetUserLinkedAccountsQuery,
    useGetUserSessionsQuery,
    useTerminateUserSessionMutation,

    // useRegisterUserMutation,
    // useAuthenticateUserMutation,
    // useChangePasswordMutation,
    // useForgotPasswordMutation,
    // useCreateUserGroupMutation,
    // useGetUserGroupQuery,
    // useJoinUserGroupMutation,
    // useLeaveUserGroupMutation,
} = endpoint;
