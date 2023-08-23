// import { Plugin, PipelineErrorHandler } from "./pipeline.type";
// import { createPipeline } from "./pipeline";

// interface PipelineState {
//   count?: number;
// }

// interface PipelineArgs {
//   count: number;
//   label: string;
// }

// interface PluginConfig {
//   prefix?: string;
//   suffix?: string;
// }

// export function createPlugin<
//   TArgs extends [PipelineArgs],
//   TState extends unknown,
//   TConfig = PluginConfig
// >({
//   name,
//   config,
//   dependencies,
// }: {
//   name: string;
//   config?: TConfig;
//   dependencies?: Plugin<any[], TState, any>[];
// }): Plugin<TArgs, TState, TConfig> {
//   return {
//     name,
//     config: config ?? ({} as TConfig),
//     dependencies,
//     handler: async (state, config, ...arg) => {
//       //console.log(`Plugin 1: ${config.prefix} ${arg}`);
//       const nextState = { ...state, count: (state.count ?? 0) + 10 };
//       pipeline.emit("plugin1Executed", nextState);
//       return nextState;
//     },
//   };
// }

// // Define your plugin functions
// const plugin1: Plugin<[PipelineArgs], PipelineState, PluginConfig> =
//   createPlugin({
//     name: "Plugin 1",
//     config: { prefix: "one" } as PluginConfig,
//   });

// const plugin2: Plugin<[PipelineArgs], PipelineState, PluginConfig> =
//   createPlugin({
//     name: "Plugin 2",
//     config: { suffix: "two" } as PluginConfig,
//   });

// const plugin3: Plugin<[PipelineArgs], PipelineState, PluginConfig> =
//   createPlugin({
//     name: "Plugin 3",
//     config: { suffix: "three" } as PluginConfig,
//   });

// // define error handler
// const errorHandler: PluginErrorHandler<[PipelineArgs], PipelineState, Error> = (
//   error,
//   state,
//   ...args
// ) => {
//   console.error("Error occurred:", error);
//   console.log("Current state:", state);
// };

// // Create an instance of Pipeline
// const initialState: PipelineState = { count: 0 };
// const pipeline = createPipeline<[PipelineArgs], PipelineState, Error>({
//   initialState,
//   beforeExecute: (state) => {
//     //console.log('Before execution');
//     //console.log('Initial state:', state);
//   },
//   afterExecute: (state) => {
//     //console.log('After execution');
//     //console.log('Final state:', state);
//   },
// });

// // Use plugins
// pipeline.use(plugin1);
// pipeline.use(plugin2);
// pipeline.use(plugin3);
// pipeline.onError(errorHandler);

// // Subscribe to events
// pipeline.subscribe("plugin1Executed", (state, nextState) => {
//   console.log("[event] Plugin 1:", state, nextState);
// });

// pipeline.subscribe("plugin2Executed", (state, nextState) => {
//   console.log("[event] Plugin 2:", state, nextState);
// });

// pipeline.subscribe("Plugin3Executed", (state, nextState) => {
//   console.log("[event] Plugin 3:", state, nextState);
// });

// // Execute the plugins
// pipeline.execute({ count: 42, label: "ChatGPT" });
