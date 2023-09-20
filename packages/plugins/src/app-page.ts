import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
    GetAppPageByIdPluginArgs,
    GetAppPageByAppInstanceIdPluginArgs,
    CreateAppPagePluginArgs,
    UpdateAppPagePluginArgs,
    DeleteAppPagePluginArgs,
} from "@quarkloop/types";

/// GetAppPageByAppInstanceId Plugin
export const GetAppPageByAppInstanceIdPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "GetAppPageByAppInstanceIdPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appPage == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[GetAppPageByAppInstanceIdPlugin]"
                ),
            };
        }

        const getArgs = args[0].appPage as GetAppPageByAppInstanceIdPluginArgs;

        const record = await prisma.appPage.findFirst({
            where: {
                appInstance: {
                    id: getArgs.instanceId,
                },
            },
        });

        if (record == null) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND(
                    "[GetAppPageByAppInstanceIdPlugin]"
                ),
            };
        }

        return {
            ...state,
            database: {
                ...state.database,
                appPage: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// GetAppPageById Plugin
export const GetAppPageByIdPlugin = createPlugin<PipelineState, PipelineArgs[]>(
    {
        name: "GetAppPageByIdPlugin",
        config: {},
        handler: async (state, config, ...args): Promise<PipelineState> => {
            if (state.status) {
                return state;
            }

            if (args.length === 0 || args[0].appPage == null) {
                return {
                    ...state,
                    status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                        "[GetAppPageByIdPlugin]"
                    ),
                };
            }

            const getArgs = args[0].appPage as GetAppPageByIdPluginArgs;

            const record = await prisma.appPage.findUnique({
                where: {
                    id: getArgs.pageId,
                },
            });

            if (record == null) {
                return {
                    ...state,
                    status: PluginStatusEntry.NOT_FOUND(
                        "[GetAppPageByIdPlugin]"
                    ),
                };
            }

            return {
                ...state,
                database: {
                    ...state.database,
                    appPage: {
                        records: record,
                        totalRecords: 1,
                    },
                },
            };
        },
    }
);

/// CreateAppPage Plugin
export const CreateAppPagePlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "CreateAppPagePlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appPage == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[CreateAppPagePlugin]"
                ),
            };
        }

        const createArgs = args[0].appPage as CreateAppPagePluginArgs;
        const pageId = generateId();

        const record = await prisma.appPage.create({
            data: {
                id: pageId,
                name: createArgs.page.name!,
                updatedAt: new Date(),
                appInstance: {
                    connect: {
                        id: createArgs.instanceId,
                    },
                },
            },
        });

        return {
            ...state,
            database: {
                ...state.database,
                appPage: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// UpdateAppPage Plugin
export const UpdateAppPagePlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "UpdateAppPagePlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appPage == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[UpdateAppPagePlugin]"
                ),
            };
        }

        const updateArgs = args[0].appPage as UpdateAppPagePluginArgs;

        const record = await prisma.appPage.update({
            where: {
                id: updateArgs.page.id,
            },
            data: {
                name: updateArgs.page.name,
                updatedAt: new Date(),
            },
        });

        if (record == null) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[UpdateAppPagePlugin]"),
            };
        }

        return state;
    },
});

/// DeleteAppPage Plugin
export const DeleteAppPagePlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "DeleteAppPagePlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appPage == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[DeleteAppPagePlugin]"
                ),
            };
        }

        const deleteArgs = args[0].appPage as DeleteAppPagePluginArgs;

        const record = await prisma.appPage.deleteMany({
            where: {
                id: deleteArgs.pageId,
                appInstance: {
                    id: deleteArgs.instanceId,
                },
            },
        });

        if (record.count === 0) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[DeleteAppPagePlugin]"),
            };
        }

        return state;
    },
});
