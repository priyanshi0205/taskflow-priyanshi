import { useMemo, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { ArrowLeft, LoaderCircle, Pencil, Plus, Trash2 } from "lucide-react";
import { toast } from "sonner";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { Skeleton } from "@/components/ui/skeleton";
import { useAuth } from "@/features/auth/use-auth";
import { ProjectEditDialog } from "@/features/projects/components/project-edit-dialog";
import { useDeleteProjectMutation, useProjectDetailQuery } from "@/features/projects/hooks";
import { TaskCard } from "@/features/tasks/components/task-card";
import { TaskEmptyState } from "@/features/tasks/components/task-empty-state";
import { TaskFormModal } from "@/features/tasks/components/task-form-modal";
import { useDeleteTaskMutation, useProjectTasksQuery } from "@/features/tasks/hooks/use-task-mutations";
import { useOptimisticTaskStatusMutation } from "@/features/tasks/hooks/use-task-status-mutation";
import { getErrorMessage } from "@/lib/errors";
import type { Task, TaskStatus } from "@/types/entities";

type TaskFilter = "all" | TaskStatus;

function ProjectDetailSkeleton(): React.JSX.Element {
  return (
    <div className="space-y-6">
      <Skeleton className="h-9 w-64" />
      <div className="grid gap-6 xl:grid-cols-[2fr_1fr]">
        <div className="space-y-4">
          <Skeleton className="h-16 w-full" />
          <Skeleton className="h-48 w-full" />
          <Skeleton className="h-48 w-full" />
        </div>
        <Skeleton className="h-80 w-full" />
      </div>
    </div>
  );
}

export function ProjectDetailPage(): React.JSX.Element {
  const navigate = useNavigate();
  const { user } = useAuth();
  const { id } = useParams<{ id: string }>();

  const [statusFilter, setStatusFilter] = useState<TaskFilter>("all");
  const [assigneeFilter, setAssigneeFilter] = useState("");
  const [isTaskFormOpen, setIsTaskFormOpen] = useState(false);
  const [selectedTask, setSelectedTask] = useState<Task | undefined>(undefined);
  const [isEditProjectOpen, setIsEditProjectOpen] = useState(false);

  const projectId = id ?? "";

  const projectQuery = useProjectDetailQuery(projectId);
  const tasksQuery = useProjectTasksQuery(projectId, {
    status: statusFilter,
    assignee: assigneeFilter,
  });

  const statusMutation = useOptimisticTaskStatusMutation(projectId);
  const deleteTaskMutation = useDeleteTaskMutation(projectId);
  const deleteProjectMutation = useDeleteProjectMutation();

  const canManageProject = user?.id === projectQuery.data?.owner_id;

  const taskStats = useMemo(() => {
    const tasks = projectQuery.data?.tasks ?? [];

    return {
      total: tasks.length,
      todo: tasks.filter((task) => task.status === "todo").length,
      inProgress: tasks.filter((task) => task.status === "in_progress").length,
      done: tasks.filter((task) => task.status === "done").length,
    };
  }, [projectQuery.data?.tasks]);

  const openCreateTaskForm = (): void => {
    setSelectedTask(undefined);
    setIsTaskFormOpen(true);
  };

  const openEditTaskForm = (task: Task): void => {
    const canManageTask = canManageProject || task.created_by_id === user?.id;
    if (!canManageTask) {
      toast.error("Only project owners or task creators can edit tasks");
      return;
    }

    setSelectedTask(task);
    setIsTaskFormOpen(true);
  };

  const handleStatusChange = (taskId: string, status: TaskStatus): void => {
    const task = (projectQuery.data?.tasks ?? []).find((item) => item.id == taskId);
    const canManageTask = canManageProject || task?.created_by_id === user?.id;
    if (!canManageTask) {
      toast.error("Only project owners or task creators can update task status");
      return;
    }

    statusMutation.mutate({ taskId, status });
  };

  const handleDeleteTask = async (taskId: string): Promise<void> => {
    const task = (projectQuery.data?.tasks ?? []).find((item) => item.id == taskId);
    const canManageTask = canManageProject || task?.created_by_id === user?.id;
    if (!canManageTask) {
      toast.error("Only project owners or task creators can delete tasks");
      return;
    }

    try {
      await deleteTaskMutation.mutateAsync(taskId);
      toast.success("Task deleted");
    } catch (error) {
      toast.error(getErrorMessage(error));
    }
  };

  const handleDeleteProject = async (): Promise<void> => {
    if (!projectQuery.data) {
      return;
    }

    const confirmed = window.confirm("Are you sure you want to delete this project?");
    if (!confirmed) {
      return;
    }

    try {
      await deleteProjectMutation.mutateAsync(projectQuery.data.id);
      toast.success("Project deleted");
      navigate("/projects", { replace: true });
    } catch (error) {
      toast.error(getErrorMessage(error));
    }
  };

  if (!projectId) {
    return (
      <Card className="p-8 text-center">
        <CardTitle>Invalid project id</CardTitle>
      </Card>
    );
  }

  if (projectQuery.isLoading) {
    return <ProjectDetailSkeleton />;
  }

  if (projectQuery.isError || !projectQuery.data) {
    return (
      <Card className="p-8 text-center">
        <CardTitle className="mb-2">Could not load project</CardTitle>
        <CardDescription>
          {projectQuery.error ? getErrorMessage(projectQuery.error) : "Project not found"}
        </CardDescription>
        <Button className="mx-auto mt-4" onClick={() => navigate("/projects")}>
          Back to projects
        </Button>
      </Card>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex flex-wrap items-center justify-between gap-3">
        <Button variant="outline" onClick={() => navigate("/projects")}>
          <ArrowLeft className="h-4 w-4" />
          Back
        </Button>

        <div className="flex flex-wrap items-center gap-2">
          {canManageProject ? (
            <>
              <Button variant="secondary" onClick={() => setIsEditProjectOpen(true)}>
                <Pencil className="h-4 w-4" />
                Edit project
              </Button>
              <Button
                variant="destructive"
                onClick={handleDeleteProject}
                disabled={deleteProjectMutation.isPending}
              >
                <Trash2 className="h-4 w-4" />
                Delete project
              </Button>
            </>
          ) : (
            <Badge variant="secondary">Read-only</Badge>
          )}

          <Button onClick={openCreateTaskForm}>
            <Plus className="h-4 w-4" />
            Add task
          </Button>
        </div>
      </div>

      <section className="space-y-2">
        <h1 className="text-3xl font-bold">{projectQuery.data.name}</h1>
        <p className="max-w-2xl text-sm text-muted-foreground">
          {projectQuery.data.description || "No project description was provided."}
        </p>
      </section>

      <div className="grid gap-6 xl:grid-cols-[2fr_1fr]">
        <section className="space-y-4">
          <Card>
            <CardContent className="flex flex-col gap-3 p-4 sm:flex-row sm:items-center sm:justify-between">
              <div>
                <p className="text-sm font-medium">Task filter</p>
                <p className="text-xs text-muted-foreground">Switch between statuses while keeping updates instant.</p>
              </div>

              <div className="grid w-full grid-cols-1 gap-2 sm:w-auto sm:grid-cols-[220px_220px]">
                <Select value={statusFilter} onValueChange={(value) => setStatusFilter(value as TaskFilter)}>
                  <SelectTrigger>
                    <SelectValue placeholder="Filter tasks" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">All tasks</SelectItem>
                    <SelectItem value="todo">To do</SelectItem>
                    <SelectItem value="in_progress">In progress</SelectItem>
                    <SelectItem value="done">Done</SelectItem>
                  </SelectContent>
                </Select>

                <Input
                  value={assigneeFilter}
                  onChange={(event) => setAssigneeFilter(event.target.value)}
                  placeholder="Filter by assignee UUID"
                />
              </div>
            </CardContent>
          </Card>

          {tasksQuery.isLoading ? (
            <div className="space-y-3">
              <Skeleton className="h-44 w-full" />
              <Skeleton className="h-44 w-full" />
            </div>
          ) : null}

          {!tasksQuery.isLoading && tasksQuery.data?.length === 0 ? (
            <TaskEmptyState onCreateClick={openCreateTaskForm} />
          ) : null}

          <div className="grid gap-4">
            {tasksQuery.data?.map((task) => {
              const canManageTask = Boolean(canManageProject || task.created_by_id === user?.id);
              return (
                <TaskCard
                  key={task.id}
                  task={task}
                  canManage={canManageTask}
                  onDelete={handleDeleteTask}
                  onEdit={openEditTaskForm}
                  onStatusChange={handleStatusChange}
                  isDeletingTask={deleteTaskMutation.isPending}
                  isUpdatingStatus={statusMutation.isPending}
                />
              );
            })}
          </div>
        </section>

        <aside className="space-y-4">
          <Card className="xl:sticky xl:top-24">
            <CardHeader>
              <CardTitle>Progress snapshot</CardTitle>
              <CardDescription>Monitor flow at a glance while you prioritize execution.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              <div className="flex items-center justify-between">
                <span>Total tasks</span>
                <Badge variant="secondary">{taskStats.total}</Badge>
              </div>
              <Separator />
              <div className="flex items-center justify-between">
                <span>To do</span>
                <Badge>{taskStats.todo}</Badge>
              </div>
              <div className="flex items-center justify-between">
                <span>In progress</span>
                <Badge>{taskStats.inProgress}</Badge>
              </div>
              <div className="flex items-center justify-between">
                <span>Done</span>
                <Badge variant="success">{taskStats.done}</Badge>
              </div>
              <Separator />
              <p className="text-xs text-muted-foreground">
                Status changes are optimistic, so updates appear instantly and roll back if the API rejects them.
              </p>
            </CardContent>
          </Card>
        </aside>
      </div>

      <ProjectEditDialog
        open={isEditProjectOpen}
        onOpenChange={setIsEditProjectOpen}
        projectId={projectQuery.data.id}
        defaultValues={{
          name: projectQuery.data.name,
          description: projectQuery.data.description ?? "",
        }}
      />

      <TaskFormModal
        open={isTaskFormOpen}
        onOpenChange={setIsTaskFormOpen}
        projectId={projectId}
        task={selectedTask}
      />

      {(tasksQuery.isFetching || statusMutation.isPending) && (
        <div className="fixed bottom-4 right-4 inline-flex items-center gap-2 rounded-md border border-border bg-card px-3 py-2 text-sm shadow-soft">
          <LoaderCircle className="h-4 w-4 animate-spin" />
          Syncing updates...
        </div>
      )}
    </div>
  );
}
