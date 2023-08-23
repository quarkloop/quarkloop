import {
  Plan as PrismaPlan,
  PlanFeatures as PrismaPlanFeatures,
  PlanPricing as PrismaPlanPricing,
  PlanBilling as PrismaPlanBilling,
  PlanSubscription as PrismaPlanSubscription,
  PlanMetrics as PrismaPlanMetrics,
  PlanType,
  PlanMetricsType,
} from "@prisma/client";
import { ApiResponse } from "@/lib/core/pipeline";

/// Types
export interface Plan extends Partial<PrismaPlan> {}
export interface PlanFeatures extends Partial<PrismaPlanFeatures> {}
export interface PlanPricing extends Partial<PrismaPlanPricing> {}
export interface PlanBilling extends Partial<PrismaPlanBilling> {}
export interface PlanSubscription extends Partial<PrismaPlanSubscription> {}
export interface PlanMetrics extends Partial<PrismaPlanMetrics> {}

export type Usage = { used: number; max: number; usagePercentage: number };
export type Metrics = Record<PlanMetricsType, Usage>;
export type Subscription = PlanSubscription & {
  plan: Plan & { features: PlanFeatures };
  metrics: Metrics;
};

export type PlanPluginArgs = GetPlans;
export type PlanSubscriptionPluginArgs =
  | GetPlanSubscriptionByIdPluginArgs
  | GetPlanSubscriptionByUserSessionPluginArgs
  | CreatePlanSubscriptionPluginArgs
  | UpdatePlanSubscriptionPluginArgs
  | DeletePlanSubscriptionPluginArgs;

export type PlanSubscriptionMetricsPluginArgs =
  UpdatePlanSubscriptionMetricsPluginArgs;

/// GetPlans
export interface GetPlans {}
export interface GetPlansApiResponse extends ApiResponse {}
export interface GetPlansApiArgs {
  planType?: PlanType;
}
export interface GetPlansPluginArgs extends GetPlansApiArgs {}

/// GetPlanSubscriptionById
export interface GetPlanSubscriptionById {}
export interface GetPlanSubscriptionByIdApiResponse extends ApiResponse {}
export interface GetPlanSubscriptionByIdApiArgs {
  id?: string;
}
export interface GetPlanSubscriptionByIdPluginArgs
  extends GetPlanSubscriptionByIdApiArgs {}

/// GetPlanSubscriptionByUserSession
export interface GetPlanSubscriptionByUserSession {}
export interface GetPlanSubscriptionByUserSessionApiResponse
  extends ApiResponse {}
export type GetPlanSubscriptionByUserSessionApiArgs = void;
export interface GetPlanSubscriptionByUserSessionPluginArgs {}

/// CreatePlanSubscription
export interface CreatePlanSubscription {}
export interface CreatePlanSubscriptionApiResponse extends ApiResponse {}
export interface CreatePlanSubscriptionApiArgs
  extends Partial<PlanSubscription> {}
export interface CreatePlanSubscriptionPluginArgs
  extends CreatePlanSubscriptionApiArgs {}

/// UpdatePlanSubscription
export interface UpdatePlanSubscription {}
export interface UpdatePlanSubscriptionApiResponse extends ApiResponse {}
export interface UpdatePlanSubscriptionApiArgs
  extends Partial<PlanSubscription> {
  id: string;
  planId: string;
}
export interface UpdatePlanSubscriptionPluginArgs
  extends UpdatePlanSubscriptionApiArgs {}

/// DeletePlanSubscription
export interface DeletePlanSubscription {}
export interface DeletePlanSubscriptionApiResponse extends ApiResponse {}
export interface DeletePlanSubscriptionApiArgs {
  id: string;
  planId: string;
}
export interface DeletePlanSubscriptionPluginArgs
  extends DeletePlanSubscriptionApiArgs {}

/// UpdatePlanSubscriptionMetrics
export interface UpdatePlanSubscriptionMetrics {}
export interface UpdatePlanSubscriptionMetricsApiResponse extends ApiResponse {}
export interface UpdatePlanSubscriptionMetricsApiArgs {
  incOs?: boolean;
  decOs?: boolean;
  incWorkspace?: boolean;
  decWorkspace?: boolean;
  incApp?: boolean;
  decApp?: boolean;
  incForm?: boolean;
  decForm?: boolean;
  incPage?: boolean;
  decPage?: boolean;
  incFileStorage?: boolean;
  decFileStorage?: boolean;
}

export interface UpdatePlanSubscriptionMetricsPluginArgs
  extends UpdatePlanSubscriptionMetricsApiArgs {}
