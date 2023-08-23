import {
    GET_GetUser,
    PUT_UpdateUser,
    PATCH_DeleteUser,
} from "@/engines/user/user.api";

export const dynamic = "force-dynamic";

export {
    GET_GetUser as GET,
    PUT_UpdateUser as PUT,
    PATCH_DeleteUser as PATCH,
};