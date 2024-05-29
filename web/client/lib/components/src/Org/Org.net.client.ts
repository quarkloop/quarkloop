"use client";

import { enpointApi } from "@/api/net";
import {
    GetOrgByIdApiResponse,
    GetOrgByIdApiArgs,
    GetOrgMembersApiResponse,
    GetOrgMembersApiArgs,
    CreateOrgApiResponse,
    CreateOrgApiArgs,
    UpdateOrgByIdApiResponse,
    UpdateOrgByIdApiArgs,
    DeleteOrgByIdApiResponse,
    DeleteOrgByIdApiArgs,
    createOrgApiArgsSchema,
    createOrgApiResponseSchema,
    GetOrgsApiResponse,
    GetOrgsApiArgs,
    getOrgsSchema,
    getOrgByIdSchema,
    updateOrgByIdApiArgsSchema,
    deleteOrgByIdApiArgsSchema,
    GetOrgWorkspaceListApiResponse,
    GetOrgWorkspaceListApiArgs,
    getOrgWorkspaceListSchema,
    ChangeOrgVisibilityApiResponse,
    ChangeOrgVisibilityApiArgs,
    changeOrgVisibilityApiArgsSchema,
    getOrgMembersApiArgsSchema,
    getOrgMembersApiResponseSchema,
} from "./Org.net.schema";

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["org"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getOrgById endpoint
            getOrgById: builder.query<GetOrgByIdApiResponse, GetOrgByIdApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/manage/${queryArg.orgSid}`,
                        method: "GET",
                    }),
                    providesTags: ["org"],
                    extraOptions: {
                        responseSchema: getOrgByIdSchema,
                    },
                }
            ),

            // getOrgWorkspaceList endpoint
            getOrgWorkspaceList: builder.query<
                GetOrgWorkspaceListApiResponse,
                GetOrgWorkspaceListApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/workspaces`,
                    method: "GET",
                }),
                providesTags: ["org"],
                extraOptions: {
                    responseSchema: getOrgWorkspaceListSchema,
                },
            }),

            // getOrgs endpoint
            getOrgs: builder.query<GetOrgsApiResponse, GetOrgsApiArgs>({
                query: (queryArg) => ({
                    url: "/manage/orgs",
                    method: "GET",
                }),
                providesTags: ["org"],
                extraOptions: {
                    responseSchema: getOrgsSchema,
                },
            }),

            // createOrg endpoint
            createOrg: builder.mutation<CreateOrgApiResponse, CreateOrgApiArgs>(
                {
                    query: (queryArg) => ({
                        url: `/manage/orgs`,
                        method: "POST",
                        body: queryArg,
                    }),
                    invalidatesTags: ["org"],
                    extraOptions: {
                        argSchema: createOrgApiArgsSchema,
                        responseSchema: createOrgApiResponseSchema,
                    },
                }
            ),

            // updateOrgById endpoint
            updateOrgById: builder.mutation<
                UpdateOrgByIdApiResponse,
                UpdateOrgByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["org"],
                extraOptions: {
                    argSchema: updateOrgByIdApiArgsSchema,
                },
            }),

            // deleteOrgById endpoint
            deleteOrgById: builder.mutation<
                DeleteOrgByIdApiResponse,
                DeleteOrgByIdApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}`,
                    method: "DELETE",
                }),
                invalidatesTags: ["org"],
                extraOptions: {
                    argSchema: deleteOrgByIdApiArgsSchema,
                },
            }),

            // changeOrgVisibility endpoint
            changeOrgVisibility: builder.mutation<
                ChangeOrgVisibilityApiResponse,
                ChangeOrgVisibilityApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/visibility`,
                    method: "PUT",
                    body: queryArg,
                }),
                invalidatesTags: ["org"],
                extraOptions: {
                    argSchema: changeOrgVisibilityApiArgsSchema,
                },
            }),

            // getOrgMembers endpoint
            getOrgMembers: builder.query<
                GetOrgMembersApiResponse,
                GetOrgMembersApiArgs
            >({
                query: (queryArg) => ({
                    url: `/manage/${queryArg.orgSid}/members`,
                    method: "GET",
                }),
                providesTags: ["org"],
                extraOptions: {
                    argSchema: getOrgMembersApiArgsSchema,
                    responseSchema: getOrgMembersApiResponseSchema,
                },
            }),

            // // backupOrg endpoint
            // backupOrg: builder.mutation<
            //     BackupOrgApiResponse,
            //     BackupOrgApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/backupOrg`,
            //         method: "POST",
            //     }),
            // }),

            // // restoreOrg endpoint
            // restoreOrg: builder.mutation<
            //     RestoreOrgApiResponse,
            //     RestoreOrgApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/restoreOrg`,
            //         method: "POST",
            //     }),
            // }),

            // // joinOrg endpoint
            // joinOrg: builder.mutation<
            //     JoinOrgApiResponse,
            //     JoinOrgApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/joinOrg`,
            //         method: "PUT",
            //     }),
            // }),

            // // inviteUserToOrg endpoint
            // inviteUserToOrg: builder.mutation<
            //     InviteUserToOrgApiResponse,
            //     InviteUserToOrgApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/inviteUserToOrg`,
            //         method: "PUT",
            //     }),
            // }),

            // // grantOrgPermission endpoint
            // grantOrgPermission: builder.mutation<
            //     GrantOrgPermissionApiResponse,
            //     GrantOrgPermissionApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/grantOrgPermission`,
            //         method: "PUT",
            //     }),
            // }),

            // // RevokeOrgPermission endpoint
            // RevokeOrgPermission: builder.mutation<
            //     RevokeOrgPermissionApiResponse,
            //     RevokeOrgPermissionApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/RevokeOrgPermission`,
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

            // // showOrgDetails endpoint
            // showOrgDetails: builder.query<
            //     ShowOrgDetailsApiResponse,
            //     ShowOrgDetailsApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/showOrgDetails`,
            //         method: "GET",
            //     }),
            // }),

            // // searchOrg endpoint
            // searchOrg: builder.query<
            //     SearchOrgApiResponse,
            //     SearchOrgApiArgs
            // >({
            //     query: (queryArg) => ({
            //         url: `/orgs/searchOrg`,
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
    useGetOrgByIdQuery,
    useLazyGetOrgByIdQuery,
    useGetOrgWorkspaceListQuery,
    useGetOrgsQuery,
    useLazyGetOrgsQuery,
    useCreateOrgMutation,
    useUpdateOrgByIdMutation,
    useDeleteOrgByIdMutation,
    useChangeOrgVisibilityMutation,
    useGetOrgMembersQuery,

    // useBackupOrgMutation,
    // useRestoreOrgMutation,
    // useGrantOrgPermissionMutation,
    // useInviteUserToOrgMutation,
    // useJoinOrgMutation,
    // useRevokeOrgPermissionMutation,
    // useShowDashboardQuery,
    // useShowAnalyticsQuery,
    // useShowOrgDetailsQuery,
    // useSearchOrgQuery,
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
