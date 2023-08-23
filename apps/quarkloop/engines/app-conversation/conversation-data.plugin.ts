import { prisma } from "@/prisma/client";

import { generateId } from "@/lib/core/core.utilities";
import { createPlugin } from "@/lib/pipeline";
import {
  PipelineArgs,
  PipelineState,
  PluginStatusEntry,
} from "@/lib/core/pipeline";

import {
  GetAppConversationDataByIdPluginArgs,
  GetAppConversationDataByAppSubmissionIdPluginArgs,
  CreateAppConversationDataPluginArgs,
  UpdateAppConversationDataPluginArgs,
  DeleteAppConversationDataPluginArgs,
} from "./conversation-data.type";

/// GetAppConversationDataById Plugin
export const GetAppConversationDataByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppConversationDataByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationDataByIdPlugin]"
        ),
      };
    }
    if (args[0].appConversationData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationDataByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appConversationData as GetAppConversationDataByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appConversationData.findFirst({
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
        status: PluginStatusEntry.NOT_FOUND(
          "[GetAppConversationDataByIdPlugin]"
        ),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appConversationData: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppConversationDataByAppSubmissionId Plugin
export const GetAppConversationDataByAppSubmissionIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppConversationDataByAppSubmissionIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationDataByAppSubmissionIdPlugin]"
        ),
      };
    }
    if (args[0].appConversationData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppConversationDataByAppSubmissionIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .appConversationData as GetAppConversationDataByAppSubmissionIdPluginArgs;

    const records = await prisma.appConversationData.findMany({
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
        appConversationData: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// CreateAppConversationData Plugin
export const CreateAppConversationDataPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppConversationDataPlugin",
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
          "[CreateAppConversationDataPlugin]"
        ),
      };
    }
    if (args[0].appConversationData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppConversationDataPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .appConversationData as CreateAppConversationDataPluginArgs;

    const { user } = state.user;
    const { subscription } = state.user;

    const record = await prisma.appConversationData.create({
      data: {
        type: createArgs.type!,
        message: createArgs.message!,
        lastUpdate: new Date(),
        appSubmission: {
          connect: {
            id: createArgs.appSubmissionId,
          },
        },
        // planMetrics: {
        //   create: {
        //     type: "AppConversationData",
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
        appConversationData: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppConversationData Plugin
export const UpdateAppConversationDataPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppConversationDataPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppConversationDataPlugin]"
        ),
      };
    }
    if (args[0].appConversationData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppConversationDataPlugin]"
        ),
      };
    }
    const updateArgs = args[0]
      .appConversationData as UpdateAppConversationDataPluginArgs;
    const { user } = state.user;

    // const record = await prisma.appConversationData.updateMany({
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
    //     status: PluginStatusEntry.NOT_FOUND("[UpdateAppConversationDataPlugin]"),
    //   };
    // }

    return state;
  },
});

/// DeleteAppConversationData Plugin
export const DeleteAppConversationDataPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteAppConversationDataPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppConversationDataPlugin]"
        ),
      };
    }
    if (args[0].appConversationData == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteAppConversationDataPlugin]"
        ),
      };
    }
    const deleteArgs = args[0]
      .appConversationData as DeleteAppConversationDataPluginArgs;
    const { user } = state.user;

    const record = await prisma.appConversationData.deleteMany({
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
        status: PluginStatusEntry.NOT_FOUND(
          "[DeleteAppConversationDataPlugin]"
        ),
      };
    }

    return state;
  },
});
