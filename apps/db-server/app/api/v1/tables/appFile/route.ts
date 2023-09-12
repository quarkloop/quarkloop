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
    GetAppFileByIdPlugin,
    CreateAppFilePlugin,
    UpdateAppFilePlugin,
    DeleteAppFilePlugin,
    GetAppFileByAppInstanceIdPlugin,
} from "@quarkloop/plugins";

// GetAppFileById
// GetAppFileByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const id = searchParams.get("id");
    const appInstanceId = searchParams.get("instanceId");

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    if (id) {
        const finalState = await pipeline
            .use(GetAppFileByIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appFile: {
                    id: id,
                } as GetAppFileByIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    if (appInstanceId) {
        const finalState = await pipeline
            .use(GetAppFileByAppInstanceIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appFile: {
                    appInstanceId: appInstanceId,
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
                ...body,
            } as CreateAppFilePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// UpdateAppFile
export async function PUT(request: Request, { params }: { params: any }) {
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
                ...body,
            } as UpdateAppFilePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// DeleteAppFile
export async function DELETE(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const id = searchParams.get("id");
    const appInstanceId = searchParams.get("appInstanceId");

    if (id == null && appInstanceId == null) {
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
                id: id,
                appInstanceId: appInstanceId,
            } as DeleteAppFilePluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}
