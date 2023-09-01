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

export async function GET_GetAppPageSettingsById(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, pageId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppPageSettingsByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPageSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: pageId,
      } as GetAppPageSettingsByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function GET_GetAppPagesSettingsByAppId(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppPagesSettingsByAppIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPageSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as GetAppPagesSettingsByAppIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function POST_CreateAppPageSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as CreateAppPageSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function PUT_UpdateAppPageSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as UpdateAppPageSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function PATCH_DeleteAppPageSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, pageId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppPageSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appPageSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: pageId,
      } as DeleteAppPageSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
