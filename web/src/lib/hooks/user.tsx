import { Item, ItemWithoutId } from "@/models/item";
import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from './config';


const storeURL = config.serverURL + "user/";


export function HasStore() {
  // const queryClient = useQueryClient()
  const authFetch = useFetch();
  console.log(storeURL + "store")
  return useQuery({
    queryKey: ["hasStore"],
    queryFn: () => authFetch(storeURL + "store", {}, {}, true),
  });
}

export { config };

