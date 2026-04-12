import { apiClient } from "@/lib/api-client";

export interface UserDropdownOption {
  name: string;
  uuid: string;
}

export async function getUsers(): Promise<UserDropdownOption[]> {
  const response = await apiClient.get<UserDropdownOption[]>("/users");
  return response.data;
}
