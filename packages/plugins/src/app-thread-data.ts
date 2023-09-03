import { prisma } from "@quarkloop/prisma/client";

import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppThreadByIdPluginArgs,
  GetAppThreadByAppInstanceIdPluginArgs,
  CreateAppThreadPluginArgs,
  UpdateAppThreadPluginArgs,
  DeleteAppThreadPluginArgs,
} from "@quarkloop/types";

/// GetAppThreadById Plugin
export const GetAppThreadByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppThreadByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadByIdPlugin]"
        ),
      };
    }

    if (args[0].appThread == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appThread as GetAppThreadByIdPluginArgs;

    const record = await prisma.appThread.findFirst({
      where: {
        id: getArgs.id,
        appInstance: {
          id: getArgs.appInstanceId,
        },
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppThreadByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appThread: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppThreadByAppInstanceId Plugin
export const GetAppThreadByAppInstanceIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppThreadByAppInstanceIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadByAppInstanceIdPlugin]"
        ),
      };
    }

    if (args[0].appThread == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadByAppInstanceIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appThread as GetAppThreadByAppInstanceIdPluginArgs;

    const records = await prisma.appThread.findMany({
      where: {
        appInstance: {
          id: getArgs.appInstanceId,
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        appThread: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// CreateAppThread Plugin
export const CreateAppThreadPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppThreadPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppThreadPlugin]"
        ),
      };
    }

    if (args[0].appThread == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppThreadPlugin]"
        ),
      };
    }

    const createArgs = args[0].appThread as CreateAppThreadPluginArgs;

    const record = await prisma.appThread.create({
      data: {
        type: createArgs.type!,
        message: createArgs.message!,
        updatedAt: new Date(),
        appInstance: {
          connect: {
            id: createArgs.appInstanceId,
          },
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        appThread: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppThread Plugin
export const UpdateAppThreadPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppThreadPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadPlugin]"
        ),
      };
    }

    if (args[0].appThread == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadPlugin]"
        ),
      };
    }
    const updateArgs = args[0].appThread as UpdateAppThreadPluginArgs;

    const record = await prisma.appThread.updateMany({
      where: {
        id: updateArgs.id,
      },
      data: {
        // once an app created, only name and icon can be updated.
        name: updateArgs.name,
        updatedAt: updateArgs.updatedAt,
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppThreadPlugin]"),
      };
    }

    return state;
  },
});

/// DeleteAppThread Plugin
export const DeleteAppThreadPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppThreadPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadPlugin]"
        ),
      };
    }

    if (args[0].appThread == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadPlugin]"
        ),
      };
    }
    const deleteArgs = args[0].appThread as DeleteAppThreadPluginArgs;

    const record = await prisma.appThread.deleteMany({
      where: {
        id: deleteArgs.id,
        appInstance: {
          id: deleteArgs.appInstanceId,
        },
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppThreadPlugin]"),
      };
    }

    return state;
  },
});
