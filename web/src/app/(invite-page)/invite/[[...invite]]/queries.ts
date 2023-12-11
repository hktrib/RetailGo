import { config } from "@/lib/hooks/config";
import { revalidatePath } from "next/cache";

export const GetStoreByUUID = async ({
  uuid,
}: {
  uuid: string;
}) => {

  let serverUrl = `${config.serverURL}/store/uuid/${uuid}`;

  console.log(`Attempting to get store by uuid ${uuid}`);

  try {
    let response = await fetch(serverUrl, {
      method: "GET",
    });

    if (!response.ok) return {
      store: {},
      success: false
    };
    const store = JSON.parse(await response.text()) ?? {};
    console.log(`response: ${response.status} + ${response.statusText}`);
    return {
      store: store,
      success: true
    }
  } catch (err) {
    console.log("Failed to retrieve recommendations:", err)
    return {
      store: {},
      success: false
    }
  }
};
