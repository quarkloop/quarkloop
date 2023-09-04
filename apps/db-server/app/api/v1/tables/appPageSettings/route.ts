import { NextResponse } from "next/server";

import {
  CreateAppPageSettingsPluginArgs,
  DeleteAppPageSettingsPluginArgs,
  GetAppPageSettingsByIdPluginArgs,
  GetAppPagesSettingsByAppIdPluginArgs,
  UpdateAppPageSettingsPluginArgs,
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
  GetAppPageSettingsByIdPlugin,
  GetAppPagesSettingsByAppIdPlugin,
  CreateAppPageSettingsPlugin,
  UpdateAppPageSettingsPlugin,
  DeleteAppPageSettingsPlugin,
} from "@quarkloop/plugins";

// GetAppPageSettingsById
// GetAppPagesSettingsByAppId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appId = searchParams.get("appId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppPageSettingsByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appPageSettings: {
          id: id,
        } as GetAppPageSettingsByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appId) {
    const finalState = await pipeline
      .use(GetAppPagesSettingsByAppIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appPageSettings: {
          appId: appId,
        } as GetAppPagesSettingsByAppIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppPageSettings
export async function POST(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppPageSettingsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPageSettings: {
        ...body,
      } as CreateAppPageSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppPageSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppPageSettingsPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPageSettings: {
        ...body,
      } as UpdateAppPageSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppPageSettings
export async function DELETE(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appId = searchParams.get("appId");

  if (id == null && appId == null) {
    return NextResponse.json({ status: "Bad request" }, { status: 400 });
  }

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppPageSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPageSettings: {
        id: id,
        appId: appId,
      } as DeleteAppPageSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
