import { apiClient } from "@/lib/api-client";

interface RegisterPayload {
  name: string;
  email: string;
  password: string;
}

interface LoginPayload {
  email: string;
  password: string;
}

interface LoginResponse {
  token: string;
}

export async function register(payload: RegisterPayload): Promise<void> {
  await apiClient.post("/auth/register", payload);
}

export async function login(payload: LoginPayload): Promise<LoginResponse> {
  const response = await apiClient.post<LoginResponse>("/auth/login", payload);
  return response.data;
}

