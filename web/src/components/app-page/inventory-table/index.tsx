import { columns } from "./columns";
import { DataTable } from "./data-table";

// Dummy inventory data
const initialInventory = [
  {
    id: 1,
    name: "Item 1",
    description: "Description 1",
    price: 10,
    quantity: 5,
  },
  {
    id: 2,
    name: "Item 2",
    description: "Description 2",
    price: 20,
    quantity: 3,
  },
  {
    id: 3,
    name: "Item 3",
    description: "Description 3",
    price: 15,
    quantity: 8,
  },
  {
    id: 4,
    name: "Item 4",
    description: "Description 4",
    price: 30,
    quantity: 2,
  },
  {
    id: 5,
    name: "Item 5",
    description: "Description 5",
    price: 25,
    quantity: 6,
  },
];

const callAPI = async () => {
  try {
      const res = await fetch(
          `https://famous-quotes4.p.rapidapi.com/random?category=all&count=2`,
          {
              method: 'GET',
              /*headers: {
                  'X-RapidAPI-Key': 'your-rapidapi-key',
                  'X-RapidAPI-Host': 'famous-quotes4.p.rapidapi.com',
              },*/
          }
      );
      const data = await res.json();
      console.log(data);
  } catch (err) {
      console.log(err);
  }

export default function InventoryTable() {
  return (
    <div>
      <DataTable columns={columns} data={initialInventory} />
    </div>
  );
}
