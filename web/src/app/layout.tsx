import type { Metadata } from "next";
import { Inter as FontSans } from "next/font/google";
import "./globals.css";
import { ClerkProvider } from "@clerk/nextjs";
import { cn } from "@/lib/utils";
import { HydrationBoundary, QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {useState} from "react"
import { AppProps } from "next/app";

const fontSans = FontSans({
  subsets: ["latin"],
  variable: "--font-sans",
});

export const metadata: Metadata = {
  title: "RetailGo",
  description:
    "RetailGo is a point-of-sale and inventory management solution designed to give businesses end-to-end control over their internal operations.",
};

// const queryClient = new QueryClient(
//   {
//     defaultOptions: {
//       queries: {
//         staleTime: 120000
//       }
//     }
//   }
// )

export default function RootLayout({children}: {children: React.ReactNode}) {

  return (
    <ClerkProvider>
          <html lang="en">
            <body
              className={cn(
                "min-h-screen bg-background font-sans antialiased",
                fontSans.variable
              )}
            >
              {children}
            </body>
          </html>
    </ClerkProvider>
  );
}
