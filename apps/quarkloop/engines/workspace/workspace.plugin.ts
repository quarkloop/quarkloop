import { prisma } from "@/prisma/client";

import { createPlugin } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";
import { PluginStatusEntry } from "@/lib/core/pipeline";
import {
  GetWorkspaceByIdPluginArgs,
  GetWorkspacesByOsIdPluginArgs,
  CreateWorkspacePluginArgs,
  UpdateWorkspacePluginArgs,
  DeleteWorkspacePluginArgs,
  //GetWorkspaceByNamePluginArgs,
} from "./workspace.type";

/// GetWorkspacesByOsId Plugin
export const GetWorkspacesByOsIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetWorkspacesByOsIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }
    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetWorkspacesByOsIdPlugin]"
        ),
      };
    }
    if (args[0].workspace == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetWorkspacesByOsIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].workspace as GetWorkspacesByOsIdPluginArgs;

    const { user } = state.user;
    const records = await prisma.workspace.findMany({
      where: {
        userRoleAssignment: {
          every: {
            user: { id: user?.id },
            os: { id: getArgs.osId },
          },
        },
        os: {
          id: getArgs.osId,
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        workspace: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// GetWorkspaceById Plugin
export const GetWorkspaceByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetWorkspaceByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }
    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetWorkspaceByIdPlugin]"
        ),
      };
    }
    if (args[0].workspace == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetWorkspaceByIdPlugin]"
        ),
      };
    }
    const getArgs = args[0].workspace as GetWorkspaceByIdPluginArgs;

    const { user } = state.user;
    const record = await prisma.workspace.findFirst({
      where: {
        id: getArgs.id,
        userRoleAssignment: {
          every: {
            user: { id: user?.id },
          },
        },
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetWorkspaceByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        workspace: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

// /// GetWorkspaceByName Plugin
// export const GetWorkspaceByNamePlugin = createPlugin<
//   PipelineState,
//   PipelineArgs[]
// >({
//   name: "GetWorkspaceByNamePlugin",
//   config: {},
//   handler: async (state, config, ...args): Promise<PipelineState> => {
//     if (state.status || state.session == null) {
//       return state;
//     }
//     if (args.length === 0) {
//       return {
//         ...state,
//         status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
//           "[GetWorkspaceByNamePlugin]"
//         ),
//       };
//     }
//     if (args[0].workspace == null) {
//       return {
//         ...state,
//         status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
//           "[GetWorkspaceByNamePlugin]"
//         ),
//       };
//     }
//     const getArgs = args[0].workspace as GetWorkspaceByNamePluginArgs;

//     const record = await prisma.workspace.findUnique({
//       where: {
//         osId_id: {
//           id: getArgs.id!,
//           osId: getArgs.osId,
//         },
//       },
//     });

//     if (record == null) {
//       return {
//         ...state,
//         status: PluginStatusEntry.NOT_FOUND("[GetWorkspaceByNamePlugin]"),
//       };
//     }

//     return {
//       ...state,
//       database: {
//         ...state.database,
//         workspace: {
//           records: record,
//           totalRecords: 1,
//         },
//       },
//     };
//   },
// });

/// CreateWorkspace Plugin
export const CreateWorkspacePlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateWorkspacePlugin",
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
          "[CreateWorkspacePlugin]"
        ),
      };
    }
    if (args[0].workspace == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateWorkspacePlugin]"
        ),
      };
    }

    const createArgs = args[0].workspace as CreateWorkspacePluginArgs;

    const { user } = state.user;
    const { subscription } = state.user;

    const record = await prisma.workspace.create({
      data: {
        id: createArgs.id!,
        name: createArgs.name!,
        path: `/os/${createArgs.osId}/${createArgs.id}`,
        // description: createArgs.description!,
        // accessType: createArgs.accessType!,
        os: {
          connect: {
            id: createArgs.osId,
          },
        },
        userRoleAssignment: {
          create: {
            type: "Workspace",
            role: "Owner",
            user: {
              connect: { id: user?.id },
            },
            os: {
              connect: { id: createArgs.osId },
            },
          },
        },
        planMetrics: {
          create: {
            type: "Workspace",
            subscription: { connect: { id: subscription.id } },
            os: { connect: { id: createArgs.osId } },
          },
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        workspace: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateWorkspace Plugin
export const UpdateWorkspacePlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateWorkspacePlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }
    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateWorkspacePlugin]"
        ),
      };
    }
    if (args[0].workspace == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateWorkspacePlugin]"
        ),
      };
    }
    const updateArgs = args[0].workspace as UpdateWorkspacePluginArgs;

    const record = await prisma.workspace.update({
      where: {
        osId_id: {
          id: updateArgs.id!,
          osId: updateArgs.osId,
        },
      },
      data: {
        name: updateArgs.name!,
        description: updateArgs.description!,
        accessType: updateArgs.accessType!,
        ...(updateArgs.name && {
          path: `/os/${updateArgs.osId}/${updateArgs.name}`,
        }),
        imageUrl: updateArgs.imageUrl,
        os: {
          connect: {
            id: updateArgs.osId,
          },
        },
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateWorkspacePlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        workspace: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// DeleteWorkspace Plugin
export const DeleteWorkspacePlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteWorkspacePlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }
    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteWorkspacePlugin]"
        ),
      };
    }
    if (args[0].workspace == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteWorkspacePlugin]"
        ),
      };
    }
    const deleteArgs = args[0].workspace as DeleteWorkspacePluginArgs;
    const { user } = state.user;

    const record = await prisma.workspace.deleteMany({
      where: {
        id: deleteArgs.id,
        osId: deleteArgs.osId,
        userRoleAssignment: {
          every: {
            user: { id: user?.id },
            os: { id: deleteArgs.osId },
          },
        },
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteWorkspacePlugin]"),
      };
    }

    return state;
  },
});

