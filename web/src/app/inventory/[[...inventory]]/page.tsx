"use client"; // This is a client component 👈🏽

import React, { useState } from 'react';
import { Trash } from 'lucide-react';
import './Table.css';
import Search from "./search";
import Infocard from "../../../components/inventory-page/Infocard";

export default function Inventory() {
  // Dummy inventory data
  const initialInventory = [
    { name: "Item 1", description: "Description 1", price: 10, quantity: 5 },
    { name: "Item 2", description: "Description 2", price: 20, quantity: 3 },
    { name: "Item 3", description: "Description 3", price: 15, quantity: 8 },
    { name: "Item 4", description: "Description 4", price: 30, quantity: 2 },
    { name: "Item 5", description: "Description 5", price: 25, quantity: 6 },
  ];

  const [search, setSearch] = useState("");
  const [inventory, setInventory] = useState(initialInventory);
  const [showAddItemForm, setShowAddItemForm] = useState(false);
  const [newItem, setNewItem] = useState({ name: '', description: '', price: 0, quantity: 0 });

  const handleAddItem = () => {
    setShowAddItemForm(true);
  };

  const handleCancelAddItem = () => {
    setShowAddItemForm(false);
    setNewItem({ name: '', description: '', price: 0, quantity: 0 });
  };

  const handleSaveItem = () => {
    setInventory([...inventory, newItem]);
    setShowAddItemForm(false);
    setNewItem({ name: '', description: '', price: 0, quantity: 0 });
  };

  const handleDeleteItem = (index: number) => {
    const updatedInventory = [...inventory];
    updatedInventory.splice(index, 1);
    setInventory(updatedInventory);
  };

  return (
    <main>
      <div className='py-24 sm:py-40 lg:py-48'>
        <h1 className="text-2xl font-bold mx-auto max-w-7xl">Inventory Stats</h1>
        <div className=" flex mx-auto max-w-7xl">
          <Infocard bgColor={"red"} title={"Test"} count={24} icon={"undefined"}></Infocard>
          <div className="bg-orange-100 rounded-lg p-6 mr-4">
            <h2 className="text-lg font-bold mb-2">Total Items</h2>
            <p className="text-3xl font-bold">{inventory.length}</p>
          </div>
          <div className="bg-green-100 rounded-lg p-6 mr-4">
            <h2 className="text-lg font-bold mb-2">TO BE ADDED</h2>
            <p className="text-3xl font-bold">0</p>
          </div>
          <div className="bg-blue-100 rounded-lg p-6 mr-4">
            <h2 className="text-lg font-bold mb-2">TO BE ADDED</h2>
            <p className="text-3xl font-bold">0</p>
          </div>
          <div className="bg-red-100 rounded-lg p-6">
            <h2 className="text-lg font-bold mb-2">TO BE ADDED</h2>
            <p className="text-3xl font-bold">0</p>
          </div>
        </div>

        <div className="mx-auto max-w-7xl ">
          <button className="bg-blue-500 hover-bg-blue-700 text-white font-bold py-2 px-4 rounded mb-4" onClick={handleAddItem}>
            Add Item 
          </button>
          {showAddItemForm && (
            <div className="p-4 border border-gray-300 mb-4">
              <h2>Add New Item</h2>
              <div className="mb-4">
                <label>Name:</label>
                <input
                  type="text"
                  value={newItem.name}
                  onChange={(e) => setNewItem({ ...newItem, name: e.target.value })}
                  className="border border-gray-300 rounded-md px-2 py-1"
                />
              </div>
              <div className="mb-4">
                <label>Description:</label>
                <input
                  type="text"
                  value={newItem.description}
                  onChange={(e) => setNewItem({ ...newItem, description: e.target.value })}
                  className="border border-gray-300 rounded-md px-2 py-1"

                />
              </div>
              <div className="mb-4">
                <label>Price:</label>
                <input
                  type="number"
                  value={newItem.price}
                  onChange={(e) => setNewItem({ ...newItem, price: Number(e.target.value) })}
                  className="border border-gray-300 rounded-md px-2 py-1"

                />
              </div>
              <div className="mb-4">
                <label>Quantity:</label>
                <input
                  type="number"
                  value={newItem.quantity}
                  onChange={(e) => setNewItem({ ...newItem, quantity: Number(e.target.value) })}
                  className="border border-gray-300 rounded-md px-2 py-1"

                />
              </div>
              <button
                className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mr-2"
                onClick={handleSaveItem}
              >
                Save
              </button>
              <button
                className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
                onClick={handleCancelAddItem}
              >
                Cancel
              </button>
            </div>
          )}
          <div className='product-list'>
            <div className="--flex-between --flex-dir-column">
              <span>
                <h3>Inventory Items</h3>
              </span>
            </div>
            <div className="ml-auto">
              <span>
                <Search
                  value={search}
                  onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setSearch(e.target.value)}
                />
              </span>
            </div>
            <table className="table">
              <thead>
                <tr>
                  <th className="px-4 py-2">Name</th>
                  <th className="px-4 py-2">Description</th>
                  <th className="px-4 py-2">Price</th>
                  <th className="px-4 py-2">Quantity</th>
                  <th className="px-4 py-2">Actions</th>
                </tr>
              </thead>
              <tbody>
                {inventory.map((item, index) => (
                  <tr key={index}>
                    <td className="border px-4 py-2">{item.name}</td>
                    <td className="border px-4 py-2">{item.description}</td>
                    <td className="border px-4 py-2">${item.price}</td>
                    <td className="border px-4 py-2">{item.quantity}</td>
                    <td className="border px-4 py-2">
                      <button
                        className="text-red-500 hover:text-red-700"
                        onClick={() => handleDeleteItem(index)}
                      >
                        <Trash size={16} />
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

        </div>
      </div>
    </main>
  );
}
