import { getStoreItemCategories, getStoreItems } from "../../queries";

import { cx } from "class-variance-authority";
import InventoryTable from "./inventory-table";
import AddItemDialog from "./item-dialog";
import { Package2 } from "lucide-react";

const stats = [
  {
    name: "Coming soon 1",
    value: "0",
    change: "+0%",
    changeType: "positive",
  },
  {
    name: "Coming soon 2",
    value: "0",
    change: "+0%",
    changeType: "positive",
  },
  {
    name: "Coming soon 3",
    value: "0",
    change: "+0%",
    changeType: "positive",
  },
];

export default async function Inventory({
  params,
}: {
  params: { store_id: string };
}) {
  const itemsData = await getStoreItems({ store_id: params.store_id });
  if (!itemsData.success) return <div>failed to fetch inventory items</div>;

  const items = itemsData.items;

  const storeCategoryData = await getStoreItemCategories({
    store_id: params.store_id,
  });
  if (!storeCategoryData.success) return <div>failed to fetch categories</div>;

  const categories = storeCategoryData.categories;

  return (
    <main className="h-full flex-grow">
      <div className="flex h-16 items-center justify-between px-4 py-5 md:px-6 xl:px-8">
        <div className="mt-1 flex items-center gap-x-3">
          <Package2 className="h-5 w-5 text-gray-800 dark:text-zinc-200" />
          <h1 className="text-xl font-medium tracking-wide">Inventory</h1>
        </div>

        <AddItemDialog categories={categories} />
      </div>

      <hr className="mb-4 border-gray-100 dark:border-zinc-800" />

      <div className="px-4 md:px-6 xl:px-8">
        <dl className="mx-auto mt-6 grid grid-cols-1 gap-px rounded-2xl border border-gray-100 bg-gray-900/5 shadow-sm dark:border-zinc-700 dark:bg-zinc-700/90 sm:grid-cols-2 lg:grid-cols-4">
          <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 rounded-t-2xl bg-white px-4 py-10 dark:bg-zinc-800 sm:rounded-tl-2xl sm:rounded-tr-none sm:px-6 lg:rounded-bl-2xl lg:rounded-tl-2xl xl:px-8">
            <dt className="text-sm font-medium leading-6 text-gray-500 dark:text-zinc-300">
              Total items
            </dt>
            {/* <dd className="text-gray-700 text-sm font-medium">
                  {stat.change}
                </dd> */}
            <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900 dark:text-white">
              {items.length}
            </dd>
          </div>

          {stats.map((stat, idx) => (
            <div
              key={stat.name}
              className={cx(
                "flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 dark:bg-zinc-800 sm:px-6 xl:px-8",
                idx === 0 && "sm:rounded-tr-2xl lg:rounded-none",
                idx === 1 && "sm:rounded-bl-2xl lg:rounded-none",
                idx === 2 &&
                  "rounded-b-2xl sm:rounded-bl-none sm:rounded-br-2xl lg:rounded-br-2xl lg:rounded-tr-2xl",
              )}
            >
              <dt className="text-sm font-medium leading-6 text-gray-500 dark:text-zinc-300">
                {stat.name}
              </dt>
              <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900 dark:text-white">
                {stat.value}
              </dd>
            </div>
          ))}
        </dl>

        <div className="mt-6">
          <InventoryTable items={items} categories={categories} />
        </div>
      </div>
    </main>
  );
}
