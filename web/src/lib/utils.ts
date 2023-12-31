import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { useAuth } from "@clerk/nextjs";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

// https://clerk.com/docs/backend-requests/making/cross-origin
export function useFetch() {
  const { getToken } = useAuth();
  const authenticatedFetch = async (
    url: string | Request,
    init?: RequestInit,
    headers?: any,
    toJSON?: boolean,
  ) => {
    const fetchHeaders =
      headers !== undefined
        ? { ...headers, Authorization: `Bearer ${await getToken()}` }
        : { Authorization: `Bearer ${await getToken()}` };

    const response = fetch(url, {
      ...init,
      headers: fetchHeaders,
    });
    if (toJSON) {
      return response.then((res) => res.json());
    } else {
      return await response;
    }
  };

  return authenticatedFetch;
}

export function useAuthenticated(): boolean {
  const { isSignedIn } = useAuth();
  if (!isSignedIn) {
    return false;
  }
  return true;
}

export const wait = () => new Promise((resolve) => setTimeout(resolve, 1000));
