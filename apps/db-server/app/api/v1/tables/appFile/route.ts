import { NextResponse } from "next/server";

import {
  GetAppFileByIdPluginArgs,
  CreateAppFilePluginArgs,
  DeleteAppFilePluginArgs,
  UpdateAppFilePluginArgs,
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
  GetAppFileByIdPlugin,
  CreateAppFilePlugin,
  UpdateAppFilePlugin,
  DeleteAppFilePlugin,
} from "@quarkloop/plugins";

// GetAppFileById
export async function GET(request: Request, { params }: { params: any }) {
  const { submissionId, fileId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppFileByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFile: {
        appSubmissionId: submissionId,
        id: fileId,
      } as GetAppFileByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// CreateAppFile
export async function POST(request: Request, { params }: { params: any }) {
  const { submissionId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppFilePlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFile: {
        ...body,
        appSubmissionId: submissionId,
      } as CreateAppFilePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppFile
export async function PUT(request: Request, { params }: { params: any }) {
  const { submissionId, fileId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppFilePlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFile: {
        ...body,
        appSubmissionId: submissionId,
        id: fileId,
      } as UpdateAppFilePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppFile
export async function DELETE(request: Request, { params }: { params: any }) {
  const { submissionId, fileId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppFilePlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFile: {
        appSubmissionId: submissionId,
        id: fileId,
      } as DeleteAppFilePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
