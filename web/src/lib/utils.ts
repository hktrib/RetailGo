import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"
import { useAuth } from '@clerk/nextjs';
 
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// From: https://clerk.com/docs/backend-requests/making/cross-origin

export function useFetch() {
  const { getToken } = useAuth();
 
  const authenticatedFetch = async (url: string | Request, init?: RequestInit, headers?: any, toJSON?: boolean) => {
    const response = fetch(url, {
      ...init,
      headers: headers !== undefined? {...headers, ...{ Authorization: `Bearer ${await getToken()}` }} : {...headers, ...{ Authorization: `Bearer ${await getToken()}` }}}
    )
    if (toJSON){
      return response.then(res => res.json())
    }
    else{
      return await response
    }
  };

// const authenticatedFetch = async (url: string | Request, init?: RequestInit, headers?: any) => {
//     const response = await fetch(url, {
//       ...init,
//       headers: headers !== undefined? {...headers, ...{ Authorization: `Bearer ${await getToken()}` }} : {...headers, ...{ Authorization: `Bearer ${await getToken()}` }}}
//     );

//     if (!response.ok){
//       return Promise.reject()
//     }else{
//       return response
//     }
//   };
  

  return authenticatedFetch;
}
