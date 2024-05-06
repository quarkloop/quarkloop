"use client";

import { z } from "zod";
import { enpointApi } from "@quarkloop/lib";
import { apiResponseV2Schema, workspaceSchema } from "@quarkloop/types";

export const getNameSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.object({
            workspace: workspaceSchema,
        }),
    })
);
export type GetNameResponse = z.infer<typeof getNameSchema>;
export type GetNameArgs = {
    orgSid: string;
    workspaceSid?: string;
    serviceId?: string;
};

const endpoint = enpointApi
    .enhanceEndpoints({ addTagTypes: ["pagedata"] })
    .injectEndpoints({
        endpoints: (builder) => ({
            // getName endpoint
            getName: builder.query<GetNameResponse, GetNameArgs>({
                query: (q) => {
                    let url = "";
                    if (q.serviceId) {
                        url = `/workspaces/${q.workspaceSid}/services/${q.serviceId}`;
                    } else if (q.workspaceSid) {
                        url = `/workspaces/${q.workspaceSid}`;
                    } else if (q.orgSid) {
                        url = `/orgs/${q.orgSid}`;
                    }

                    return {
                        url: url,
                        method: "GET",
                    };
                },
                providesTags: ["pagedata"],
                extraOptions: {
                    responseSchema: getNameSchema,
                },
            }),
        }),
    });

export const { useGetNameQuery } = endpoint;
