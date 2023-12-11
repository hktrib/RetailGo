import { getPOSData } from "../../queries";
import POSController from "./pos-controller";

export default async function POSPage({
  params,
}: {
  params: { store_id: string };
}) {
  const res = await getPOSData({ store_id: params.store_id });

  return (
    <main className="h-full flex-grow bg-gray-50 dark:bg-zinc-900">
      <POSController
        categories={res.data.categories}
        items={res.data.items}
        storeId={params.store_id as string}
      />
    </main>
  );
}
