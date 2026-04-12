import { useState } from "react";
import { FolderOpen, Pencil } from "lucide-react";
import { Link } from "react-router-dom";
import { toast } from "sonner";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useAuth } from "@/features/auth/use-auth";
import { ProjectCreateDialog } from "@/features/projects/components/project-create-dialog";
import { ProjectEditDialog } from "@/features/projects/components/project-edit-dialog";
import { useDeleteProjectMutation, useProjectsQuery } from "@/features/projects/hooks";
import { getErrorMessage } from "@/lib/errors";
import type { Project } from "@/types/entities";

function ProjectsSkeleton(): React.JSX.Element {
  return (
    <div className="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
      {Array.from({ length: 6 }).map((_, index) => (
        <Card key={index}>
          <CardHeader>
            <Skeleton className="h-6 w-3/4" />
            <Skeleton className="h-4 w-full" />
          </CardHeader>
          <CardContent>
            <Skeleton className="h-4 w-5/6" />
          </CardContent>
          <CardFooter>
            <Skeleton className="h-9 w-28" />
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}

export function ProjectsPage(): React.JSX.Element {
  const { user } = useAuth();
  const projectsQuery = useProjectsQuery();
  const deleteMutation = useDeleteProjectMutation();
  const [projectToEdit, setProjectToEdit] = useState<Project | null>(null);

  const handleDelete = async (projectId: string): Promise<void> => {
    const confirmed = window.confirm("Are you sure you want to delete this project?");
    if (!confirmed) return;

    try {
      await deleteMutation.mutateAsync(projectId);
      toast.success("Project deleted");
    } catch (error) {
      toast.error(getErrorMessage(error));
    }
  };

  return (
    <div className="space-y-6">
      <section className="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
        <div className="space-y-1">
          <h1 className="text-3xl font-bold">Your Projects</h1>
          <p className="text-muted-foreground">
            Plan, prioritize, and execute every green initiative with one clear system.
          </p>
        </div>

        <ProjectCreateDialog />
      </section>

      {projectsQuery.isLoading ? <ProjectsSkeleton /> : null}

      {!projectsQuery.isLoading && projectsQuery.data?.length === 0 ? (
        <Card className="p-10 text-center">
          <FolderOpen className="mx-auto mb-3 h-10 w-10 text-muted-foreground" />
          <h2 className="text-xl font-semibold">No projects yet</h2>
        </Card>
      ) : null}

      <div className="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        {projectsQuery.data?.map((project) => {
          const canManage = user?.id === project.owner_id;

          return (
            <Card key={project.id}>
              <CardHeader>
                <div className="flex justify-between gap-2">
                  <CardTitle>{project.name}</CardTitle>
                  <Badge>Active</Badge>
                </div>
                <CardDescription>{project.description || "No description"}</CardDescription>
              </CardHeader>

              <CardContent>
                <p className="text-xs text-muted-foreground">
                  Created {new Date(project.created_at).toLocaleDateString()}
                </p>
              </CardContent>

              <CardFooter className="flex flex-wrap gap-2">
                <Button asChild variant="outline">
                  <Link to={`/projects/${project.id}`}>Open</Link>
                </Button>

                {canManage ? (
                  <>
                    <Button variant="secondary" onClick={() => setProjectToEdit(project)}>
                      <Pencil className="h-4 w-4" />
                      Edit
                    </Button>
                    <Button
                      variant="destructive"
                      onClick={() => handleDelete(project.id)}
                      disabled={deleteMutation.isPending}
                    >
                      Delete
                    </Button>
                  </>
                ) : (
                  <p className="text-xs text-muted-foreground">Read-only access</p>
                )}
              </CardFooter>
            </Card>
          );
        })}
      </div>

      {projectToEdit ? (
        <ProjectEditDialog
          open={Boolean(projectToEdit)}
          onOpenChange={(open) => {
            if (!open) {
              setProjectToEdit(null);
            }
          }}
          projectId={projectToEdit.id}
          defaultValues={{
            name: projectToEdit.name,
            description: projectToEdit.description ?? "",
          }}
        />
      ) : null}
    </div>
  );
}
