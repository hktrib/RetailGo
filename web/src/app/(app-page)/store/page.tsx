"use client";

import Link from "next/link";
import { redirect } from "next/navigation";
import { useUser } from "@clerk/nextjs";
import { Store } from "lucide-react";

type StoreMetadata = {
  Permission_level: number;
  id: number;
  storename: string;
};

export default function StoreViewPage() {
  const { isLoaded, user } = useUser();

  // Reloades Clerk Metadata
  if (!isLoaded) return <p>Loading...</p>;

  // if (user) {
  //   user.reload();
  // }

  // Stall til not Loaded
  if (!user || !user.publicMetadata){
    return redirect("/register-store");
  }

  const publicMetadata = user.publicMetadata as {
    stores?: StoreMetadata[];
  };
  if (!publicMetadata.stores) redirect("/register-store");

  return (
    <main className="h-full flex-grow">
      <div className="flex h-16 items-center justify-between px-4 py-5 md:px-6 xl:px-8">
        <div className="mt-1 flex items-center gap-x-3">
          <Store className="h-5 w-5 text-gray-800 dark:text-zinc-200" />
          <h1 className="text-xl font-medium tracking-wide">My Stores</h1>
        </div>
      </div>

      <hr className="mb-4 border-gray-100 dark:border-zinc-800" />

      <div className="px-4 md:px-6 xl:px-8">
        <section className="flex flex-row flex-wrap gap-4">
          {publicMetadata.stores.map((store: StoreMetadata) => (
            <StoreCard
              key={store.id}
              id={store.id}
              storeName={store.storename}
            />
          ))}
        </section>
      </div>
    </main>
  );
}

const StoreCard = ({ id, storeName }: { id: number; storeName: string }) => {
  const createUrl = (id: number) => {
    const baseUrl = `/store/${id}`;
    return `${baseUrl}`;
  };

  return (
    <Link href={createUrl(id)}>
      <article className="w-52 rounded-md bg-white shadow-sm dark:bg-zinc-800">
        <div className="h-40 w-full rounded-t-md bg-gray-200 dark:bg-zinc-950" />
        <div className="p-4">
          <span className="text-sm font-medium">{storeName}</span>
        </div>
      </article>
    </Link>
  );
};
