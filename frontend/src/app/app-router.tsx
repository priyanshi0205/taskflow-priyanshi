import { Navigate, Route, Routes } from "react-router-dom";

import { AppLayout } from "@/components/layout/app-layout";
import { LoginPage } from "@/features/auth/pages/login-page";
import { RegisterPage } from "@/features/auth/pages/register-page";
import { useAuth } from "@/features/auth/use-auth";
import { ProjectDetailPage } from "@/features/projects/pages/project-detail-page";
import { ProjectsPage } from "@/features/projects/pages/projects-page";
import { ProtectedRoute } from "@/routes/protected-route";

function RootRedirect(): React.JSX.Element {
  const { isAuthenticated } = useAuth();
  return <Navigate replace to={isAuthenticated ? "/projects" : "/login"} />;
}

export function AppRouter(): React.JSX.Element {
  return (
    <Routes>
      <Route element={<RootRedirect />} path="/" />
      <Route element={<LoginPage />} path="/login" />
      <Route element={<RegisterPage />} path="/register" />

      <Route element={<ProtectedRoute />}>
        <Route element={<AppLayout />}>
          <Route element={<ProjectsPage />} path="/projects" />
          <Route element={<ProjectDetailPage />} path="/projects/:id" />
        </Route>
      </Route>

      <Route element={<RootRedirect />} path="*" />
    </Routes>
  );
}

