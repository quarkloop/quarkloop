import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppInstanceByIdPluginArgs,
  CreateAppInstancePluginArgs,
  GetAppInstancesByAppIdPluginArgs,
  UpdateAppInstancePluginArgs,
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

    const record = await prisma.appInstance.findFirst({
      where: {
        id: getArgs.id,
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
    if (state.status) {
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

    const record = await prisma.appInstance.findMany({
      where: {
        app: {
          id: getArgs.appId,
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

    const record = await prisma.appInstance.updateMany({
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
