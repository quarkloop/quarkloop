import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppFileSettingsByIdPluginArgs,
  GetAppFileSettingsByAppIdPluginArgs,
  CreateAppFileSettingsPluginArgs,
  UpdateAppFileSettingsPluginArgs,
  DeleteAppFileSettingsPluginArgs,
} from "@quarkloop/types";

/// GetAppFileSettingsById Plugin
export const GetAppFileSettingsByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppFileSettingsByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appFileSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppFileSettingsByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appFileSettings as GetAppFileSettingsByIdPluginArgs;

    const record = await prisma.appFileSettings.findFirst({
      where: {
        id: getArgs.id,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppFileSettingsByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appFileSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppFileSettingsByAppId Plugin
export const GetAppFileSettingsByAppIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppFileSettingsByAppIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appFileSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppFileSettingsByAppIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appFileSettings as GetAppFileSettingsByAppIdPluginArgs;

    const record = await prisma.appFileSettings.findFirst({
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
          "[GetAppFileSettingsByAppIdPlugin]"
        ),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appFileSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// CreateAppFileSettings Plugin
export const CreateAppFileSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppFileSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appFileSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppFileSettingsPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .appFileSettings as CreateAppFileSettingsPluginArgs;

    const app = await prisma.app.findFirst({
      where: {
        id: createArgs.appId,
      },
    });

    if (app == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND(
          "[CreateAppFileSettingsPlugin] app not found"
        ),
      };
    }

    const appFileSettingsId = generateId();

    const record = await prisma.appFileSettings.create({
      data: {
        id: appFileSettingsId,
        enable: createArgs.enable!,
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
        appFileSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppFileSettings Plugin
export const UpdateAppFileSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppFileSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appFileSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppFileSettingsPlugin]"
        ),
      };
    }

    const updateArgs = args[0]
      .appFileSettings as UpdateAppFileSettingsPluginArgs;

    const record = await prisma.appFileSettings.updateMany({
      where: {
        id: updateArgs.id,
      },
      data: {
        // once an app created, only name and icon can be updated.
        // ...(updateArgs.name && { name: updateArgs.name }),
        // ...(updateArgs.status != null && { status: updateArgs.status }),
        // ...(updateArgs.icon && { icon: updateArgs.icon }),

        // ...(updateArgs.metadata && { metadata: updateArgs.metadata }),
        // ...(updateArgs.pages && { pages: updateArgs.pages }),
        // ...(updateArgs.forms && { forms: updateArgs.forms }),

        updatedAt: updateArgs.updatedAt,
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppFileSettingsPlugin]"),
      };
    }

    return state;
  },
});

/// DeleteAppFileSettings Plugin
export const DeleteAppFileSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppFileSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0 || args[0].appFileSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppFileSettingsPlugin]"
        ),
      };
    }

    const deleteArgs = args[0]
      .appFileSettings as DeleteAppFileSettingsPluginArgs;

    const record = await prisma.appFileSettings.deleteMany({
      where: {
        id: deleteArgs.id,
        appId: deleteArgs.appId,
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppFileSettingsPlugin]"),
      };
    }

    return state;
  },
});
