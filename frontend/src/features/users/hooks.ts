import { useQuery } from "@tanstack/react-query";

import { getUsers } from "@/features/users/api";
import { queryKeys } from "@/lib/query-keys";

export function useUsersQuery() {
  return useQuery({
    queryKey: queryKeys.users.all,
    queryFn: getUsers,
  });
}
