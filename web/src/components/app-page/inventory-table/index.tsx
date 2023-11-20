'use client'
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils"
import { useItems } from "@/app/(app-page)/app/hooks/items";
import { Client } from "@clerk/nextjs/server";

// Dummy inventory data
const initialInventory = [
  {
    id: 1,
    name: "Item 1",
    description: "Description 1",
    category: "Test",
    price: 10,
    quantity: 5,
  },
  {
    id: 2,
    name: "Item 2",
    description: "Description 2",
    category: "Test",
    price: 20,
    quantity: 3,
  },
  {
    id: 3,
    name: "Item 3",
    description: "Description 3",
    category: "Test",
    price: 15,
    quantity: 8,
  },
  {
    id: 4,
    name: "Item 4",
    description: "Description 4",
    category: "Test",
    price: 30,
    quantity: 2,
  },
  {
    id: 5,
    name: "Item 5",
    description: "Description 5",
    category: "Test",
    price: 25,
    quantity: 6,
  },
];

export default  function InventoryTable(){
  let data = initialInventory;

  const itemQuery = useItems("1")

  if (itemQuery.isLoading){
    return (<div>
      LOADING...
    </div>)
  } else if (itemQuery.isError){
    console.log("Error loading items:", itemQuery.error)
    return (<div>
      There was an error loading your items. Please try again!
    </div>)
  }

  return(
    <div>
      <DataTable columns = {columns} data = {itemQuery.data} />
    </div>
  );
}