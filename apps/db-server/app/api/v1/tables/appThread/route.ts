import { NextResponse } from "next/server";

import {
  CreateAppThreadDataPluginArgs,
  DeleteAppThreadDataPluginArgs,
  GetAppThreadDataByIdPluginArgs,
  GetAppThreadDataByAppSubmissionIdPluginArgs,
  UpdateAppThreadDataPluginArgs,
} from "@quarkloop/types";
import { createPipeline } from "@quarkloop/plugin";
import {
  PipelineState,
  PipelineArgs,
  DefaultErrorHandler,
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
  GetAppThreadDataByIdPlugin,
  GetAppThreadDataByAppSubmissionIdPlugin,
  CreateAppThreadDataPlugin,
  UpdateAppThreadDataPlugin,
  DeleteAppThreadDataPlugin,
} from "@quarkloop/plugins";

// GetAppThreadDataById
export async function GET(request: Request, { params }: { params: any }) {
  const { submissionId, threadId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppThreadDataByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadData: {
        appSubmissionId: submissionId,
        id: threadId,
      } as GetAppThreadDataByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// // GetAppThreadDataByAppSubmissionId
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { submissionId } = params;

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppThreadDataByAppSubmissionIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appThreadData: {
//         appSubmissionId: submissionId,
//       } as GetAppThreadDataByAppSubmissionIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateAppThreadData
export async function POST(request: Request, { params }: { params: any }) {
  const { submissionId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppThreadDataPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadData: {
        ...body,
        appSubmissionId: submissionId,
      } as CreateAppThreadDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppThreadData
export async function PUT(request: Request, { params }: { params: any }) {
  const { submissionId, threadId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppThreadDataPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadData: {
        ...body,
        appSubmissionId: submissionId,
        id: threadId,
      } as UpdateAppThreadDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppThreadData
export async function PATCH(request: Request, { params }: { params: any }) {
  const { submissionId, threadId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppThreadDataPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadData: {
        appSubmissionId: submissionId,
        id: threadId,
      } as DeleteAppThreadDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
