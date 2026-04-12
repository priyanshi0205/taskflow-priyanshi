import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "react-router-dom";
import { LoaderCircle } from "lucide-react";
import { toast } from "sonner";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { login, register } from "@/features/auth/api";
import { AuthShell } from "@/features/auth/components/auth-shell";
import { useAuth } from "@/features/auth/use-auth";
import { getErrorMessage, getFieldErrors } from "@/lib/errors";

const registerSchema = z.object({
  name: z.string().min(2, "Name must be at least 2 characters"),
  email: z.string().email("Enter a valid email"),
  password: z.string().min(6, "Password must be at least 6 characters"),
});

type RegisterFormValues = z.infer<typeof registerSchema>;

export function RegisterPage(): React.JSX.Element {
  const navigate = useNavigate();
  const { isAuthenticated, login: storeLogin } = useAuth();

  const form = useForm<RegisterFormValues>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      name: "",
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
      await register(values);
      const response = await login({ email: values.email, password: values.password });
      storeLogin(response.token);
      toast.success("Account created");
      navigate("/projects", { replace: true });
    } catch (error) {
      const fieldErrors = getFieldErrors(error);
      if (fieldErrors?.name) {
        form.setError("name", { message: fieldErrors.name });
      }
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
      title="CREATE ACCOUNT"
      footerText="Already have an account?"
      footerLinkLabel="Sign in"
      footerLinkTo="/login"
    >
      <form className="space-y-4" onSubmit={onSubmit}>
        <div className="space-y-2">
          <Label htmlFor="name">Name</Label>
          <Input id="name" placeholder="Priya Sharma" {...form.register("name")} />
          <p className="min-h-5 text-xs text-destructive">{form.formState.errors.name?.message}</p>
        </div>

        <div className="space-y-2">
          <Label htmlFor="email">Email</Label>
          <Input id="email" type="email" placeholder="you@example.com" {...form.register("email")} />
          <p className="min-h-5 text-xs text-destructive">{form.formState.errors.email?.message}</p>
        </div>

        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input id="password" type="password" placeholder="Create a secure password" {...form.register("password")} />
          <p className="min-h-5 text-xs text-destructive">{form.formState.errors.password?.message}</p>
        </div>

        <Button className="w-full" disabled={form.formState.isSubmitting} type="submit">
          {form.formState.isSubmitting ? <LoaderCircle className="h-4 w-4 animate-spin" /> : null}
          Create account
        </Button>
      </form>
    </AuthShell>
  );
}

