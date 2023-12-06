"use client";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils";
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

export default function InventoryTable({ items }: { items: Item[] }) {
  if (!items || !items.length) return <div>You have no items.</div>;

  return (
    <div>
      <DataTable columns={columns} data={items} />
    </div>
  );
}
