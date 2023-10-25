// import {
//     Plan as PrismaPlan,
//     PlanFeatures as PrismaPlanFeatures,
//     PlanPricing as PrismaPlanPricing,
//     PlanBilling as PrismaPlanBilling,
//     PlanSubscription as PrismaPlanSubscription,
//     PlanMetrics as PrismaPlanMetrics,
//     PlanType,
//     PlanMetricsType,
// } from "@quarkloop/prisma/types";
import { ApiResponse } from "./api-response";

export interface Plan {
    id: string;
    type: PlanType;
    status: PlanStatus;
    name: string;
    description: string;
    isDefault: boolean;
    createdAt: Date;
    updatedAt: Date | null;
    briefFeaturesLabel: string;
    briefFeatures: any;
    buttonLabel: string;
    buttonSignupPath: string;
}

export interface PlanFeatures {
    id: string;
    maxOs: number;
    maxOsLabel: string;
    maxWorkspaces: number;
    maxWorkspacesLabel: string;
    maxApps: number;
    maxAppsLabel: string;
    maxAppConversations: number;
    maxAppConversationsLabel: string;
    maxAppPages: number;
    maxAppPagesLabel: string;
    maxAppForms: number;
    maxAppFormsLabel: string;
    maxAppFormFields: number;
    maxAppFormFieldsLabel: string;
    maxAppFiles: number;
    maxAppFilesLabel: string;
    planId: string;
}

export interface PlanPricing {
    id: string;
    price: string;
    currency: PlanPricingCurreny;
    cycle: PlanPricingCycle;
    createdAt: Date;
    updatedAt: Date | null;
    planId: string;
}

export interface PlanBilling {
    id: string;
    status: number;
    dueDate: Date;
    amount: number;
    provider: string;
    subscriptionId: string;
}

export interface PlanSubscription {
    id: string;
    status: PlanSubscriptionStatus;
    startDate: Date;
    endDate: Date | null;
    endReason: string | null;
    createdAt: Date;
    updatedAt: Date | null;
    userId: string;
    planId: string;
}

export interface PlanMetrics {
    id: number;
    type: PlanMetricsType;
    value: any | null;
    subscriptionId: string;
    osId: string | null;
    workspaceId: number | null;
}

export enum PlanType {
    Free = "Free",
    Standard = "Standard",
    Professional = "Professional",
    Enterprise = "Enterprise",
}

export enum PlanStatus {
    Active = "Active",
    Archived = "Archived",
}

export enum PlanPricingCurreny {
    Dollar = "Dollar",
    Euro = "Euro",
    TurkishLira = "TurkishLira",
}

export enum PlanPricingCycle {
    Monthly = "Monthly",
    Yearly = "Yearly",
}

export enum PlanSubscriptionStatus {
    Active = "Active",
    Expired = "Expired",
    Upgraded = "Upgraded",
    DownGraded = "DownGraded",
    Canceled = "Canceled",
}

export enum PlanMetricsType {
    Os = "Os",
    Workspace = "Workspace",
    App = "App",
    AppConversation = "AppConversation",
    AppForm = "AppForm",
    AppFormField = "AppFormField",
    AppPage = "AppPage",
    AppFile = "AppFile",
}

export enum PermissionRole {
    Owner = "Owner",
    Admin = "Admin",
    Member = "Member",
    Viewer = "Viewer",
    Contributer = "Contributer",
    Guest = "Guest",
}

export enum PermissionType {
    OS = "OS",
    Workspace = "Workspace",
    App = "App",
}

export type Usage = { used: number; max: number; usagePercentage: number };

export type Metrics = Record<PlanMetricsType, Usage>;

export type Subscription = PlanSubscription & {
    plan: Plan & { features: PlanFeatures };
    metrics: Metrics;
};

// export type PlanPluginArgs = GetPlans;
// export type PlanSubscriptionPluginArgs =
//   | GetPlanSubscriptionByIdPluginArgs
//   | GetPlanSubscriptionByUserSessionPluginArgs
//   | CreatePlanSubscriptionPluginArgs
//   | UpdatePlanSubscriptionPluginArgs
//   | DeletePlanSubscriptionPluginArgs;

// export type PlanSubscriptionMetricsPluginArgs =
//   UpdatePlanSubscriptionMetricsPluginArgs;

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
