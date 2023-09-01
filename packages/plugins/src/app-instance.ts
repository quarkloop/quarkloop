import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAdminAppInstanceByIdPluginArgs,
  GetAppInstanceByIdPluginArgs,
  CreateAppInstancePluginArgs,
  GetAppInstancesByAppIdPluginArgs,
  UpdateAppInstancePluginArgs,
} from "@quarkloop/types";

/// GetAdminAppInstanceById Plugin
export const GetAdminAppInstanceByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAdminAppInstanceByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAdminAppInstanceByIdPlugin]"
        ),
      };
    }
    if (args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAdminAppInstanceByIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].appInstance as GetAdminAppInstanceByIdPluginArgs;

    const record = await prisma.appInstance.findFirst({
      where: {
        id: getArgs.id,
        app: {
          id: getArgs.appId,
          userRoleAssignment: {
            every: {
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
        status: PluginStatusEntry.NOT_FOUND("[GetAdminAppInstanceByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appInstance: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppInstanceById Plugin
export const GetAppInstanceByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppInstanceByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppInstanceByIdPlugin]"
        ),
      };
    }
    if (args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppInstanceByIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].appInstance as GetAppInstanceByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appInstance.findFirst({
      where: {
        id: getArgs.id,
        appInstanceUser: {
          every: {
            user: {
              id: user?.id,
            },
          },
        },
      },
      include: {
        app: true,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppInstanceByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appInstance: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppInstancesByAppId Plugin
export const GetAppInstancesByAppIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppInstancesByAppIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppInstancesByAppIdPlugin]"
        ),
      };
    }
    if (args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppInstancesByAppIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].appInstance as GetAppInstancesByAppIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appInstance.findMany({
      where: {
        app: {
          ...(getArgs.appId && { id: getArgs.appId }),
          ...(getArgs.workspaceId &&
            getArgs.osId && {
              workspace: {
                id: getArgs.workspaceId,
                os: {
                  id: getArgs.osId,
                },
              },
            }),
        },
        ...(getArgs.appId == null && {
          appInstanceUser: {
            every: {
              user: {
                id: user?.id,
              },
            },
          },
        }),
      },
      include: {
        app: true,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppInstanceByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appInstance: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// CreateAppInstance Plugin
export const CreateAppInstancePlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppInstancePlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppInstancePlugin]"
        ),
      };
    }
    if (args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppInstancePlugin]"
        ),
      };
    }

    const createArgs = args[0].appInstance as CreateAppInstancePluginArgs;
    const { user } = state.user;

    const submissionId = generateId();

    const record = await prisma.appInstance.create({
      data: {
        id: submissionId,
        title: createArgs.title!,
        stage: createArgs.stage!,
        // ...(createArgs.conversations && {
        //   conversations: createArgs.conversations,
        // }),
        // ...(createArgs.files && { files: createArgs.files }),
        // ...(createArgs.forms && { forms: createArgs.forms }),
        // ...(createArgs.payments && { payments: createArgs.payments }),
        app: {
          connect: {
            id: createArgs.appId,
          },
        },
        appInstanceUser: {
          create: {
            role: "Owner",
            user: {
              connect: {
                id: user?.id,
              },
            },
          },
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        appInstance: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppInstance Plugin
export const UpdateAppInstancePlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppInstancePlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppInstancePlugin]"
        ),
      };
    }
    if (args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppInstancePlugin]"
        ),
      };
    }
    const updateArgs = args[0].appInstance as UpdateAppInstancePluginArgs;
    const { user } = state.user;

    const record = await prisma.appInstance.updateMany({
      where: {
        id: updateArgs.id,
        app: {
          id: updateArgs.appId,
        },
        appInstanceUser: {
          every: {
            user: {
              id: user?.id,
            },
          },
        },
      },
      data: {
        ...(updateArgs.title && { title: updateArgs.title }),
        ...(updateArgs.stage && { stage: updateArgs.stage }),
        lastUpdate: new Date(),
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppInstancePlugin]"),
      };
    }

    return state;
  },
});
