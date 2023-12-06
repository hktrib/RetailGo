"use server";

import { revalidatePath } from "next/cache";

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
  const serverUrl = `https://retailgo-production.up.railway.app/store/${store_id}/inventory/create`;

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
};
