import { NextResponse } from "next/server";

import {
  CreateAppFileSettingsPluginArgs,
  DeleteAppFileSettingsPluginArgs,
  GetAppFileSettingsByIdPluginArgs,
  GetAppFileSettingsByAppIdPluginArgs,
  UpdateAppFileSettingsPluginArgs,
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
  GetAppFileSettingsByIdPlugin,
  GetAppFileSettingsByAppIdPlugin,
  CreateAppFileSettingsPlugin,
  UpdateAppFileSettingsPlugin,
  DeleteAppFileSettingsPlugin,
} from "@quarkloop/plugins";

// GetAppFileSettingsById
// GetAppFileSettingsByAppId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appId = searchParams.get("appId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppFileSettingsByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appFileSettings: {
          id: id,
        } as GetAppFileSettingsByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appId) {
    const finalState = await pipeline
      .use(GetAppFileSettingsByAppIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appFileSettings: {
          appId: appId,
        } as GetAppFileSettingsByAppIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppFileSettings
export async function POST(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppFileSettingsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFileSettings: {
        ...body,
      } as CreateAppFileSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppFileSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppFileSettingsPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFileSettings: {
        ...body,
      } as UpdateAppFileSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppFileSettings
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
    .use(DeleteAppFileSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFileSettings: {
        id: id,
        appId: appId,
      } as DeleteAppFileSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
