import { Outlet } from "react-router-dom";

import { Navbar } from "@/components/layout/navbar";

export function AppLayout(): React.JSX.Element {
  return (
    <div className="min-h-screen">
      <Navbar />
      <main className="container py-6 sm:py-8">
        <Outlet />
      </main>
    </div>
  );
}

