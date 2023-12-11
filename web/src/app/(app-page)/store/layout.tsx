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
      <div className="min-h-screen h-full flex flex-col dark:bg-zinc-800">
        <MobileNav stores={publicMetadata.stores} />
        <Sidebar stores={publicMetadata.stores} />
        <div className="xl:pl-64 h-full flex-grow flex flex-col">
          {children}
        </div>
      </div>

      <ToastContainer />
    </Providers>
  );
}
