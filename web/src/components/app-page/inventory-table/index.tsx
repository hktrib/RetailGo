"use client";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useItems } from "@/lib/hooks/items";

// dummy inventory data
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

export default function InventoryTable() {
  const itemQuery = useItems("10");

  // let data = initialInventory;

  if (itemQuery.isLoading) {
    return <div>LOADING...</div>;
  }

  if (!itemQuery.isSuccess) {
    console.error("Error loading items:", itemQuery.error);
    return <div>There was an error loading your items. Please try again!</div>;
  }

  return (
    <div>
      <DataTable columns={columns} data={itemQuery.data} />
    </div>
  );
}
