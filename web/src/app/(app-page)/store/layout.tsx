import { auth, clerkClient } from "@clerk/nextjs";
import { notFound, redirect } from "next/navigation";

import { Toaster } from "react-hot-toast";
import MobileNav from "@/components/app-page/mobile-nav";
import Sidebar from "@/components/app-page/sidebar";
import Providers from "../providers";
import { GetStoresByClerkId } from "./queries";

export type StoreMetadata = {
  Permission_level: number;
  id: number;
  storename: string;
};

export default async function AppLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { userId, sessionClaims } = auth();

  if (!sessionClaims || !sessionClaims.publicMetadata) return notFound();

  const publicMetadata = sessionClaims.publicMetadata as {
    stores?: StoreMetadata[];
  };

  if (!publicMetadata.stores) {
    console.log("No public metadata getting from db");
    const db_stores = await GetStoresByClerkId();
    console.log("Stores are: ", db_stores.stores);

    if (db_stores.success) {
      await clerkClient.users.updateUserMetadata(userId, {
        publicMetadata: {
          stores: db_stores.stores,
        },
      });
    } else {
      console.log("No public metadata -> redirecting to register-store");
      redirect("/register-store");
    }
  }

  return (
    <Providers>
      <div className="flex h-full min-h-screen flex-col dark:bg-zinc-800">
        <MobileNav stores={publicMetadata.stores ?? []} />
        <Sidebar stores={publicMetadata.stores ?? []} />
        <div className="flex h-full flex-grow flex-col xl:pl-64">
          <div className="flex h-full flex-grow flex-col p-2">
            <div className="flex h-full flex-grow flex-col rounded-3xl border border-gray-100 bg-gray-50 shadow-sm dark:border-zinc-800 dark:bg-zinc-900">
              {children}
            </div>
          </div>
        </div>
      </div>

      <Toaster position="bottom-right" reverseOrder={false} />
    </Providers>
  );
}
