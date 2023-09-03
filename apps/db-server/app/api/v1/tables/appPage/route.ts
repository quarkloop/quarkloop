import { NextResponse } from "next/server";

import {
  GetAppPageByIdPluginArgs,
  CreateAppPagePluginArgs,
  DeleteAppPagePluginArgs,
  UpdateAppPagePluginArgs,
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
  GetAppPageByIdPlugin,
  CreateAppPagePlugin,
  UpdateAppPagePlugin,
  DeleteAppPagePlugin,
} from "@quarkloop/plugins";

// GetAppPageById
export async function GET(request: Request, { params }: { params: any }) {
  const { submissionId, pageId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppPageByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPage: {
        appSubmissionId: submissionId,
        id: pageId,
      } as GetAppPageByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// CreateAppPage
export async function POST(request: Request, { params }: { params: any }) {
  const { submissionId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppPagePlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPage: {
        ...body,
        appSubmissionId: submissionId,
      } as CreateAppPagePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppPage
export async function PUT(request: Request, { params }: { params: any }) {
  const { submissionId, pageId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppPagePlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPage: {
        ...body,
        appSubmissionId: submissionId,
        id: pageId,
      } as UpdateAppPagePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppPage
export async function PATCH(request: Request, { params }: { params: any }) {
  const { submissionId, pageId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppPagePlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPage: {
        appSubmissionId: submissionId,
        id: pageId,
      } as DeleteAppPagePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
