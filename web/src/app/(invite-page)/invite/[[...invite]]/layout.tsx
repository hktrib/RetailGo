"use client";

import Header from "@/components/landing-page/Header";
import Footer from "@/components/landing-page/footer";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useState } from "react";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000,
          },
        },
      })
  );
  return (
    <div className="min-h-screen h-full flex flex-col relative">
      <div className="flex-1 flex flex-col">
        <QueryClientProvider client={queryClient}>
          {children}
        </QueryClientProvider>
      </div>
    </div>


  );
}
