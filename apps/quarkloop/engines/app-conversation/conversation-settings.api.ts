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
  GetAppConversationSettingsByIdPlugin,
  GetAppConversationSettingsByAppIdPlugin,
  CreateAppConversationSettingsPlugin,
  UpdateAppConversationSettingsPlugin,
  DeleteAppConversationSettingsPlugin,
} from "./conversation-settings.plugin";

import { GetUserPlugin } from "@/engines/user";
import {
  CreateAppConversationSettingsPluginArgs,
  DeleteAppConversationSettingsPluginArgs,
  GetAppConversationSettingsByIdPluginArgs,
  GetAppConversationSettingsByAppIdPluginArgs,
  UpdateAppConversationSettingsPluginArgs,
} from "./conversation-settings.type";
import { GetPlanSubscriptionByUserSessionPlugin } from "@/engines/plan";

/// GetAppConversationSettingsById
export async function GET_GetAppConversationSettingsById(
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
    .use(GetAppConversationSettingsByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: fileId,
      } as GetAppConversationSettingsByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetAppConversationSettingsByAppId
export async function GET_GetAppConversationSettingsByAppId(
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
    .use(GetAppConversationSettingsByAppIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as GetAppConversationSettingsByAppIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// CreateAppConversationSettings
export async function POST_CreateAppConversationSettings(
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
    .use(CreateAppConversationSettingsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationSettings: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as CreateAppConversationSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateAppConversationSettings
export async function PUT_UpdateAppConversationSettings(
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
    .use(UpdateAppConversationSettingsPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationSettings: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as UpdateAppConversationSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// DeleteAppConversationSettings
export async function PATCH_DeleteAppConversationSettings(
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
    .use(DeleteAppConversationSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: fileId,
      } as DeleteAppConversationSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
