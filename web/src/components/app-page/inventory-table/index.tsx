"use client";
import { InventoryItem, columns } from "./columns";
import { DataTable } from "./data-table";

export default function InventoryTable({
  items,
  categories,
}: {
  items: InventoryItem[];
  categories: Category[];
}) {
  if (!items || !items.length) return <div>You have no items.</div>;

  return (
    <div>
      <DataTable columns={columns} data={items} categories={categories} />
    </div>
  );
}
