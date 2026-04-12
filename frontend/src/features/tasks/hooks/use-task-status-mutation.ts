import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

import { updateTask } from "@/features/tasks/api";
import { getErrorMessage } from "@/lib/errors";
import { queryKeys } from "@/lib/query-keys";
import type { Task, TaskStatus } from "@/types/entities";

interface StatusMutationInput {
  taskId: string;
  status: TaskStatus;
}

interface RollbackContext {
  previousTaskLists: Array<[readonly unknown[], Task[] | undefined]>;
  previousProjectTasks: Task[] | undefined;
}

export function useOptimisticTaskStatusMutation(projectId: string) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ taskId, status }: StatusMutationInput) => updateTask(taskId, { status }),
    onMutate: async ({ taskId, status }) => {
      await queryClient.cancelQueries({ queryKey: queryKeys.tasks.byProject(projectId) });

      const previousTaskLists = queryClient.getQueriesData<Task[]>({
        queryKey: queryKeys.tasks.byProject(projectId),
      });

      const projectDetailKey = queryKeys.projects.detail(projectId);
      const previousProjectDetail = queryClient.getQueryData<{ tasks: Task[] }>(projectDetailKey);

      previousTaskLists.forEach(([queryKey, currentTasks]) => {
        if (!currentTasks) {
          return;
        }

        const filter = String(queryKey[2] ?? "all");
        const updated = currentTasks
          .map((task) => (task.id === taskId ? { ...task, status } : task))
          .filter((task) => (filter === "all" ? true : task.status === filter));

        queryClient.setQueryData(queryKey, updated);
      });

      if (previousProjectDetail?.tasks) {
        queryClient.setQueryData(projectDetailKey, {
          ...previousProjectDetail,
          tasks: previousProjectDetail.tasks.map((task) =>
            task.id === taskId ? { ...task, status } : task,
          ),
        });
      }

      return {
        previousTaskLists,
        previousProjectTasks: previousProjectDetail?.tasks,
      } satisfies RollbackContext;
    },
    onError: (error, _, context) => {
      context?.previousTaskLists.forEach(([queryKey, previousValue]) => {
        queryClient.setQueryData(queryKey, previousValue);
      });

      if (context?.previousProjectTasks) {
        queryClient.setQueryData(queryKeys.projects.detail(projectId), (previousDetail?: { tasks: Task[] }) => {
          if (!previousDetail) {
            return previousDetail;
          }

          return {
            ...previousDetail,
            tasks: context.previousProjectTasks,
          };
        });
      }

      toast.error(getErrorMessage(error));
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.tasks.byProject(projectId) });
      queryClient.invalidateQueries({ queryKey: queryKeys.projects.detail(projectId) });
    },
  });
}

