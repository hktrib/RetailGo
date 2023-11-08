import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"
import { useAuth } from '@clerk/nextjs';

 
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// From: https://clerk.com/docs/backend-requests/making/cross-origin

export function useFetch() {
  const { getToken } = useAuth();
 
  const authenticatedFetch = async (url: string | Request, init?: RequestInit) => {
    return fetch(url, {
      headers: { Authorization: `Bearer ${await getToken()}` }
    }).then(res => res.json());
  };
 
  return authenticatedFetch;
}
