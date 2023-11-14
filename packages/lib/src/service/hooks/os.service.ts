"use client";

import { enpointApi } from "../api";
import {
    GetOperatingSystemByIdApiResponse,
    GetOperatingSystemByIdApiArgs,
    GetOperatingSystemUsersApiResponse,
    GetOperatingSystemUsersApiArgs,
    GetOperatingSystemsByUserIdApiResponse,
    GetOperatingSystemsByUserIdApiArgs,
    CreateOperatingSystemApiResponse,
    CreateOperatingSystemApiArgs,
    UpdateOperatingSystemApiResponse,
    UpdateOperatingSystemApiArgs,
    DeleteOperatingSystemApiResponse,
    DeleteOperatingSystemApiArgs,
    getOperatingSystemsByUserIdSchema,
} from "./os.type";

import {
    getOperatingSystemsSchema,
    getOperatingSystemByIdSchema,
    createOperatingSystemSchema,
    updateOperatingSystemSchema,
    deleteOperatingSystemSchema,
} from "./os.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["os"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getOperatingSystemById endpoint
            getOperatingSystemById: builder.query<
                GetOperatingSystemByIdApiResponse,
                GetOperatingSystemByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os/${queryArg.id}`,
                    method: "GET",
                }),
                providesTags: ["os"],
                extraOptions: {
                    dataSchema: getOperatingSystemByIdSchema,
                },
            }),

            // getOperatingSystemUsers endpoint
            getOperatingSystemUsers: builder.query<
                GetOperatingSystemUsersApiResponse,
                GetOperatingSystemUsersApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os/${queryArg.id}/users`,
                    method: "GET",
                    params: {
                        workspaceId: queryArg.workspaceId,
                    },
                }),
                providesTags: ["os"],
            }),

            // getOperatingSystemsByUserId endpoint
            getOperatingSystemsByUserId: builder.query<
                GetOperatingSystemsByUserIdApiResponse,
                GetOperatingSystemsByUserIdApiArgs
            >({
                query: (queryArg) => ({
                    url: "/os",
                    method: "GET",
                }),
                providesTags: ["os"],
                extraOptions: {
                    dataSchema: getOperatingSystemsByUserIdSchema,
                },
            }),

            // createOperatingSystem endpoint
            createOperatingSystem: builder.mutation<
                CreateOperatingSystemApiResponse,
                CreateOperatingSystemApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["os"],
                extraOptions: {
                    dataSchema: createOperatingSystemSchema,
                },
            }),

            // updateOperatingSystem endpoint
            updateOperatingSystem: builder.mutation<
                UpdateOperatingSystemApiResponse,
                UpdateOperatingSystemApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os/${queryArg.id}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["os"],
                extraOptions: {
                    dataSchema: updateOperatingSystemSchema,
                },
            }),

            // deleteOperatingSystem endpoint
            deleteOperatingSystem: builder.mutation<
                DeleteOperatingSystemApiResponse,
                DeleteOperatingSystemApiArgs
            >({
                query: (queryArg) => ({
                    url: `/os/${queryArg.id}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["os"],
                extraOptions: {
                    dataSchema: deleteOperatingSystemSchema,
                },
            }),

            // // backupOperatingSystem endpoint
            // backupOperatingSystem: builder.mutation<
            //     BackupOperatingSystemApiResponse,
            //     BackupOperatingSystemApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/backupOperatingSystem`,
            //         method: "POST",
            //     }),
            // }),

            // // restoreOperatingSystem endpoint
            // restoreOperatingSystem: builder.mutation<
            //     RestoreOperatingSystemApiResponse,
            //     RestoreOperatingSystemApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/restoreOperatingSystem`,
            //         method: "POST",
            //     }),
            // }),

            // // joinOperatingSystem endpoint
            // joinOperatingSystem: builder.mutation<
            //     JoinOperatingSystemApiResponse,
            //     JoinOperatingSystemApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/joinOperatingSystem`,
            //         method: "PUT",
            //     }),
            // }),

            // // inviteUserToOperatingSystem endpoint
            // inviteUserToOperatingSystem: builder.mutation<
            //     InviteUserToOperatingSystemApiResponse,
            //     InviteUserToOperatingSystemApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/inviteUserToOperatingSystem`,
            //         method: "PUT",
            //     }),
            // }),

            // // grantOperatingSystemPermission endpoint
            // grantOperatingSystemPermission: builder.mutation<
            //     GrantOperatingSystemPermissionApiResponse,
            //     GrantOperatingSystemPermissionApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/grantOperatingSystemPermission`,
            //         method: "PUT",
            //     }),
            // }),

            // // RevokeOperatingSystemPermission endpoint
            // RevokeOperatingSystemPermission: builder.mutation<
            //     RevokeOperatingSystemPermissionApiResponse,
            //     RevokeOperatingSystemPermissionApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/RevokeOperatingSystemPermission`,
            //         method: "PUT",
            //     }),
            // }),

            // // showDashboard endpoint
            // showDashboard: builder.query<ShowDashboardApiResponse, ShowDashboardApiArgs>({
            //     query: (queryArg) => ({
            //         url: `/os/showDashboard`,
            //         method: "GET",
            //     }),
            // }),

            // // showAnalytics endpoint
            // showAnalytics: builder.query<ShowAnalyticsApiResponse, ShowAnalyticsApiArgs>({
            //     query: (queryArg) => ({
            //         url: `/os/showAnalytics`,
            //         method: "GET",
            //     }),
            // }),

            // // showOperatingSystemDetails endpoint
            // showOperatingSystemDetails: builder.query<
            //     ShowOperatingSystemDetailsApiResponse,
            //     ShowOperatingSystemDetailsApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/showOperatingSystemDetails`,
            //         method: "GET",
            //     }),
            // }),

            // // searchOperatingSystem endpoint
            // searchOperatingSystem: builder.query<
            //     SearchOperatingSystemApiResponse,
            //     SearchOperatingSystemApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/searchOperatingSystem`,
            //         method: "GET",
            //     }),
            // }),

            // // sendNotification endpoint
            // sendNotification: builder.mutation<
            //     SendNotificationApiResponse,
            //     SendNotificationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/sendNotification`,
            //         method: "POST",
            //     }),
            // }),

            // // integrateExternalService endpoint
            // integrateExternalService: builder.mutation<
            //     IntegrateExternalServiceApiResponse,
            //     IntegrateExternalServiceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/integrateExternalService`,
            //         method: "POST",
            //     }),
            // }),

            // // moderateContent endpoint
            // moderateContent: builder.mutation<
            //     ModerateContentApiResponse,
            //     ModerateContentApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/moderateContent`,
            //         method: "POST",
            //     }),
            // }),

            // // logActivity endpoint
            // logActivity: builder.mutation<
            //     LogActivityApiResponse,
            //     LogActivityApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/logActivity`,
            //         method: "POST",
            //     }),
            // }),

            // // defineRole endpoint
            // defineRole: builder.mutation<
            //     DefineRoleApiResponse,
            //     DefineRoleApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/defineRole`,
            //         method: "POST",
            //     }),
            // }),

            // // assignUserRole endpoint
            // assignUserRole: builder.mutation<
            //     AssignUserRoleApiResponse,
            //     AssignUserRoleApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/assignUserRole`,
            //         method: "PUT",
            //     }),
            // }),

            // // specifyAccessControl endpoint
            // specifyAccessControl: builder.mutation<
            //     SpecifyAccessControlApiResponse,
            //     SpecifyAccessControlApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/specifyAccessControl`,
            //         method: "PUT",
            //     }),
            // }),

            // // getWorkspaces endpoint
            // getWorkspaces: builder.query<GetWorkspacesApiResponse, GetWorkspacesApiArgs>({
            //     query: (queryArg) => ({
            //         url: `/os/getWorkspaces`,
            //         method: "GET",
            //     }),
            // }),

            // // createWorkspace endpoint
            // createWorkspace: builder.mutation<
            //     CreateWorkspaceApiResponse,
            //     CreateWorkspaceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/createWorkspace`,
            //         method: "POST",
            //     }),
            // }),

            // // updateWorkspace endpoint
            // updateWorkspace: builder.mutation<
            //     UpdateWorkspaceApiResponse,
            //     UpdateWorkspaceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/updateWorkspace`,
            //         method: "PUT",
            //     }),
            // }),

            // // deleteWorkspace endpoint
            // deleteWorkspace: builder.mutation<
            //     DeleteWorkspaceApiResponse,
            //     DeleteWorkspaceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/os/deleteWorkspace`,
            //         method: "DELETE",
            //     }),
            // }),
        }),
    });

