"use client";

import Link from "next/link";
import { useParams, usePathname } from "next/navigation";
import { UserButton } from "@clerk/nextjs";

import { cx } from "class-variance-authority";
import StoreSelector from "./store-selector";
import ThemeSwitcher from "./theme-switcher";
import {
  HelpCircle,
  HomeIcon,
  Package2,
  Settings,
  Users2,
  ShoppingBag,
} from "lucide-react";

import type { StoreMetadata } from "@/app/(app-page)/store/layout";

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
      <div className="flex grow flex-col border-r px-6 dark:border-zinc-700">
        <div className="flex items-center justify-between py-5">
          <Link href="/store" className="text-lg font-semibold">
            RetailGo
          </Link>

          <div className="flex items-center">
            <UserButton afterSignOutUrl="/" appearance={{}} />
          </div>
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
  const pathname = usePathname();

  const buildUrl = (href: string) => {
    if (!currentStoreId) return "/store";

    const baseUrl = `/store/${currentStoreId}`;
    return `${baseUrl}${href}`;
  };

  return (
    <nav className="flex h-full flex-1 flex-col">
      <ul role="list" className="flex flex-1 flex-col">
        {currentStoreId && (
          <li className="py-5">
            <div className="mb-2.5 px-3 text-xs text-gray-600 dark:text-zinc-400">
              Navigation
            </div>
            <ul role="list" className="space-y-1.5">
              {navigation.map((item) => (
                <li key={item.name} className="group">
                  <Link
                    href={buildUrl(item.href)}
                    className={cx(
                      "flex items-center gap-x-3 rounded-md px-3 py-1.5",
                      pathname.endsWith(item.href) ||
                        (item.href === "/" &&
                          !pathname.endsWith("/employees") &&
                          !pathname.endsWith("/inventory") &&
                          !pathname.endsWith("/pos"))
                        ? "bg-gray-100 font-medium text-gray-900 shadow-sm transition duration-150 ease-in-out dark:bg-zinc-700 dark:text-white"
                        : "text-gray-700 transition duration-150 ease-in-out hover:bg-gray-100 dark:text-zinc-300 dark:hover:bg-zinc-700",
                    )}
                  >
                    <item.icon className="h-4 w-4" aria-hidden="true" />
                    <span className="text-sm">{item.name}</span>
                  </Link>
                </li>
              ))}
            </ul>
          </li>
        )}

        <li className="mt-auto pb-14 xl:pb-4">
          <ul role="list" className="space-y-1.5">
            <ThemeSwitcher />

            <li className="flex items-center gap-x-3 px-3 py-1.5 text-gray-700 dark:text-zinc-300">
              <Settings className="h-4 w-4" aria-hidden="true" />
              <span className="text-sm">Settings</span>
            </li>

            <li className="flex items-center gap-x-3 px-3 py-1.5 text-gray-700 dark:text-zinc-300">
              <HelpCircle className="h-4 w-4" aria-hidden="true" />
              <span className="text-sm">Help</span>
            </li>
          </ul>
        </li>
      </ul>
    </nav>
  );
};
