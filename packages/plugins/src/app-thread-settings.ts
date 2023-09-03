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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadSettingsByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appThreadSettings as GetAppThreadSettingsByIdPluginArgs;

    const record = await prisma.appThreadSettings.findFirst({
      where: {
        id: getArgs.id,
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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadSettingsByAppIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appThreadSettings as GetAppThreadSettingsByAppIdPluginArgs;

    const record = await prisma.appThreadSettings.findFirst({
      where: {
        app: {
          id: getArgs.appId,
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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppThreadSettingsPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .appThreadSettings as CreateAppThreadSettingsPluginArgs;

    const app = await prisma.app.findFirst({
      where: {
        id: createArgs.appId,
      },
    });

    if (app == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[CreateAppThreadSettingsPlugin]"),
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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadSettingsPlugin]"
        ),
      };
    }

    const updateArgs = args[0]
      .appThreadSettings as UpdateAppThreadSettingsPluginArgs;

    const record = await prisma.appThreadSettings.updateMany({
      where: {
        id: updateArgs.id,
      },
      data: {
        updatedAt: updateArgs.updatedAt,
      },
    });

    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppThreadSettingsPlugin]"),
      };
    }

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
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appThreadSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadSettingsPlugin]"
        ),
      };
    }

    const deleteArgs = args[0]
      .appThreadSettings as DeleteAppThreadSettingsPluginArgs;

    const record = await prisma.appThreadSettings.deleteMany({
      where: {
        id: deleteArgs.id,
        appId: deleteArgs.appId,
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
