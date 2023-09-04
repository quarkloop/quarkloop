import { NextResponse } from "next/server";

import {
  GetAppFormByIdPluginArgs,
  CreateAppFormPluginArgs,
  DeleteAppFormPluginArgs,
  UpdateAppFormPluginArgs,
  GetAppFormByAppInstanceIdPluginArgs,
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
  GetAppFormByAppInstanceIdPlugin,
} from "@quarkloop/plugins";

// GetAppFormById
// GetAppFormByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appInstanceId = searchParams.get("appInstanceId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppFormByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appForm: {
          id: id,
        } as GetAppFormByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appInstanceId) {
    const finalState = await pipeline
      .use(GetAppFormByAppInstanceIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appForm: {
          appInstanceId: appInstanceId,
        } as GetAppFormByAppInstanceIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppForm
export async function POST(request: Request, { params }: { params: any }) {
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
      } as CreateAppFormPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppForm
export async function PUT(request: Request, { params }: { params: any }) {
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
      } as UpdateAppFormPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppForm
export async function DELETE(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appInstanceId = searchParams.get("appInstanceId");

  if (id == null && appInstanceId == null) {
    return NextResponse.json({ status: "Bad request" }, { status: 400 });
  }

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppFormPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appForm: {
        id: id,
        appInstanceId: appInstanceId,
      } as DeleteAppFormPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
