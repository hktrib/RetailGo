"use server";

import { revalidatePath, revalidateTag } from "next/cache";
import { config } from "@/lib/hooks/config";

export const createItem = async ({
  item,
  store_id,
}: {
  item: {
    name: string;
    price: number;
    quantity: number;
    category_name: string;
  };
  store_id: string;
}) => {
  const serverUrl = `${config.serverURL}/store/${store_id}/inventory/create`;
  console.log(serverUrl);
  console.log(item);
  console.log(`attempting to create item for store ${store_id}`);

  try {
    await fetch(serverUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });
  } catch (err) {
    console.error(`error creating item for store ${store_id}: ${err}`);
  }

  revalidatePath("/store/[store]", "page");
  revalidateTag("storeCategories");
};

export const updateItem = async ({
  item,
  store_id,
}: {
  item: {
    id: number;
    name: string;
    price: number;
    quantity: number;
    category_name: string;
  };
  store_id: string;
}) => {
  const serverUrl = `${config.serverURL}/store/${store_id}/inventory/update`;
  console.log(serverUrl);

  console.log(item);
  console.log(`attempting to update item for store ${store_id}`);

  try {
    await fetch(serverUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });
  } catch (err) {
    console.error(`error updating item for store ${store_id}: ${err}`);
  }

  revalidatePath("/store/[store]", "page");
};

export const deleteItem = async ({
  storeId,
  itemId,
}: {
  storeId: string;
  itemId: number;
}) => {
  const serverUrl = `${config.serverURL}/store/${storeId}/inventory?id=${itemId}`;
  console.log(serverUrl);

  console.log(`attempting to delete item ${itemId} for store ${storeId}`);

  try {
    await fetch(serverUrl, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
  } catch (err) {
    console.error(`error deleting item ${itemId} for store ${storeId}: ${err}`);
  }

  revalidatePath("/store/[store]", "page");
};
