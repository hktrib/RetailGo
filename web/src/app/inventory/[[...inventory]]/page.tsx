import React from 'react';

export default function Inventory() {
  // Dummy inventory data
  const inventory = [
    { name: "Item 1", description: "Description 1", price: 10, quantity: 5 },
    { name: "Item 2", description: "Description 2", price: 20, quantity: 3 },
    { name: "Item 3", description: "Description 3", price: 15, quantity: 8 },
    { name: "Item 4", description: "Description 4", price: 30, quantity: 2 },
  ];

  return (
    <main>
      <div className="mx-auto max-w-2xl py-24 sm:py-40 lg:py-48">
        <h1>Inventory</h1>
        <table className="table-auto w-full mt-8">
          <thead>
            <tr>
              <th className="px-4 py-2">Name</th>
              <th className="px-4 py-2">Description</th>
              <th className="px-4 py-2">Price</th>
              <th className="px-4 py-2">Quantity</th>
            </tr>
          </thead>
          <tbody>
            {inventory.map((item, index) => (
              <tr key={index}>
                <td className="border px-4 py-2">{item.name}</td>
                <td className="border px-4 py-2">{item.description}</td>
                <td className="border px-4 py-2">${item.price}</td>
                <td className="border px-4 py-2">{item.quantity}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  );
}
