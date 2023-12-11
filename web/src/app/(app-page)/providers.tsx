"use client";

import { useState } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ThemeProvider } from "next-themes";
import SelectedStoreProvider from "../../components/storeprovider";

export default function Providers({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000,
          },
        },
      }),
  );

  return (
    <ThemeProvider attribute="class">
      <QueryClientProvider client={queryClient}>
        <SelectedStoreProvider>{children}</SelectedStoreProvider>
      </QueryClientProvider>
    </ThemeProvider>
  );
}
