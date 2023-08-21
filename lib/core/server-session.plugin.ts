import { getServerSession } from "next-auth/next";
import { cookies } from "next/headers";

import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { createPlugin } from "@/lib/pipeline";
import { PipelineState, PluginStatusEntry } from "@/lib/core/pipeline";

import { GetServerSessionPluginArgs } from "@/engines/user/user.type";

/// GetServerSession Plugin
export const GetServerSessionPlugin = createPlugin<
  PipelineState,
  GetServerSessionPluginArgs[]
>({
  name: "GetServerSessionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    const session = await getServerSession(authOptions);

    if (session == null || session.user?.email == null) {
      return { ...state, status: PluginStatusEntry.UNAUTHORIZED() };
    }

    const authToken = cookies().get(process.env.NEXTAUTH_SESSION_TOKEN);
    if (authToken == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetServerSessionPlugin] Auth token check failed."
        ),
      };
    }

    return {
      ...state,
      session: {
        session: session,
        sessionToken: authToken.value,
      },
    };
  },
});

/// CheckUserSession Plugin
export const CheckUserSessionPlugin = createPlugin<
  PipelineState,
  GetServerSessionPluginArgs[]
>({
  name: "CheckUserSessionPlugin",
  config: {},
  handler: async (state, config, ...args): Promise<PipelineState> => {
    const session = await getServerSession(authOptions);

    if (session == null || session.user?.email == null) {
      return state;
    }

    const authToken = cookies().get(process.env.NEXTAUTH_SESSION_TOKEN);
    if (authToken == null) {
      return {
        ...state,
        status: PluginStatusEntry.INTERNAL_SERVER_ERROR(
          "[GetServerSessionPlugin] Auth token check failed."
        ),
      };
    }

    return {
      ...state,
      session: {
        session: session,
        sessionToken: authToken.value,
      },
    };
  },
});
