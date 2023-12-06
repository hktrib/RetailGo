import { auth } from "@clerk/nextjs";

export const getStoreItemCategories = async ({
  store_id,
}: {
  store_id: string;
}) => {
  const fetchUrl = `https://retailgo-production.up.railway.app/store/${store_id}/category`;

  const { sessionId } = auth();

  try {
    const res = await fetch(fetchUrl, {
      cache: "force-cache",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
    });
    if (!res.ok) return { success: false, categories: [] };

    const categories = JSON.parse(await res.text()) ?? [];

    return { success: true, categories: categories.categories };
  } catch (err) {
    console.error("error fetching store categories");
    return { success: false, categories: [] };
  }
};

export const getStoreItems = async ({ store_id }: { store_id: string }) => {
  const fetchUrl = `https://retailgo-production.up.railway.app/store/${store_id}/inventory`;

  const { sessionId } = auth();

  try {
    const res = await fetch(fetchUrl, {
      cache: "force-cache",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
    });
    if (!res.ok) return { success: false, items: [] };

    const items = JSON.parse(await res.text()) ?? [];

    return { success: true, items: items };
  } catch (err) {
    console.error("error fetching store items");
    return { success: false, items: [] };
  }
};
