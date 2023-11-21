'use client'
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils"
import { useItems } from "@/app/(app-page)/app/hooks/items";
import { Client } from "@clerk/nextjs/server";

// Dummy inventory data
const initialInventory = [
  {
    name: "Item 1",
    category: "Test",
    price: 10,
    quantity: 5,
  },
  {
    name: "Item 2",
    category: "Test",
    price: 20,
    quantity: 3,
  },
  {
    name: "Item 3",
    category: "Test",
    price: 15,
    quantity: 8,
  },
  {
    name: "Item 4",
    category: "Test",
    price: 30,
    quantity: 2,
  },
  {
    name: "Item 5",
    category: "Test",
    price: 25,
    quantity: 6,
  },
];

export default function InventoryTable(){
  const itemQuery = useItems("1")

  let data = initialInventory;

  if (itemQuery.isLoading){
    return (<div>
      LOADING...
    </div>)
  } else if (!itemQuery.isSuccess){
    console.log("Error loading items:", itemQuery.error)
    return (<div>
      There was an error loading your items. Please try again!
    </div>)
  }

  data = itemQuery.data

  // try{
  //   data = await authFetch("http://localhost:8080/store/1/inventory/");
  // }
  // catch{
  // }

  return(
    <div>
      <DataTable columns = {columns} data = {itemQuery.data} />
    </div>
  );
}