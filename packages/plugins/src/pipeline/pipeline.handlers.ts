import { PipelineErrorHandler } from "../../../plugin/src/pipeline.type";
import { PluginStatusEntry } from "./pipeline.error";

export const DefaultErrorHandler: PipelineErrorHandler<Readonly<any>> = (
  state,
  error,
  ...args
) => {
  console.error("Error occurred:", error);
  console.log("Current state:", state);

  return {
    ...state,
    status: PluginStatusEntry.INTERNAL_SERVER_ERROR({
      errorMessage: "DefaultErrorHandler",
      errorDetail: error,
    }),
  };
};

export const DefaultBeforeExecute = <T>(state: Readonly<T>) => {
  //console.log('Before execution');
  //console.log('Initial state:', state);
};

export const DefaultAfterExecute = <T>(state: Readonly<T>) => {
  //console.log('After execution');
  //console.log('Final state:', state);
};
