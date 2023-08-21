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
  GetAppFormSettingsByIdPlugin,
  GetAppFormsSettingsByAppIdPlugin,
  CreateAppFormSettingsPlugin,
  UpdateAppFormSettingsPlugin,
  DeleteAppFormSettingsPlugin,
} from "./form-settings.plugin";

import { GetUserPlugin } from "@/engines/user";
import {
  CreateAppFormSettingsPluginArgs,
  DeleteAppFormSettingsPluginArgs,
  GetAppFormSettingsByIdPluginArgs,
  GetAppFormsSettingsByAppIdPluginArgs,
  UpdateAppFormSettingsPluginArgs,
} from "./form-settings.type";
import { GetPlanSubscriptionByUserSessionPlugin } from "@/engines/plan";

/// GetAppFormSettingsById
export async function GET_GetAppFormSettingsById(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, formId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetAppFormSettingsByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: formId,
      } as GetAppFormSettingsByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetAppFormsSettingsByAppId
export async function GET_GetAppFormsSettingsByAppId(
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
    .use(GetAppFormsSettingsByAppIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as GetAppFormsSettingsByAppIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// CreateAppFormSettings
export async function POST_CreateAppFormSettings(
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
    .use(CreateAppFormSettingsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as CreateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateAppFormSettings
export async function PUT_UpdateAppFormSettings(
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
    .use(UpdateAppFormSettingsPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as UpdateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// DeleteAppFormSettings
export async function PATCH_DeleteAppFormSettings(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, formId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(DeleteAppFormSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: formId,
      } as DeleteAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
