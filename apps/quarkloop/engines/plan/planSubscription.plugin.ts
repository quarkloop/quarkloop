import { Prisma, prisma } from "@/prisma/client";
import { PlanType, PlanMetricsType } from "@prisma/client";

import { generateId } from "@/lib/core/core.utilities";
import { createPlugin } from "@/lib/pipeline";
import {
  PipelineArgs,
  PipelineState,
  PluginStatusEntry,
} from "@/lib/core/pipeline";

import {
  GetPlansPluginArgs,
  GetPlanSubscriptionByIdPluginArgs,
  CreatePlanSubscriptionPluginArgs,
  UpdatePlanSubscriptionPluginArgs,
  DeletePlanSubscriptionPluginArgs,
  GetPlanSubscriptionByUserSessionPluginArgs,
  UpdatePlanSubscriptionMetricsPluginArgs,
  PlanMetrics,
  Metrics,
} from "./planSubscription.type";

function countByType(metrics: PlanMetrics[]): Record<PlanMetricsType, number> {
  const countObj: Record<PlanMetricsType, number> = {
    Os: 0,
    Workspace: 0,
    App: 0,
    AppConversation: 0,
    AppPage: 0,
    AppForm: 0,
    AppFormField: 0,
    AppFile: 0,
  };

  metrics.forEach((metric) => {
    if (metric.type === "AppForm" && metric.value != null) {
      const { fieldCount } = metric.value as Prisma.JsonObject;

      if (fieldCount) {
        if (countObj["AppFormField"]) {
          countObj["AppFormField"] += fieldCount as number;
        } else {
          countObj["AppFormField"] = fieldCount as number;
        }
      }
    }

    if (countObj[metric.type!]) {
      countObj[metric.type!]++;
    } else {
      countObj[metric.type!] = 1;
    }
  });

  return countObj;
}

/// GetPlans Plugin
export const GetPlansPlugin = createPlugin<PipelineState, PipelineArgs[]>({
  name: "GetPlansPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status) {
      return state;
    }

    let planType: PlanType | undefined = undefined;

    if (args.length !== 0 && args[0].plan) {
      const getArgs = args[0].plan as GetPlansPluginArgs;
      planType = getArgs.planType;
    }

    const records = await prisma.plan.findMany({
      ...(planType && { where: { type: planType } }),
      include: {
        features: true,
        price: true,
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        plan: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// GetPlanSubscriptionByUserSession Plugin
export const GetPlanSubscriptionByUserSessionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetPlanSubscriptionByUserSessionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    const { user } = state.user;

    const record = await prisma.planSubscription.findFirst({
      where: {
        user: {
          id: user?.id,
        },
        status: "Active",
      },
      include: {
        plan: { include: { features: true } },
        billing: true,
        metrics: true,
      },
    });

    if (record == null) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND(
          "[GetPlanSubscriptionByUserSessionPlugin] User doesn't have any active subscription plan."
        ),
      };
    }

    const metrics = countByType(record.metrics);
    const usage: Partial<Metrics> = {};

    {
      const used = metrics.Os;
      const max = record.plan.features?.maxOs;
      usage.Os = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.Workspace;
      const max = record.plan.features?.maxWorkspaces;
      usage.Workspace = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.App;
      const max = record.plan.features?.maxApps;
      usage.App = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.AppConversation;
      const max = record.plan.features?.maxAppConversations;
      usage.AppConversation = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.AppPage;
      const max = record.plan.features?.maxAppPages;
      usage.AppPage = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.AppForm;
      const max = record.plan.features?.maxAppForms;
      usage.AppForm = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.AppFormField;
      const max = record.plan.features?.maxAppFormFields;

      usage.AppFormField = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }
    {
      const used = metrics.AppFile;
      const max = record.plan.features?.maxAppFiles;
      usage.AppFile = {
        used: used || 0,
        max: max || 1,
        usagePercentage: (used ?? 0) * (100 / (max ?? 1)),
      };
    }

    return {
      ...state,
      user: {
        ...state.user,
        subscription: { ...record, metrics: usage } as any,
      },
      database: {
        ...state.database,
        planSubscription: {
          records: { ...record, metrics: usage } as any,
          totalRecords: 1,
        },
      },
    };
  },
});

/// GetPlanSubscriptionById Plugin
export const GetPlanSubscriptionByIdPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "GetPlanSubscriptionByIdPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetPlanSubscriptionByIdPlugin]"
        ),
      };
    }
    if (args[0].planSubscription == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetPlanSubscriptionByIdPlugin]"
        ),
      };
    }

    const getArgs = args[0]
      .planSubscription as GetPlanSubscriptionByIdPluginArgs;

    const records = await prisma.planSubscription.findMany({
      where: {
        id: getArgs.id,
      },
      include: {
        plan: true,
        billing: true,
        metrics: true,
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        planSubscription: {
          records: records,
          totalRecords: records.length,
        },
      },
    };
  },
});

