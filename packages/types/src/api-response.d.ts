export type ApiResponse = {
  status: StatusState;
  error?: string;
  errorDetails?: Record<string, any>;
  data?: {
    database?: DatabaseState;
  };
};
