import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

import {
  createTask,
  deleteTask,
  getProjectTasks,
  updateTask,
  type TaskFilters,
  type UpsertTaskPayload,
  type UpdateTaskPayload,
} from "@/features/tasks/api";
import { queryKeys } from "@/lib/query-keys";

export function useProjectTasksQuery(projectId: string, filters: TaskFilters) {
  const status = filters.status ?? "all";
  const assignee = filters.assignee?.trim() ?? "";

  return useQuery({
    queryKey: queryKeys.tasks.list(projectId, status, assignee),
    queryFn: () => getProjectTasks(projectId, { status, assignee }),
    enabled: Boolean(projectId),
  });
}

export function useTaskUpsertMutation(projectId: string, taskId?: string) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (payload: UpsertTaskPayload) => {
      if (taskId) {
        return updateTask(taskId, payload satisfies UpdateTaskPayload);
      }

      return createTask(projectId, payload);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.tasks.byProject(projectId) });
      queryClient.invalidateQueries({ queryKey: queryKeys.projects.detail(projectId) });
    },
  });
}

export function useDeleteTaskMutation(projectId: string) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: deleteTask,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.tasks.byProject(projectId) });
      queryClient.invalidateQueries({ queryKey: queryKeys.projects.detail(projectId) });
    },
  });
}

