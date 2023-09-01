import { PluginStatus } from "./pipeline.error";

import {
  ApiResponse,
  App,
  AppPluginArgs,
  AppInstance,
  AppInstancePluginArgs,
  AppPageSettings,
  AppPageSettingsPluginArgs,
  AppFormSettings,
  AppFormSettingsPluginArgs,
  AppFileSettings,
  AppFileSettingsPluginArgs,
  AppThreadSettings,
  AppThreadSettingsPluginArgs,
  AppThreadData,
  AppThreadDataPluginArgs,
} from "@quarkloop/types";

type ApiResponseState = ApiResponse;

// Status handling properties
type StatusState = PluginStatus | Error;

export type TableState<T extends any> = {
  records: T | T[];
  totalRecords: number;
};

// Database specific properties
export type DatabaseState = {
  app?: TableState<App>;
  appThreadSettings?: TableState<AppThreadSettings>;
  appPageSettings?: TableState<AppPageSettings>;
  appFormSettings?: TableState<AppFormSettings>;
  appFileSettings?: TableState<AppFileSettings>;
  appThreadData?: TableState<AppThreadData>;
  appInstance?: TableState<AppInstance>;
};

export interface PipelineArgs {
  app?: AppPluginArgs;
  appThreadSettings?: AppThreadSettingsPluginArgs;
  appPageSettings?: AppPageSettingsPluginArgs;
  appFormSettings?: AppFormSettingsPluginArgs;
  appFileSettings?: AppFileSettingsPluginArgs;
  appThreadData?: AppThreadDataPluginArgs;
  appInstance?: AppInstancePluginArgs;
}

export interface PipelineState {
  status?: StatusState;
  database?: DatabaseState;
  apiResponse?: ApiResponseState;
}
