"use client";

import { ColumnDef } from "@tanstack/react-table";

import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuLabel,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
} from "@/components/ui/dropdown";
import { Button } from "@/components/ui/button";
import { ArrowUpDown, MoreHorizontal, PencilIcon, Trash2 } from "lucide-react";

export type EmployeeItem = {
  name: string;
  employeeId: string;
  position: string;
  department: string;
  hireDate: string;
  location: string;
};

export const columns: ColumnDef<EmployeeItem>[] = [
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
