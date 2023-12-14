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

  let status = false

  try {
    let response = await fetch(serverUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });

    if (
      response.status === 200 ||
      response.status === 201
      ) {
        status = true
    }

  } catch (err) {
    status = false
    console.error(`error creating item for store ${store_id}: ${err}`);
  }

  
  revalidatePath("/store/[store]", "page");
  revalidateTag("storeCategories");

  return status
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

  let status = false;

  try {
    let response = await fetch(serverUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });

    if (response.status === 200) {
      status = true
    }
  } catch (err) {
    status = false
    console.error(`error updating item for store ${store_id}: ${err}`);
  }

  revalidatePath("/store/[store]", "page");

  return status
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

  let status = false

  try {
    let response = await fetch(serverUrl, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (response.status === 200) {
      status = true
    }

  } catch (err) {
    status = false
    console.error(`error deleting item ${itemId} for store ${storeId}: ${err}`);
  }

  revalidatePath("/store/[store]", "page");
  return status
};
