import { NextResponse } from "next/server";

import {
    GetAppFormByIdPluginArgs,
    CreateAppFormPluginArgs,
    DeleteAppFormPluginArgs,
    UpdateAppFormPluginArgs,
    GetAppFormByAppInstanceIdPluginArgs,
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
    GetAppFormByIdPlugin,
    CreateAppFormPlugin,
    UpdateAppFormPlugin,
    DeleteAppFormPlugin,
    GetAppFormByAppInstanceIdPlugin,
} from "@quarkloop/plugins";

// GetAppFormById
// GetAppFormByAppInstanceId
export async function GET(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const formId = searchParams.get("formId");

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    if (appId && instanceId && formId) {
        const finalState = await pipeline
            .use(GetAppFormByIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appForm: {
                    appId,
                    instanceId,
                    formId,
                } as GetAppFormByIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    if (appId && instanceId) {
        const finalState = await pipeline
            .use(GetAppFormByAppInstanceIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                appForm: {
                    appId,
                    instanceId,
                } as GetAppFormByAppInstanceIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse, {
            status:
                (finalState.apiResponse?.status.statusCode as number) ?? 500,
        });
    }

    return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateAppForm
export async function POST(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(CreateAppFormPlugin)
        .use(CreateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appForm: {
                appId,
                instanceId,
                form: body,
            } as CreateAppFormPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// UpdateAppForm
export async function PUT(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");

    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(UpdateAppFormPlugin)
        .use(UpdateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appForm: {
                appId,
                instanceId,
                form: body,
            } as UpdateAppFormPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}

// DeleteAppForm
export async function DELETE(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const appId = searchParams.get("appId");
    const instanceId = searchParams.get("instanceId");
    const formId = searchParams.get("formId");

    if (appId == null || instanceId == null || formId == null) {
        return NextResponse.json({ status: "Bad request" }, { status: 400 });
    }

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(DeleteAppFormPlugin)
        .use(DeleteApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            appForm: {
                appId,
                instanceId,
                formId,
            } as DeleteAppFormPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}
