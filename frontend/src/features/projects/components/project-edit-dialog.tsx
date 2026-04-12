import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
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
import { Textarea } from "@/components/ui/textarea";

import { useUpdateProjectMutation } from "@/features/projects/hooks";
import { getErrorMessage } from "@/lib/errors";

const schema = z.object({
  name: z.string().min(2, "Name is required"),
  description: z.string().optional(),
});

type FormValues = z.infer<typeof schema>;

interface Props {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  projectId: string;
  defaultValues: {
    name: string;
    description: string;
  };
}

export function ProjectEditDialog({
  open,
  onOpenChange,
  projectId,
  defaultValues,
}: Props): React.JSX.Element {
  const mutation = useUpdateProjectMutation(projectId);

  const form = useForm<FormValues>({
    resolver: zodResolver(schema),
    defaultValues,
  });

  useEffect(() => {
    form.reset(defaultValues);
  }, [defaultValues, form, open]);

  const onSubmit = form.handleSubmit(async (values) => {
    try {
      await mutation.mutateAsync(values);
      toast.success("Project updated");
      onOpenChange(false);
    } catch (err) {
      toast.error(getErrorMessage(err));
    }
  });

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Project</DialogTitle>
          <DialogDescription>
            Update project name and description. Changes are saved immediately.
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={onSubmit} className="space-y-4">
          <div>
            <Label>Name</Label>
            <Input {...form.register("name")} />
          </div>

          <div>
            <Label>Description</Label>
            <Textarea {...form.register("description")} />
          </div>

          <DialogFooter>
            <Button type="submit" disabled={mutation.isPending}>
              {mutation.isPending && (
                <LoaderCircle className="h-4 w-4 animate-spin" />
              )}
              Save
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
