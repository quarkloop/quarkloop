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
export async function GET(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId, formId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
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

// // GetAppFormsSettingsByAppId
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { osId, workspaceId, appId } = params;

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppFormsSettingsByAppIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appFormSettings: {
//         osId: osId,
//         workspaceId: workspaceId,
//         appId: appId,
//       } as GetAppFormsSettingsByAppIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateAppFormSettings
export async function POST(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as CreateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppFormSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as UpdateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppFormSettings
export async function PATCH(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId, formId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
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
