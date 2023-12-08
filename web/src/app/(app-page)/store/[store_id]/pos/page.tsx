import { getPOSData } from "../../queries";
import POSController from "./controller";

export default async function POSPage({
  params,
}: {
  params: { store_id: string };
}) {
  const res = await getPOSData({ store_id: params.store_id });

  return (
    <POSController categories={res.data.categories} items={res.data.items} />
  );
}
