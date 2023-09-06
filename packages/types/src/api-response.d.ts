export type ApiResponse = {
    status: StatusState;
    error?: string;
    errorDetails?: Record<string, any>;
    database?: DatabaseState;
};
