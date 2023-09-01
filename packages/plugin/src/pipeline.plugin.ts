import {
    Plugin,
    DefaultPipelineArgs,
    DefaultPipelineState,
    DefaultPluginConfig,
} from "./pipeline.type";

export function createPlugin<
    TState extends DefaultPipelineState,
    TArgs extends DefaultPipelineArgs[] = DefaultPipelineArgs[],
    TConfig = DefaultPluginConfig
>({
    name,
    config,
    dependencies,
    handler
}: Plugin<TState, TArgs, TConfig>): Plugin<TState, TArgs, TConfig> {
    return {
        name,
        config: config ?? {} as TConfig,
        dependencies,
        handler,
    };
}