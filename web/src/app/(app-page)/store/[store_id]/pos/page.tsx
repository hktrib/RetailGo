import { getPOSData } from "../../queries";

import POSController from "./pos-controller";
import { ShoppingBag } from "lucide-react";

export default async function POSPage({
  params,
}: {
  params: { store_id: string };
}) {
  const res = await getPOSData({ store_id: params.store_id });

  return (
    <main className="h-full flex-grow">
      <div className="flex h-16 items-center justify-between px-4 py-5 md:px-6 xl:px-8">
        <div className="mt-1 flex items-center gap-x-3">
          <ShoppingBag className="h-5 w-5 text-gray-800 dark:text-zinc-200" />
          <h1 className="text-xl font-medium tracking-wide">Checkout</h1>
        </div>
      </div>

      <hr className="mb-4 border-gray-100 dark:border-zinc-800" />

      <POSController
        categories={res.data.categories}
        items={res.data.items}
        storeId={params.store_id as string}
      />
    </main>
  );
}
