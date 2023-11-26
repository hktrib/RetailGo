import { Item, ItemWithoutId } from "@/models/item";
import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from './config';


const storeURL = config.serverURL + "store/";


export function GetStaffByStore(storeId: string) {
  // const queryClient = useQueryClient()
  const authFetch = useFetch();
  console.log(storeURL + storeId + "/staff")
  return useQuery({
    queryKey: ["staff", storeId],
    queryFn: () => authFetch(storeURL + storeId + "/staff", {}, {}, true),
  });
}

