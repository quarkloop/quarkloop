export type DefaultPipelineState = unknown;
export type DefaultPipelineArgs = unknown;
export type DefaultPluginConfig = unknown;

export type Plugin<
  TState extends DefaultPipelineState = DefaultPipelineState,
  TArgs extends DefaultPipelineArgs[] = DefaultPipelineArgs[],
  TConfig extends DefaultPluginConfig = DefaultPluginConfig
> = {
  name: string;
  config: TConfig;
  dependencies?: Plugin<TState, TArgs, TConfig>[];
  handler: (
    state: TState,
    config: TConfig,
    ...args: TArgs
  ) => TState | Promise<TState>;
};

export type PipelineErrorHandler<
  TState extends DefaultPipelineState = DefaultPipelineState,
  TArgs extends DefaultPipelineArgs[] = DefaultPipelineArgs[],
  TError extends Error = Error
> = (state: TState, error: TError, ...args: TArgs) => TState;

export type PluginLifecycleHook<TState extends DefaultPipelineState> = (
  state: TState
) => void;

export type PluginEvent = string;

export type PluginEventHandler<
  TState extends DefaultPipelineState,
  TArgs extends DefaultPipelineArgs[]
> = (state: TState, ...args: TArgs) => void;

export interface Pipeline<
  TState extends DefaultPipelineState = DefaultPipelineState,
  TArgs extends DefaultPipelineArgs[] = DefaultPipelineArgs[],
  TError extends Error = Error
> {
  state: TState;
  plugins: Plugin<TState, TArgs>[];

  use<T extends Plugin<TState, TArgs, any>>(
    plugin: T
  ): Pipeline<TState, TArgs, TError>;

  execute(...args: TArgs): Promise<TState>;

  errorHandler?: PipelineErrorHandler<TState, TArgs, TError>;
  onError(
    handler: PipelineErrorHandler<TState, TArgs, TError>
  ): Pipeline<TState, TArgs, TError>;
  onBeforeExecute(hook: PluginLifecycleHook<TState>): void;
  onAfterExecute(hook: PluginLifecycleHook<TState>): void;

  emit(event: PluginEvent, ...args: TArgs): void;
  subscribe(
    event: PluginEvent,
    handler: PluginEventHandler<TState, TArgs>
  ): void;
}

export interface CreatePipelineOptions<TState extends DefaultPipelineState> {
  initialState: TState;
  beforeExecute?: PluginLifecycleHook<TState>;
  afterExecute?: PluginLifecycleHook<TState>;
}
