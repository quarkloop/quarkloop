import { Session } from "next-auth";

import { PluginStatus } from "./pipeline.error";

import {
  UserPluginArgs,
  UserAccountPluginArgs,
  UserSessionPluginArgs,
  User,
  UserAccount,
  UserSession,
  UserPermissions,
} from "@/engines/user/user.type";
import {
  OperatingSystem,
  OperatingSystemApiResponse,
  OperatingSystemPluginArgs,
  OperatingSystemUser,
} from "@/engines/os/os.type";
import {
  Workspace,
  WorkspacePluginArgs,
} from "@/engines/workspace/workspace.type";
import { App, AppPluginArgs } from "@/engines/app/app.type";
import {
  AppSubmission,
  AppSubmissionPluginArgs,
} from "@/engines/app-submission";
import {
  Plan,
  PlanPluginArgs,
  PlanSubscription,
  PlanSubscriptionMetricsPluginArgs,
  PlanSubscriptionPluginArgs,
} from "@/engines/plan/planSubscription.type";
import {
  AppPageSettings,
  AppPageSettingsPluginArgs,
} from "@/engines/app-page/page-settings.type";
import { AppFormSettings, AppFormSettingsPluginArgs } from "@/engines/app-form";
import { AppFileSettings, AppFileSettingsPluginArgs } from "@/engines/app-file";
import {
  AppConversationSettings,
  AppConversationSettingsPluginArgs,
  AppConversationData,
  AppConversationDataPluginArgs,
} from "@/engines/app-conversation";

// Status handling properties
type StatusState = PluginStatus | Error;

export type ApiResponse = ApiResponseState;

// Rest api response properties
export type ApiResponseState = {
  status: StatusState;
  error?: string;
  errorDetails?: Record<string, any>;
  isOwner?: boolean;
  user?: User;
  userPermissions?: UserPermissions;
  planSubscription?: PlanSubscription;
  data?: {
    database?: DatabaseState;
  };
};

export type TableState<T extends any> = {
  records: T | T[];
  totalRecords: number;
};

// Database specific properties
export type DatabaseState = {
  user?: TableState<User>;
  userAccount?: TableState<UserAccount>;
  userSession?: TableState<UserSession>;
  os?: TableState<OperatingSystem>;
  osUser?: TableState<OperatingSystemUser>;
  workspace?: TableState<Workspace>;
  app?: TableState<App>;
  appConversationSettings?: TableState<AppConversationSettings>;
  appPageSettings?: TableState<AppPageSettings>;
  appFormSettings?: TableState<AppFormSettings>;
  appFileSettings?: TableState<AppFileSettings>;
  appConversationData?: TableState<AppConversationData>;
  appSubmission?: TableState<AppSubmission>;
  plan?: TableState<Plan>;
  planSubscription?: TableState<PlanSubscription>;
};

// Authentication related properties
type SessionState = {
  session: Session;
  sessionToken: string;
};

type UserState = {
  user: User;
  permissions: UserPermissions;
  subscription: PlanSubscription;
};

// Metadata properties
type MetadataState = {
  createdAt: Date;
  updatedAt: Date;
  createdBy: string;
  updatedBy: string;
  version: number;
};

// Context properties
type ContextState = {
  requestId: string;
  locale: string;
  timeZone: string;
  environment: "development" | "staging" | "production";
  request: {
    ipAddress: string;
    headers: Record<string, string>;
    queryParams: Record<string, string>;
  };
  response: {
    headers: Record<string, string>;
  };
};

// Caching properties
type CacheState = {
  enabled: boolean;
  expiry: number;
  maxAge: number;
  cacheKeyPrefix: string;
  strategy: "memory" | "redis" | "custom";
};

// Logging properties
type LoggingState = {
  enabled: boolean;
  level: "info" | "debug" | "warn" | "error";
};

// Metrics properties
type MetricsState = {
  enabled: boolean;
  requests: number;
  responseTime: number;
  visits: number;
};

// Notification properties
type NotificationState = {
  email: {
    enabled: boolean;
    sender: string;
    recipients: string[];
    subject: string;
    body: string;
  };
  sms: {
    enabled: boolean;
    recipients: string[];
  };
  push: {
    enabled: boolean;
    recipients: string[];
  };
};

// External services properties
type ExternalServicesState = {
  database: {
    enabled: boolean;
    connectionString: string;
  };
  paymentGateway: {
    enabled: boolean;
    apiKey: string;
    endpoint: string;
    provider: string;
    account: string;
    method: string;
    amount: number;
    cardNumber: string;
    currency: string;
  };
  messagingService: {
    enabled: boolean;
    apiKey: string;
    endpoint: string;
  };
};

// Integrations properties
type IntegrationsState = {
  github: {
    enabled: boolean;
    apiKey: string;
  };
};

// Feature flags properties
type FlagsState = Record<string, boolean>;

// Progress tracking properties
type ProgressState = {
  currentStep: number;
  totalSteps: number;
  percentage: number;
  status: "pending" | "in-progress" | "completed";
};

export interface PipelineArgs {
  user?: UserPluginArgs;
  userAccount?: UserAccountPluginArgs;
  userSession?: UserSessionPluginArgs;
  os?: OperatingSystemPluginArgs;
  workspace?: WorkspacePluginArgs;
  app?: AppPluginArgs;
  appConversationSettings?: AppConversationSettingsPluginArgs;
  appPageSettings?: AppPageSettingsPluginArgs;
  appFormSettings?: AppFormSettingsPluginArgs;
  appFileSettings?: AppFileSettingsPluginArgs;
  appConversationData?: AppConversationDataPluginArgs;
  appSubmission?: AppSubmissionPluginArgs;
  plan?: PlanPluginArgs;
  planSubscription?: PlanSubscriptionPluginArgs;
  planSubscriptionMetrics?: PlanSubscriptionMetricsPluginArgs;
}

export interface PipelineState {
  status?: StatusState;
  session?: SessionState;
  user?: UserState;
  database?: DatabaseState;

  // api responses
  apiResponse?: ApiResponseState;
  osApiReponse?: OperatingSystemApiResponse;

  metadata?: MetadataState;
  context?: ContextState;
  cache?: CacheState;
  logging?: LoggingState;
  metrics?: MetricsState;
  notifications?: NotificationState;
  externalServices?: ExternalServicesState;
  integrations?: IntegrationsState;
  flags?: FlagsState;
  progress?: ProgressState;
}
