"use server";

import { config } from "@/lib/hooks/config";

export async function createCheckout({
  lineItems,
  store_id,
}: {
  lineItems: { id: number; quantity: number }[];
  store_id: string;
}) {
  const serverUrl = `${config.serverUrl}/store/${store_id}/pos/checkout`;

  console.log(lineItems);
  console.log(`attempting to create checkout for store: ${store_id}`);

  try {
    const res = await fetch(serverUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(lineItems),
    });

    const data = await res.text();

    return data;
  } catch (err) {
    console.error(`error creating checkout for store ${store_id}: ${err}`);
  }
}
