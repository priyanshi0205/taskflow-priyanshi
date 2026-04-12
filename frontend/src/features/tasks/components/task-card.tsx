import { CalendarDays, UserRound } from "lucide-react";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import type { Task, TaskStatus } from "@/types/entities";

interface TaskCardProps {
  task: Task;
  canManage: boolean;
  onStatusChange: (taskId: string, status: TaskStatus) => void;
  onEdit: (task: Task) => void;
  onDelete: (taskId: string) => void;
  isUpdatingStatus: boolean;
  isDeletingTask: boolean;
}

const nextStatusMap: Record<TaskStatus, TaskStatus> = {
  todo: "in_progress",
  in_progress: "done",
  done: "todo",
};

const statusLabelMap: Record<TaskStatus, string> = {
  todo: "To do",
  in_progress: "In progress",
  done: "Done",
};

export function TaskCard({
  task,
  canManage,
  onStatusChange,
  onEdit,
  onDelete,
  isUpdatingStatus,
  isDeletingTask,
}: TaskCardProps): React.JSX.Element {
  const dueDateText = task.due_date ? new Date(task.due_date).toLocaleDateString() : "No due date";

  return (
    <Card className="h-full">
      <CardHeader className="space-y-3">
        <div className="flex items-start justify-between gap-2">
          <CardTitle className="line-clamp-2 text-lg">{task.title}</CardTitle>
          <Badge variant={task.status === "done" ? "success" : "default"}>{statusLabelMap[task.status]}</Badge>
        </div>
        <p className="line-clamp-3 text-sm text-muted-foreground">{task.description || "No description"}</p>
      </CardHeader>

      <CardContent className="space-y-2 text-sm text-muted-foreground">
        <p className="inline-flex items-center gap-2">
          <CalendarDays className="h-4 w-4" />
          {dueDateText}
        </p>
        <p className="inline-flex items-center gap-2">
          <UserRound className="h-4 w-4" />
          {task.assignee_id ? `Assignee ${task.assignee_id.slice(0, 8)}...` : "Unassigned"}
        </p>
        <Badge className="capitalize" variant="secondary">
          Priority: {task.priority}
        </Badge>
      </CardContent>

      <CardFooter className="flex flex-wrap gap-2">
        {canManage ? (
          <>
            <Button
              disabled={isUpdatingStatus}
              size="sm"
              variant="secondary"
              onClick={() => onStatusChange(task.id, nextStatusMap[task.status])}
            >
              Move to {statusLabelMap[nextStatusMap[task.status]]}
            </Button>
            <Button size="sm" variant="outline" onClick={() => onEdit(task)}>
              Edit
            </Button>
            <Button
              disabled={isDeletingTask}
              size="sm"
              variant="destructive"
              onClick={() => onDelete(task.id)}
            >
              Delete
            </Button>
          </>
        ) : (
          <p className="text-xs text-muted-foreground">Read-only task for non-owner members.</p>
        )}
      </CardFooter>
    </Card>
  );
}

