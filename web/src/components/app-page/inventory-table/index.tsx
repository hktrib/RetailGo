import { columns } from "./columns";
import { DataTable } from "./data-table";

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

export default function InventoryTable() {
  return (
    <div>
      <DataTable columns={columns} data={initialInventory} />
    </div>
  );
}
