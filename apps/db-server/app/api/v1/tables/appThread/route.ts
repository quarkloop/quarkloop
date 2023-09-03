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

export async function GET_GetAppThreadDataById(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId, conversationId } = params;

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
        id: conversationId,
      } as GetAppThreadDataByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function GET_GetAppThreadDataByAppSubmissionId(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppThreadDataByAppSubmissionIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadData: {
        appSubmissionId: submissionId,
      } as GetAppThreadDataByAppSubmissionIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function POST_CreateAppThreadData(
  request: Request,
  { params }: { params: any }
) {
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

export async function PUT_UpdateAppThreadData(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId, conversationId } = params;
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
        id: conversationId,
      } as UpdateAppThreadDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function PATCH_DeleteAppThreadData(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId, conversationId } = params;

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
        id: conversationId,
      } as DeleteAppThreadDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