/// CreatePlanSubscription Plugin
export const CreatePlanSubscriptionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "CreatePlanSubscriptionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreatePlanSubscriptionPlugin]"
        ),
      };
    }
    if (args[0].planSubscription == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[CreatePlanSubscriptionPlugin]"
        ),
      };
    }

    const createArgs = args[0]
      .planSubscription as CreatePlanSubscriptionPluginArgs;
    const { user } = state.user;

    const planSubscriptionId = generateId();
    //const planSubscriptionMetricsId = generateId();

    const record = await prisma.planSubscription.create({
      data: {
        id: planSubscriptionId,
        startDate: createArgs.startDate!,
        endDate: createArgs.endDate,
        endReason: createArgs.endReason!,
        status: createArgs.status!,
        user: {
          connect: {
            id: user?.id,
          },
        },
        plan: {
          connect: {
            id: createArgs.planId,
          },
        },
        // metrics: {
        //   create: {
        //     id: planSubscriptionMetricsId,
        //     usedOs: 0,
        //     usedWorkspaces: 0,
        //     usedApps: 0,
        //     usedPages: 0,
        //     usedForms: 0,
        //     usedFileStorage: 0,
        //   },
        // },
      },
    });

    return {
      ...state,
      database: {
        ...state.database,
        planSubscription: {
          records: record,
          totalRecords: 1,
        },
      },
    };
  },
});

/// UpdatePlanSubscription Plugin
export const UpdatePlanSubscriptionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdatePlanSubscriptionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdatePlanSubscriptionPlugin]"
        ),
      };
    }
    if (args[0].planSubscription == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdatePlanSubscriptionPlugin]"
        ),
      };
    }
    const updateArgs = args[0]
      .planSubscription as UpdatePlanSubscriptionPluginArgs;
    const { user } = state.user;

    const record = await prisma.planSubscription.updateMany({
      where: {
        id: updateArgs.id,
        user: {
          id: user?.id,
        },
      },
      data: {
        startDate: updateArgs.startDate!,
        endDate: updateArgs.endDate,
        endReason: updateArgs.endReason!,
        status: updateArgs.status!,
        updatedAt: updateArgs.updatedAt,
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[UpdatePlanSubscriptionPlugin]"),
      };
    }

    return state;
  },
});

/// DeletePlanSubscription Plugin
export const DeletePlanSubscriptionPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "DeletePlanSubscriptionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeletePlanSubscriptionPlugin]"
        ),
      };
    }
    if (args[0].planSubscription == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[DeletePlanSubscriptionPlugin]"
        ),
      };
    }
    const deleteArgs = args[0]
      .planSubscription as DeletePlanSubscriptionPluginArgs;
    const { user } = state.user;

    const record = await prisma.planSubscription.deleteMany({
      where: {
        id: deleteArgs.id,
        user: {
          id: user?.id,
        },
      },
    });

    if (record.count === 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND("[DeletePlanSubscriptionPlugin]"),
      };
    }

    return state;
  },
});

/// UpdatePlanSubscriptionMetrics Plugin
export const UpdatePlanSubscriptionMetricsPlugin = createPlugin<
  PipelineState,
  PipelineArgs[]
>({
  name: "UpdatePlanSubscriptionMetricsPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    if (state.status || state.session == null || state.user == null) {
      return state;
    }

    if (args.length === 0) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdatePlanSubscriptionMetricsPlugin]"
        ),
      };
    }
    if (args[0].planSubscriptionMetrics == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[UpdatePlanSubscriptionMetricsPlugin]"
        ),
      };
    }

    const updateArgs = args[0]
      .planSubscriptionMetrics as UpdatePlanSubscriptionMetricsPluginArgs;
    const { user } = state.user;

    const os = updateArgs.incOs || updateArgs.decOs;
    const workspace = updateArgs.incWorkspace || updateArgs.decWorkspace;
    const app = updateArgs.incApp || updateArgs.decApp;
    const form = updateArgs.incForm || updateArgs.decForm;
    const page = updateArgs.incPage || updateArgs.decPage;
    const fileStorage = updateArgs.incFileStorage || updateArgs.decFileStorage;

    const record = await prisma.planMetrics.updateMany({
      where: {
        subscription: {
          status: "Active",
          user: {
            id: user?.id,
          },
        },
      },
      data: {
        ...(os && {
          usedOs: {
            ...(updateArgs.incOs && { increment: 1 }),
            ...(updateArgs.decOs && { decrement: 1 }),
          },
        }),
        ...(workspace && {
          usedWorkspaces: {
            ...(updateArgs.incWorkspace && { increment: 1 }),
            ...(updateArgs.decWorkspace && { decrement: 1 }),
          },
        }),
        ...(app && {
          usedApps: {
            ...(updateArgs.incApp && { increment: 1 }),
            ...(updateArgs.decApp && { decrement: 1 }),
          },
        }),
        ...(form && {
          usedForms: {
            ...(updateArgs.incForm && { increment: 1 }),
            ...(updateArgs.decForm && { decrement: 1 }),
          },
        }),
        ...(page && {
          usedPages: {
            ...(updateArgs.incPage && { increment: 1 }),
            ...(updateArgs.decPage && { decrement: 1 }),
          },
        }),
        ...(fileStorage && {
          usedFileStorage: {
            ...(updateArgs.incFileStorage && { increment: 1 }),
            ...(updateArgs.decFileStorage && { decrement: 1 }),
          },
        }),
      },
    });

    // TODO: following check may not work, investigate on alternatives
    if (record.count == 0) {
      return {
        ...state,
        status: PluginStatusEntry.NOT_FOUND(
          "[UpdatePlanSubscriptionMetricsPlugin]"
        ),
      };
    }

    return state;
  },
});
