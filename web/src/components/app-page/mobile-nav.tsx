"use client";

import { useParams } from "next/navigation";
import { UserButton } from "@clerk/nextjs";
import { Navigation } from "./sidebar";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import { Menu } from "lucide-react";

import type { StoreMetadata } from "@/app/(app-page)/app/layout";

export default function MobileNav({ stores }: { stores: StoreMetadata[] }) {
  return (
    <div className="xl:hidden block bg-white/75 border-b">
      <div className="px-6 md:px-8 h-12 flex items-center justify-between">
        <div className="flex items-center gap-x-2">
          <SidebarNav stores={stores} />

          <span className="font-bold">RetailGo</span>
        </div>
        <div>
          <UserButton afterSignOutUrl="/" />
        </div>
      </div>
    </div>
  );
}

const SidebarNav = ({ stores }: { stores: StoreMetadata[] }) => {
  const { store_id } = useParams();

  let currentStore;
  if (store_id) {
    const store = stores.filter((store) => store.id === Number(store_id));
    currentStore = store[0];
  }

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="ghost" className="p-1 -ml-2 w-fit h-fit">
          <Menu className="h-5 w-5" />
        </Button>
      </SheetTrigger>
      <SheetContent side="left" className="h-full w-[300px]">
        <SheetHeader className={`${store_id && "border-b pb-5"}`}>
          {currentStore && (
            <div className="px-3 py-1.5 flex items-center gap-x-3 bg-gray-100 rounded-md shadow">
              <div className="h-4 w-4 bg-white flex items-center justify-center rounded-md">
                <span className="text-xs">
                  {currentStore.storename.charAt(0)}
                </span>
              </div>
              <SheetTitle className="text-sm">
                {currentStore.storename}
              </SheetTitle>
            </div>
          )}
        </SheetHeader>

        <Navigation userStores={stores} currentStoreId={store_id as string} />
      </SheetContent>
    </Sheet>
  );
};
