// https://clerk.com/docs/backend-requests/making/cross-origin

import { useAuth } from "@clerk/nextjs";

export default function useFetch() {
  const { getToken } = useAuth();

  const authenticatedFetch = async ({
    url,
    init,
    headers,
    toJSON,
  }: {
    url: string | Request;
    init?: RequestInit;
    headers?: any;
    toJSON?: boolean;
  }) => {
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
  };

  return authenticatedFetch;
}
