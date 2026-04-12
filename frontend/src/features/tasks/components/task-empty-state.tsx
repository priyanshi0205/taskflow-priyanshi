import { ClipboardList } from "lucide-react";

import { Button } from "@/components/ui/button";

interface TaskEmptyStateProps {
  onCreateClick?: () => void;
}

export function TaskEmptyState({ onCreateClick }: TaskEmptyStateProps): React.JSX.Element {
  return (
    <div className="rounded-xl border border-dashed border-border bg-card/70 px-6 py-12 text-center">
      <ClipboardList className="mx-auto mb-3 h-9 w-9 text-muted-foreground" />
      <h3 className="text-xl font-semibold">No tasks yet</h3>
      <p className="mt-1 text-sm text-muted-foreground">
        Add your first task to start tracking work for this project.
      </p>
      {onCreateClick ? (
        <Button className="mt-5" onClick={onCreateClick}>
          Create task
        </Button>
      ) : (
        <p className="mt-3 text-xs text-muted-foreground">Only project owners can create tasks.</p>
      )}
    </div>
  );
}

