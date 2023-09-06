import { NextResponse } from "next/server";

import {
    GetAppByIdPluginArgs,
    CreateAppPluginArgs,
    UpdateAppPluginArgs,
    DeleteAppPluginArgs,
} from "@quarkloop/types";
import { createPipeline } from "@quarkloop/plugin";
import {
    PipelineArgs,
    PipelineState,
    DefaultErrorHandler,
    GetApiResponsePlugin,
    CreateApiResponsePlugin,
    UpdateApiResponsePlugin,
    DeleteApiResponsePlugin,
    GetAppByIdPlugin,
    CreateAppPlugin,
    UpdateAppPlugin,
    DeleteAppPlugin,
} from "@quarkloop/plugins";

// GetAppById
export async function GET(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const id = searchParams.get("id");

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    if (id) {
        const finalState = await pipeline
            .use(GetAppByIdPlugin)
            .use(GetApiResponsePlugin)
            .onError(DefaultErrorHandler)
            .execute({
                app: {
                    id: id,
                } as GetAppByIdPluginArgs,
            });

        return NextResponse.json(finalState.apiResponse);
    }

    return NextResponse.json({ status: "Bad request" }, { status: 400 });
}

// CreateApp
export async function POST(request: Request, { params }: { params: any }) {
    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(CreateAppPlugin)
        .use(CreateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            app: {
                ...body,
            } as CreateAppPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse);
}

// UpdateApp
export async function PUT(request: Request, { params }: { params: any }) {
    const body = await request.json();

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(UpdateAppPlugin)
        .use(UpdateApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            app: {
                ...body,
            } as UpdateAppPluginArgs,
        });

    return NextResponse.json(finalState.apiResponse);
}

// DeleteApp
export async function DELETE(request: Request, { params }: { params: any }) {
    const { searchParams } = new URL(request.url);
    const id = searchParams.get("id");

    if (id == null) {
        return NextResponse.json({ status: "Bad request" }, { status: 400 });
    }

    const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
        initialState: {},
    });

    const finalState = await pipeline
        .use(DeleteAppPlugin)
        .use(DeleteApiResponsePlugin)
        .onError(DefaultErrorHandler)
        .execute({
            app: {
                id: id,
            } as DeleteAppPluginArgs,
        });

    // https://github.com/vercel/next.js/issues/49005#issuecomment-1708756641
    return new Response(null, {
        status: (finalState.apiResponse?.status.statusCode as number) ?? 500,
    });
}
