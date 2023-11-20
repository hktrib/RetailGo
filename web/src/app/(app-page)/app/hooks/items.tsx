import { Item } from "@/models/item"
import {useFetch} from "../../../../lib/utils"
import { auth } from "@clerk/nextjs"
import {useQuery, QueryFunctionContext, useQueryClient} from "@tanstack/react-query"

const serverURL = "http://localhost:8080"

const storeURL = serverURL + "/store/"

// function readItems(store: string){
//     return authFetch(inventoryURL + store + "/inventory/")
// }

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

    const queryClient = useQueryClient()

    const authFetch = useFetch()

    return useQuery(
        {
            queryKey: ["items", store], 
            queryFn: () => authFetch(storeURL + store + "/inventory/"),
            staleTime: 2592000000
        }
    )
}