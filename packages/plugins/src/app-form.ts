import { prisma } from "@quarkloop/prisma/client";

import { generateId } from "@quarkloop/lib";
import { createPlugin } from "@quarkloop/plugin";
import { PipelineArgs, PipelineState, PluginStatusEntry } from "./pipeline";

import {
    GetAppFormByIdPluginArgs,
    GetAppFormByAppInstanceIdPluginArgs,
    CreateAppFormPluginArgs,
    UpdateAppFormPluginArgs,
    DeleteAppFormPluginArgs,
} from "@quarkloop/types";

/// GetAppFormByAppInstanceId Plugin
export const GetAppFormByAppInstanceIdPlugin = createPlugin<
    PipelineState,
    PipelineArgs[]
>({
    name: "GetAppFormByAppInstanceIdPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appForm == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[GetAppFormByAppInstanceIdPlugin]"
                ),
            };
        }

        const getArgs = args[0].appForm as GetAppFormByAppInstanceIdPluginArgs;

        const record = await prisma.appForm.findFirst({
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
                    "[GetAppFormByAppInstanceIdPlugin]"
                ),
            };
        }

        return {
            ...state,
            database: {
                ...state.database,
                appForm: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// GetAppFormById Plugin
export const GetAppFormByIdPlugin = createPlugin<PipelineState, PipelineArgs[]>(
    {
        name: "GetAppFormByIdPlugin",
        config: {},
        handler: async (state, config, ...args): Promise<PipelineState> => {
            if (state.status) {
                return state;
            }

            if (args.length === 0 || args[0].appForm == null) {
                return {
                    ...state,
                    status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                        "[GetAppFormByIdPlugin]"
                    ),
                };
            }

            const getArgs = args[0].appForm as GetAppFormByIdPluginArgs;

            const record = await prisma.appForm.findUnique({
                where: {
                    id: getArgs.formId,
                },
            });

            if (record == null) {
                return {
                    ...state,
                    status: PluginStatusEntry.NOT_FOUND(
                        "[GetAppFormByIdPlugin]"
                    ),
                };
            }

            return {
                ...state,
                database: {
                    ...state.database,
                    appForm: {
                        records: record,
                        totalRecords: 1,
                    },
                },
            };
        },
    }
);

/// CreateAppForm Plugin
export const CreateAppFormPlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "CreateAppFormPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appForm == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[CreateAppFormPlugin]"
                ),
            };
        }

        const createArgs = args[0].appForm as CreateAppFormPluginArgs;
        const formId = generateId();

        const record = await prisma.appForm.create({
            data: {
                id: formId,
                name: createArgs.form.name!,
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
                appForm: {
                    records: record,
                    totalRecords: 1,
                },
            },
        };
    },
});

/// UpdateAppForm Plugin
export const UpdateAppFormPlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "UpdateAppFormPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appForm == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[UpdateAppFormPlugin]"
                ),
            };
        }

        const updateArgs = args[0].appForm as UpdateAppFormPluginArgs;

        const record = await prisma.appForm.update({
            where: {
                id: updateArgs.form.id,
            },
            data: {
                name: updateArgs.form.name,
                updatedAt: new Date(),
            },
        });

        if (record == null) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[UpdateAppFormPlugin]"),
            };
        }

        return state;
    },
});

/// DeleteAppForm Plugin
export const DeleteAppFormPlugin = createPlugin<PipelineState, PipelineArgs[]>({
    name: "DeleteAppFormPlugin",
    config: {},
    handler: async (state, config, ...args): Promise<PipelineState> => {
        if (state.status) {
            return state;
        }

        if (args.length === 0 || args[0].appForm == null) {
            return {
                ...state,
                status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
                    "[DeleteAppFormPlugin]"
                ),
            };
        }

        const deleteArgs = args[0].appForm as DeleteAppFormPluginArgs;

        const record = await prisma.appForm.deleteMany({
            where: {
                id: deleteArgs.formId,
                appInstance: {
                    id: deleteArgs.instanceId,
                },
            },
        });

        if (record.count === 0) {
            return {
                ...state,
                status: PluginStatusEntry.NOT_FOUND("[DeleteAppFormPlugin]"),
            };
        }

        return state;
    },
});
