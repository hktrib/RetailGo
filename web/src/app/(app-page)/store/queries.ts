import { auth } from "@clerk/nextjs";
import { config } from "@/lib/hooks/config";


// Dashboard:

export const getItemRecommendations = async({
  store_id
}: {store_id: string}) => {
  const fetchUrl = `https://recommendation-server-production.up.railway.app/recommend/${store_id}`;
  console.log("FetchURL:", fetchUrl)
  try{
    const res = await fetch(fetchUrl,
      {headers: {'Access-Control-Allow-Origin': "no-cors"}}
      );

    if (!res.ok) return {
      items: [],
      success: false
    };

    const items = JSON.parse(await res.text()) ?? [];
    console.log("Fetched item recommendations:", items)
    return {
      items: items,
      success: true
    }
  }
  catch (err){
    console.log("Failed to retrieve recommendations:", err)
    return {
      items: [],
      success: false
    }
  }
}

// Inventory/POS Items

export const getStoreItemCategories = async ({
  store_id,
}: {
  store_id: string;
}) => {
  const fetchUrl = `${config.serverURL}/store/${store_id}/category`;
  console.log(fetchUrl);
  console.log(`fetching categories for store ${store_id}`);

  const { sessionId } = auth();

  try {
    const res = await fetch(fetchUrl, {
      cache: "force-cache",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
      next: { tags: ["storeCategories"] },
    });
    if (!res.ok) return { success: false, categories: [] };

    const data = JSON.parse(await res.text()) ?? [];
    console.log("fetched categories:", data);

    return {
      success: true,
      categories: data,
    };
  } catch (err) {
    console.error("error fetching store categories");
    return { success: false, categories: [] };
  }
};

export const getStoreItems = async ({ store_id }: { store_id: string }) => {
  const fetchUrl = `${config.serverURL}/store/${store_id}/inventory`;
  console.log(fetchUrl);
  console.log(`fetching items for store ${store_id}`);

  const { sessionId } = auth();

  try {
    const res = await fetch(fetchUrl, {
      cache: "force-cache",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
      next: { tags: ["storeItems"] },
    });
    if (!res.ok) return { success: false, items: [] };

    const data = JSON.parse(await res.text()) ?? [];
    console.log("fetched items:", data);

    return { success: true, items: data };
  } catch (err) {
    console.error("error fetching store items");
    return { success: false, items: [] };
  }
};

export const getPOSData = async ({ store_id }: { store_id: string }) => {
  console.log(`fetching POS data for store ${store_id}`);

  try {
    const categories = await getStoreItemCategories({ store_id: store_id });
    const items = await getStoreItems({ store_id: store_id });

    let categoriesSuccess = true;
    let itemsSuccess = true;
    if (!categories.success) categoriesSuccess = false;
    if (!items.success) itemsSuccess = false;

    return {
      success: categoriesSuccess && itemsSuccess,
      data: { categories: categories.categories, items: items.items },
    };
  } catch (err) {
    console.error("error fetching POS data");

    return { success: false, data: { categories: [], items: [] } };
  }
};


//Employee Items



export const GetStaffByStore = async ({
  store_id,
}: {
  store_id: string;
}) => {
  const fetchUrl = `${config.serverURL}/store/${store_id}/staff`;
  console.log(`fetching employees: ${fetchUrl}`);

  const { sessionId } = auth();

  try {
    const res = await fetch(fetchUrl, {
      cache: "force-cache",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
      next: { tags: ["storeEmployees"] },
    });
    if (!res.ok) return { success: false, employees: [] };

    const data = JSON.parse(await res.text()) ?? [];
    console.log("fetched categories:", data);
    console.log("fetched employees:", data);
    return {
      success: true,
      employees: data,
    };
  } catch (err) {
    console.error("error fetching store employees");
    return { success: false, employees: [] };
  }
};