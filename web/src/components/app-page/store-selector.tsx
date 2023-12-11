"use client";

import { useRouter } from "next/navigation";
import type { StoreMetadata } from "@/app/(app-page)/store/layout";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
} from "@/components/ui/select";
import { cx } from "class-variance-authority";
import Link from "next/link";
import { Plus } from "lucide-react";

export default function StoreSelector({
  stores,
  currentStoreId,
  type,
}: {
  stores: StoreMetadata[];
  currentStoreId: string;
  type?: string;
}) {
  const router = useRouter();

  let currentStore;
  if (currentStoreId) {
    const store = stores.filter((store) => store.id === Number(currentStoreId));
    currentStore = store[0];
  }

  const routeStore = (storeId: string) => {
    if (!storeId) return;

    router.push(`/store/${storeId}`);
  };

  return (
    <Select value={currentStoreId} onValueChange={routeStore}>
      <SelectTrigger className="relative flex items-center gap-x-2 rounded-md border-none bg-gray-100 px-3 py-1.5 shadow-sm outline-none ring-0 dark:bg-zinc-700 dark:focus:ring-0">
        {currentStore ? (
          <>
            <div className="flex h-6 w-6 items-center justify-center rounded-md bg-white dark:bg-zinc-800">
              <span className="text-xs font-medium">
                {currentStore.storename.charAt(0)}
              </span>
            </div>
            <p
              className={cx(
                "truncate text-left text-sm font-medium",
                type === "mobile" ? "w-44" : "w-28",
              )}
            >
              {currentStore.storename}
            </p>
          </>
        ) : (
          <span>Select a store</span>
        )}
      </SelectTrigger>

      <SelectContent
        sideOffset={5}
        className="dark:border-zinc-800 dark:bg-zinc-900"
      >
        <SelectGroup>
          <SelectLabel>My stores</SelectLabel>
          {stores.map((store) => (
            <SelectItem
              key={store.id}
              value={store.id.toString()}
              className="dark:focus:bg-zinc-800"
            >
              {store.storename}
            </SelectItem>
          ))}
        </SelectGroup>

        <div className="mt-0.5 px-0.5 py-1">
          <Link
            href="/register-store"
            className="flex items-center justify-center gap-x-2 rounded-md bg-sky-500 py-2 pr-2 shadow-inner dark:bg-zinc-800 dark:shadow-zinc-700"
          >
            <Plus className="h-4 w-4 text-sky-50 dark:text-zinc-300" />
            <span className="mr-2 text-sm font-medium text-white">
              New store
            </span>
          </Link>
        </div>
      </SelectContent>
    </Select>
  );
}
