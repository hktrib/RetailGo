"use client";

import Link from "next/link";
import { useParams } from "next/navigation";
import { UserButton } from "@clerk/nextjs";
import { Navigation } from "./sidebar";
import StoreSelector from "./store-selector";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import { Menu } from "lucide-react";

import type { StoreMetadata } from "@/app/(app-page)/store/layout";

export default function MobileNav({ stores }: { stores: StoreMetadata[] }) {
  return (
    <div className="xl:hidden block bg-white/75 border-b">
      <div className="px-6 md:px-8 h-12 flex items-center justify-between">
        <div className="flex items-center gap-x-2">
          <SidebarNav stores={stores} />

          <Link href="/store" className="font-semibold text-lg">
            RetailGo
          </Link>
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
          <StoreSelector
            stores={stores}
            currentStoreId={store_id as string}
            type="mobile"
          />
        </SheetHeader>

        <Navigation userStores={stores} currentStoreId={store_id as string} />
      </SheetContent>
    </Sheet>
  );
};
