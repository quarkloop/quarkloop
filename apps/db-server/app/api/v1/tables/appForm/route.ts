import { NextResponse } from "next/server";

import {
  GetAppFormByIdPluginArgs,
  CreateAppFormPluginArgs,
  DeleteAppFormPluginArgs,
  UpdateAppFormPluginArgs,
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
  GetAppFormByIdPlugin,
  CreateAppFormPlugin,
  UpdateAppFormPlugin,
  DeleteAppFormPlugin,
} from "@quarkloop/plugins";

// GetAppFormById
export async function GET(request: Request, { params }: { params: any }) {
  const { submissionId, formId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppFormByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appForm: {
        appSubmissionId: submissionId,
        id: formId,
      } as GetAppFormByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// CreateAppForm
export async function POST(request: Request, { params }: { params: any }) {
  const { submissionId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppFormPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appForm: {
        ...body,
        appSubmissionId: submissionId,
      } as CreateAppFormPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppForm
export async function PUT(request: Request, { params }: { params: any }) {
  const { submissionId, formId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppFormPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appForm: {
        ...body,
        appSubmissionId: submissionId,
        id: formId,
      } as UpdateAppFormPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppForm
export async function DELETE(request: Request, { params }: { params: any }) {
  const { submissionId, formId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppFormPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appForm: {
        appSubmissionId: submissionId,
        id: formId,
      } as DeleteAppFormPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
