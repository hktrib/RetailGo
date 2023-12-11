import MobileNav from "@/components/app-page/mobile-nav";
import Sidebar from "@/components/app-page/sidebar";
import Providers from "../providers";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { auth } from "@clerk/nextjs";
import { notFound, redirect } from "next/navigation";

export type StoreMetadata = {
  Permission_level: number;
  id: number;
  storename: string;
};

export default function AppLayout({ children }: { children: React.ReactNode }) {
  const { sessionClaims } = auth();

  if (!sessionClaims || !sessionClaims.publicMetadata) return notFound();

  const publicMetadata = sessionClaims.publicMetadata as {
    stores?: StoreMetadata[];
  };
  if (!publicMetadata.stores) {
    console.log("No public metadata -> redirecting to main page");
    redirect("/");
  }

  return (
    <Providers>
      <div className="flex h-full min-h-screen flex-col dark:bg-zinc-800">
        <MobileNav stores={publicMetadata.stores} />
        <Sidebar stores={publicMetadata.stores} />
        <div className="flex h-full flex-grow flex-col xl:pl-64">
          <div className="flex h-full flex-grow flex-col p-2">
            <div className="flex h-full flex-grow flex-col rounded-3xl border border-gray-100 shadow-sm dark:border-zinc-800 dark:bg-zinc-900">
              {children}
            </div>
          </div>
        </div>
      </div>

      <ToastContainer />
    </Providers>
  );
}
