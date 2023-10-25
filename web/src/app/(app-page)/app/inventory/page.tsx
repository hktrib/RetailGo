"use client"; // This is a client component 👈🏽

import React, { useState } from "react";
import AddItemDialog from "@/components/app-page/add-item-dialog";
import { DataTable } from "./data-table";
import { columns } from "./columns";

// Dummy inventory data
const initialInventory = [
  { name: "Item 1", description: "Description 1", price: 10, quantity: 5 },
  { name: "Item 2", description: "Description 2", price: 20, quantity: 3 },
  { name: "Item 3", description: "Description 3", price: 15, quantity: 8 },
  { name: "Item 4", description: "Description 4", price: 30, quantity: 2 },
  { name: "Item 5", description: "Description 5", price: 25, quantity: 6 },
];

export default function Inventory() {
  const stats = [
    {
      name: "Total items",
      value: "5",
      change: "+0%",
      changeType: "positive",
    },
    {
      name: "Coming soon",
      value: "0",
      change: "+0%",
      changeType: "positive",
    },
    {
      name: "Coming soon",
      value: "0",
      change: "+0%",
      changeType: "positive",
    },
    {
      name: "Coming soon",
      value: "0",
      change: "+0%",
      changeType: "positive",
    },
  ];

  return (
    <main className="bg-gray-50 h-full flex-grow">
      <div className="py-6 px-8 max-w-6xl mx-auto lg:ml-0">
        <div className="flex items-center justify-between ">
          <h1 className="text-2xl font-bold">Inventory</h1>

          <div>
            <AddItemDialog />
          </div>
        </div>
        <hr className="my-2" />

        <div className="mt-6">
          <dl className="border rounded-md mx-auto grid grid-cols-1 gap-px bg-gray-900/5 sm:grid-cols-2 lg:grid-cols-4">
            {stats.map((stat, idx) => (
              <div
                key={stat.name}
                className={`${idx === 0 ? "rounded-tl-md rounded-bl-md" : ""} ${
                  idx === stats.length - 1 ? "rounded-tr-md rounded-br-md" : ""
                } flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8`}
              >
                <dt className="text-sm font-medium leading-6 text-gray-500">
                  {stat.name}
                </dt>
                {/* <dd className="text-gray-700 text-sm font-medium">
                  {stat.change}
                </dd> */}
                <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900">
                  {stat.value}
                </dd>
              </div>
            ))}
          </dl>
        </div>

        <div className="mt-6">
          <DataTable columns={columns} data={initialInventory} />
        </div>
      </div>
    </main>
  );
}
