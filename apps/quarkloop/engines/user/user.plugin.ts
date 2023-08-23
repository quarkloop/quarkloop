import { PermissionRole } from "@prisma/client";
import { Prisma, prisma } from "@/prisma/client";

import { createPlugin } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";
import { PluginStatusEntry } from "@/lib/core/pipeline";
import {
  GetUserPluginArgs,
  UpdateUserPluginArgs,
  DeleteUserPluginArgs,
  GetUserLinkedAccountsPluginArgs,
  GetUserSessionsPluginArgs,
  TerminateUserSessionPluginArgs,
  GetUserPermissionsPluginArgs,
  UserPermissions,
} from "./user.type";

/// GetUser Plugin
export const GetUserPlugin = createPlugin<PipelineState, GetUserPluginArgs[]>({
  name: "GetUserPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }

    const { session } = state.session;
    if (session.user?.email == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetUserPlugin] session is null."
        ),
      };
    }

    const record = await prisma.user.findUnique({
      where: {
        email: session.user.email,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetUserPlugin] User not found"),
      };
    }

    return {
      ...state,
      user: {
        ...state.user!,
        user: record as any,
      },
      database: {
        ...state.database,
        user: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetUserPermissions Plugin
export const GetUserPermissionsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetUserPermissionsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetUserPermissionsPlugin]"
        ),
      };
    }
    if (args[0].user == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetUserPermissionsPlugin]"
        ),
      };
    }

    const getArgs = args[0].user as GetUserPermissionsPluginArgs;

    const { user } = state.user;

    let whereClause: Prisma.UserRoleAssignmentMapWhereInput = {};
    if (getArgs.appId) {
      whereClause = {
        user: { id: user?.id },
        OR: [
          { type: "OS", os: { id: getArgs.osId } },
          { type: "Workspace", workspace: { id: getArgs.workspaceId } },
          { type: "App", app: { id: getArgs.appId } },
        ],
        bot: { is: null },
      };
    } else if (getArgs.workspaceId) {
      whereClause = {
        user: { id: user?.id },
        OR: [
          { type: "OS", os: { id: getArgs.osId } },
          { type: "Workspace", workspace: { id: getArgs.workspaceId } },
        ],
        app: { is: null },
        bot: { is: null },
      };
    } else if (getArgs.osId) {
      whereClause = {
        type: "OS",
        user: { id: user?.id },
        os: { id: getArgs.osId },
        workspace: { is: null },
        app: { is: null },
        bot: { is: null },
      };
    }

    const records = await prisma.userRoleAssignmentMap.findMany({
      where: whereClause,
    });

    console.log("---------", records);

    function canCreate(role: PermissionRole) {
      return role === "Owner" || role === "Admin";
    }

    function canUpdate(role: PermissionRole) {
      return role === "Owner" || role === "Admin" || role === "Contributer";
    }

    function canDelete(role: PermissionRole) {
      return role === "Owner" || role === "Admin";
    }

    const osPermissions = records.find((perm) => perm.type === "OS");
    const wsPermissions = records.find((perm) => perm.type === "Workspace");
    const appPermissions = records.find((perm) => perm.type === "App");

    let userPermissions: UserPermissions | undefined;
    userPermissions = {
      os: {
        canRead: osPermissions ? true : false,
        canCreate: osPermissions ? canCreate(osPermissions.role) : false,
        canUpdate: osPermissions ? canUpdate(osPermissions.role) : false,
        canDelete: osPermissions ? canDelete(osPermissions.role) : false,
      },
      workspace: {
        canRead: wsPermissions ? true : false,
        canCreate: wsPermissions ? canCreate(wsPermissions.role) : false,
        canUpdate: wsPermissions ? canUpdate(wsPermissions.role) : false,
        canDelete: wsPermissions ? canDelete(wsPermissions.role) : false,
      },
      app: {
        canRead: appPermissions ? true : false,
        canCreate: appPermissions ? canCreate(appPermissions.role) : false,
        canUpdate: appPermissions ? canUpdate(appPermissions.role) : false,
        canDelete: appPermissions ? canDelete(appPermissions.role) : false,
      },
    };

    return {
      ...state,
      user: {
        ...state.user,
        permissions: userPermissions,
      },
    };
  },
});

/// UpdateUser Plugin
export const UpdateUserPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "UpdateUserPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[UpdateUserPlugin]"),
      };
    }
    if (args[0].user == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[UpdateUserPlugin]"),
      };
    }

    const updateArgs = args[0].user as UpdateUserPluginArgs;
    const { session } = state.session;

    const record = await prisma.user.update({
      where: {
        email: session.user?.email!,
      },
      data: {
        ...updateArgs,
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateUserPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        user: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// DeleteUser Plugin
export const DeleteUserPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "DeleteUserPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[DeleteUserPlugin]"),
      };
    }
    if (args[0].user == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[DeleteUserPlugin]"),
      };
    }

    const deleteArgs = args[0].user as DeleteUserPluginArgs;
    const { user } = state.user;

    if (deleteArgs.email !== user?.email) {
      return {
        ...state,
        status: PluginStatusEntry.BAD_REQUEST("[DeleteUserPlugin]"),
      };
    }

    // const appSubmissions = await prisma.appSubmission.findMany({
    //   where: {
    //     app: {

    //     }
    //   }
    // });

    const record = await prisma.user.deleteMany({
      where: {
        id: user?.id,
        email: user?.email,
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteUserPlugin]"),
      };
    }

    return state;
  },
});

/// GetUserLinkedAccounts Plugin
export const GetUserLinkedAccountsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetUserLinkedAccountsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    const { user } = state.user;
    const records = await prisma.account.findMany({
      where: {
        userId: user?.id,
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        userAccount: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// GetUserSessions Plugin
export const GetUserSessionsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetUserSessionsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    const { sessionToken } = state.session;
    const { user } = state.user;

    const records = await prisma.session.findMany({
      where: {
        userId: user?.id,
      },
    });

    const newRecords = records.map((record) => ({
      ...record,
      isCurrent: record.sessionToken === sessionToken,
    }));

    return {
      ...state,
      database: {
        ...state.database,
        userSession: {
          records: newRecords,
          totalRecords: newRecords.length,
        },
      },
    };
  },
});

/// TerminateUserSession Plugin
export const TerminateUserSessionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "TerminateUserSessionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }
    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[TerminateUserSessionPlugin]"
        ),
      };
    }
    if (args[0].user == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[TerminateUserSessionPlugin]"
        ),
      };
    }

    const terminateArgs = args[0].user as TerminateUserSessionPluginArgs;
    const { sessionToken } = state.session;
    const { user } = state.user;

    if (sessionToken === terminateArgs.sessionToken) {
      return {
        ...state,
        status: PluginStatusEntry.BAD_REQUEST("[TerminateUserSessionPlugin]"),
      };
    }

    const record = await prisma.session.deleteMany({
      where: {
        userId: user?.id,
        sessionToken: terminateArgs.sessionToken,
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[TerminateUserSessionPlugin]"),
      };
    }

    return state;
  },
});
