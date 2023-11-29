import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { useAuth } from "@clerk/nextjs";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

// https://clerk.com/docs/backend-requests/making/cross-origin
export async function useFetch({
  url,
  init,
  headers,
  toJSON,
}: {
  url: string | Request;
  init?: RequestInit;
  headers?: any;
  toJSON?: boolean;
}) {
  const { getToken } = useAuth();

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
  }

  return await response;
}
