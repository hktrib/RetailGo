import { columns } from "./columns";
import { DataTable } from "./data-table";

// Dummy inventory data
const initialInventory = [
  { name: "Item 1", description: "Description 1", price: 10, quantity: 5 },
  { name: "Item 2", description: "Description 2", price: 20, quantity: 3 },
  { name: "Item 3", description: "Description 3", price: 15, quantity: 8 },
  { name: "Item 4", description: "Description 4", price: 30, quantity: 2 },
  { name: "Item 5", description: "Description 5", price: 25, quantity: 6 },
];

const getInventory = async () => {

  try{
    const res = await fetch(`https://jsonplaceholder.typicode.com/posts/1`);
		const data = await res.json();
		console.log(data);
  }catch{
    console.log("Ther was an error fetching the data please try again!");
  }


};



export default function InventoryTable() {
  return (
    <div>
      <DataTable columns={columns} data={initialInventory} />
    </div>
  );
}
