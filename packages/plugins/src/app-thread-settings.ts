import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppThreadSettingsByIdPluginArgs,
  GetAppThreadSettingsByAppIdPluginArgs,
  CreateAppThreadSettingsPluginArgs,
  UpdateAppThreadSettingsPluginArgs,
  DeleteAppThreadSettingsPluginArgs,
} from "@quarkloop/types";

/// GetAppThreadSettingsById Plugin
export const GetAppThreadSettingsByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppThreadSettingsByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadSettingsByIdPlugin]"
        ),
      };
    }
    if (args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadSettingsByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appThreadSettings as GetAppThreadSettingsByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appThreadSettings.findFirst({
      where: {
        id: getArgs.id,
        app: {
          userRoleAssignment: {
            every: {
              user: {
                id: user?.id,
              },
              workspace: {
                id: getArgs.workspaceId,
              },
              os: {
                id: getArgs.osId,
              },
            },
          },
        },
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppThreadSettingsByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appThreadSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppThreadSettingsByAppId Plugin
export const GetAppThreadSettingsByAppIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppThreadSettingsByAppIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadSettingsByAppIdPlugin]"
        ),
      };
    }
    if (args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadSettingsByAppIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appThreadSettings as GetAppThreadSettingsByAppIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appThreadSettings.findFirst({
      where: {
        app: {
          id: getArgs.appId,
          userRoleAssignment: {
            every: {
              user: {
                id: user?.id,
              },
              workspace: {
                id: getArgs.workspaceId,
              },
              os: {
                id: getArgs.osId,
              },
            },
          },
        },
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND(
          "[GetAppThreadSettingsByAppIdPlugin]"
        ),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appThreadSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// CreateAppThreadSettings Plugin
export const CreateAppThreadSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppThreadSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (
      state.status ||
      state.session == null ||
      state.user == null ||
      state.user.subscription == null
    ) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppThreadSettingsPlugin]"
        ),
      };
    }
    if (args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppThreadSettingsPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .appThreadSettings as CreateAppThreadSettingsPluginArgs;

    const { user } = state.user;
    const { subscription } = state.user;

    const app = await prisma.app.findFirst({
      where: {
        id: createArgs.appId,
        userRoleAssignment: {
          every: {
            user: {
              id: user?.id,
            },
            workspace: {
              id: createArgs.workspaceId,
            },
            os: {
              id: createArgs.osId,
            },
          },
        },
      },
    });

    if (app == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND(
          "[CreateAppThreadSettingsPlugin] app not found"
        ),
      };
    }

    const appThreadSettingsId = generateId();

    const record = await prisma.appThreadSettings.create({
      data: {
        id: appThreadSettingsId,
        app: {
          connect: {
            id: createArgs.appId,
          },
        },
        planMetrics: {
          create: {
            type: "AppThread",
            subscription: { connect: { id: subscription.id } },
            os: { connect: { id: createArgs.osId } },
            workspace: {
              connect: {
                osId_id: {
                  id: createArgs.workspaceId!,
                  osId: createArgs.osId,
                },
              },
            },
            app: { connect: { id: createArgs.appId } },
          },
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        appThreadSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppThreadSettings Plugin
export const UpdateAppThreadSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppThreadSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadSettingsPlugin]"
        ),
      };
    }
    if (args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadSettingsPlugin]"
        ),
      };
    }
    const updateArgs = args[0]
      .appThreadSettings as UpdateAppThreadSettingsPluginArgs;
    const { user } = state.user;

    // const record = await prisma.appThreadSettings.updateMany({
    //   where: {
    //     id: updateArgs.id,
    //     workspace: {
    //       id: updateArgs.workspaceId,
    //       os: {
    //         id: updateArgs.osId,
    //         user: {
    //           id: user.id,
    //         },
    //       },
    //     },
    //   },
    //   data: {
    //     // once an app created, only name and icon can be updated.
    //     ...(updateArgs.name && { name: updateArgs.name }),
    //     ...(updateArgs.status != null && { status: updateArgs.status }),
    //     ...(updateArgs.icon && { icon: updateArgs.icon }),

    //     ...(updateArgs.metadata && { metadata: updateArgs.metadata }),
    //     ...(updateArgs.pages && { pages: updateArgs.pages }),
    //     ...(updateArgs.forms && { forms: updateArgs.forms }),

    //     updatedAt: updateArgs.updatedAt,
    //   },
    // });

    // // TODO: following check may not work, investigate on alternatives
    // if (record.count == 0) {
    //   return {
    //     ...state,
    //     status: PluginStatusEntry.NOT_FOUND("[UpdateAppThreadSettingsPlugin]"),
    //   };
    // }

    return state;
  },
});

/// DeleteAppThreadSettings Plugin
export const DeleteAppThreadSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppThreadSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadSettingsPlugin]"
        ),
      };
    }
    if (args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadSettingsPlugin]"
        ),
      };
    }
    const deleteArgs = args[0]
      .appThreadSettings as DeleteAppThreadSettingsPluginArgs;
    const { user } = state.user;

    const record = await prisma.appThreadSettings.deleteMany({
      where: {
        id: deleteArgs.id,
        appId: deleteArgs.appId,
        app: {
          userRoleAssignment: {
            every: {
              user: {
                id: user?.id,
              },
              workspace: {
                id: deleteArgs.workspaceId,
              },
              os: {
                id: deleteArgs.osId,
              },
            },
          },
        },
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppThreadSettingsPlugin]"),
      };
    }

    return state;
  },
});
