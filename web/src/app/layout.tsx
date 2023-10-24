import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/landing-page/Header";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "RetailGo",
  description:
    "RetailGo is a point-of-sale and inventory management solution designed to give businesses end-to-end control over their internal operations.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="h-full">
      <body className={`${inter.className} h-full`}>
        <Header />
        {children}
      </body>
    </html>
  );
}
