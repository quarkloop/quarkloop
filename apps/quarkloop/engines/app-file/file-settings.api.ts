import { NextResponse } from "next/server";

import { DefaultErrorHandler, createPipeline } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";

import { GetServerSessionPlugin } from "@/lib/core/server-session.plugin";
import {
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
} from "@/lib/core/core.plugins";

import {
  GetAppFileSettingsByIdPlugin,
  GetAppFileSettingsByAppIdPlugin,
  CreateAppFileSettingsPlugin,
  UpdateAppFileSettingsPlugin,
  DeleteAppFileSettingsPlugin,
} from "./file-settings.plugin";

import { GetUserPlugin } from "@/engines/user";
import {
  CreateAppFileSettingsPluginArgs,
  DeleteAppFileSettingsPluginArgs,
  GetAppFileSettingsByIdPluginArgs,
  GetAppFileSettingsByAppIdPluginArgs,
  UpdateAppFileSettingsPluginArgs,
} from "./file-settings.type";
import { GetPlanSubscriptionByUserSessionPlugin } from "@/engines/plan";

/// GetAppFileSettingsById
export async function GET_GetAppFileSettingsById(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, fileId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
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

/// GetAppFileSettingsByAppId
export async function GET_GetAppFileSettingsByAppId(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
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

/// CreateAppFileSettings
export async function POST_CreateAppFileSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetPlanSubscriptionByUserSessionPlugin)
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

/// UpdateAppFileSettings
export async function PUT_UpdateAppFileSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
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

/// DeleteAppFileSettings
export async function PATCH_DeleteAppFileSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, fileId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
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
