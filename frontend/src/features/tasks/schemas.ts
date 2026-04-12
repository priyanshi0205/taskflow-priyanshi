import { z } from "zod";

const statusSchema = z.enum(["todo", "in_progress", "done"]);
const prioritySchema = z.enum(["low", "medium", "high"]);

export const taskFormSchema = z.object({
  title: z.string().min(1, "Task title is required"),
  description: z.string().max(300, "Description can be up to 300 characters").optional(),
  status: statusSchema,
  priority: prioritySchema,
  assignee_id: z.string().optional(),
  due_date: z
    .string()
    .optional()
    .refine((value) => !value || /^\d{4}-\d{2}-\d{2}$/.test(value), "Due date must be YYYY-MM-DD"),
});

export type TaskFormValues = z.infer<typeof taskFormSchema>;

