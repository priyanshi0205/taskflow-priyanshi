import { useCallback, useEffect, useMemo, useState } from "react";

import { AuthContext, type AuthContextValue } from "@/features/auth/auth-context";
import { AUTH_TOKEN_KEY } from "@/features/auth/constants";
import { isTokenExpired, tokenToUser } from "@/lib/jwt";
import type { User } from "@/types/entities";

function loadSession(): { token: string; user: User } | null {
  const token = localStorage.getItem(AUTH_TOKEN_KEY);
  if (!token || isTokenExpired(token)) {
    localStorage.removeItem(AUTH_TOKEN_KEY);
    return null;
  }

  const user = tokenToUser(token);
  if (!user) {
    localStorage.removeItem(AUTH_TOKEN_KEY);
    return null;
  }

  return { token, user };
}

export function AuthProvider({ children }: { children: React.ReactNode }): React.JSX.Element {
  const [initialSession] = useState(() => loadSession());
  const [token, setToken] = useState<string | null>(initialSession?.token ?? null);
  const [user, setUser] = useState<User | null>(initialSession?.user ?? null);
  const [isHydrated] = useState(true);

  const logout = useCallback(() => {
    localStorage.removeItem(AUTH_TOKEN_KEY);
    setToken(null);
    setUser(null);
  }, []);

  const login = useCallback((newToken: string) => {
    const mappedUser = tokenToUser(newToken);
    if (!mappedUser) {
      throw new Error("Received invalid token from server");
    }

    localStorage.setItem(AUTH_TOKEN_KEY, newToken);
    setToken(newToken);
    setUser(mappedUser);
  }, []);

  useEffect(() => {
    const handler = (): void => {
      logout();
    };

    window.addEventListener("auth:unauthorized", handler);
    return () => window.removeEventListener("auth:unauthorized", handler);
  }, [logout]);

  const value = useMemo<AuthContextValue>(
    () => ({
      user,
      token,
      isHydrated,
      isAuthenticated: Boolean(token && user),
      login,
      logout,
    }),
    [user, token, isHydrated, login, logout],
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