// /// WorkspacePageData Plugin
// export const WorkspacePageDataPlugin = createPlugin<
//     PipelineState,
//     PipelineArgs[]
// >({
//     name: "WorkspacePageDataPlugin",
//     config: {},
//     handler: async (state, config, ...args): Promise<PipelineState> => {
//         if (state.session == null) {
//             return state;
//         }
//         if (state.route == null || state.route.mapData == null || state.route.parseData == null) {
//             return state;
//         }
//         if (!state.route.mapData.isRouteValid) {
//             return state;
//         }

//         const currentNodeId = state.route.route?.[state.route.parseData.fullPath!].currentNodeId;
//         if (currentNodeId == null) {
//             return {
//                 ...state,
//                 status: PluginStatusEntry.NOT_FOUND("[WorkspacePageDataPlugin]"),
//             };
//         }

//         const {
//             osId,
//             slug,
//             mapData: {
//                 isRouteValid,
//             },
//             parseData: {
//                 fullPath,
//                 osBasePath,
//                 wsBasePath,
//                 wsTreeBasePath,
//                 wsAppBasePath,
//             },
//         }: RouteData = state.route;

//         const os = state.database?.os?.records as OperatingSystem | undefined;
//         const ws = state.database?.workspace?.records as Workspace | undefined;
//         const apps = state.database?.appMarket?.records as AppMarket[];

//         if (os == null || ws == null) {
//             return {
//                 ...state,
//                 status: PluginStatusEntry.INTERNAL_SERVER_ERROR("[WorkspacePageDataPlugin]"),
//             };
//         }

//         // currentLinkPathList
//         let basePath = `/os/${osId}/ws`;
//         const currentLinkPathList = slug.map((segment: string) => {
//             basePath = `${basePath}/${segment}`;
//             return {
//                 name: segment.toLowerCase(),
//                 path: basePath,
//             };
//         }).filter((segment: any) => segment.name !== "tree");

//         // createAppButtun
//         const createAppButtun: CreateAppButtonData[] = apps.map(app => {
//             const params = {
//                 workspaceId: ws.id,
//                 nodeParentId: currentNodeId,
//                 appId: app.id,
//                 basePath: wsAppBasePath,
//             };
//             const generatedParams = Object.entries(params).map(([key, value]) => (`${key}=${value}&`)).join("");
//             const generatedLink = `/os/${osId}/new/app?${generatedParams}`;

//             return {
//                 name: app.name!,
//                 category: app.category!,
//                 icon: app.icon!,
//                 path: generatedLink,
//             };
//         });

//         return {
//             ...state,
//             ui: {
//                 ...state.ui,
//                 workspace: {
//                     description: ws.description!,
//                     osName: os.name!,
//                     osId: os.nameId!,
//                     osPath: os.path!,
//                     workspaceId: ws.id!,
//                     workspaceName: ws.name!,
//                     workspacePath: ws.path!,
//                     workspaceAccessType: ws.accessType!,
//                     currentNodeId,
//                     createAppButtun,
//                     currentLinkPathList,
//                     overview: {},
//                 }
//             }
//         };
//     },
// });
