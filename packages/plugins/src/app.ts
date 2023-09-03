import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppByIdPluginArgs,
  CreateAppPluginArgs,
  UpdateAppPluginArgs,
  DeleteAppPluginArgs,
} from "@quarkloop/types";

/// GetAppById Plugin
export const GetAppByIdPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "GetAppByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[GetAppByIdPlugin]"),
      };
    }

    const getArgs = args[0].app as GetAppByIdPluginArgs;

    const record = await prisma.app.findUnique({
      where: {
        id: getArgs.id,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        app: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// CreateApp Plugin
export const CreateAppPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "CreateAppPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[CreateAppPlugin]"),
      };
    }

    const createArgs = args[0].app as CreateAppPluginArgs;

    const appId = generateId();

    const record = await prisma.app.create({
      data: {
        id: appId,
        name: createArgs.name!,
        type: createArgs.type!,
        icon: createArgs.icon,
        metadata: createArgs.metadata!,
        updatedAt: new Date(),
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        app: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateApp Plugin
export const UpdateAppPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "UpdateAppPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[UpdateAppPlugin]"),
      };
    }

    const updateArgs = args[0].app as UpdateAppPluginArgs;

    const record = await prisma.app.update({
      where: {
        id: updateArgs.id,
      },
      data: {
        // once an app created, only name, metadata and icon can be updated.
        ...(updateArgs.name && { name: updateArgs.name }),
        ...(updateArgs.icon && { icon: updateArgs.icon }),
        ...(updateArgs.metadata && { metadata: updateArgs.metadata }),
        updatedAt: new Date(),
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppPlugin]"),
      };
    }

    return state;
  },
});

/// DeleteApp Plugin
export const DeleteAppPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "DeleteAppPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[DeleteAppPlugin]"),
      };
    }

    const deleteArgs = args[0].app as DeleteAppPluginArgs;

    const record = await prisma.app.delete({
      where: {
        id: deleteArgs.id,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppPlugin]"),
      };
    }

    return state;
  },
});
