import type { Metadata } from "next";
import { Inter as FontSans } from "next/font/google";
import "./globals.css";
import { ClerkProvider } from "@clerk/nextjs";
import { cn } from "@/lib/utils";

const fontSans = FontSans({
  subsets: ["latin"],
  variable: "--font-sans",
});

export const metadata: Metadata = {
  metadataBase: new URL("https://retail-go.vercel.app"),
  title: { default: "RetailGo", template: "%s | RetailGo" },
  description:
    "RetailGo is a point-of-sale and inventory management solution designed to give businesses end-to-end control over their internal operations.",
  openGraph: {
    title: "RetailGo",
    description:
      "RetailGo is a point-of-sale and inventory management solution designed to give businesses end-to-end control over their internal operations.",
    url: "https://retail-go.vercel.app",
    siteName: "RetailGo",
    images: {
      url: "https://retail-go.vercel.app/retailgo-banner.png",
      width: 1200,
      height: 630,
    },
    locale: "en-US",
    type: "website",
  },
  twitter: {
    title: "RetailGo",
    card: "summary_large_image",
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ClerkProvider>
      <html lang="en">
        <body
          className={cn(
            "min-h-screen bg-background font-sans antialiased",
            fontSans.variable,
          )}
        >
          {children}
        </body>
      </html>
    </ClerkProvider>
  );
}
