import axios from "axios";

import type { ApiError } from "@/types/api";

export function getErrorMessage(error: unknown): string {
  if (axios.isAxiosError<ApiError>(error)) {
    return error.response?.data?.error ?? error.message;
  }

  if (error instanceof Error) {
    return error.message;
  }

  return "Something went wrong. Please try again.";
}

export function getFieldErrors(error: unknown): Record<string, string> | undefined {
  if (!axios.isAxiosError<ApiError>(error)) {
    return undefined;
  }

  return error.response?.data?.fields;
}

