import { prisma, Prisma } from "@/prisma/client";

import { generateId } from "@/lib/core/core.utilities";
import { createPlugin } from "@/lib/pipeline";
import {
  PipelineArgs,
  PipelineState,
  PluginStatusEntry,
} from "@/lib/core/pipeline";

import {
  GetOperatingSystemByIdPluginArgs,
  CreateOperatingSystemPluginArgs,
  UpdateOperatingSystemPluginArgs,
  DeleteOperatingSystemPluginArgs,
  GetOperatingSystemUsersPluginArgs,
  OperatingSystemUser,
} from "./os.type";

/// GetOperatingSystemApiResponse Plugin
export const GetOperatingSystemApiResponsePlugin = createPlugin<PipelineState>({
  name: "GetOperatingSystemApiResponsePlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    // Error Response
    if (state.status) {
      return {
        ...state,
        osApiReponse: {
          status: state.status,
        },
      };
    }

    // Success Response
    return {
      ...state,
      osApiReponse: {
        status: PluginStatusEntry.OK(),
        isOwner: false,
        user: state.user?.user,
        planSubscription: state.database?.planSubscription?.records as any,
        mydata: state.database?.osUser?.records as OperatingSystemUser[],
      },
    };
  },
});

/// GetOperatingSystems Plugin
export const GetOperatingSystemsPlugin = createPlugin<PipelineState>({
  name: "GetOperatingSystemsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    const records = await prisma.operatingSystem.findMany();

    return {
      ...state,
      database: {
        ...state.database,
        os: {
          records: records.map((record) => ({
            ...record,
            isOwner: false,
          })),
          totalRecords: records.length,
        },
      },
    };
  },
});

