'use server'

import {config} from "@/lib/hooks/config";

export const createStore = async ({
    postData
}: {
    postData: {
        store_name: string,
        store_phone: string,
        store_address: string,
        store_type: string,
        owner_email: string
    }
}) => {
    try {
        console.log("POST Data: ", JSON.stringify(postData));
        const response = await fetch(config.serverURL + "create/store", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(postData),
        });
        
        console.log(`response: ${response.status} + ${response.statusText}`);
        if (response.status === 200 || response.status === 201) {
            return true
        }
  
      } catch (error) {
        console.error("Error making create store request:", error);
      }
      return false
}