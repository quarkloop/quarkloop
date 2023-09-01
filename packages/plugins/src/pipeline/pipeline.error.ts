export interface PluginStatus {
    statusCode: number;
    statusCodeString: string;
    timestamp: Date;
    message?: string;
    details?: any;
}

function formatStatusLog(status: PluginStatus): string {
    let log = `[${status.timestamp.toLocaleDateString()} ${status.timestamp.toLocaleTimeString()}] `;
    log += `[STATUS] Code: ${status.statusCode} >> ${status.statusCodeString}`;

    if (status.message) {
        log += ` | Message: ${status.message} `;
    }

    if (status.details) {
        log += ` | Details: ${JSON.stringify(status.details)}\n`;
    }

    return log;
}

export type PluginStatusFactory = {
    (details?: any): PluginStatus;
}

type CreatePluginStatusProps = {
    statusCode: number;
    statusCodeString: string;
    defaultMessage: string;
}

function createPluginStatus(props: CreatePluginStatusProps): PluginStatusFactory {
    return (details?: any) => {
        const status = {
            ...props,
            timestamp: new Date(),
            details,
        };

        console.log(formatStatusLog(status));
        return status;
    };
}

export enum StatusCode {
    OK = 200,
    CREATED = 201,
    NO_CONTENT = 204,
    BAD_REQUEST = 400,
    UNAUTHORIZED = 401,
    FORBIDDEN = 403,
    NOT_FOUND = 404,
    METHOD_NOT_ALLOWED = 405,
    CONFLICT = 409,
    INTERNAL_SERVER_ERROR = 500,
    BAD_GATEWAY = 502,
    SERVICE_UNAVAILABLE = 503,
    GATEWAY_TIMEOUT = 504,
}

type StatusCodeStringValues = {
    [key in keyof typeof StatusCode]: string;
}

const STATUS_CODES: StatusCodeStringValues = {
    OK: "OK",
    CREATED: "CREATED",
    NO_CONTENT: "NO_CONTENT",
    BAD_REQUEST: "BAD_REQUEST",
    UNAUTHORIZED: "UNAUTHORIZED",
    FORBIDDEN: "FORBIDDEN",
    NOT_FOUND: "NOT_FOUND",
    METHOD_NOT_ALLOWED: "METHOD_NOT_ALLOWED",
    CONFLICT: "CONFLICT",
    INTERNAL_SERVER_ERROR: "INTERNAL_SERVER_ERROR",
    BAD_GATEWAY: "BAD_GATEWAY",
    SERVICE_UNAVAILABLE: "SERVICE_UNAVAILABLE",
    GATEWAY_TIMEOUT: "GATEWAY_TIMEOUT",
};

export const PluginStatusEntry = {
    OK: createPluginStatus({
        statusCode: StatusCode.OK,
        statusCodeString: STATUS_CODES.OK,
        defaultMessage: 'The request has succeeded.',
    }),
    CREATED: createPluginStatus({
        statusCode: StatusCode.CREATED,
        statusCodeString: STATUS_CODES.CREATED,
        defaultMessage: 'The request has been fulfilled and a new resource has been created.',
    }),
    NO_CONTENT: createPluginStatus({
        statusCode: StatusCode.NO_CONTENT,
        statusCodeString: STATUS_CODES.NO_CONTENT,
        defaultMessage: 'The request has been successfully processed, but no response body is needed.',
    }),
    BAD_REQUEST: createPluginStatus({
        statusCode: StatusCode.BAD_REQUEST,
        statusCodeString: STATUS_CODES.BAD_REQUEST,
        defaultMessage: 'The server cannot process the request due to a client error, such as invalid syntax or missing parameters.',
    }),
    UNAUTHORIZED: createPluginStatus({
        statusCode: StatusCode.UNAUTHORIZED,
        statusCodeString: STATUS_CODES.UNAUTHORIZED,
        defaultMessage: 'The request requires user authentication. Please provide valid credentials.',
    }),
    FORBIDDEN: createPluginStatus({
        statusCode: StatusCode.FORBIDDEN,
        statusCodeString: STATUS_CODES.FORBIDDEN,
        defaultMessage: 'The server understood the request, but the client does not have the necessary permissions to access the requested resource.',
    }),
    NOT_FOUND: createPluginStatus({
        statusCode: StatusCode.NOT_FOUND,
        statusCodeString: STATUS_CODES.NOT_FOUND,
        defaultMessage: 'The requested resource could not be found on the server.',
    }),
    METHOD_NOT_ALLOWED: createPluginStatus({
        statusCode: StatusCode.METHOD_NOT_ALLOWED,
        statusCodeString: STATUS_CODES.METHOD_NOT_ALLOWED,
        defaultMessage: 'The request method used is not supported for the requested resource.',
    }),
    CONFLICT: createPluginStatus({
        statusCode: StatusCode.CONFLICT,
        statusCodeString: STATUS_CODES.CONFLICT,
        defaultMessage: 'The request could not be completed due to a conflict with the current state of the target resource.',
    }),
    INTERNAL_SERVER_ERROR: createPluginStatus({
        statusCode: StatusCode.INTERNAL_SERVER_ERROR,
        statusCodeString: STATUS_CODES.INTERNAL_SERVER_ERROR,
        defaultMessage: 'The server encountered an unexpected condition that prevented it from fulfilling the request.',
    }),
    BAD_GATEWAY: createPluginStatus({
        statusCode: StatusCode.BAD_GATEWAY,
        statusCodeString: STATUS_CODES.BAD_GATEWAY,
        defaultMessage: 'The server acting as a gateway or proxy received an invalid response from an upstream server.',
    }),
    SERVICE_UNAVAILABLE: createPluginStatus({
        statusCode: StatusCode.SERVICE_UNAVAILABLE,
        statusCodeString: STATUS_CODES.SERVICE_UNAVAILABLE,
        defaultMessage: 'The server is currently unable to handle the request due to temporary overload or maintenance.',
    }),
    GATEWAY_TIMEOUT: createPluginStatus({
        statusCode: StatusCode.GATEWAY_TIMEOUT,
        statusCodeString: STATUS_CODES.GATEWAY_TIMEOUT,
        defaultMessage: 'The server acting as a gateway or proxy did not receive a timely response from an upstream server.',
    }),
};


// test

// describe('PluginStatus', () => {
//     it('should create an error object correctly', () => {
//         const err = errors.BAD_REQUEST('Invalid user input.');
//         expect(err).toBeDefined();
//         expect(err.statusCode).toBe(400);
//         expect(err.statusCodeString).toBe('BAD_REQUEST');
//         expect(err.details).toBeUndefined();
//         expect(err.message).toBe('Invalid user input.');
//     }),
//         it('should override default message with provided message', () => {
//             const err = errors.FORBIDDEN('Invalid Token.', 'Auth error');
//             expect(err.message).toBe('Invalid Token.');
//         }),
//         it('should include details field in error object if provided as an arg.', () => {
//             const err = errors.INTERNAL_SERVER_ERROR('Some error occurred', { statusCode: 500, reason: "Internal Status" });
//             expect(err.details).toEqual({ statusCode: 500, reason: "Internal Status" });
//         });
// });
