import { NextResponse } from "next/server";

import {
    GetAppPageByIdPluginArgs,
    CreateAppPagePluginArgs,
    DeleteAppPagePluginArgs,
    UpdateAppPagePluginArgs,
    GetAppPageByAppInstanceIdPluginArgs,
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
    GetAppPageByIdPlugin,
    CreateAppPagePlugin,
    UpdateAppPagePlugin,
    DeleteAppPagePlugin,
    GetAppPageByAppInstanceIdPlugin,
} from "@quarkloop/plugins";

// GetAppPageById
// GetAppPageByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const pageId = searchParams.get("pageId");

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    if (appId && instanceId && pageId) {
        const finalState = await pipeline
            .use(GetAppPageByIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appPage: {
                    appId,
                    instanceId,
                    pageId,
                } as GetAppPageByIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    if (appId && instanceId) {
        const finalState = await pipeline
            .use(GetAppPageByAppInstanceIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appPage: {
                    appId,
                    instanceId,
                } as GetAppPageByAppInstanceIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppPage
export async function POST(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(CreateAppPagePlugin)
        .use(CreateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appPage: {
                appId,
                instanceId,
                page: body,
            } as CreateAppPagePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// UpdateAppPage
export async function PUT(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(UpdateAppPagePlugin)
        .use(UpdateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appPage: {
                appId,
                instanceId,
                page: body,
            } as UpdateAppPagePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// DeleteAppPage
export async function DELETE(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const pageId = searchParams.get("pageId");

    if (appId == null || instanceId == null || pageId == null) {
        return NextResponse.json({ status: "Bad request" }, { status: 400 });
    }

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(DeleteAppPagePlugin)
        .use(DeleteApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appPage: {
                appId,
                instanceId,
                pageId,
            } as DeleteAppPagePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}
