"use client";

import { ColumnDef } from "@tanstack/react-table";
import React, { useState } from 'react';

import { Button } from "@/components/ui/button";
import { ArrowUpDown, MoreHorizontal, PencilIcon, Trash2 } from "lucide-react";

export type InventoryItem = {
  id: number;
  name: string;
  description: string;
  category: string;
  price: number;
  quantity: number;
};


// Inside your component
const [isEditing, setIsEditing] = useState(false);
const [selectedEmployee, setSelectedEmployee] = useState(null);


export const columns: ColumnDef<InventoryItem>[] = [
  {
    accessorKey: "id",
    header: () => <div className="text-xs">ID</div>,
  },
  {
    accessorKey: "name",
    header: () => <div className="text-xs">Name</div>,
  },
  {
    accessorKey: "category",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          size="sm"
          className="-ml-3 h-8 data-[state=open]:bg-accent"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Category
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      );
    },
  },
  {
    accessorKey: "price",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          size="sm"
          className="-ml-3 h-8 data-[state=open]:bg-accent"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Price
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => {
      const price = parseFloat(row.getValue("price"));
      const formatted = new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
      }).format(price);

      return <div>{formatted}</div>;
    },
  },
  {
    accessorKey: "quantity",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          size="sm"
          className="-ml-3 h-8 data-[state=open]:bg-accent"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Quantity
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      );
    },
  },
  {
    id: "actions",
    enableHiding: false,
    cell: ({ row }) => {
      return (
        <div className="text-right flex flex-row space-x-2 h-8 w-8 p-0">
          <button onClick={() => console.log("Edit button clicked")}>
            <PencilIcon style={{ color: "orange" }} className="h-5 w-5 p-0"></PencilIcon>
          </button>
          <button onClick={() => console.log("Delete button clicked")}>
            <Trash2 style={{ color: "red" }} className="h-5 w-5 p-0"></Trash2>
          </button>
        </div>
      );
    },
  },
];
