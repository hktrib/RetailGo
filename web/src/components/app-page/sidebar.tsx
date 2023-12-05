import Link from "next/link";
import { UserButton, useUser } from "@clerk/nextjs";
import {
  HelpCircle,
  HomeIcon,
  Package2,
  Settings,
  Users2,
  ShoppingBag,
  ArrowDown,
  BookUser,
} from "lucide-react";
import { useSelectedStore } from "../storeprovider";
import { IStore } from "@/models/store";
import { useEffect, useState } from "react";

const navigation = [
  { name: "Dashboard", href: "/", icon: HomeIcon },
  { name: "Employees", href: "/employees", icon: Users2 },
  { name: "Inventory", href: "/inventory", icon: Package2 },
  { name: "POS", href: "/pos", icon: ShoppingBag },
];

const stores: IStore[] = [];

export default function Sidebar() {
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
        const metadataStores = mapMetadataStores(
          user.publicMetadata.stores as any[]
        );
        setStores(metadataStores);
      }
    }
    if (user && stores) {
      fetchStores();
    }
  }, [user, selectStore]);

  return (
    <div className="hidden xl:fixed xl:inset-y-0 xl:z-50 xl:flex xl:w-64 xl:flex-col">
      <div className="flex grow flex-col overflow-y-auto px-6 border-r">
        <div className="flex items-center justify-between py-5 border-b">
          <span className="text-lg font-semibold">RetailGo</span>
          <UserButton />
        </div>

        <div className="py-5 border-b">
          <div className="px-3 py-1.5 flex items-center gap-x-3 bg-gray-100 rounded-md shadow">
            <div className="h-4 w-4 bg-white flex items-center justify-center rounded-md">
              <span className="text-xs">R</span>
            </div>
            <span className="text-sm">{selectedStore?.storename}</span>
          </div>
        </div>

        <nav className="flex flex-1 flex-col">
          <ul role="list" className="flex flex-1 flex-col">
            <li className="py-5 border-b">
              <div className="text-xs text-gray-600 mb-2.5 px-3">
                Navigation
              </div>
              <ul role="list" className="space-y-1.5">
                {navigation.map((item) => (
                  <li key={item.name} className="group">
                    <Link
                      href={`/app${item.href}`}
                      className="text-gray-900 hover:text-black flex items-center gap-x-3 px-3 py-1.5 rounded-md"
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
                    <Link
                      href={"#"}
                      onClick={() => {
                        selectStore(store);
                      }}
                    >
                      <div className="group text-gray-900 hover:text-black flex items-center gap-x-3 px-3 hover:bg-gray-100 py-1.5 rounded-md">
                        <div className="bg-gray-100 group-hover:bg-white h-4 w-4 flex items-center justify-center rounded-sm">
                          <span className="text-xs">
                            {store.storename.charAt(0)}
                          </span>
                        </div>
                        <span className="text-sm">{store.storename}</span>
                      </div>
                    </Link>
                  </li>
                ))}
              </ul>
            </li>

            <li className="mt-auto pb-5">
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
      </div>
    </div>
  );
}
