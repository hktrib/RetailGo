import { Item, ItemWithoutId } from "@/models/item";
import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from './config';
import { PostJoinStoreModel } from "@/models/user";


const storeURL = config.serverURL + "user/";


export function HasStore() {
  // const queryClient = useQueryClient()
  const authFetch = useFetch();
  return useQuery({
    queryKey: ["hasStore"],
    queryFn: () => authFetch(storeURL + "store", {}, {}, true),
  });
}

export function PostJoinStore(storeId: string) {
  const authFetch = useFetch();
  return useMutation({
    mutationFn: (clerkId: string) =>
      authFetch(
        `${storeURL}/join/`,
        {
          method: "POST",
          body: JSON.stringify({"ClerkId": clerkId, "StoreId": storeId}),
        },
        {
          "Content-Type": "application/json",
        }
        
      ),
    onError: (err, clerkId, context) => {
    },
    onSuccess: (clerkId) => {
      // You can add any additional logic you need on successful invite here
    },
  });
}


export { config };

