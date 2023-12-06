import POSController from "./controller";

export default async function POSPage({
  params,
}: {
  params: { store_id: string };
}) {
  const res = await fetch(
    `https://retailgo-production.up.railway.app/store/${params.store_id}/pos/info`,
    { cache: "force-cache" }
  );
  if (!res.ok) return <div>failed to fetch pos stuff</div>;

  const data: { categories: Category[]; items: Item[] } = JSON.parse(
    await res.text()
  );

  return <POSController categories={data.categories} items={data.items} />;
}