export const {
    useGetOperatingSystemByIdQuery,
    useLazyGetOperatingSystemByIdQuery,
    useGetOperatingSystemUsersQuery,
    useGetOperatingSystemsByUserIdQuery,
    useLazyGetOperatingSystemsByUserIdQuery,
    useCreateOperatingSystemMutation,
    useUpdateOperatingSystemMutation,
    useDeleteOperatingSystemMutation,

    // useBackupOperatingSystemMutation,
    // useRestoreOperatingSystemMutation,
    // useGrantOperatingSystemPermissionMutation,
    // useInviteUserToOperatingSystemMutation,
    // useJoinOperatingSystemMutation,
    // useRevokeOperatingSystemPermissionMutation,
    // useShowDashboardQuery,
    // useShowAnalyticsQuery,
    // useShowOperatingSystemDetailsQuery,
    // useSearchOperatingSystemQuery,
    // useSendNotificationMutation,
    // useIntegrateExternalServiceMutation,
    // useModerateContentMutation,
    // useLogActivityMutation,
    // useDefineRoleMutation,
    // useAssignUserRoleMutation,
    // useSpecifyAccessControlMutation,
    // useGetWorkspacesQuery,
    // useCreateWorkspaceMutation,
    // useUpdateWorkspaceMutation,
    // useDeleteWorkspaceMutation,
} = endpoint;
