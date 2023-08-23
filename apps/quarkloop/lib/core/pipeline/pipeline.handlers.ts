import { PipelineErrorHandler } from "@/lib/pipeline";

export const NextApiErrorHandler: PipelineErrorHandler = (
    state,
    error,
    ...args
) => {
    console.error('Error occurred:', error);
    console.log('Current state:', state);
    return state;
};