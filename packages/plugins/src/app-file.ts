import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
    GetAppFileByAppInstanceIdPluginArgs,
    GetAppFileByIdPluginArgs,
    CreateAppFilePluginArgs,
    UpdateAppFilePluginArgs,
    DeleteAppFilePluginArgs,
} from "@quarkloop/types";

/// GetAppFileByAppInstanceId Plugin
export const GetAppFileByAppInstanceIdPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "GetAppFileByAppInstanceIdPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appFile == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[GetAppFileByAppInstanceIdPlugin]"
                ),
            };
        }

        const getArgs = args[0].appFile as GetAppFileByAppInstanceIdPluginArgs;

        const record = await prisma.appFile.findFirst({
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
                    "[GetAppFileByAppInstanceIdPlugin]"
                ),
            };
        }

        return {
            ...state,
            database: {
                ...state.database,
                appFile: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// GetAppFileById Plugin
export const GetAppFileByIdPlugin = createPlugin<PipelineState, PipelineArgs[]>(
    {
        name: "GetAppFileByIdPlugin",
        config: {},
        handler: async (state, config, ...args): Promise<PipelineState> => {
            if (state.status) {
                return state;
            }

            if (args.length === 0 || args[0].appFile == null) {
                return {
                    ...state,
                    status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                        "[GetAppFileByIdPlugin]"
                    ),
                };
            }

            const getArgs = args[0].appFile as GetAppFileByIdPluginArgs;

            const record = await prisma.appFile.findUnique({
                where: {
                    id: getArgs.fileId,
                },
            });

            if (record == null) {
                return {
                    ...state,
                    status: PluginStatusEntry.NOT_FOUND(
                        "[GetAppFileByIdPlugin]"
                    ),
                };
            }

            return {
                ...state,
                database: {
                    ...state.database,
                    appFile: {
                        records: record,
                        totalRecords: 1,
                    },
                },
            };
        },
    }
);

/// CreateAppFile Plugin
export const CreateAppFilePlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "CreateAppFilePlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appFile == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[CreateAppFilePlugin]"
                ),
            };
        }

        const createArgs = args[0].appFile as CreateAppFilePluginArgs;
        // TODO: should get from uploader
        const fileId = generateId();

        const record = await prisma.appFile.create({
            data: {
                id: fileId,
                enable: createArgs.file.enable!,
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
                appFile: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// UpdateAppFile Plugin
export const UpdateAppFilePlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "UpdateAppFilePlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appFile == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[UpdateAppFilePlugin]"
                ),
            };
        }

        const updateArgs = args[0].appFile as UpdateAppFilePluginArgs;

        const record = await prisma.appFile.update({
            where: {
                id: updateArgs.file.id,
            },
            data: {
                enable: updateArgs.file.enable,
                updatedAt: new Date(),
            },
        });

        if (record == null) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[UpdateAppFilePlugin]"),
            };
        }

        return state;
    },
});

/// DeleteAppFile Plugin
export const DeleteAppFilePlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "DeleteAppFilePlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appFile == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[DeleteAppFilePlugin]"
                ),
            };
        }

        const deleteArgs = args[0].appFile as DeleteAppFilePluginArgs;

        const record = await prisma.appFile.deleteMany({
            where: {
                id: deleteArgs.fileId,
                appInstance: {
                    id: deleteArgs.instanceId,
                },
            },
        });

        if (record.count === 0) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[DeleteAppFilePlugin]"),
            };
        }

        return state;
    },
});
