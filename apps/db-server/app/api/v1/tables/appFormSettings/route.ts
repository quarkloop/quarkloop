import { NextResponse } from "next/server";

import {
  CreateAppFormSettingsPluginArgs,
  DeleteAppFormSettingsPluginArgs,
  GetAppFormSettingsByIdPluginArgs,
  GetAppFormsSettingsByAppIdPluginArgs,
  UpdateAppFormSettingsPluginArgs,
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
  GetAppFormSettingsByIdPlugin,
  GetAppFormsSettingsByAppIdPlugin,
  CreateAppFormSettingsPlugin,
  UpdateAppFormSettingsPlugin,
  DeleteAppFormSettingsPlugin,
} from "@quarkloop/plugins";

// GetAppFormSettingsById
// GetAppFormsSettingsByAppId
export async function GET(request: Request, { params }: { params: any }) {
  const { searchParams } = new URL(request.url);
  const id = searchParams.get("id");
  const appId = searchParams.get("appId");

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  if (id) {
    const finalState = await pipeline
      .use(GetAppFormSettingsByIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appFormSettings: {
          id: id,
        } as GetAppFormSettingsByIdPluginArgs,
      });

    return NextResponse.json(finalState.apiResponse);
  }

  if (appId) {
    const finalState = await pipeline
      .use(GetAppFormsSettingsByAppIdPlugin)
      .use(GetApiResponsePlugin)
      .onError(DefaultErrorHandler)
      .execute({
        appFormSettings: {
          appId: appId,
        } as GetAppFormsSettingsByAppIdPluginArgs,
      });
    return NextResponse.json(finalState.apiResponse);
  }

  return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppFormSettings
export async function POST(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppFormSettingsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        ...body,
      } as CreateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppFormSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppFormSettingsPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        ...body,
      } as UpdateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppFormSettings
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
    .use(DeleteAppFormSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        id: id,
        appId: appId,
      } as DeleteAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
