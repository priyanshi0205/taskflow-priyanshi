import { Navigate, Outlet, useLocation } from "react-router-dom";

import { useAuth } from "@/features/auth/use-auth";

export function ProtectedRoute(): React.JSX.Element {
  const { isAuthenticated, isHydrated } = useAuth();
  const location = useLocation();

  if (!isHydrated) {
    return (
      <div className="flex min-h-screen items-center justify-center text-sm text-muted-foreground">
        Validating session...
      </div>
    );
  }

  if (!isAuthenticated) {
    return <Navigate replace state={{ from: location }} to="/login" />;
  }

  return <Outlet />;
}

