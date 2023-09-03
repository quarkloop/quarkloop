import { NextResponse } from "next/server";

import {
  GetAdminAppInstanceByIdPluginArgs,
  GetAppInstanceByIdPluginArgs,
  CreateAppInstancePluginArgs,
  GetAppInstancesByAppIdPluginArgs,
  UpdateAppInstancePluginArgs,
} from "@quarkloop/types";
import { createPipeline } from "@quarkloop/plugin";
import {
  PipelineArgs,
  PipelineState,
  DefaultErrorHandler,
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
  GetAdminAppInstanceByIdPlugin,
  GetAppInstanceByIdPlugin,
  CreateAppInstancePlugin,
  GetAppInstancesByAppIdPlugin,
  UpdateAppInstancePlugin,
} from "@quarkloop/plugins";

// GetAdminAppInstanceById
export async function GET(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId, id } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAdminAppInstanceByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appInstance: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: id,
      } as GetAdminAppInstanceByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// // GetAppInstanceById
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { userId, appId, id } = params;

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppInstanceByIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appInstance: {
//         userId: userId,
//         appId: appId,
//         id: id,
//       } as GetAppInstanceByIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// // GetAppInstancesByAppId
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { searchParams } = new URL(request.url);

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppInstancesByAppIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appInstance: {
//         osId: searchParams.get("osId"),
//         workspaceId: searchParams.get("workspaceId"),
//         appId: searchParams.get("appId"),
//       } as GetAppInstancesByAppIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateAppInstance
export async function POST(request: Request, { params }: { params: any }) {
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppInstancePlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appInstance: body as CreateAppInstancePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppInstance
export async function PUT(request: Request, { params }: { params: any }) {
  const { submissionId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppInstancePlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appInstance: {
        ...body,
        id: submissionId,
      } as UpdateAppInstancePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
