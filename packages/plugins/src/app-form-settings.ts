import { Prisma, prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppFormSettingsByIdPluginArgs,
  GetAppFormsSettingsByAppIdPluginArgs,
  CreateAppFormSettingsPluginArgs,
  UpdateAppFormSettingsPluginArgs,
  DeleteAppFormSettingsPluginArgs,
  AppFormField,
} from "@quarkloop/types";

/// GetAppFormSettingsById Plugin
export const GetAppFormSettingsByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppFormSettingsByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppFormSettingsByIdPlugin]"
        ),
      };
    }
    if (args[0].appFormSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppFormSettingsByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appFormSettings as GetAppFormSettingsByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appFormSettings.findFirst({
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
        status: PluginStatusEntry.NOT_FOUND("[GetAppFormSettingsByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appFormSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppFormsSettingsByAppId Plugin
export const GetAppFormsSettingsByAppIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppFormsSettingsByAppIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppFormsSettingsByAppIdPlugin]"
        ),
      };
    }
    if (args[0].appFormSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppFormsSettingsByAppIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appFormSettings as GetAppFormsSettingsByAppIdPluginArgs;
    const { user } = state.user;

    const records = await prisma.appFormSettings.findMany({
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

    return {
      ...state,
      database: {
        ...state.database,
        appFormSettings: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// CreateAppFormSettings Plugin
export const CreateAppFormSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppFormSettingsPlugin",
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
          "[CreateAppFormSettingsPlugin]"
        ),
      };
    }
    if (args[0].appFormSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppFormSettingsPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .appFormSettings as CreateAppFormSettingsPluginArgs;

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
          "[CreateAppFormSettingsPlugin] app not found"
        ),
      };
    }

    const appFormSettingsId = generateId();

    const record = await prisma.appFormSettings.create({
      data: {
        id: appFormSettingsId,
        name: createArgs.name!,
        app: {
          connect: {
            id: createArgs.appId,
          },
        },
        planMetrics: {
          create: {
            type: "AppForm",
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
        appFormSettings: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppFormSettings Plugin
export const UpdateAppFormSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppFormSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppFormSettingsPlugin]"
        ),
      };
    }
    if (args[0].appFormSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppFormSettingsPlugin]"
        ),
      };
    }
    const updateArgs = args[0]
      .appFormSettings as UpdateAppFormSettingsPluginArgs;
    const { user } = state.user;

    let newFields: Prisma.JsonValue[] = [];

    if (updateArgs.formFieldCreate && updateArgs.fields) {
      const fields = await prisma.appFormSettings.findUnique({
        where: {
          id: updateArgs.id,
        },
        select: {
          fields: true,
        },
      });

      if (fields == null) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[UpdateAppFormSettingsPlugin] Cannot find appFormSettings"
          ),
        };
      }

      newFields = [
        ...fields.fields,
        ...updateArgs.fields.map((field: any) => ({
          ...field,
          id: generateId(),
        })),
      ];
    } else if (updateArgs.formFieldDelete && updateArgs.fields) {
      const fields = await prisma.appFormSettings.findUnique({
        where: {
          id: updateArgs.id,
        },
        select: {
          fields: true,
        },
      });

      if (fields == null) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[UpdateAppFormSettingsPlugin] Cannot find appFormSettings"
          ),
        };
      }

      const fieldForDelete: AppFormField | null = updateArgs.fields.find(
        (field) => field
      ) as any;

      newFields = fields.fields.filter((jsonField) => {
        const field: AppFormField = jsonField as any;
        return field.id !== fieldForDelete?.id;
      });
    }

    const record = await prisma.appFormSettings.updateMany({
      where: {
        id: updateArgs.id,
        app: {
          id: updateArgs.appId,
          userRoleAssignment: {
            every: {
              user: {
                id: user?.id,
              },
              workspace: {
                id: updateArgs.workspaceId,
              },
              os: {
                id: updateArgs.osId,
              },
            },
          },
        },
      },
      data: {
        ...(updateArgs.name && { name: updateArgs.name }),
        fields: [], // TODO
        //...(updateArgs.fields && { fields: newFields }),
        ...(updateArgs.fieldCount && { fieldCount: newFields.length }),
        updatedAt: new Date(),
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppFormSettingsPlugin]"),
      };
    }

    if (updateArgs.formFieldCreate || updateArgs.formFieldDelete) {
      const metrics = await prisma.planMetrics.updateMany({
        where: {
          appFormSettingsId: updateArgs.id,
        },
        data: {
          value: {
            fieldCount: newFields.length,
          },
        },
      });

      if (metrics == null) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[UpdateAppFormSettingsPlugin] Failed to update plan metrics."
          ),
        };
      }
    }

    return state;
  },
});

/// DeleteAppFormSettings Plugin
export const DeleteAppFormSettingsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppFormSettingsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppFormSettingsPlugin]"
        ),
      };
    }
    if (args[0].appFormSettings == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppFormSettingsPlugin]"
        ),
      };
    }
    const deleteArgs = args[0]
      .appFormSettings as DeleteAppFormSettingsPluginArgs;
    const { user } = state.user;

    const record = await prisma.appFormSettings.deleteMany({
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
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppFormSettingsPlugin]"),
      };
    }

    return state;
  },
});
