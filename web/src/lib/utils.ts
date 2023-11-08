import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"
import { useAuth } from '@clerk/nextjs';
import {merge} from 'lodash';

 
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// From: https://clerk.com/docs/backend-requests/making/cross-origin

export function useFetch() {
  const { getToken } = useAuth();
 
  const authenticatedFetch = async (url: string | Request, init?: RequestInit, headers?: any) => {
    return fetch(url, {
      ...init,
      headers: headers !== undefined? merge(headers, { Authorization: `Bearer ${await getToken()}` }) : { Authorization: `Bearer ${await getToken()}` }
    }).then(res => res.json());
  };
 
  return authenticatedFetch;
}
