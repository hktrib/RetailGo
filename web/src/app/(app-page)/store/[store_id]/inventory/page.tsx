import { useState, useEffect } from 'react';
import InventoryTable from "@/components/app-page/inventory-table";
import AddItemDialog from "@/components/app-page/item-dialog";
import { cx } from "class-variance-authority";
import { getStoreItemCategories, getStoreItems } from "../../queries";

// To prevent any type error for item and category variables in renderCategoryView
type Item = {
  id: number;
  name: string;
  categoryId: number;
}

type Category = {
  id: number;
  name: string;  
}

//Props for items in inventory
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
  const [items, setItems] = useState([]);
  const [categories, setCategories] = useState([]);
  const [viewType, setViewType] = useState('default');

  useEffect(() => {
    async function fetchData() {
      try {
        const itemsResponse = await getStoreItems({ store_id: params.store_id });
        if (itemsResponse.success) {
          setItems(itemsResponse.items);
        } else {
          console.error('Failed to fetch inventory items');
        }
  
        const categoriesResponse = await getStoreItemCategories({ store_id: params.store_id });
        if (categoriesResponse.success) {
          setCategories(categoriesResponse.categories);
        } else {
          console.error('Failed to fetch categories');
        }
      } catch (error) {
        console.error('An error occurred while fetching data:', error);
      }
    }
  
    fetchData();
  }, [params.store_id]); 
  const renderCategoryView = () => {
    return categories.map((category: Category) => {
      const categoryItems = items.filter(((item: Item) => item.categoryId === category.id));
      return (
        <div key={category.id} className="mt-6">
          <h2 className="text-xl font-semibold tracking-wide">{category.name}</h2>
          <div className="rounded-md bg-white px-4 py-10 dark:bg-zinc-800">
            <dt className="text-sm font-medium leading-6 text-gray-500 dark:text-zinc-300">
              Total items
            </dt>
            <dd className="text-3xl font-medium">{categoryItems.length}</dd>
          </div>
        </div>
      );
    });
  };

  return (
    <main className="h-full flex-grow bg-gray-50 dark:bg-zinc-900">
      <div className="mx-auto max-w-6xl px-6 py-6 md:px-8 lg:ml-0">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-semibold tracking-wide dark:text-white">
            Inventory
          </h1>
          <div>
            <select 
              value={viewType} 
              onChange={e => setViewType(e.target.value)}
              className="ml-4"
            >
              <option value="default">Default View</option>
              <option value="category">Category View</option>
            </select>
            <AddItemDialog categories={categories} />
          </div>
        </div>
        <hr className="my-4 dark:border-zinc-800" />

        {viewType === 'default' ? (
          <>
            <div className="mt-6">
              <dl className="border rounded-md mx-auto grid grid-cols-1 gap-px bg-gray-900/5 sm:grid-cols-2 lg:grid-cols-4">
                <div className="rounded-tl-md rounded-bl-md flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8">
                  <dt className="text-sm font-medium leading-6 text-gray-500">
                    Total items
                  </dt>
                  {/* <dd className="text-gray-700 text-sm font-medium">
                      {stat.change}
                    </dd> */}
                  <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900">
                    {items.length}
                  </dd>
                </div>

                {stats.map((stat, idx) => (
                  <div
                    key={stat.name}
                    className={cx(
                      "flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8",
                      idx === stats.length - 1 && "rounded-tr-md rounded-br-md"
                    )}
                  >
                    <dt className="text-sm font-medium leading-6 text-gray-500">
                      {stat.name}
                    </dt>
                    <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900">
                      {stat.value}
                    </dd>
                  </div>
                ))}
              </dl>
            </div>


        <div className="mt-6">
          <InventoryTable items={items} categories={categories} />
        </div>
        </>
        ) : renderCategoryView()}

        <div className="mt-6">
          <InventoryTable items={items} categories={categories} />
        </div>
      </div>
    </main>
  );
}