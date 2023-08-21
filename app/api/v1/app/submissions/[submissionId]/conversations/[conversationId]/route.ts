import {
  GET_GetAppConversationDataById,
  PUT_UpdateAppConversationData,
  PATCH_DeleteAppConversationData,
} from "@/engines/app-conversation/conversation-data.api";

export const dynamic = "force-dynamic";

export {
  GET_GetAppConversationDataById as GET,
  PUT_UpdateAppConversationData as PUT,
  PATCH_DeleteAppConversationData as PATCH,
};
