import { Link } from "react-router-dom";
import { LogOut, Sprout } from "lucide-react";

import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { useAuth } from "@/features/auth/use-auth";

function getInitials(name: string): string {
  return name
    .split(" ")
    .filter(Boolean)
    .map((value) => value[0]?.toUpperCase())
    .join("")
    .slice(0, 2);
}

export function Navbar(): React.JSX.Element {
  const { user, logout } = useAuth();

  return (
    <header className="sticky top-0 z-40 border-b border-border/70 bg-background/85 backdrop-blur">
      <div className="container flex h-16 items-center justify-between">
        <Link className="inline-flex items-center gap-2 text-foreground" to="/projects">
          <span className="inline-flex h-9 w-9 items-center justify-center rounded-full bg-primary/15 text-primary">
            <Sprout className="h-5 w-5" />
          </span>
          <span className="text-lg font-bold">TaskFlow <p className="text-xs text-muted-foreground">Green India by Zomato</p></span>
        </Link>

        <div className="flex items-center gap-3">
          <div className="hidden text-right sm:block">
            <p className="text-sm font-semibold">{user?.name ?? "Member"}</p>
            <p className="text-xs text-muted-foreground">{user?.email}</p>
          </div>

          <Avatar>
            <AvatarFallback>{getInitials(user?.name ?? "Member")}</AvatarFallback>
          </Avatar>

          <Button size="sm" variant="outline" onClick={logout}>
            <LogOut className="h-4 w-4" />
            <span className="hidden sm:inline">Logout</span>
          </Button>
        </div>
      </div>
    </header>
  );
}

