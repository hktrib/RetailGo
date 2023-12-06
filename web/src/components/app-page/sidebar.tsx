"use client";

import Link from "next/link";
import { useParams } from "next/navigation";
import { UserButton } from "@clerk/nextjs";
import {
  HelpCircle,
  HomeIcon,
  Package2,
  Settings,
  Users2,
  ShoppingBag,
} from "lucide-react";
import type { StoreMetadata } from "@/app/(app-page)/store/layout";
import StoreSelector from "./store-selector";
import { cx } from "class-variance-authority";

const navigation = [
  { name: "Dashboard", href: "/", icon: HomeIcon },
  { name: "Employees", href: "/employees", icon: Users2 },
  { name: "Inventory", href: "/inventory", icon: Package2 },
  { name: "POS", href: "/pos", icon: ShoppingBag },
];

export default function Sidebar({ stores }: { stores: StoreMetadata[] }) {
  const { store_id } = useParams();

  return (
    <div className="hidden xl:fixed xl:inset-y-0 xl:z-50 xl:flex xl:w-64 xl:flex-col">
      <div className="flex grow flex-col px-6 border-r">
        <div className="flex items-center justify-between py-5">
          <Link href="/store" className="text-lg font-semibold">
            RetailGo
          </Link>
          <UserButton afterSignOutUrl="/" />
        </div>

        <div className="h-full pb-12">
          <StoreSelector stores={stores} currentStoreId={store_id as string} />

          <Navigation userStores={stores} currentStoreId={store_id as string} />
        </div>
      </div>
    </div>
  );
}

export const Navigation = ({
  currentStoreId,
}: {
  userStores: StoreMetadata[];
  currentStoreId: string;
}) => {
  const buildUrl = (href: string) => {
    if (!currentStoreId) return "/store";

    const baseUrl = `/store/${currentStoreId}`;
    return `${baseUrl}${href}`;
  };

  return (
    <nav className="flex flex-1 flex-col h-full">
      <ul role="list" className="flex flex-1 flex-col">
        {currentStoreId && (
          <li className="py-5">
            <div className="text-xs text-gray-600 mb-2.5 px-3">Navigation</div>
            <ul role="list" className="space-y-1.5">
              {navigation.map((item) => (
                <li key={item.name} className="group">
                  <Link
                    href={buildUrl(item.href)}
                    className={cx(
                      "text-gray-900 hover:text-black flex items-center gap-x-3 px-3 py-1.5 rounded-md"
                    )}
                  >
                    <item.icon className="w-4 h-4" aria-hidden="true" />
                    <span className="text-sm">{item.name}</span>
                  </Link>
                </li>
              ))}
            </ul>
          </li>
        )}

        <li className="mt-auto pb-14 xl:pb-4">
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
  );
};
