"use client";

import Link from "next/link";
import { redirect } from "next/navigation";
import { useUser } from "@clerk/nextjs";

type StoreMetadata = {
  Permission_level: number;
  id: number;
  storename: string;
};

export default function StoreViewPage() {
  const { isLoaded, user } = useUser();

  // Reloades Clerk Metadata
  if (!isLoaded) {
    return <p>Loading...</p>;
  }

  if (user) {
    user.reload();
  }

  // Stall til not Loaded
  if (!user || !user.publicMetadata) return redirect("/registrationForm");

  const publicMetadata = user.publicMetadata as {
    stores?: StoreMetadata[];
  };
  if (!publicMetadata.stores) redirect("/registrationForm");

  return (
    <main className="flex h-full flex-grow bg-gray-50 dark:bg-zinc-900">
      <div className="mx-auto max-w-6xl flex-grow px-6 py-6 md:px-8 lg:ml-0">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-semibold tracking-wide">My stores</h1>
        </div>

        <hr className="my-4 dark:border-zinc-800" />

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
      <article className="w-52 rounded-md bg-white shadow-sm dark:bg-zinc-950">
        <div className="h-40 w-full rounded-t-md bg-gray-200 dark:bg-zinc-800" />
        <div className="p-4">
          <span className="text-sm font-medium">{storeName}</span>
        </div>
      </article>
    </Link>
  );
};
