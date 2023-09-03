import { NextResponse } from "next/server";

import {
  GetAppByIdPluginArgs,
  CreateAppPluginArgs,
  UpdateAppPluginArgs,
  DeleteAppPluginArgs,
  GetAppsByOsIdPluginArgs,
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
  GetAppByIdPlugin,
  GetAppsByOsIdPlugin,
  CreateAppPlugin,
  UpdateAppPlugin,
  DeleteAppPlugin,
} from "@quarkloop/plugins";

// GetAppById
export async function GET(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        osId: osId,
        workspaceId: workspaceId,
        id: appId,
      } as GetAppByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// // GetAppsByOsId
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { osId } = params;

//   const { searchParams } = new URL(request.url);
//   const workspaceId = searchParams.get("workspaceId");

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppsByOsIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       app: {
//         osId: osId,
//         workspaceId: workspaceId,
//       } as GetAppsByOsIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateApp
export async function POST(request: Request, { params }: { params: any }) {
  const { osId, workspaceId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
      } as CreateAppPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateApp
export async function PUT(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
      } as UpdateAppPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteApp
export async function PATCH(request: Request, { params }: { params: any }) {
  const { osId, workspaceId, appId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        osId: osId,
        workspaceId: workspaceId,
        id: appId,
      } as DeleteAppPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
