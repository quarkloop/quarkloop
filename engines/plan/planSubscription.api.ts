import { NextResponse } from "next/server";

import { DefaultErrorHandler, createPipeline } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";

import { GetServerSessionPlugin } from "@/lib/core/server-session.plugin";
import { GetUserPlugin } from "@/engines/user";
import {
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
} from "@/lib/core/core.plugins";

import {
  GetPlanSubscriptionByIdPluginArgs,
  CreatePlanSubscriptionPluginArgs,
  UpdatePlanSubscriptionPluginArgs,
  DeletePlanSubscriptionPluginArgs,
  GetPlanSubscriptionByUserSessionPluginArgs,
  GetPlansPluginArgs,
} from "./planSubscription.type";
import {
  GetPlanSubscriptionByIdPlugin,
  CreatePlanSubscriptionPlugin,
  UpdatePlanSubscriptionPlugin,
  DeletePlanSubscriptionPlugin,
  GetPlanSubscriptionByUserSessionPlugin,
  GetPlansPlugin,
} from "./planSubscription.plugin";

/// GET_GetPlans
export async function GET_GetPlans(
  request: Request,
  { params }: { params: any }
) {
  const { searchParams } = new URL(request.url);
  const planType = searchParams.get("planType");

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetPlansPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      plan: {
        planType,
      } as GetPlansPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetPlanSubscriptionById
export async function GET_GetPlanSubscriptionById(
  request: Request,
  { params }: { params: any }
) {
  const { id } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetPlanSubscriptionByIdPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      planSubscription: {
        id,
      } as GetPlanSubscriptionByIdPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// GetPlanSubscriptionByUserSession
export async function GET_GetPlanSubscriptionByUserSession(
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
    .use(GetPlanSubscriptionByUserSessionPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute();

  return NextResponse.json(finalState.apiResponse);
}

/// CreatePlanSubscription
export async function POST_CreatePlanSubscription(
  request: Request,
  { params }: { params: any }
) {
  const { planId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(CreatePlanSubscriptionPlugin)
    .use(CreateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      planSubscription: {
        ...body,
        planId,
      } as CreatePlanSubscriptionPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// UpdatePlanSubscription
export async function PUT_UpdatePlanSubscription(
  request: Request,
  { params }: { params: any }
) {
  const { id, planId } = params;
  const body = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(UpdatePlanSubscriptionPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      planSubscription: {
        ...body,
        id,
        planId,
      } as UpdatePlanSubscriptionPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}

/// DeletePlanSubscription
export async function PATCH_DeletePlanSubscription(
  request: Request,
  { params }: { params: any }
) {
  const { id, planId } = params;

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(DeletePlanSubscriptionPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({
      planSubscription: {
        id,
        planId,
      } as DeletePlanSubscriptionPluginArgs,
    });

  return NextResponse.json(finalState.apiResponse);
}
