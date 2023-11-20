"use client";

import { ColumnDef } from "@tanstack/react-table";

import { Button } from "@/components/ui/button";
import { ArrowUpDown, MoreHorizontal, PencilIcon, Trash2 } from "lucide-react";
import React, { useState } from 'react';
import { Employee } from "@/models/employee";


export const columns: ColumnDef<Employee>[] = [
  {
    accessorKey: "employeeId",
    header: () => <div className="text-xs">ID</div>,
  },
  {
    accessorKey: "name",
    header: () => <div className="text-xs">Name</div>,
  },
  {
    accessorKey: "position",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          size="sm"
          className="-ml-3 h-8 data-[state=open]:bg-accent"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Position
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      );
    },
  },
  {
    accessorKey: "department",
    header: () => <div className="text-xs">ID</div>,
  },
  {
    accessorKey: "hireDate",
    header: () => <div className="text-xs">ID</div>,
  },
];

