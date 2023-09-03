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
export async function GET(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId, fileId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppThreadSettingsByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: fileId,
      } as GetAppThreadSettingsByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// // GetAppThreadSettingsByAppId
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { osId, workspaceId, appId } = params;

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppThreadSettingsByAppIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appThreadSettings: {
//         osId: osId,
//         workspaceId: workspaceId,
//         appId: appId,
//       } as GetAppThreadSettingsByAppIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateAppThreadSettings
export async function POST(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as CreateAppThreadSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppThreadSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;
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
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
      } as UpdateAppThreadSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppThreadSettings
export async function PATCH(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId, fileId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppThreadSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThreadSettings: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: fileId,
      } as DeleteAppThreadSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
