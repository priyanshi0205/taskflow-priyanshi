import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { LoaderCircle, Plus } from "lucide-react";
import { toast } from "sonner";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useCreateProjectMutation } from "@/features/projects/hooks";
import { getErrorMessage } from "@/lib/errors";

const createProjectSchema = z.object({
  name: z.string().min(2, "Project name must be at least 2 characters"),
  description: z.string().max(200, "Keep description under 200 characters").optional(),
});

type CreateProjectValues = z.infer<typeof createProjectSchema>;

export function ProjectCreateDialog(): React.JSX.Element {
  const createProjectMutation = useCreateProjectMutation();

  const form = useForm<CreateProjectValues>({
    resolver: zodResolver(createProjectSchema),
    defaultValues: {
      name: "",
      description: "",
    },
  });

  const onSubmit = form.handleSubmit(async (values) => {
    try {
      await createProjectMutation.mutateAsync({
        name: values.name,
        description: values.description ?? "",
      });
      toast.success("Project created");
      form.reset();
    } catch (error) {
      toast.error(getErrorMessage(error));
    }
  });

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>
          <Plus className="h-4 w-4" />
          New project
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create project</DialogTitle>
          <DialogDescription>
            Name a project and give your team a short context before adding tasks.
          </DialogDescription>
        </DialogHeader>

        <form className="space-y-4" id="create-project-form" onSubmit={onSubmit}>
          <div className="space-y-2">
            <Label htmlFor="project-name">Name</Label>
            <Input id="project-name" placeholder="Urban tree census" {...form.register("name")} />
            <p className="min-h-5 text-xs text-destructive">{form.formState.errors.name?.message}</p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="project-description">Description</Label>
            <Textarea id="project-description" placeholder="High-impact goal, stakeholders, timeline..." {...form.register("description")} />
            <p className="min-h-5 text-xs text-destructive">{form.formState.errors.description?.message}</p>
          </div>
        </form>

        <DialogFooter>
          <Button disabled={createProjectMutation.isPending} form="create-project-form" type="submit">
            {createProjectMutation.isPending ? <LoaderCircle className="h-4 w-4 animate-spin" /> : null}
            Create
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

