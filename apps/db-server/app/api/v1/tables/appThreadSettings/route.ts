import { NextResponse } from "next/server";

import {
  CreateAppThreadSettingsPluginArgs,
  DeleteAppThreadSettingsPluginArgs,
  GetAppThreadSettingsByIdPluginArgs,
  GetAppThreadSettingsByAppIdPluginArgs,
  UpdateAppThreadSettingsPluginArgs,
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
  GetAppThreadSettingsByIdPlugin,
  GetAppThreadSettingsByAppIdPlugin,
  CreateAppThreadSettingsPlugin,
  UpdateAppThreadSettingsPlugin,
  DeleteAppThreadSettingsPlugin,
} from "@quarkloop/plugins";

// GetAppThreadSettingsById
// GetAppThreadSettingsByAppId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appId = searchParams.get("appId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppThreadSettingsByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appThreadSettings: {
          id: id,
        } as GetAppThreadSettingsByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appId) {
    const finalState = await pipeline
      .use(GetAppThreadSettingsByAppIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appThreadSettings: {
          appId: appId,
        } as GetAppThreadSettingsByAppIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppThreadSettings
export async function POST(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppThreadSettingsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadSettings: {
        ...body,
      } as CreateAppThreadSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppThreadSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppThreadSettingsPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadSettings: {
        ...body,
      } as UpdateAppThreadSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppThreadSettings
export async function DELETE(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appId = searchParams.get("appId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppThreadSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadSettings: {
        id: id,
        appId: appId,
      } as DeleteAppThreadSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
