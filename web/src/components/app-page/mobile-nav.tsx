import { UserButton, useUser } from "@clerk/nextjs";
import {
  HelpCircle,
  HomeIcon,
  Menu,
  Package2,
  Settings,
  ShoppingBag,
  Users2,
} from "lucide-react";

import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { useSelectedStore } from "../storeprovider";
import { useEffect, useState } from "react";
import { IStore } from "@/models/store";

const navigation = [
  { name: "Dashboard", href: "/", icon: HomeIcon },
  { name: "Employees", href: "/employees", icon: Users2 },
  { name: "Inventory", href: "/inventory", icon: Package2 },
  { name: "POS", href: "/pos", icon: ShoppingBag },
];

const fakeStores = [{ name: "RetailGo", id: "123" }];

export default function MobileNav() {
  return (
    <div className="xl:hidden block bg-white/75 border-b">
      <div className="px-6 md:px-8 h-12 flex items-center justify-between">
        <div className="flex items-center gap-x-2">
          <SidebarNav />

          <span className="font-bold">RetailGo</span>
        </div>
        <div>
          <UserButton />
        </div>
      </div>
    </div>
  );
}

const SidebarNav = () => {
  const { user } = useUser();
  const { selectedStore, selectStore } = useSelectedStore();
  const [stores, setStores] = useState<IStore[]>([]);

  const mapMetadataStores = (storesData: any[]) =>
  storesData.map((store) => ({
    id: store.id,
    storename: store.storename,
    Permission_level: store.Permission_level,
  }));
  
  useEffect(() => {
    async function fetchStores() {
      if (user && user.publicMetadata.stores) {
        const metadataStores = mapMetadataStores(user.publicMetadata.stores as any[]);
        setStores(metadataStores);
      }
    }
    if(user && stores){
      fetchStores();
    }
  }, [user, selectStore]);
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="ghost" className="p-1 -ml-2 w-fit h-fit">
          <Menu className="h-5 w-5" />
        </Button>
      </SheetTrigger>
      <SheetContent side="left" className="h-full w-[300px]">
        <SheetHeader className="border-b pb-5">
          <div className="px-3 py-1.5 flex items-center gap-x-3 bg-gray-100 rounded-md shadow">
            <div className="h-4 w-4 bg-white flex items-center justify-center rounded-md">
              <span className="text-xs">{selectedStore?.storename.charAt(0)}</span>
            </div>
            <SheetTitle className="text-sm"> {selectedStore?.storename} </SheetTitle>
          </div>
        </SheetHeader>

        <nav className="flex flex-1 flex-col h-full">
          <ul role="list" className="flex flex-1 flex-col">
            <li className="py-5 border-b">
              <div className="text-xs text-gray-600 mb-2.5 px-3">
                Navigation
              </div>
              <ul role="list" className="space-y-1.5">
                {navigation.map((item) => (
                  <li key={item.name}>
                    <Link
                      href={`/app/${item.href}`}
                      className="text-gray-900 hover:text-black flex items-center gap-x-3 px-3 hover:bg-gray-100 py-1.5 rounded-md"
                    >
                      <item.icon className="w-4 h-4" aria-hidden="true" />
                      <span className="text-sm">{item.name}</span>
                    </Link>
                  </li>
                ))}
              </ul>
            </li>
            <li className="py-5 border-b">
              <div className="text-xs text-gray-600 mb-2.5 px-3">
                Your stores
              </div>
              <ul role="list" className="space-y-1.5">
                {stores.map((store) => (
                  <li key={store.storename}>
                    <div className="group text-gray-900 hover:text-black flex items-center gap-x-3 px-3 hover:bg-gray-100 py-1.5 rounded-md">
                      <div className="bg-gray-100 group-hover:bg-white h-4 w-4 flex items-center justify-center rounded-md">
                        <span className="text-xs">{store.storename.charAt(0)}</span>
                      </div>
                      <span className="text-sm">{store.storename}</span>
                    </div>
                  </li>
                ))}
              </ul>
            </li>

            <li className="mt-auto pb-12">
              <ul role="list" className="space-y-1.5">
                <li className="flex items-center px-3 py-1.5 gap-x-3">
                  <Settings className="w-4 h-4" aria-hidden="true" />
                  <span className="text-sm text-gray-900">Settings</span>
                </li>

                <li className="flex items-center px-3 py-1.5 gap-x-3">
                  <HelpCircle className="w-4 h-4" aria-hidden="true" />
                  <span className="text-sm text-gray-900">Help</span>
                </li>
              </ul>
            </li>
          </ul>
        </nav>
      </SheetContent>
    </Sheet>
  );
};