/// GetOperatingSystemById Plugin
export const GetOperatingSystemByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetOperatingSystemByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetOperatingSystemByIdPlugin]"
        ),
      };
    }
    if (args[0].os == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetOperatingSystemByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0].os as GetOperatingSystemByIdPluginArgs;

    // const record = await prisma.userRoleAssignmentMap.findFirst({
    //   where: {
    //     ...(state.user && { user: { id: state.user.user.id } }),
    //     os: { id: getArgs.id },
    //     workspace: {
    //       isNot: null,
    //     },
    //   },
    //   include: {
    //     os: true,
    //     ...(state.user && {
    //       workspace: true,
    //       app: true,
    //     }),
    //   },
    // });

    const record = await prisma.operatingSystem.findFirst({
      where: {
        id: getArgs.id,
        ...(state.user && {
          userRoleAssignment: {
            every: {
              user: {
                id: state.user.user?.id,
              },
            },
          },
        }),
      },
      include: state.user
        ? {
            workspaces: {
              include: {
                userRoleAssignment: {
                  include: {
                    app: true,
                  },
                },
              },
              where: {
                os: {
                  userRoleAssignment: {
                    every: {
                      userId: state.user.user?.id,
                    },
                  },
                },
              },
            },
          }
        : undefined,
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[GetOperatingSystemByIdPlugin]"),
      };
    }

    return {
      ...state,
      database: {
        ...state.database,
        os: {
          records: {
            ...record,
            isOwner: true, //TODO: state.user == null ? false : record.userId === state.user.user.id,
          },
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetOperatingSystemUsers Plugin
export const GetOperatingSystemUsersPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetOperatingSystemUsersPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetOperatingSystemUsersPlugin]"
        ),
      };
    }
    if (args[0].os == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetOperatingSystemUsersPlugin]"
        ),
      };
    }

    const getArgs = args[0].os as GetOperatingSystemUsersPluginArgs;

    let whereClause: Prisma.UserRoleAssignmentMapWhereInput[] = [
      {
        type: "OS",
        user: { isNot: null },
        os: { id: getArgs.id },
      },
    ];

    if (getArgs.workspaceId) {
      whereClause = [
        ...whereClause,
        {
          type: "Workspace",
          user: { isNot: null },
          os: { id: getArgs.id },
          workspace: { id: getArgs.workspaceId },
        },
      ];
    }

    const records = await prisma.userRoleAssignmentMap.findMany({
      where: {
        OR: whereClause,
      },
      select: {
        type: true,
        role: true,
        createdAt: true,
        osId: true,
        ...(getArgs.workspaceId && { workspaceId: true }),
        user: {
          select: {
            id: true,
            name: true,
            email: true,
            image: true,
          },
        },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        osUser: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// GetOperatingSystemsByUserId Plugin
export const GetOperatingSystemsByUserIdPlugin = createPlugin<PipelineState>({
  name: "GetOperatingSystemsByUserIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    const { user } = state.user;

    const records = await prisma.userRoleAssignmentMap.findMany({
      where: {
        user: { id: user?.id },
        os: { isNot: null },
        workspace: { is: null },
        app: { is: null },
      },
      select: {
        //role: true,
        user: true,
        os: true,
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        os: {
          records: records.map((record) => ({
            ...record.os,
            isOwner:
              state.user == null
                ? false
                : record.user.id === state.user.user?.id,
          })),
          totalRecords: records.length,
        },
      },
    };
  },
});

/// CreateOperatingSystem Plugin
export const CreateOperatingSystemPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreateOperatingSystemPlugin",
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
          "[CreateOperatingSystemPlugin]"
        ),
      };
    }
    if (args[0].os == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreateOperatingSystemPlugin]"
        ),
      };
    }

    const createArgs = args[0].os as CreateOperatingSystemPluginArgs;

    const { user } = state.user;
    const { subscription } = state.user;

    const id = generateId();

    const record = await prisma.operatingSystem.create({
      data: {
        id: id,
        name: createArgs.name!,
        description: createArgs.description!,
        path: `/os/${id}`,
        overview: createArgs.overview,
        websiteUrl: createArgs.websiteUrl,
        url1: createArgs.url1,
        url2: createArgs.url2,
        url3: createArgs.url3,
        url4: createArgs.url4,
        userRoleAssignment: {
          create: {
            type: "OS",
            role: "Owner",
            user: {
              connect: { id: user?.id },
            },
          },
        },
        planMetrics: {
          create: {
            type: "Os",
            subscription: { connect: { id: subscription.id } },
          },
        },
      },
      include: {
        userRoleAssignment: {
          select: { role: true, user: true },
        },
      },
    });

    const owner = record.userRoleAssignment.find(
      (u) => u.user.id === state.user?.user?.id
    );

    return {
      ...state,
      database: {
        ...state.database,
        os: {
          records: {
            ...record,
            isOwner: state.user == null ? false : owner ? true : false,
          },
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdateOperatingSystem Plugin
export const UpdateOperatingSystemPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdateOperatingSystemPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }
    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateOperatingSystemPlugin]"
        ),
      };
    }
    if (args[0].os == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdateOperatingSystemPlugin]"
        ),
      };
    }
    const updateArgs = args[0].os as UpdateOperatingSystemPluginArgs;

    const { user } = state.user;
    const record = await prisma.operatingSystem.update({
      where: {
        id: updateArgs.id,
      },
      data: {
        name: updateArgs.name,
        description: updateArgs.description,
        //path: `/os/${updateArgs.nameId}`,
        overview: updateArgs.overview,
        websiteUrl: updateArgs.websiteUrl,
        url1: updateArgs.url1,
        url2: updateArgs.url2,
        url3: updateArgs.url3,
        url4: updateArgs.url4,
      },
      include: {
        userRoleAssignment: {
          select: { role: true, user: true },
        },
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdateOperatingSystemPlugin]"),
      };
    }

    const owner = record.userRoleAssignment.find(
      (u) => u.user.id === state.user?.user?.id
    );

    return {
      ...state,
      database: {
        ...state.database,
        os: {
          records: {
            ...record,
            isOwner: state.user == null ? false : owner ? true : false,
          },
          totalRecords: 1,
        },
      },
    };
  },
});

/// DeleteOperatingSystem Plugin
export const DeleteOperatingSystemPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeleteOperatingSystemPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteOperatingSystemPlugin]"
        ),
      };
    }
    if (args[0].os == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeleteOperatingSystemPlugin]"
        ),
      };
    }

    const deleteArgs = args[0].os as DeleteOperatingSystemPluginArgs;
    const { user } = state.user;

    const record = await prisma.operatingSystem.deleteMany({
      where: {
        id: deleteArgs.id,
        userRoleAssignment: {
          every: {
            user: {
              id: user?.id,
            },
          },
        },
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeleteOperatingSystemPlugin]"),
      };
    }

    return state;
  },
});
