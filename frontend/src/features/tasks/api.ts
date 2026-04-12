import { apiClient } from "@/lib/api-client";
import type { Task, TaskPriority, TaskStatus } from "@/types/entities";

export interface TaskFilters {
  status?: TaskStatus | "all";
  assignee?: string;
}

export interface UpsertTaskPayload {
  title: string;
  description: string;
  status: TaskStatus;
  priority: TaskPriority;
  assignee_id?: string;
  due_date?: string;
}

export interface UpdateTaskPayload {
  title?: string;
  description?: string;
  status?: TaskStatus;
  priority?: TaskPriority;
  assignee_id?: string;
  due_date?: string;
}

export async function getProjectTasks(projectId: string, filters: TaskFilters): Promise<Task[]> {
  const response = await apiClient.get<Task[]>(`/projects/${projectId}/tasks`, {
    params: {
      status: filters.status && filters.status !== "all" ? filters.status : undefined,
      assignee: filters.assignee?.trim() ? filters.assignee.trim() : undefined,
    },
  });

  return response.data;
}

export async function createTask(projectId: string, payload: UpsertTaskPayload): Promise<Task> {
  const response = await apiClient.post<Task>(`/projects/${projectId}/tasks`, payload);
  return response.data;
}

export async function updateTask(taskId: string, payload: UpdateTaskPayload): Promise<Task> {
  const response = await apiClient.patch<Task>(`/tasks/${taskId}`, payload);
  return response.data;
}

export async function deleteTask(taskId: string): Promise<void> {
  await apiClient.delete(`/tasks/${taskId}`);
}

