import { NextResponse } from "next/server";

import {
  CreateAppThreadPluginArgs,
  DeleteAppThreadPluginArgs,
  GetAppThreadByIdPluginArgs,
  GetAppThreadByAppInstanceIdPluginArgs,
  UpdateAppThreadPluginArgs,
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
  GetAppThreadByIdPlugin,
  GetAppThreadByAppInstanceIdPlugin,
  CreateAppThreadPlugin,
  UpdateAppThreadPlugin,
  DeleteAppThreadPlugin,
} from "@quarkloop/plugins";

// GetAppThreadById
// GetAppThreadByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appInstanceId = searchParams.get("appInstanceId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppThreadByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appThread: {
          id: id,
        } as GetAppThreadByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appInstanceId) {
    const finalState = await pipeline
      .use(GetAppThreadByAppInstanceIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appThread: {
          appInstanceId: appInstanceId,
        } as GetAppThreadByAppInstanceIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppThread
export async function POST(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppThreadPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        ...body,
      } as CreateAppThreadPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppThread
export async function PUT(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppThreadPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        ...body,
      } as UpdateAppThreadPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppThread
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
    .use(DeleteAppThreadPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        id: id,
        appInstanceId: appInstanceId,
      } as DeleteAppThreadPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
