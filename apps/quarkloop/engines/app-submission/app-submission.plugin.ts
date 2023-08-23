import { prisma } from "@/prisma/client";

import { generateId } from "@/lib/core/core.utilities";
import { createPlugin } from "@/lib/pipeline";
import {
  PipelineArgs,
  PipelineState,
  PluginStatusEntry,
} from "@/lib/core/pipeline";

import {
  GetAdminAppSubmissionByIdPluginArgs,
  GetAppSubmissionByIdPluginArgs,
  CreateAppSubmissionPluginArgs,
  GetAppSubmissionsByAppIdPluginArgs,
  UpdateAppSubmissionPluginArgs,
} from "./app-submission.type";

/// GetAdminAppSubmissionById Plugin
export const GetAdminAppSubmissionByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAdminAppSubmissionByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAdminAppSubmissionByIdPlugin]"
        ),
      };
    }
    if (args[0].appSubmission == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAdminAppSubmissionByIdPlugin]"
        ),
      };
    }
    const getArgs = args[0]
      .appSubmission as GetAdminAppSubmissionByIdPluginArgs;

    const record = await prisma.appSubmission.findFirst({
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
        status: PluginStatusEntry.NOT_FOUND(
          "[GetAdminAppSubmissionByIdPlugin]"
        ),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appSubmission: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppSubmissionById Plugin
export const GetAppSubmissionByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppSubmissionByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppSubmissionByIdPlugin]"
        ),
      };
    }
    if (args[0].appSubmission == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppSubmissionByIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].appSubmission as GetAppSubmissionByIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appSubmission.findFirst({
      where: {
        id: getArgs.id,
        appSubmissionUser: {
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
        status: PluginStatusEntry.NOT_FOUND("[GetAppSubmissionByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appSubmission: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetAppSubmissionsByAppId Plugin
export const GetAppSubmissionsByAppIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetAppSubmissionsByAppIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppSubmissionsByAppIdPlugin]"
        ),
      };
    }
    if (args[0].appSubmission == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppSubmissionsByAppIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].appSubmission as GetAppSubmissionsByAppIdPluginArgs;
    const { user } = state.user;

    const record = await prisma.appSubmission.findMany({
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
          appSubmissionUser: {
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
        status: PluginStatusEntry.NOT_FOUND("[GetAppSubmissionByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        appSubmission: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// CreateAppSubmission Plugin
export const CreateAppSubmissionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateAppSubmissionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppSubmissionPlugin]"
        ),
      };
    }
    if (args[0].appSubmission == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateAppSubmissionPlugin]"
        ),
      };
    }

    const createArgs = args[0].appSubmission as CreateAppSubmissionPluginArgs;
    const { user } = state.user;

    const submissionId = generateId();

    const record = await prisma.appSubmission.create({
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
        appSubmissionUser: {
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
        appSubmission: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateAppSubmission Plugin
export const UpdateAppSubmissionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateAppSubmissionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppSubmissionPlugin]"
        ),
      };
    }
    if (args[0].appSubmission == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateAppSubmissionPlugin]"
        ),
      };
    }
    const updateArgs = args[0].appSubmission as UpdateAppSubmissionPluginArgs;
    const { user } = state.user;

    const record = await prisma.appSubmission.updateMany({
      where: {
        id: updateArgs.id,
        app: {
          id: updateArgs.appId,
        },
        appSubmissionUser: {
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
        status: PluginStatusEntry.NOT_FOUND("[UpdateAppSubmissionPlugin]"),
      };
    }

    return state;
  },
});
