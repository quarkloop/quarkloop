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
  GetAdminAppSubmissionByIdPluginArgs,
  GetAppSubmissionByIdPluginArgs,
  CreateAppSubmissionPluginArgs,
  GetAppSubmissionsByAppIdPluginArgs,
  UpdateAppSubmissionPluginArgs,
} from "./app-submission.type";
import {
  GetAdminAppSubmissionByIdPlugin,
  GetAppSubmissionByIdPlugin,
  CreateAppSubmissionPlugin,
  GetAppSubmissionsByAppIdPlugin,
  UpdateAppSubmissionPlugin,
} from "./app-submission.plugin";
import {
  GetPlanSubscriptionByUserSessionPlugin,
  UpdatePlanSubscriptionMetricsPlugin,
  UpdatePlanSubscriptionMetricsPluginArgs,
} from "@/engines/plan";
import { GetUserPlugin } from "@/engines/user";

/// GetAdminAppSubmissionById
export async function GET_GetAdminAppSubmissionById(
  request: Request,
  { params }: { params: any }
) {
  const { osId, workspaceId, appId, id } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetAdminAppSubmissionByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appSubmission: {
        osId: osId,
        workspaceId: workspaceId,
        appId: appId,
        id: id,
      } as GetAdminAppSubmissionByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetAppSubmissionById
export async function GET_GetAppSubmissionById(
  request: Request,
  { params }: { params: any }
) {
  const { userId, appId, id } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetAppSubmissionByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appSubmission: {
        userId: userId,
        appId: appId,
        id: id,
      } as GetAppSubmissionByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetAppSubmissionsByAppId
export async function GET_GetAppSubmissionsByAppId(
  request: Request,
  { params }: { params: any }
) {
  const { searchParams } = new URL(request.url);

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetAppSubmissionsByAppIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appSubmission: {
        osId: searchParams.get("osId"),
        workspaceId: searchParams.get("workspaceId"),
        appId: searchParams.get("appId"),
      } as GetAppSubmissionsByAppIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// CreateAppSubmission
export async function POST_CreateAppSubmission(
  request: Request,
  { params }: { params: any }
) {
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(CreateAppSubmissionPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appSubmission: body as CreateAppSubmissionPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateAppSubmission
export async function PUT_UpdateAppSubmission(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(UpdateAppSubmissionPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appSubmission: {
        ...body,
        id: submissionId,
      } as UpdateAppSubmissionPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
