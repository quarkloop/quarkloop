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
  const { appId, formId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppFormSettingsByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
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
//   const { appId } = params;

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppFormsSettingsByAppIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appFormSettings: {
//         appId: appId,
//       } as GetAppFormsSettingsByAppIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateAppFormSettings
export async function POST(request: Request, { params }: { params: any }) {
  const { appId } = params;
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
        appId: appId,
      } as CreateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppFormSettings
export async function PUT(request: Request, { params }: { params: any }) {
  const { appId } = params;
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
        appId: appId,
      } as UpdateAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppFormSettings
export async function DELETE(request: Request, { params }: { params: any }) {
  const { appId, formId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppFormSettingsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appFormSettings: {
        appId: appId,
        id: formId,
      } as DeleteAppFormSettingsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
