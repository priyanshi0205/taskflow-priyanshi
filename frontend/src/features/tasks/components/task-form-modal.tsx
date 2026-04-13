import { useEffect, useMemo } from "react";
import { useForm, useWatch } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { LoaderCircle } from "lucide-react";
import { toast } from "sonner";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import { Textarea } from "@/components/ui/textarea";
import { useIsMobile } from "@/hooks/use-mobile";
import { getErrorMessage } from "@/lib/errors";
import type { Task } from "@/types/entities";

import { useTaskUpsertMutation } from "@/features/tasks/hooks/use-task-mutations";
import { taskFormSchema, type TaskFormValues } from "@/features/tasks/schemas";
import type { UserDropdownOption } from "@/features/users/api";

import { PRIORITY_OPTIONS, STATUS_OPTIONS, UNASSIGNED_VALUE } from "../constant";

interface TaskFormModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  projectId: string;
  assigneeOptions: UserDropdownOption[];
  isAssigneeOptionsLoading: boolean;
  task?: Task;
}

function toDateInputValue(value: string | null): string {
  if (!value) {
    return "";
  }
  return value.slice(0, 10);
}

export function TaskFormModal({
  open,
  onOpenChange,
  projectId,
  assigneeOptions,
  isAssigneeOptionsLoading,
  task,
}: TaskFormModalProps): React.JSX.Element {
  const isMobile = useIsMobile();
  const taskMutation = useTaskUpsertMutation(projectId, task?.id);

  const form = useForm<TaskFormValues>({
    resolver: zodResolver(taskFormSchema),
    defaultValues: {
      title: task?.title ?? "",
      description: task?.description ?? "",
      status: task?.status ?? "todo",
      priority: task?.priority ?? "medium",
      assignee_id: task?.assignee_id ?? "",
      due_date: toDateInputValue(task?.due_date ?? null),
    },
  });
  const selectedStatus = useWatch({ control: form.control, name: "status" });
  const selectedPriority = useWatch({ control: form.control, name: "priority" });
  const selectedAssignee = useWatch({ control: form.control, name: "assignee_id" }) ?? "";

  const assigneeOptionsForSelect = useMemo(() => {
    if (!selectedAssignee) {
      return assigneeOptions;
    }

    const selectedExists = assigneeOptions.some((option) => option.uuid === selectedAssignee);
    if (selectedExists) {
      return assigneeOptions;
    }

    return [{ name: "Current assignee", uuid: selectedAssignee }, ...assigneeOptions];
  }, [assigneeOptions, selectedAssignee]);

  useEffect(() => {
    form.reset({
      title: task?.title ?? "",
      description: task?.description ?? "",
      status: task?.status ?? "todo",
      priority: task?.priority ?? "medium",
      assignee_id: task?.assignee_id ?? "",
      due_date: toDateInputValue(task?.due_date ?? null),
    });
  }, [form, task, open]);

  const onSubmit = form.handleSubmit(async (values) => {
    try {
      await taskMutation.mutateAsync({
        title: values.title,
        description: values.description ?? "",
        status: values.status,
        priority: values.priority,
        assignee_id: values.assignee_id?.trim() ? values.assignee_id.trim() : undefined,
        due_date: values.due_date?.trim() ? values.due_date : undefined,
      });

      toast.success(task ? "Task updated" : "Task created");
      onOpenChange(false);
    } catch (error) {
      toast.error(getErrorMessage(error));
    }
  });

  const formBody = (
    <form className="space-y-4" id="task-form" onSubmit={onSubmit}>
      <div className="space-y-2">
        <Label htmlFor="task-title">Title</Label>
        <Input id="task-title" placeholder="Collect permit approvals" {...form.register("title")} />
        <p className="min-h-5 text-xs text-destructive">{form.formState.errors.title?.message}</p>
      </div>

      <div className="space-y-2">
        <Label htmlFor="task-description">Description</Label>
        <Textarea
          id="task-description"
          placeholder="Capture key context, handoff details, and blockers..."
          {...form.register("description")}
        />
        <p className="min-h-5 text-xs text-destructive">{form.formState.errors.description?.message}</p>
      </div>

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
        <div className="space-y-2">
          <Label>Status</Label>
          <Select
            value={selectedStatus}
            onValueChange={(value) => form.setValue("status", value as TaskFormValues["status"])}
          >
            <SelectTrigger>
              <SelectValue placeholder="Status" />
            </SelectTrigger>
            <SelectContent>
              {STATUS_OPTIONS.map((status) => (
                <SelectItem key={status.value} value={status.value}>
                  {status.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="space-y-2">
          <Label>Priority</Label>
          <Select
            value={selectedPriority}
            onValueChange={(value) => form.setValue("priority", value as TaskFormValues["priority"])}
          >
            <SelectTrigger>
              <SelectValue placeholder="Priority" />
            </SelectTrigger>
            <SelectContent>
              {PRIORITY_OPTIONS.map((priority) => (
                <SelectItem key={priority.value} value={priority.value}>
                  {priority.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
        <div className="space-y-2">
          <Label>Assignee (optional)</Label>
          <Select
            disabled={isAssigneeOptionsLoading}
            value={selectedAssignee || UNASSIGNED_VALUE}
            onValueChange={(value) => {
              form.setValue("assignee_id", value === UNASSIGNED_VALUE ? "" : value);
            }}
          >
            <SelectTrigger>
              <SelectValue placeholder="Select assignee" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value={UNASSIGNED_VALUE}>Unassigned</SelectItem>
              {assigneeOptionsForSelect.map((option) => (
                <SelectItem key={option.uuid} value={option.uuid}>
                  {option.name} - {option.uuid}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="space-y-2">
          <Label htmlFor="task-due-date">Due date</Label>
          <Input id="task-due-date" type="date" {...form.register("due_date")} />
          <p className="min-h-5 text-xs text-destructive">{form.formState.errors.due_date?.message}</p>
        </div>
      </div>
    </form>
  );

  if (isMobile) {
    return (
      <Sheet open={open} onOpenChange={onOpenChange}>
        <SheetContent side="bottom" className="max-h-[94vh] overflow-y-auto rounded-t-2xl">
          <SheetHeader>
            <SheetTitle>{task ? "Edit task" : "Create task"}</SheetTitle>
            <SheetDescription>Fill details and keep your project workflow clear.</SheetDescription>
          </SheetHeader>
          <div className="mt-4">{formBody}</div>
          <SheetFooter>
            <Button className="w-full" disabled={taskMutation.isPending} form="task-form" type="submit">
              {taskMutation.isPending ? <LoaderCircle className="h-4 w-4 animate-spin" /> : null}
              {task ? "Save changes" : "Create task"}
            </Button>
          </SheetFooter>
        </SheetContent>
      </Sheet>
    );
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{task ? "Edit task" : "Create task"}</DialogTitle>
          <DialogDescription>Fill details and keep your project workflow clear.</DialogDescription>
        </DialogHeader>

        {formBody}

        <DialogFooter>
          <Button disabled={taskMutation.isPending} form="task-form" type="submit">
            {taskMutation.isPending ? <LoaderCircle className="h-4 w-4 animate-spin" /> : null}
            {task ? "Save changes" : "Create task"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
