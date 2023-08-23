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
  GetAppByIdPluginArgs,
  //GetAppsByWorkspaceIdPluginArgs,
  CreateAppPluginArgs,
  UpdateAppPluginArgs,
  DeleteAppPluginArgs,
  GetAppsByOsIdPluginArgs,
} from "./app.type";
import {
  GetAppByIdPlugin,
  GetAppsByOsIdPlugin,
  CreateAppPlugin,
  UpdateAppPlugin,
  DeleteAppPlugin,
} from "./app.plugin";
import {
  GetPlanSubscriptionByUserSessionPlugin,
  UpdatePlanSubscriptionMetricsPlugin,
  UpdatePlanSubscriptionMetricsPluginArgs,
} from "@/engines/plan";
import { GetUserPlugin } from "@/engines/user";

/// GetAppById
export async function GET_GetAppById(
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

/// GetAppsByOsId
export async function GET_GetAppsByOsId(
  request: Request,
  { params }: { params: any }
) {
  const { osId } = params;

  const { searchParams } = new URL(request.url);
  const workspaceId = searchParams.get("workspaceId");

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetAppsByOsIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        osId: osId,
        workspaceId: workspaceId,
      } as GetAppsByOsIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

// /// GetAppsByWorkspaceId
// export async function GET_GetAppsByWorkspaceId(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { osId, workspaceId } = params;

//   // pipeline
//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetServerSessionPlugin)
//     .use(GetAppsByWorkspaceIdPlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       app: {
//         osId: osId,
//         workspaceId: workspaceId,
//       } as GetAppsByWorkspaceIdPluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

/// CreateApp
export async function POST_CreateApp(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetPlanSubscriptionByUserSessionPlugin)
    .use(CreateAppPlugin)
    //.use(UpdatePlanSubscriptionMetricsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        ...body,
        osId: osId,
        workspaceId: workspaceId,
      } as CreateAppPluginArgs,
      // planSubscriptionMetrics: {
      //   incApp: true,
      // } as UpdatePlanSubscriptionMetricsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateApp
export async function PUT_UpdateApp(
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

/// DeleteApp
export async function PATCH_DeleteApp(
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
    .use(DeleteAppPlugin)
    //.use(UpdatePlanSubscriptionMetricsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      app: {
        osId: osId,
        workspaceId: workspaceId,
        id: appId,
      } as DeleteAppPluginArgs,
      // planSubscriptionMetrics: {
      //   decApp: true,
      // } as UpdatePlanSubscriptionMetricsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
