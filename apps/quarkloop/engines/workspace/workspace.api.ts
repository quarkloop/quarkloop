import { NextResponse } from "next/server";

import { DefaultErrorHandler, createPipeline } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";

import {
  GetWorkspaceByIdPluginArgs,
  GetWorkspacesByOsIdPluginArgs,
  CreateWorkspacePluginArgs,
  UpdateWorkspacePluginArgs,
  DeleteWorkspacePluginArgs,
  //GetWorkspaceByNamePluginArgs,
} from "./workspace.type";

import { GetServerSessionPlugin } from "@/lib/core/server-session.plugin";

import {
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
} from "@/lib/core/core.plugins";

import {
  GetWorkspaceByIdPlugin,
  GetWorkspacesByOsIdPlugin,
  CreateWorkspacePlugin,
  UpdateWorkspacePlugin,
  DeleteWorkspacePlugin,
  //GetWorkspaceByNamePlugin,
} from "./workspace.plugin";
import {
  GetPlanSubscriptionByUserSessionPlugin,
  UpdatePlanSubscriptionMetricsPlugin,
  UpdatePlanSubscriptionMetricsPluginArgs,
} from "@/engines/plan";
import { GetUserPlugin } from "@/engines/user";

/// GetWorkspaceById
export async function GET_GetWorkspaceById(
  request: Request,
  { params }: { params: any }
) {
  const args: GetWorkspaceByIdPluginArgs = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetWorkspaceByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ workspace: args });

  return NextResponse.json(finalState.apiResponse);
}

// /// GetWorkspaceByName
// export async function GET_GetWorkspaceByName(
//   request: Request,
//   { params }: { params: any }
// ) {
//   const { osId } = params;
//   const { searchParams } = new URL(request.url);

//   // pipeline
//   const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
//     initialState: {},
//   });

//   const finalState = await pipeline
//     .use(GetServerSessionPlugin)
//     .use(GetWorkspaceByNamePlugin)
//     .use(GetApiResponsePlugin)
//     .onError(DefaultErrorHandler)
//     .execute({
//       workspace: {
//         osId: osId,
//         name: searchParams.get("name"),
//       } as GetWorkspaceByNamePluginArgs,
//     });

//   return NextResponse.json(finalState.apiResponse);
// }

/// GetWorkspacesByOsId
export async function GET_GetWorkspacesByOsId(
  request: Request,
  { params }: { params: any }
) {
  const args: GetWorkspacesByOsIdPluginArgs = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetPlanSubscriptionByUserSessionPlugin)
    .use(GetWorkspacesByOsIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ workspace: args });

  return NextResponse.json(finalState.apiResponse);
}

/// CreateWorkspace
export async function POST_CreateWorkspace(
  request: Request,
  { params }: { params: any }
) {
  const args: CreateWorkspacePluginArgs = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetPlanSubscriptionByUserSessionPlugin)
    .use(CreateWorkspacePlugin)
    //.use(UpdatePlanSubscriptionMetricsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      workspace: args,
      // planSubscriptionMetrics: {
      //   incWorkspace: true,
      // } as UpdatePlanSubscriptionMetricsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateWorkspace
export async function PUT_UpdateWorkspace(
  request: Request,
  { params }: { params: any }
) {
  const args = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(UpdateWorkspacePlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      workspace: args as UpdateWorkspacePluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// DeleteWorkspace
export async function PATCH_DeleteWorkspace(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(DeleteWorkspacePlugin)
    //.use(UpdatePlanSubscriptionMetricsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      workspace: {
        id: workspaceId,
        osId: osId,
      } as DeleteWorkspacePluginArgs,
      // planSubscriptionMetrics: {
      //   decWorkspace: true,
      // } as UpdatePlanSubscriptionMetricsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
