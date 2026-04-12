import { apiClient } from "@/lib/api-client";
import type { Project, ProjectDetail } from "@/types/entities";

interface CreateProjectPayload {
  name: string;
  description: string;
}

interface UpdateProjectPayload {
  name?: string;
  description?: string;
}

export async function getProjects(): Promise<Project[]> {
  const response = await apiClient.get<Project[]>("/projects");
  return response.data;
}

export async function createProject(payload: CreateProjectPayload): Promise<Project> {
  const response = await apiClient.post<Project>("/projects", payload);
  return response.data;
}

export async function getProjectById(projectId: string): Promise<ProjectDetail> {
  const response = await apiClient.get<ProjectDetail>(`/projects/${projectId}`);
  return response.data;
}

export async function updateProject(projectId: string, payload: UpdateProjectPayload): Promise<Project> {
  const response = await apiClient.patch<Project>(`/projects/${projectId}`, payload);
  return response.data;
}

export async function deleteProject(projectId: string): Promise<void> {
  await apiClient.delete(`/projects/${projectId}`);
}

