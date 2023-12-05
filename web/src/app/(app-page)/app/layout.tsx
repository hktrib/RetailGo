"use client";
import MobileNav from "@/components/app-page/mobile-nav";
import Sidebar from "@/components/app-page/sidebar";
import { AppProps } from "next/app";
import { useState } from "react";
import { QueryClient } from "@tanstack/react-query";
import Providers from "../../providers";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function AppLayout({
  children,
  pageProps,
}: {
  children: React.ReactNode;
  pageProps: AppProps;
}) {
  const [queryClient] = useState(() => new QueryClient());
  const [infoMessage, setInfoMessage] = useState(null);
  return (
    <Providers>
      <div className="min-h-screen h-full flex flex-col">
      <ToastContainer />
        <MobileNav />
        <Sidebar />
        <div className="xl:pl-64 h-full flex-grow flex flex-col">
          {children}
        </div>
      </div>
      {/* <ReactQueryDevtools/> */}
    </Providers>
  );
}
