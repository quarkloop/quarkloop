import { NextResponse } from "next/server";

import {
  CreateAppThreadPluginArgs,
  DeleteAppThreadPluginArgs,
  GetAppThreadByIdPluginArgs,
  GetAppThreadByAppInstanceIdPluginArgs,
  UpdateAppThreadPluginArgs,
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
  GetAppThreadByIdPlugin,
  GetAppThreadByAppInstanceIdPlugin,
  CreateAppThreadPlugin,
  UpdateAppThreadPlugin,
  DeleteAppThreadPlugin,
} from "@quarkloop/plugins";

// GetAppThreadById
export async function GET(request: Request, { params }: { params: any }) {
  const { appInstanceId, threadId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetAppThreadByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        appInstanceId: appInstanceId,
        id: threadId,
      } as GetAppThreadByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// // GetAppThreadByAppInstanceId
// export async function GET(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { appInstanceId } = params;

//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetAppThreadByAppInstanceIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       appThread: {
//         appInstanceId: appInstanceId,
//       } as GetAppThreadByAppInstanceIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

// CreateAppThread
export async function POST(request: Request, { params }: { params: any }) {
  const { appInstanceId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(CreateAppThreadPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        ...body,
        appInstanceId: appInstanceId,
      } as CreateAppThreadPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// UpdateAppThread
export async function PUT(request: Request, { params }: { params: any }) {
  const { appInstanceId, threadId } = params;
  const body = await request.json();

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(UpdateAppThreadPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        ...body,
        appInstanceId: appInstanceId,
        id: threadId,
      } as UpdateAppThreadPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// DeleteAppThread
export async function DELETE(request: Request, { params }: { params: any }) {
  const { appInstanceId, threadId } = params;

  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(DeleteAppThreadPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appThread: {
        appInstanceId: appInstanceId,
        id: threadId,
      } as DeleteAppThreadPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
