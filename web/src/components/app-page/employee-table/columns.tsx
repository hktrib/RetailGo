"use client";

import React from "react";
import { ColumnDef } from "@tanstack/react-table";

import { Button } from "@/components/ui/button";
import { ArrowUpDown } from "lucide-react";

export type EmployeeData = {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  position: string;
};

export const columns: ColumnDef<EmployeeData>[] = [
  {
    accessorKey: "id",
    header: () => <div className="text-xs">ID</div>,
  },
  {
    accessorKey: "first_name",
    header: () => <div className="text-xs">First Name</div>,
  },
  {
    accessorKey: "last_name",
    header: () => <div className="text-xs">Last Name</div>,
  },
  {
    accessorKey: "email",
    header: () => <div className="text-xs">Email</div>,
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
];
