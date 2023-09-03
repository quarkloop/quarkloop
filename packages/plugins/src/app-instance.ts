import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppInstanceByIdPluginArgs,
  CreateAppInstancePluginArgs,
  GetAppInstancesByAppIdPluginArgs,
  UpdateAppInstancePluginArgs,
  DeleteAppInstancePluginArgs,
} from "@quarkloop/types";

/// GetAppInstanceById Plugin
export const GetAppInstanceByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppInstanceByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppInstanceByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appInstance as GetAppInstanceByIdPluginArgs;

    const record = await prisma.appInstance.findFirst({
      where: {
        id: getArgs.id,
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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppInstancesByAppIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appInstance as GetAppInstancesByAppIdPluginArgs;

    const records = await prisma.appInstance.findMany({
      where: {
        app: {
          id: getArgs.appId,
        },
      },
      include: {
        app: true,
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        appInstance: {
          records: records,
          totalRecords: records.length,
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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppInstancePlugin]"
        ),
      };
    }

    const createArgs = args[0].appInstance as CreateAppInstancePluginArgs;

    const instanceId = generateId();

    const record = await prisma.appInstance.create({
      data: {
        id: instanceId,
        name: createArgs.name!,
        app: {
          connect: {
            id: createArgs.appId,
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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppInstancePlugin]"
        ),
      };
    }

    const updateArgs = args[0].appInstance as UpdateAppInstancePluginArgs;

    const record = await prisma.appInstance.update({
      where: {
        id: updateArgs.id,
        app: {
          id: updateArgs.appId,
        },
      },
      data: {
        ...(updateArgs.name && { title: updateArgs.name }),
        updatedAt: new Date(),
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppInstancePlugin]"),
      };
    }

    return state;
  },
});

/// DeleteAppInstance Plugin
export const DeleteAppInstancePlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppInstancePlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appInstance == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppInstancePlugin]"
        ),
      };
    }

    const deleteArgs = args[0].appInstance as DeleteAppInstancePluginArgs;

    const record = await prisma.appInstance.delete({
      where: {
        id: deleteArgs.id,
        app: {
          id: deleteArgs.appId,
        },
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppInstancePlugin]"),
      };
    }

    return state;
  },
});
