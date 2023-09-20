import { NextResponse } from "next/server";

import {
    GetAppFileByIdPluginArgs,
    CreateAppFilePluginArgs,
    DeleteAppFilePluginArgs,
    UpdateAppFilePluginArgs,
    GetAppFileByAppInstanceIdPluginArgs,
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
    GetAppFileByAppInstanceIdPlugin,
    GetAppFileByIdPlugin,
    CreateAppFilePlugin,
    UpdateAppFilePlugin,
    DeleteAppFilePlugin,
} from "@quarkloop/plugins";

// GetAppFileById
// GetAppFileByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const fileId = searchParams.get("fileId");

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    if (appId && instanceId && fileId) {
        const finalState = await pipeline
            .use(GetAppFileByIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appFile: {
                    appId,
                    instanceId,
                    fileId,
                } as GetAppFileByIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    if (appId && instanceId) {
        const finalState = await pipeline
            .use(GetAppFileByAppInstanceIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appFile: {
                    appId,
                    instanceId,
                } as GetAppFileByAppInstanceIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppFile
export async function POST(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(CreateAppFilePlugin)
        .use(CreateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appFile: {
                appId,
                instanceId,
                file: body,
            } as CreateAppFilePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// UpdateAppFile
export async function PUT(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(UpdateAppFilePlugin)
        .use(UpdateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appFile: {
                appId,
                instanceId,
                file: body,
            } as UpdateAppFilePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// DeleteAppFile
export async function DELETE(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const fileId = searchParams.get("fileId");

    if (appId == null || instanceId == null || fileId == null) {
        return NextResponse.json({ status: "Bad request" }, { status: 400 });
    }

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(DeleteAppFilePlugin)
        .use(DeleteApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appFile: {
                appId,
                instanceId,
                fileId,
            } as DeleteAppFilePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}
