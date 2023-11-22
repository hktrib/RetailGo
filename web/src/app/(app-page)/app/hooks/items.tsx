import { Item } from "@/models/item"
import {useFetch} from "../../../../lib/utils"
import { auth } from "@clerk/nextjs"
import {useQuery, useMutation, useQueryClient} from "@tanstack/react-query"

const serverURL = "http://localhost:8080"

const storeURL = serverURL + "/store/"

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

export function useItems(store: string){

    // const queryClient = useQueryClient()
    const authFetch = useFetch()

    return useQuery(
        {
            queryKey: ["items", store], 
            queryFn: () => authFetch(storeURL + store + "/inventory/"),
        }
    )
}

export function useCreateItem(store: string){

    const authFetch = useFetch()

    const queryClient = useQueryClient()

    return useMutation(
        {
            mutationFn: (newItem: Item) => authFetch(
                      storeURL + store + "/inventory/create", 
                      {
                        method: 'POST',
                        body: JSON.stringify(newItem, (key, value) => key === "quantity" || key === "price" ? parseFloat(value) : value)
                      },
                      {
                        'Content-Type': 'application/json'
                      }
                    ),
            onMutate: (newItem: Item) => {
                // await queryClient.cancelQueries({ queryKey: ['todos'] })
                const prevItems = queryClient.getQueryData(["items", store]) as Item[];
                queryClient.setQueryData(["items", store], (old: Item[]) => {
                    return [...prevItems, newItem]
                })   
                return {prevItems}
            },
            onError: (err, newItem: Item, context) => {
                console.log("Error while creating", newItem.name, ":", err)
                queryClient.setQueryData(["items", store], context?.prevItems)
            },
            // onSettled: () => {
            //     queryClient.invalidateQueries({ queryKey: ['todos'] })
            //   }
        }
    )

}