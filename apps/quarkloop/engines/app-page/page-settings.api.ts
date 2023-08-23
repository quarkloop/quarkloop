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
  GetAppPageSettingsByIdPlugin,
  GetAppPagesSettingsByAppIdPlugin,
  CreateAppPageSettingsPlugin,
  UpdateAppPageSettingsPlugin,
  DeleteAppPageSettingsPlugin,
} from "./page-settings.plugin";

import { GetUserPlugin } from "@/engines/user";
import {
  CreateAppPageSettingsPluginArgs,
  DeleteAppPageSettingsPluginArgs,
  GetAppPageSettingsByIdPluginArgs,
  GetAppPagesSettingsByAppIdPluginArgs,
  UpdateAppPageSettingsPluginArgs,
} from "./page-settings.type";
import { GetPlanSubscriptionByUserSessionPlugin } from "@/engines/plan";

/// GetAppPageSettingsById
export async function GET_GetAppPageSettingsById(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, pageId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
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

/// GetAppPagesSettingsByAppId
export async function GET_GetAppPagesSettingsByAppId(
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

/// CreateAppPageSettings
export async function POST_CreateAppPageSettings(
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

/// UpdateAppPageSettings
export async function PUT_UpdateAppPageSettings(
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

/// DeleteAppPageSettings
export async function PATCH_DeleteAppPageSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, pageId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
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
