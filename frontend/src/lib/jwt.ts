import type { User } from "@/types/entities";

interface JwtPayload {
  user_id?: string;
  email?: string;
  exp?: number;
}

function decodeBase64Url(value: string): string {
  const base64 = value.replace(/-/g, "+").replace(/_/g, "/");
  const padded = base64.padEnd(Math.ceil(base64.length / 4) * 4, "=");
  return atob(padded);
}

export function parseToken(token: string): JwtPayload | null {
  try {
    const payloadPart = token.split(".")[1];
    if (!payloadPart) {
      return null;
    }

    const decoded = decodeBase64Url(payloadPart);
    return JSON.parse(decoded) as JwtPayload;
  } catch {
    return null;
  }
}

export function tokenToUser(token: string): User | null {
  const payload = parseToken(token);
  if (!payload?.user_id || !payload.email) {
    return null;
  }

  const localName = payload.email.split("@")[0] ?? "Member";

  return {
    id: payload.user_id,
    email: payload.email,
    name: localName,
  };
}

export function isTokenExpired(token: string): boolean {
  const payload = parseToken(token);
  if (!payload?.exp) {
    return true;
  }

  return payload.exp * 1000 <= Date.now();
}

