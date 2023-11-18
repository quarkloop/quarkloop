import {
    BaseQueryFn,
    FetchArgs,
    FetchBaseQueryError,
    FetchBaseQueryMeta,
    createApi,
    fetchBaseQuery,
    retry,
} from "@reduxjs/toolkit/query/react";
import { ZodSchema } from "zod";

type TBaseQuery = BaseQueryFn<
    string | FetchArgs,
    unknown,
    FetchBaseQueryError,
    { dataSchema?: ZodSchema },
    FetchBaseQueryMeta
>;

const baseQuery = fetchBaseQuery({
    baseUrl: `${process.env.NEXT_PUBLIC_API_ENDPOINT_URL}/api/v1/`,
    prepareHeaders: (headers, { getState }) => {
        // const token = (getState() as AppState).auth.token
        // if (token) {
        //     headers.set('authentication', `Bearer ${token}`)
        // }
        return headers;
    },
});

const baseQueryWithZodValidation: (baseQuery: TBaseQuery) => TBaseQuery =
    (baseQuery: TBaseQuery) => async (args, api, extraOptions) => {
        const returnValue = await baseQuery(args, api, extraOptions);

        const zodSchema = extraOptions?.dataSchema;
        const { data } = returnValue;

        if (data && zodSchema) {
            try {
                zodSchema.parse(data);
            } catch (error: any) {
                console.error(error);
                return { error };
            }
        }

        return returnValue;
    };

const baseQueryWithRetry = retry(baseQueryWithZodValidation(baseQuery), {
    maxRetries: 3,
});

export const enpointApi = createApi({
    reducerPath: "enpointApi",
    baseQuery: baseQueryWithRetry,
    endpoints: () => ({}),
});
