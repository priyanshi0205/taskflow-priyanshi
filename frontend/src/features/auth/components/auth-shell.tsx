import { Leaf } from "lucide-react";
import { Link } from "react-router-dom";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

interface AuthShellProps {
  title: string;
  footerText: string;
  footerLinkLabel: string;
  footerLinkTo: string;
  children: React.ReactNode;
}

export function AuthShell({
  title,
  footerText,
  footerLinkLabel,
  footerLinkTo,
  children,
}: AuthShellProps): React.JSX.Element {
  return (
    <div className="relative flex min-h-screen items-center justify-center overflow-hidden px-4 py-10">
      <div className="absolute -left-28 top-12 h-80 w-80 rounded-full bg-primary/20 blur-3xl" />
      <div className="absolute -right-20 bottom-8 h-72 w-72 rounded-full bg-accent/30 blur-3xl" />

      <Card className="relative z-10 w-full max-w-md border-primary/15 !bg-white ">
        <CardHeader className="text-center">
          <div className="mx-auto mb-2 flex h-12 w-12 items-center justify-center rounded-full bg-primary/15 text-primary">
            <Leaf className="h-6 w-6" />
            
          </div>
          <img src="../public/logo.png" alt="Greening India by Zomato" className="h-28 ml-auto mr-auto" />
          <CardTitle className="text-xl ">{title}</CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          {children}
          <p className="text-center text-sm text-muted-foreground">
            {footerText}{" "}
            <Link className="font-semibold text-primary hover:underline" to={footerLinkTo}>
              {footerLinkLabel}
            </Link>
          </p>
        </CardContent>
      </Card>
    </div>
  );
}

