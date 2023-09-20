import { prisma } from "@quarkloop/prisma/client";

import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
    GetAppThreadByIdPluginArgs,
    GetAppThreadByAppInstanceIdPluginArgs,
    CreateAppThreadPluginArgs,
    UpdateAppThreadPluginArgs,
    DeleteAppThreadPluginArgs,
} from "@quarkloop/types";
import { generateId } from "@quarkloop/lib";

/// GetAppThreadByAppInstanceId Plugin
export const GetAppThreadByAppInstanceIdPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "GetAppThreadByAppInstanceIdPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appThread == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[GetAppThreadByAppInstanceIdPlugin]"
                ),
            };
        }

        const getArgs = args[0]
            .appThread as GetAppThreadByAppInstanceIdPluginArgs;

        const record = await prisma.appThread.findFirst({
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
                    "[GetAppThreadByAppInstanceIdPlugin]"
                ),
            };
        }

        return {
            ...state,
            database: {
                ...state.database,
                appThread: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// GetAppThreadById Plugin
export const GetAppThreadByIdPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "GetAppThreadByIdPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appThread == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[GetAppThreadByIdPlugin]"
                ),
            };
        }

        const getArgs = args[0].appThread as GetAppThreadByIdPluginArgs;

        const record = await prisma.appThread.findUnique({
            where: {
                id: getArgs.threadId,
            },
        });

        if (record == null) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[GetAppThreadByIdPlugin]"),
            };
        }

        return {
            ...state,
            database: {
                ...state.database,
                appThread: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// CreateAppThread Plugin
export const CreateAppThreadPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "CreateAppThreadPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appThread == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[CreateAppThreadPlugin]"
                ),
            };
        }

        const createArgs = args[0].appThread as CreateAppThreadPluginArgs;
        const threadId = generateId();

        const record = await prisma.appThread.create({
            data: {
                id: threadId,
                type: createArgs.thread.type!,
                message: createArgs.thread.message!,
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
                appThread: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// UpdateAppThread Plugin
export const UpdateAppThreadPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "UpdateAppThreadPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appThread == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[UpdateAppThreadPlugin]"
                ),
            };
        }

        const updateArgs = args[0].appThread as UpdateAppThreadPluginArgs;

        const record = await prisma.appThread.update({
            where: {
                id: updateArgs.thread.id,
            },
            data: {
                message: updateArgs.thread.message,
                updatedAt: new Date(),
            },
        });

        if (record == null) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[UpdateAppThreadPlugin]"),
            };
        }

        return state;
    },
});

/// DeleteAppThread Plugin
export const DeleteAppThreadPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "DeleteAppThreadPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appThread == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[DeleteAppThreadPlugin]"
                ),
            };
        }

        const deleteArgs = args[0].appThread as DeleteAppThreadPluginArgs;

        const record = await prisma.appThread.deleteMany({
            where: {
                id: deleteArgs.threadId,
                appInstance: {
                    id: deleteArgs.instanceId,
                },
            },
        });

        if (record.count === 0) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[DeleteAppThreadPlugin]"),
            };
        }

        return state;
    },
});
