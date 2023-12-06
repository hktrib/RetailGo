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
      <SelectTrigger className="shadow-none outline-none ring-0 border-none px-3 py-1.5 flex items-center gap-x-2 bg-gray-100 rounded-md relative">
        {currentStore ? (
          <>
            <div className="h-5 w-5 bg-white flex items-center justify-center rounded-md">
              <span className="text-xs font-medium">
                {currentStore.storename.charAt(0)}
              </span>
            </div>
            <p
              className={cx(
                "text-sm font-medium truncate text-left",
                type === "mobile" ? "w-44" : "w-28"
              )}
            >
              {currentStore.storename}
            </p>
          </>
        ) : (
          <span>Select a store</span>
        )}
      </SelectTrigger>

      <SelectContent sideOffset={5}>
        <SelectGroup>
          <SelectLabel>My stores</SelectLabel>
          {stores.map((store) => (
            <SelectItem key={store.id} value={store.id.toString()}>
              {store.storename}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
