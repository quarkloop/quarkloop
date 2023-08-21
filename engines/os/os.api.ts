import { NextResponse } from "next/server";

import { DefaultErrorHandler, createPipeline } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";

import {
  CheckUserSessionPlugin,
  GetServerSessionPlugin,
} from "@/lib/core/server-session.plugin";
import { GetUserPlugin } from "@/engines/user";
import {
  GetPlanSubscriptionByUserSessionPlugin,
  UpdatePlanSubscriptionMetricsPlugin,
  UpdatePlanSubscriptionMetricsPluginArgs,
} from "@/engines/plan";
import {
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
} from "@/lib/core/core.plugins";

import {
  GetOperatingSystemByIdPluginArgs,
  GetOperatingSystemsByUserIdPluginArgs,
  CreateOperatingSystemPluginArgs,
  UpdateOperatingSystemPluginArgs,
  DeleteOperatingSystemPluginArgs,
  GetOperatingSystemByIdApiArgs,
  GetOperatingSystemUsersPluginArgs,
} from "./os.type";
import {
  GetOperatingSystemByIdPlugin,
  GetOperatingSystemsByUserIdPlugin,
  CreateOperatingSystemPlugin,
  UpdateOperatingSystemPlugin,
  DeleteOperatingSystemPlugin,
  GetOperatingSystemUsersPlugin,
  GetOperatingSystemApiResponsePlugin,
} from "./os.plugin";

/// GetOperatingSystemById
export async function GET_GetOperatingSystemById(
  request: Request,
  { params }: { params: any }
) {
  const { osId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    //.use(GetServerSessionPlugin)
    .use(CheckUserSessionPlugin)
    .use(GetUserPlugin)
    .use(GetOperatingSystemByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ os: { id: osId } as GetOperatingSystemByIdPluginArgs });

  return NextResponse.json(finalState.apiResponse);
}

// GetOperatingSystemUsers
export async function GET_GetOperatingSystemUsers(
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
    .use(GetUserPlugin)
    .use(GetOperatingSystemUsersPlugin)
    .use(GetOperatingSystemApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      os: {
        id: osId,
        workspaceId,
      } as GetOperatingSystemUsersPluginArgs,
    });

  return NextResponse.json(finalState.osApiReponse);
}

/// GetOperatingSystemsByUserId
export async function GET_GetOperatingSystemsByUserId(
  request: Request,
  { params }: { params: any }
) {
  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetOperatingSystemsByUserIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute();

  return NextResponse.json(finalState.apiResponse);
}

/// CreateOperatingSystem
export async function POST_CreateOperatingSystem(
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
    .use(GetUserPlugin)
    .use(GetPlanSubscriptionByUserSessionPlugin)
    .use(CreateOperatingSystemPlugin)
    //.use(UpdatePlanSubscriptionMetricsPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      os: { ...args } as CreateOperatingSystemPluginArgs,
      // planSubscriptionMetrics: {
      //   incOs: true,
      // } as UpdatePlanSubscriptionMetricsPluginArgs,
    });

  console.log("////////", finalState);

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateOperatingSystem
export async function PUT_UpdateOperatingSystem(
  request: Request,
  { params }: { params: any }
) {
  //const { osId: id } = params;
  const args = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(UpdateOperatingSystemPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ os: { ...args } as UpdateOperatingSystemPluginArgs });

  return NextResponse.json(finalState.apiResponse);
}

/// DeleteOperatingSystem
export async function PATCH_DeleteOperatingSystem(
  request: Request,
  { params }: { params: any }
) {
  const { osId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(DeleteOperatingSystemPlugin)
    //.use(UpdatePlanSubscriptionMetricsPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      os: { id: osId } as DeleteOperatingSystemPluginArgs,
      // planSubscriptionMetrics: {
      //   decOs: true,
      // } as UpdatePlanSubscriptionMetricsPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
