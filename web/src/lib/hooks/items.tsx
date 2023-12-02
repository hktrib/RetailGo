import { Item, ItemWithoutId } from "@/models/item";
import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from './config';


const storeURL = config.serverURL + "store/";

// function createItem(store: string, item: JSON){
//     return authFetch(inventoryURL + "create",
//     {
//         method: 'POST',
//         body: JSON.stringify(item, (key, value) => key === "quantity" || key === "price" ? parseFloat(value) : value)
//       },
//       {
//         'Content-Type': 'application/json'
//       })
// }

// function deleteItem(store: string, item: JSON)

// function updateItem(store: string, item: JSON)

export function useItems(store: string) {
  // const queryClient = useQueryClient()
  const authFetch = useFetch();

  return useQuery({
    queryKey: ["items", store],
    queryFn: () => authFetch(storeURL + store + "/inventory/", {}, {}, true),
  });
}

export function useCreateItem(store: string) {
  const authFetch = useFetch();

  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (newItem: ItemWithoutId) =>
      authFetch(
        storeURL + store + "/inventory/create",
        {
          method: "POST",
          body: JSON.stringify(newItem, (key, value) =>
            key === "quantity" || key === "price" ? parseFloat(value) : value
          ),
        },
        {
          "Content-Type": "application/json",
        }
      ),
    onMutate: (newItem: ItemWithoutId) => {
      // await queryClient.cancelQueries({ queryKey: ['todos'] })
      const prevItems = queryClient.getQueryData(["items", store]) as Item[];
      queryClient.setQueryData(["items", store], (old: Item[]) => {
        return prevItems.length > 0
          ? [
              ...prevItems,
              { ...newItem, id: prevItems.length + prevItems[0].id },
            ]
          : [{ ...newItem, id: 0 }];
      });
      return { prevItems };
    },
    onError: (err, newItem: ItemWithoutId, context) => {
      console.log("Error while creating", newItem.name, ":", err);
      queryClient.setQueryData(["items", store], context?.prevItems);
    },
    onSuccess: (newItem: ItemWithoutId) => {
      queryClient.invalidateQueries({ queryKey: ["items", store] });
    },
  });
}

export function useEditItem(store: string) {
  const authFetch = useFetch();

  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (newItem: Item) =>
      authFetch(storeURL + store + "/inventory/update", {
        method: "POST",
        body: JSON.stringify(newItem, (key, value) =>
          key === "quantity" || key === "price" ? parseFloat(value) : value
        ),
      }),
    onMutate: (newItem: Item) => {
      console.log("EDIT");
      // await queryClient.cancelQueries({ queryKey: ['todos'] })
      const prevItems = queryClient.getQueryData(["items", store]) as Item[];
      queryClient.setQueryData(["items", store], (old: Item[]) => {
        return prevItems.map((item) =>
          item.id !== newItem.id ? item : { ...item, ...newItem }
        );
      });
      return { prevItems };
    },
    onError: (err, newItem: Item, context) => {
      console.log("Error while creating", newItem.name, ":", err);
      queryClient.setQueryData(["items", store], context?.prevItems);
    },
  });
}

export function useDeleteItem(store: string) {
  const authFetch = useFetch();

  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) =>
      authFetch(storeURL + store + "/inventory/?id=" + id, {
        method: "DELETE",
        body: JSON.stringify({ id: id }),
      }),
    onMutate: (id: number) => {
      // await queryClient.cancelQueries({ queryKey: ['todos'] })
      const prevItems = queryClient.getQueryData(["items", store]) as Item[];
      queryClient.setQueryData(["items", store], (old: Item[]) => {
        return prevItems.filter((item) => item.id !== id);
      });
      return { prevItems };
    },
    onError: (err, id: number, context) => {
      console.log("Error while deleting", id, ":", err);
      queryClient.setQueryData(["items", store], context?.prevItems);
    },
  });
}
