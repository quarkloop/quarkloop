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
  AppThread,
  AppThreadPluginArgs,
  AppPagePluginArgs,
  AppFormPluginArgs,
  AppFilePluginArgs,
  AppPage,
  AppForm,
  AppFile,
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
  appInstance?: TableState<AppInstance>;

  appThread?: TableState<AppThread>;
  appThreadSettings?: TableState<AppThreadSettings>;

  appPage?: TableState<AppPage>;
  appPageSettings?: TableState<AppPageSettings>;

  appForm?: TableState<AppForm>;
  appFormSettings?: TableState<AppFormSettings>;

  appFile?: TableState<AppFile>;
  appFileSettings?: TableState<AppFileSettings>;
};

export interface PipelineArgs {
  app?: AppPluginArgs;
  appInstance?: AppInstancePluginArgs;

  appThread?: AppThreadPluginArgs;
  appThreadSettings?: AppThreadSettingsPluginArgs;

  appPage?: AppPagePluginArgs;
  appPageSettings?: AppPageSettingsPluginArgs;

  appForm?: AppFormPluginArgs;
  appFormSettings?: AppFormSettingsPluginArgs;

  appFile?: AppFilePluginArgs;
  appFileSettings?: AppFileSettingsPluginArgs;
}

export interface PipelineState {
  status?: StatusState;
  database?: DatabaseState;
  apiResponse?: ApiResponseState;
}
