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
    GetAppThreadByAppInstanceIdPlugin,
    GetAppThreadByIdPlugin,
    CreateAppThreadPlugin,
    UpdateAppThreadPlugin,
    DeleteAppThreadPlugin,
} from "@quarkloop/plugins";

// GetAppThreadById
// GetAppThreadByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const threadId = searchParams.get("threadId");

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    if (appId && instanceId && threadId) {
        const finalState = await pipeline
            .use(GetAppThreadByIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appThread: {
                    appId,
                    instanceId,
                    threadId,
                } as GetAppThreadByIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    if (appId && instanceId) {
        const finalState = await pipeline
            .use(GetAppThreadByAppInstanceIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appThread: {
                    appId,
                    instanceId,
                } as GetAppThreadByAppInstanceIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppThread
export async function POST(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

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
                appId,
                instanceId,
                thread: body,
            } as CreateAppThreadPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// UpdateAppThread
export async function PUT(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

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
                appId,
                instanceId,
                thread: body,
            } as UpdateAppThreadPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// DeleteAppThread
export async function DELETE(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const threadId = searchParams.get("threadId");

    if (appId == null || instanceId == null || threadId == null) {
        return NextResponse.json({ status: "Bad request" }, { status: 400 });
    }

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(DeleteAppThreadPlugin)
        .use(DeleteApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appThread: {
                appId,
                instanceId,
                threadId,
            } as DeleteAppThreadPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}
