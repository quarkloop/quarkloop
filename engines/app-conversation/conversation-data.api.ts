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
  GetAppConversationDataByIdPlugin,
  GetAppConversationDataByAppSubmissionIdPlugin,
  CreateAppConversationDataPlugin,
  UpdateAppConversationDataPlugin,
  DeleteAppConversationDataPlugin,
} from "./conversation-data.plugin";

import { GetUserPlugin } from "@/engines/user";
import {
  CreateAppConversationDataPluginArgs,
  DeleteAppConversationDataPluginArgs,
  GetAppConversationDataByIdPluginArgs,
  GetAppConversationDataByAppSubmissionIdPluginArgs,
  UpdateAppConversationDataPluginArgs,
} from "./conversation-data.type";
import { GetPlanSubscriptionByUserSessionPlugin } from "@/engines/plan";

/// GetAppConversationDataById
export async function GET_GetAppConversationDataById(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId, conversationId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetAppConversationDataByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationData: {
        appSubmissionId: submissionId,
        id: conversationId,
      } as GetAppConversationDataByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetAppConversationDataByAppSubmissionId
export async function GET_GetAppConversationDataByAppSubmissionId(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(GetAppConversationDataByAppSubmissionIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationData: {
        appSubmissionId: submissionId,
      } as GetAppConversationDataByAppSubmissionIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// CreateAppConversationData
export async function POST_CreateAppConversationData(
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
    .use(GetPlanSubscriptionByUserSessionPlugin)
    .use(CreateAppConversationDataPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationData: {
        ...body,
        appSubmissionId: submissionId,
      } as CreateAppConversationDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateAppConversationData
export async function PUT_UpdateAppConversationData(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId, conversationId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(UpdateAppConversationDataPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationData: {
        ...body,
        appSubmissionId: submissionId,
        id: conversationId,
      } as UpdateAppConversationDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// DeleteAppConversationData
export async function PATCH_DeleteAppConversationData(
  request: Request,
  { params }: { params: any }
) {
  const { submissionId, conversationId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(DeleteAppConversationDataPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      appConversationData: {
        appSubmissionId: submissionId,
        id: conversationId,
      } as DeleteAppConversationDataPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
