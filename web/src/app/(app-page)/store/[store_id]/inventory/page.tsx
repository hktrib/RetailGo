import InventoryTable from "./inventory-table";
import AddItemDialog from "./item-dialog";
import { getStoreItemCategories, getStoreItems } from "../../queries";
import { cx } from "class-variance-authority";

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
    <main className="h-full flex-grow bg-gray-50 dark:bg-zinc-900">
      <div className="mx-auto max-w-6xl px-6 py-6 md:px-8 lg:ml-0">
        <div className="flex items-center justify-between ">
          <h1 className="text-2xl font-semibold tracking-wide dark:text-white">
            Inventory
          </h1>

          <>
            <AddItemDialog categories={categories} />
          </>
        </div>

        <hr className="my-4 dark:border-zinc-800" />

        <div className="mt-6">
          <dl className="mx-auto grid grid-cols-1 gap-px rounded-md border bg-gray-900/5 shadow-sm dark:border-zinc-700 dark:bg-zinc-700/90 sm:grid-cols-2 lg:grid-cols-4">
            <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 rounded-bl-md rounded-tl-md bg-white px-4 py-10 dark:bg-zinc-800 sm:px-6 xl:px-8">
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
                  idx === stats.length - 1 && "rounded-br-md rounded-tr-md",
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
        </div>

        <div className="mt-6">
          <InventoryTable items={items} categories={categories} />
        </div>
      </div>
    </main>
  );
}
