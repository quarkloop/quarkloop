import { prisma } from "@/prisma/client";

import { generateId } from "@/lib/core/core.utilities";
import { createPlugin } from "@/lib/pipeline";
import {
  PipelineArgs,
  PipelineState,
  PluginStatusEntry,
} from "@/lib/core/pipeline";

import {
  GetAppConversationSettingsByIdPluginArgs,
  GetAppConversationSettingsByAppIdPluginArgs,
  CreateAppConversationSettingsPluginArgs,
  UpdateAppConversationSettingsPluginArgs,
  DeleteAppConversationSettingsPluginArgs,
} from "./conversation-settings.type";

/// GetAppConversationSettingsById Plugin
export const GetAppConversationSettingsByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppConversationSettingsByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationSettingsByIdPlugin]"
        ),
      };
    }
    if (args[0].appConversationSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationSettingsByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appConversationSettings as GetAppConversationSettingsByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appConversationSettings.findFirst({
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
        status: PluginStatusEntry.NOT_FOUND(
          "[GetAppConversationSettingsByIdPlugin]"
        ),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appConversationSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppConversationSettingsByAppId Plugin
export const GetAppConversationSettingsByAppIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppConversationSettingsByAppIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationSettingsByAppIdPlugin]"
        ),
      };
    }
    if (args[0].appConversationSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationSettingsByAppIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appConversationSettings as GetAppConversationSettingsByAppIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appConversationSettings.findFirst({
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
          "[GetAppConversationSettingsByAppIdPlugin]"
        ),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appConversationSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// CreateAppConversationSettings Plugin
export const CreateAppConversationSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppConversationSettingsPlugin",
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
          "[CreateAppConversationSettingsPlugin]"
        ),
      };
    }
    if (args[0].appConversationSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppConversationSettingsPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .appConversationSettings as CreateAppConversationSettingsPluginArgs;

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
          "[CreateAppConversationSettingsPlugin] app not found"
        ),
      };
    }

    const appConversationSettingsId = generateId();

    const record = await prisma.appConversationSettings.create({
      data: {
        id: appConversationSettingsId,
        app: {
          connect: {
            id: createArgs.appId,
          },
        },
        planMetrics: {
          create: {
            type: "AppConversation",
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
        appConversationSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppConversationSettings Plugin
export const UpdateAppConversationSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppConversationSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppConversationSettingsPlugin]"
        ),
      };
    }
    if (args[0].appConversationSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppConversationSettingsPlugin]"
        ),
      };
    }
    const updateArgs = args[0]
      .appConversationSettings as UpdateAppConversationSettingsPluginArgs;
    const { user } = state.user;

    // const record = await prisma.appConversationSettings.updateMany({
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

    //     lastUpdate: updateArgs.lastUpdate,
    //   },
    // });

    // // TODO: following check may not work, investigate on alternatives
    // if (record.count == 0) {
    //   return {
    //     ...state,
    //     status: PluginStatusEntry.NOT_FOUND("[UpdateAppConversationSettingsPlugin]"),
    //   };
    // }

    return state;
  },
});

/// DeleteAppConversationSettings Plugin
export const DeleteAppConversationSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppConversationSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppConversationSettingsPlugin]"
        ),
      };
    }
    if (args[0].appConversationSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppConversationSettingsPlugin]"
        ),
      };
    }
    const deleteArgs = args[0]
      .appConversationSettings as DeleteAppConversationSettingsPluginArgs;
    const { user } = state.user;

    const record = await prisma.appConversationSettings.deleteMany({
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
        status: PluginStatusEntry.NOT_FOUND(
          "[DeleteAppConversationSettingsPlugin]"
        ),
      };
    }

    return state;
  },
});
