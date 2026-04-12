import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {useNavigate } from "react-router-dom";
import { LoaderCircle } from "lucide-react";
import { toast } from "sonner";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { login } from "@/features/auth/api";
import { AuthShell } from "@/features/auth/components/auth-shell";
import { useAuth } from "@/features/auth/use-auth";
import { getErrorMessage, getFieldErrors } from "@/lib/errors";

const loginSchema = z.object({
  email: z.string().email("Enter a valid email"),
  password: z.string().min(1, "Password is required"),
});

type LoginFormValues = z.infer<typeof loginSchema>;

export function LoginPage(): React.JSX.Element {
  const navigate = useNavigate();
  const { isAuthenticated, login: storeLogin } = useAuth();

  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  useEffect(() => {
    if (isAuthenticated) {
      navigate("/projects", { replace: true });
    }
  }, [isAuthenticated, navigate]);

  const onSubmit = form.handleSubmit(async (values) => {
    try {
      const response = await login(values);
      storeLogin(response.token);
      toast.success("Welcome back");
      navigate("/projects", { replace: true });
    } catch (error) {
      const fieldErrors = getFieldErrors(error);
      if (fieldErrors?.email) {
        form.setError("email", { message: fieldErrors.email });
      }

      if (fieldErrors?.password) {
        form.setError("password", { message: fieldErrors.password });
      }

      toast.error(getErrorMessage(error));
    }
  });

  return (
    <AuthShell
      title="SIGN IN"
      footerText="New here?"
      footerLinkLabel="Create an account"
      footerLinkTo="/register"
    >
      <form className="space-y-4" onSubmit={onSubmit}>
        <div className="space-y-2">
          <Label htmlFor="email">Email</Label>
          <Input id="email" type="email" placeholder="you@example.com" {...form.register("email")} />
          <p className="min-h-5 text-xs text-destructive">{form.formState.errors.email?.message}</p>
        </div>

        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input id="password" type="password" placeholder="••••••••" {...form.register("password")} />
          <p className="min-h-5 text-xs text-destructive">{form.formState.errors.password?.message}</p>
        </div>

        <Button className="w-full" disabled={form.formState.isSubmitting} type="submit">
          {form.formState.isSubmitting ? <LoaderCircle className="h-4 w-4 animate-spin" /> : null}
          Continue
        </Button>
      </form>

      <p className="text-center text-xs text-muted-foreground">
        By continuing, you agree to keep project activity aligned with your team standards.
      </p>

    </AuthShell>
  );
}

