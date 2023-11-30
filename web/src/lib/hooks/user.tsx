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
    mutationFn: (body: PostJoinStoreModel) =>
      authFetch(
        `${storeURL}/join/`,
        {
          method: "POST",
          body: JSON.stringify(body),
        },
        {
          "Content-Type": "application/json",
        }
        
      ),
    onError: (err, email, context) => {
      console.log("Error while sending invite to", email, ":", err);
    },
    onSuccess: (email) => {
      console.log("Invite sent successfully to", email);
      // You can add any additional logic you need on successful invite here
    },
  });
}


export { config };

