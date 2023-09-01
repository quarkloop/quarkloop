import { prisma } from "@quarkloop/prisma/client";

import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppThreadDataByIdPluginArgs,
  GetAppThreadDataByAppSubmissionIdPluginArgs,
  CreateAppThreadDataPluginArgs,
  UpdateAppThreadDataPluginArgs,
  DeleteAppThreadDataPluginArgs,
} from "@quarkloop/types";

/// GetAppThreadDataById Plugin
export const GetAppThreadDataByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppThreadDataByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadDataByIdPlugin]"
        ),
      };
    }
    if (args[0].appThreadData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadDataByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].appThreadData as GetAppThreadDataByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appThreadData.findFirst({
      where: {
        id: getArgs.id,
        appSubmission: {
          id: getArgs.appSubmissionId,
          appSubmissionUser: {
            every: {
              user: {
                id: user?.id,
              },
            },
          },
        },
      },
      include: {
        appSubmission: {
          include: {
            appSubmissionUser: {
              include: {
                user: true,
              },
            },
          },
        },
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppThreadDataByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appThreadData: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppThreadDataByAppSubmissionId Plugin
export const GetAppThreadDataByAppSubmissionIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppThreadDataByAppSubmissionIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadDataByAppSubmissionIdPlugin]"
        ),
      };
    }
    if (args[0].appThreadData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppThreadDataByAppSubmissionIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appThreadData as GetAppThreadDataByAppSubmissionIdPluginArgs;

    const records = await prisma.appThreadData.findMany({
      where: {
        appSubmission: {
          id: getArgs.appSubmissionId,
        },
      },
      include: {
        appSubmission: {
          include: {
            appSubmissionUser: {
              include: {
                user: true,
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
        appThreadData: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// CreateAppThreadData Plugin
export const CreateAppThreadDataPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppThreadDataPlugin",
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
          "[CreateAppThreadDataPlugin]"
        ),
      };
    }
    if (args[0].appThreadData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppThreadDataPlugin]"
        ),
      };
    }

    const createArgs = args[0].appThreadData as CreateAppThreadDataPluginArgs;

    const { user } = state.user;
    const { subscription } = state.user;

    const record = await prisma.appThreadData.create({
      data: {
        type: createArgs.type!,
        message: createArgs.message!,
        updatedAt: new Date(),
        appSubmission: {
          connect: {
            id: createArgs.appSubmissionId,
          },
        },
        // planMetrics: {
        //   create: {
        //     type: "AppThreadData",
        //     subscription: { connect: { id: subscription.id } },
        //     os: { connect: { id: createArgs.osId } },
        //     workspace: { connect: { id: createArgs.workspaceId } },
        //     app: { connect: { id: createArgs.appId } },
        //   },
        // },
      },
      include: {
        appSubmission: {
          include: {
            appSubmissionUser: {
              include: {
                user: true,
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
        appThreadData: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppThreadData Plugin
export const UpdateAppThreadDataPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppThreadDataPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadDataPlugin]"
        ),
      };
    }
    if (args[0].appThreadData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppThreadDataPlugin]"
        ),
      };
    }
    const updateArgs = args[0].appThreadData as UpdateAppThreadDataPluginArgs;
    const { user } = state.user;

    // const record = await prisma.appThreadData.updateMany({
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
    //     status: PluginStatusEntry.NOT_FOUND("[UpdateAppThreadDataPlugin]"),
    //   };
    // }

    return state;
  },
});

/// DeleteAppThreadData Plugin
export const DeleteAppThreadDataPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppThreadDataPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadDataPlugin]"
        ),
      };
    }
    if (args[0].appThreadData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppThreadDataPlugin]"
        ),
      };
    }
    const deleteArgs = args[0].appThreadData as DeleteAppThreadDataPluginArgs;
    const { user } = state.user;

    const record = await prisma.appThreadData.deleteMany({
      where: {
        id: deleteArgs.id,
        appSubmission: {
          id: deleteArgs.appSubmissionId,
        },
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteAppThreadDataPlugin]"),
      };
    }

    return state;
  },
});
