import { NextResponse } from "next/server";

import { DefaultErrorHandler, createPipeline } from "@/lib/pipeline";
import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";

import {
  GetUserPluginArgs,
  DeleteUserPluginArgs,
  UpdateUserPluginArgs,
  GetUserLinkedAccountsPluginArgs,
  GetUserSessionsPluginArgs,
  TerminateUserSessionPluginArgs,
} from "./user.type";
import {
  GetUserPlugin,
  DeleteUserPlugin,
  UpdateUserPlugin,
  GetUserLinkedAccountsPlugin,
  GetUserSessionsPlugin,
  TerminateUserSessionPlugin,

  // RegisterUserPlugin,
  // AuthenticateUserPlugin,
  // ChangePasswordPlugin,
  // ForgotPasswordPlugin,
  // CreateUserGroupPlugin,
  // GetUserGroupPlugin,
  // JoinUserGroupPlugin,
  // LeaveUserGroupPlugin,
} from "./user.plugin";

import {
  GetApiResponsePlugin,
  CreateApiResponsePlugin,
  UpdateApiResponsePlugin,
  DeleteApiResponsePlugin,
} from "@/lib/core/core.plugins";
import { GetServerSessionPlugin } from "@/lib/core/server-session.plugin";

/// GetUser
export async function GET_GetUser(
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
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute();

  return NextResponse.json(finalState.apiResponse);
}

/// UpdateUser
export async function PUT_UpdateUser(
  request: Request,
  { params }: { params: any }
) {
  const args: UpdateUserPluginArgs = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(UpdateUserPlugin)
    .use(UpdateApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ user: args });

  return NextResponse.json(finalState.apiResponse);
}

/// DeleteUser
export async function PATCH_DeleteUser(
  request: Request,
  { params }: { params: any }
) {
  const args: DeleteUserPluginArgs = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(DeleteUserPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ user: args });

  return NextResponse.json(finalState.apiResponse);
}

/// GetUserLinkedAccounts
export async function GET_GetUserLinkedAccounts(
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
    .use(GetUserLinkedAccountsPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute();

  return NextResponse.json(finalState.apiResponse);
}

/// GetUserSessions
export async function GET_GetUserSessions(
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
    .use(GetUserSessionsPlugin)
    .use(GetApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute();

  return NextResponse.json(finalState.apiResponse);
}

/// TerminateUserSession
export async function PATCH_TerminateUserSession(
  request: Request,
  { params }: { params: any }
) {
  const args: TerminateUserSessionPluginArgs = await request.json();

  // pipeline
  const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    initialState: {},
  });

  const finalState = await pipeline
    .use(GetServerSessionPlugin)
    .use(GetUserPlugin)
    .use(TerminateUserSessionPlugin)
    .use(DeleteApiResponsePlugin)
    .onError(DefaultErrorHandler)
    .execute({ userSession: args });

  return NextResponse.json(finalState.apiResponse);
}
