import {
  CreatePipelineOptions,
  Pipeline,
  PipelineErrorHandler,
  Plugin,
  PluginEvent,
  PluginEventHandler,
  PluginLifecycleHook,
  DefaultPipelineArgs,
  DefaultPipelineState,
} from "./pipeline.type";

export function createPipeline<
  TState extends DefaultPipelineState = DefaultPipelineState,
  TArgs extends DefaultPipelineArgs[] = DefaultPipelineArgs[],
  TError extends Error = Error
>(options: CreatePipelineOptions<TState>): Pipeline<TState, TArgs, TError> {
  const { initialState, beforeExecute, afterExecute } = options;

  const plugins: Plugin<TState, TArgs>[] = [];

  let currentState: TState = initialState;
  let nextState: TState | null = null;

  const beforeExecuteHooks: PluginLifecycleHook<TState>[] = [];
  const afterExecuteHooks: PluginLifecycleHook<TState>[] = [];
  const eventSubscriptions: Record<
    PluginEvent,
    PluginEventHandler<TState, TArgs>[]
  > = {};

  let errorHandler: PipelineErrorHandler<TState, TArgs, TError> | undefined;

  function use<T extends Plugin<TState, TArgs>>(
    plugin: T
  ): Pipeline<TState, TArgs, TError> {
    const dependencies = plugin.dependencies ?? [];
    const missingDependencies = dependencies.filter(
      (dependency) => !plugins.includes(dependency)
    );

    if (missingDependencies.length > 0) {
      throw new Error(
        `[Pipeline] Dependency not found: ${missingDependencies[0].name}`
      );
    }

    plugins.push(plugin);
    console.log(`[Pipeline] Added plugin: ${plugin.name}`);
    return pipeline;
  }

  function onError(
    handler: PipelineErrorHandler<TState, TArgs, TError>
  ): Pipeline<TState, TArgs, TError> {
    errorHandler = handler;
    return pipeline;
  }

  async function execute(...args: TArgs): Promise<TState> {
    try {
      nextState = currentState;

      for (const hook of beforeExecuteHooks) {
        hook(nextState);
      }

      for (const plugin of plugins) {
        nextState = await plugin.handler(nextState, plugin.config, ...args);
      }

      for (const hook of afterExecuteHooks) {
        hook(nextState);
      }

      currentState = nextState;

      console.log("[Pipeline] Execution completed");
      return currentState;
    } catch (error: any) {
      if (errorHandler) {
        return errorHandler(nextState ?? currentState, error, ...args);
      } else {
        console.error("[Pipeline] Plugin execution error:", error);
      }
      throw error;
    }
  }

  function onBeforeExecute(hook: PluginLifecycleHook<TState>): void {
    beforeExecuteHooks.push(hook);
  }

  function onAfterExecute(hook: PluginLifecycleHook<TState>): void {
    afterExecuteHooks.push(hook);
  }

  function emit(event: PluginEvent, ...args: TArgs): void {
    const subscriptions = eventSubscriptions[event];
    if (subscriptions) {
      for (const handler of subscriptions) {
        handler(nextState ?? currentState, ...args);
      }
    }
  }

  function subscribe(
    event: PluginEvent,
    handler: PluginEventHandler<TState, TArgs>
  ): void {
    if (!eventSubscriptions[event]) {
      eventSubscriptions[event] = [];
    }
    eventSubscriptions[event].push(handler);
  }

  const pipeline: Pipeline<TState, TArgs, TError> = {
    plugins,
    errorHandler,
    state: currentState,
    use,
    onError,
    execute,
    onBeforeExecute,
    onAfterExecute,
    emit,
    subscribe,
  };

  if (beforeExecute) {
    onBeforeExecute(beforeExecute);
  }

  if (afterExecute) {
    onAfterExecute(afterExecute);
  }

  console.log("[Pipeline] Creating plugin...");

  return pipeline;
}
