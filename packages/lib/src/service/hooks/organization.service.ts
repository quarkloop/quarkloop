"use client";

import { enpointApi } from "../api";
import {
    GetOrganizationByIdApiResponse,
    GetOrganizationByIdApiArgs,
    GetOrganizationUsersApiResponse,
    GetOrganizationUsersApiArgs,
    GetOrganizationsByUserIdApiResponse,
    GetOrganizationsByUserIdApiArgs,
    CreateOrganizationApiResponse,
    CreateOrganizationApiArgs,
    UpdateOrganizationApiResponse,
    UpdateOrganizationApiArgs,
    DeleteOrganizationApiResponse,
    DeleteOrganizationApiArgs,
    getOrganizationsByUserIdSchema,
} from "./organization.type";

import {
    getOrganizationsSchema,
    getOrganizationByIdSchema,
    createOrganizationSchema,
    updateOrganizationSchema,
    deleteOrganizationSchema,
} from "./organization.type";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["org"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getOrganizationById endpoint
            getOrganizationById: builder.query<
                GetOrganizationByIdApiResponse,
                GetOrganizationByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/orgs/${queryArg.id}`,
                    method: "GET",
                }),
                providesTags: ["org"],
                extraOptions: {
                    dataSchema: getOrganizationByIdSchema,
                },
            }),

            // getOrganizationUsers endpoint
            getOrganizationUsers: builder.query<
                GetOrganizationUsersApiResponse,
                GetOrganizationUsersApiArgs
            >({
                query: (queryArg) => ({
                    url: `/orgs/${queryArg.id}/users`,
                    method: "GET",
                    params: {
                        workspaceId: queryArg.workspaceId,
                    },
                }),
                providesTags: ["org"],
            }),

            // getOrganizationsByUserId endpoint
            getOrganizationsByUserId: builder.query<
                GetOrganizationsByUserIdApiResponse,
                GetOrganizationsByUserIdApiArgs
            >({
                query: (queryArg) => ({
                    url: "/orgs",
                    method: "GET",
                }),
                providesTags: ["org"],
                extraOptions: {
                    dataSchema: getOrganizationsByUserIdSchema,
                },
            }),

            // createOrganization endpoint
            createOrganization: builder.mutation<
                CreateOrganizationApiResponse,
                CreateOrganizationApiArgs
            >({
                query: (queryArg) => ({
                    url: `/orgs`,
                    method: "POST",
                    body: queryArg,
                }),
                invalidatesTags: ["org"],
                extraOptions: {
                    dataSchema: createOrganizationSchema,
                },
            }),

            // updateOrganization endpoint
            updateOrganization: builder.mutation<
                UpdateOrganizationApiResponse,
                UpdateOrganizationApiArgs
            >({
                query: (queryArg) => ({
                    url: `/orgs/${queryArg.id}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["org"],
                extraOptions: {
                    dataSchema: updateOrganizationSchema,
                },
            }),

            // deleteOrganization endpoint
            deleteOrganization: builder.mutation<
                DeleteOrganizationApiResponse,
                DeleteOrganizationApiArgs
            >({
                query: (queryArg) => ({
                    url: `/orgs/${queryArg.id}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["org"],
                extraOptions: {
                    dataSchema: deleteOrganizationSchema,
                },
            }),

            // // backupOrganization endpoint
            // backupOrganization: builder.mutation<
            //     BackupOrganizationApiResponse,
            //     BackupOrganizationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/backupOrganization`,
            //         method: "POST",
            //     }),
            // }),

            // // restoreOrganization endpoint
            // restoreOrganization: builder.mutation<
            //     RestoreOrganizationApiResponse,
            //     RestoreOrganizationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/restoreOrganization`,
            //         method: "POST",
            //     }),
            // }),

            // // joinOrganization endpoint
            // joinOrganization: builder.mutation<
            //     JoinOrganizationApiResponse,
            //     JoinOrganizationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/joinOrganization`,
            //         method: "PUT",
            //     }),
            // }),

            // // inviteUserToOrganization endpoint
            // inviteUserToOrganization: builder.mutation<
            //     InviteUserToOrganizationApiResponse,
            //     InviteUserToOrganizationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/inviteUserToOrganization`,
            //         method: "PUT",
            //     }),
            // }),

            // // grantOrganizationPermission endpoint
            // grantOrganizationPermission: builder.mutation<
            //     GrantOrganizationPermissionApiResponse,
            //     GrantOrganizationPermissionApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/grantOrganizationPermission`,
            //         method: "PUT",
            //     }),
            // }),

            // // RevokeOrganizationPermission endpoint
            // RevokeOrganizationPermission: builder.mutation<
            //     RevokeOrganizationPermissionApiResponse,
            //     RevokeOrganizationPermissionApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/RevokeOrganizationPermission`,
            //         method: "PUT",
            //     }),
            // }),

            // // showDashboard endpoint
            // showDashboard: builder.query<ShowDashboardApiResponse, ShowDashboardApiArgs>({
            //     query: (queryArg) => ({
            //         url: `/orgs/showDashboard`,
            //         method: "GET",
            //     }),
            // }),

            // // showAnalytics endpoint
            // showAnalytics: builder.query<ShowAnalyticsApiResponse, ShowAnalyticsApiArgs>({
            //     query: (queryArg) => ({
            //         url: `/orgs/showAnalytics`,
            //         method: "GET",
            //     }),
            // }),

            // // showOrganizationDetails endpoint
            // showOrganizationDetails: builder.query<
            //     ShowOrganizationDetailsApiResponse,
            //     ShowOrganizationDetailsApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/showOrganizationDetails`,
            //         method: "GET",
            //     }),
            // }),

            // // searchOrganization endpoint
            // searchOrganization: builder.query<
            //     SearchOrganizationApiResponse,
            //     SearchOrganizationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/searchOrganization`,
            //         method: "GET",
            //     }),
            // }),

            // // sendNotification endpoint
            // sendNotification: builder.mutation<
            //     SendNotificationApiResponse,
            //     SendNotificationApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/sendNotification`,
            //         method: "POST",
            //     }),
            // }),

            // // integrateExternalService endpoint
            // integrateExternalService: builder.mutation<
            //     IntegrateExternalServiceApiResponse,
            //     IntegrateExternalServiceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/integrateExternalService`,
            //         method: "POST",
            //     }),
            // }),

            // // moderateContent endpoint
            // moderateContent: builder.mutation<
            //     ModerateContentApiResponse,
            //     ModerateContentApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/moderateContent`,
            //         method: "POST",
            //     }),
            // }),

            // // logActivity endpoint
            // logActivity: builder.mutation<
            //     LogActivityApiResponse,
            //     LogActivityApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/logActivity`,
            //         method: "POST",
            //     }),
            // }),

            // // defineRole endpoint
            // defineRole: builder.mutation<
            //     DefineRoleApiResponse,
            //     DefineRoleApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/defineRole`,
            //         method: "POST",
            //     }),
            // }),

            // // assignUserRole endpoint
            // assignUserRole: builder.mutation<
            //     AssignUserRoleApiResponse,
            //     AssignUserRoleApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/assignUserRole`,
            //         method: "PUT",
            //     }),
            // }),

            // // specifyAccessControl endpoint
            // specifyAccessControl: builder.mutation<
            //     SpecifyAccessControlApiResponse,
            //     SpecifyAccessControlApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/specifyAccessControl`,
            //         method: "PUT",
            //     }),
            // }),

            // // getWorkspaces endpoint
            // getWorkspaces: builder.query<GetWorkspacesApiResponse, GetWorkspacesApiArgs>({
            //     query: (queryArg) => ({
            //         url: `/orgs/getWorkspaces`,
            //         method: "GET",
            //     }),
            // }),

            // // createWorkspace endpoint
            // createWorkspace: builder.mutation<
            //     CreateWorkspaceApiResponse,
            //     CreateWorkspaceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/createWorkspace`,
            //         method: "POST",
            //     }),
            // }),

            // // updateWorkspace endpoint
            // updateWorkspace: builder.mutation<
            //     UpdateWorkspaceApiResponse,
            //     UpdateWorkspaceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/updateWorkspace`,
            //         method: "PUT",
            //     }),
            // }),

            // // deleteWorkspace endpoint
            // deleteWorkspace: builder.mutation<
            //     DeleteWorkspaceApiResponse,
            //     DeleteWorkspaceApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/deleteWorkspace`,
            //         method: "DELETE",
            //     }),
            // }),
        }),
    });

export const {
    useGetOrganizationByIdQuery,
    useLazyGetOrganizationByIdQuery,
    useGetOrganizationUsersQuery,
    useGetOrganizationsByUserIdQuery,
    useLazyGetOrganizationsByUserIdQuery,
    useCreateOrganizationMutation,
    useUpdateOrganizationMutation,
    useDeleteOrganizationMutation,

    // useBackupOrganizationMutation,
    // useRestoreOrganizationMutation,
    // useGrantOrganizationPermissionMutation,
    // useInviteUserToOrganizationMutation,
    // useJoinOrganizationMutation,
    // useRevokeOrganizationPermissionMutation,
    // useShowDashboardQuery,
    // useShowAnalyticsQuery,
    // useShowOrganizationDetailsQuery,
    // useSearchOrganizationQuery,
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
