export const queryKeys = {
  users: {
    all: ["users"] as const,
  },
  projects: {
    all: ["projects"] as const,
    detail: (projectId: string) => ["projects", projectId] as const,
  },
  tasks: {
    list: (projectId: string, status: string, assignee: string) =>
      ["tasks", projectId, status, assignee] as const,
    byProject: (projectId: string) => ["tasks", projectId] as const,
  },
};
