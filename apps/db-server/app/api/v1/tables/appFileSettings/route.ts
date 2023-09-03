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

export async function GET_GetAppFileSettingsById(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, fileId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppFileSettingsByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFileSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: fileId,
      } as GetAppFileSettingsByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function GET_GetAppFileSettingsByAppId(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppFileSettingsByAppIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFileSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as GetAppFileSettingsByAppIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function POST_CreateAppFileSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as CreateAppFileSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function PUT_UpdateAppFileSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as UpdateAppFileSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

export async function PATCH_DeleteAppFileSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, fileId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppFileSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFileSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: fileId,
      } as DeleteAppFileSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
