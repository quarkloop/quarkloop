import { NextResponse } from "next/server";

import {
  GetAppPageByIdPluginArgs,
  CreateAppPagePluginArgs,
  DeleteAppPagePluginArgs,
  UpdateAppPagePluginArgs,
  GetAppPageByAppInstanceIdPluginArgs,
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
  GetAppPageByAppInstanceIdPlugin,
} from "@quarkloop/plugins";

// GetAppPageById
// GetAppPageByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appInstanceId = searchParams.get("appInstanceId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppPageByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appPage: {
          id: id,
        } as GetAppPageByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appInstanceId) {
    const finalState = await pipeline
      .use(GetAppPageByAppInstanceIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appPage: {
          appInstanceId: appInstanceId,
        } as GetAppPageByAppInstanceIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppPage
export async function POST(request: Request, { params }: { params: any }) {
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
      } as CreateAppPagePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppPage
export async function PUT(request: Request, { params }: { params: any }) {
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
      } as UpdateAppPagePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppPage
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
    .use(DeleteAppPagePlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPage: {
        id: id,
        appInstanceId: appInstanceId,
      } as DeleteAppPagePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
