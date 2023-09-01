import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
  GetAppByIdPluginArgs,
  //GetAppsByWorkspaceIdPluginArgs,
  CreateAppPluginArgs,
  UpdateAppPluginArgs,
  DeleteAppPluginArgs,
  GetAppByPathPluginArgs,
  GetAppsByOsIdPluginArgs,
} from "@quarkloop/types";

/// GetAppById Plugin
export const GetAppByIdPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "GetAppByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[GetAppByIdPlugin]"),
      };
    }
    if (args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[GetAppByIdPlugin]"),
      };
    }

    const getArgs = args[0].app as GetAppByIdPluginArgs;

    const record = await prisma.app.findFirst({
      where: {
        id: getArgs.id,
        userRoleAssignment: {
          every: {
            os: {
              id: getArgs.osId,
            },
            ...(getArgs.workspaceId && {
              workspace: { id: getArgs.workspaceId! },
            }),
          },
        },
      },
      include: {
        pagesSettings: true,
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

/// GetAppByPath Plugin
export const GetAppByPathPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "GetAppByPathPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[GetAppByPathPlugin]"),
      };
    }
    if (args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[GetAppByPathPlugin]"),
      };
    }
    const getArgs = args[0].app as GetAppByPathPluginArgs;

    const record = await prisma.app.findFirst({
      where: {
        ...(getArgs.workspacePath && { workspacePath: getArgs.workspacePath }),
        ...(getArgs.profilePath && { profilePath: getArgs.profilePath }),
        userRoleAssignment: {
          every: {
            os: {
              id: getArgs.osId,
            },
            workspace: {
              id: getArgs.workspaceId,
            },
          },
        },
      },
      ...(getArgs.workspacePath && {
        include: {
          submissions: {
            where: {
              stage: {
                not: "Draft",
              },
            },
            select: {
              _count: true,
            },
          },
        },
      }),
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetAppByPathPlugin]"),
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

/// GetAppsByOsId Plugin
export const GetAppsByOsIdPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "GetAppsByOsIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppsByOsIdPlugin]"
        ),
      };
    }
    if (args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetAppsByOsIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].app as GetAppsByOsIdPluginArgs;

    const records = await prisma.app.findMany({
      where: {
        userRoleAssignment: {
          every: {
            os: {
              id: getArgs.osId,
            },
            ...(getArgs.workspaceId
              ? { workspace: { id: getArgs.workspaceId } }
              : { workspace: { isNot: null } }),
            app: { isNot: null },
            user: { isNot: undefined },
          },
        },
      },
      include: {
        userRoleAssignment: {
          include: {
            workspace: true,
          },
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        app: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

// /// GetAppsByWorkspaceId Plugin
// export const GetAppsByWorkspaceIdPlugin = createPlugin<
//   PipelineState,
//   PipelineArgs[]
// >({
//   name: "GetAppsByWorkspaceIdPlugin",
//   config: {},
//   handler: async (state, config, ...args): Promise<PipelineState> => {
//     if (state.status) {
//       return state;
//     }

//     if (args.length === 0) {
//       return {
//         ...state,
//         status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
//           "[GetAppsByWorkspaceIdPlugin]"
//         ),
//       };
//     }
//     if (args[0].app == null) {
//       return {
//         ...state,
//         status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
//           "[GetAppsByWorkspaceIdPlugin]"
//         ),
//       };
//     }
//     const getArgs = args[0].app as GetAppsByWorkspaceIdPluginArgs;

//     const records = await prisma.app.findMany({
//       where: {
//         userRoleAssignment: {
//           every: {
//             workspace: {
//               id: getArgs.workspaceId,
//             },
//             os: {
//               id: getArgs.osId,
//             },
//           },
//         },
//       },
//     });

//     return {
//       ...state,
//       database: {
//         ...state.database,
//         app: {
//           records: records,
//           totalRecords: records.length,
//         },
//       },
//     };
//   },
// });

/// CreateApp Plugin
export const CreateAppPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "CreateAppPlugin",
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
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[CreateAppPlugin]"),
      };
    }
    if (args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[CreateAppPlugin]"),
      };
    }

    const createArgs = args[0].app as CreateAppPluginArgs;

    const { user } = state.user;
    const { subscription } = state.user;

    const appId = generateId();

    const workspacePath = `/os/${createArgs.osId}/${createArgs.workspaceId}/${appId}`;
    const profilePath = `/os/${createArgs.osId}/services/${appId}`;

    const record = await prisma.app.create({
      data: {
        id: appId,
        name: createArgs.name!,
        visibility: createArgs.visibility!,
        type: createArgs.type!,
        status: createArgs.status!,
        workspacePath: workspacePath,
        profilePath: profilePath,
        icon: createArgs.icon,
        metadata: createArgs.metadata!,
        lastUpdate: new Date(),
        userRoleAssignment: {
          create: {
            type: "App",
            role: "Owner",
            user: { connect: { id: user?.id } },
            os: { connect: { id: createArgs.osId } },
            workspace: {
              connect: {
                osId_id: {
                  id: createArgs.workspaceId!,
                  osId: createArgs.osId,
                },
              },
            },
          },
        },
        planMetrics: {
          create: {
            type: "App",
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
          },
        },
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
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[UpdateAppPlugin]"),
      };
    }
    if (args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[UpdateAppPlugin]"),
      };
    }
    const updateArgs = args[0].app as UpdateAppPluginArgs;
    const { user } = state.user;

    const record = await prisma.app.updateMany({
      where: {
        id: updateArgs.id,
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
      data: {
        // once an app created, only name and icon can be updated.
        ...(updateArgs.name && { name: updateArgs.name }),
        ...(updateArgs.status != null && {
          status: updateArgs.status,
        }),
        ...(updateArgs.icon && { icon: updateArgs.icon }),

        ...(updateArgs.metadata && { metadata: updateArgs.metadata }),
        //...(updateArgs.pages && { pages: updateArgs.pages }),
        //...(updateArgs.forms && { forms: updateArgs.forms }),

        lastUpdate: new Date(),
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
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
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[DeleteAppPlugin]"),
      };
    }
    if (args[0].app == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[DeleteAppPlugin]"),
      };
    }

    const deleteArgs = args[0].app as DeleteAppPluginArgs;
    const { user } = state.user;

    const submissions = await prisma.appSubmission.findMany({
      where: {
        stage: {
          not: "Draft",
        },
        app: {
          id: deleteArgs.id,
        },
      },
    });

    if (submissions.length == 0) {
      const record = await prisma.userRoleAssignmentMap.deleteMany({
        where: {
          user: {
            id: user?.id,
          },
          os: {
            id: deleteArgs.osId,
          },
          workspace: {
            id: deleteArgs.workspaceId,
          },
          app: {
            id: deleteArgs.id,
          },
        },
      });

      if (record.count === 0) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[DeleteAppPlugin] app not found."
          ),
        };
      }
    } else {
      const record = await prisma.app.updateMany({
        where: {
          id: deleteArgs.id,
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
        data: {
          status: "Archived",
        },
      });

      if (record.count === 0) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[DeleteAppPlugin] fail to update app, app doesn't exist."
          ),
        };
      }

      const updatedSubmissions = await prisma.appSubmission.updateMany({
        where: {
          app: {
            id: deleteArgs.id,
          },
        },
        data: {
          stage: "Archived",
        },
      });

      if (updatedSubmissions.count === 0) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[DeleteAppPlugin] failed to update submission, submission doesn't exist."
          ),
        };
      }

      // appConversationData: {
      //   createMany: {
      //     data: {
      //       type: "Inline",
      //       message:
      //         "The associated app to this submission was archived.",
      //       userId: user.id!,
      //     },
      //   },
      // },

      const submissions = await prisma.appSubmissionUserMap.findMany({
        where: {
          appSubmission: {
            app: {
              id: deleteArgs.id,
              status: "Archived",
            },
          },
          user: {
            id: user?.id,
          },
        },
        select: {
          id: true,
        },
      });

      const conversation = await prisma.appConversationData.createMany({
        data: submissions.map((sub) => ({
          message: "The associated app to this submission was archived.",
          type: "Inline",
          lastUpdate: new Date(),
          userId: user?.id!,
          submissionUserId: sub.id,
        })),
      });

      if (conversation.count === 0) {
        return {
          ...state,
          status: PluginStatusEntry.NOT_FOUND(
            "[DeleteAppPlugin] failed to update appConversationData, appConversationData doesn't exist."
          ),
        };
      }
    }

    return state;
  },
});
