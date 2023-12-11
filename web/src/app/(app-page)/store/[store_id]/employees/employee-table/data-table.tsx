"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { deleteUser } from "../actions";

import EditEmployeeDialog from "../edit-employee-dialog";
import { toast } from "react-toastify";
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  useReactTable,
  getSortedRowModel,
  getPaginationRowModel,
} from "@tanstack/react-table";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown";
import { Trash2 } from "lucide-react";

import type { EmployeeData } from "./columns";
import {useRouter} from "next/navigation";

interface DataTableProps<TData extends EmployeeData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
}

export function DataTable<TData extends EmployeeData, TValue>({
  columns,
  data,
}: DataTableProps<TData, TValue>) {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const router = useRouter();
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    onColumnFiltersChange: setColumnFilters,
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
    },
  });

  const params = useParams();
  const id = params.store_id;

  const onDelete = async (userId: string) => {
    console.log("Deleting user with id:", id);
    let response = await deleteUser({
      user_id: userId,
    });
    if (response) {
      console.log("Invite Mutations Success");
      toast.success("Deleted user sucessfully!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000,
      });
      router.refresh()
    } else {
      console.log("Invite Mutations Failuire");
      toast.error("Error deleting user!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000,
      });
    }
  };

  return (
    <div>
      <div className="flex items-center justify-between pb-4">
        <Input
          placeholder="Search employees..."
          value={
            (table.getColumn("first_name")?.getFilterValue() as string) ?? ""
          }
          onChange={(event) =>
            table.getColumn("first_name")?.setFilterValue(event.target.value)
          }
          className="max-w-sm dark:border-zinc-800 dark:focus:ring-zinc-700"
        />

        <div className="flex items-center space-x-2">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                className="ml-auto dark:border-zinc-800 dark:hover:bg-zinc-800"
              >
                Customize
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent
              align="end"
              className="dark:border-zinc-700 dark:bg-zinc-800"
            >
              {table
                .getAllColumns()
                .filter((column) => column.getCanHide())
                .map((column) => {
                  return (
                    <DropdownMenuCheckboxItem
                      key={column.id}
                      className="capitalize dark:hover:bg-zinc-700"
                      checked={column.getIsVisible()}
                      onCheckedChange={(value) =>
                        column.toggleVisibility(!!value)
                      }
                    >
                      {column.id}
                    </DropdownMenuCheckboxItem>
                  );
                })}
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>

      <div className="rounded-md border dark:border-zinc-800">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow
                key={headerGroup.id}
                className="dark:border-zinc-800 dark:hover:bg-zinc-800"
              >
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext(),
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                  className="dark:hover:bg-zinc-800"
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext(),
                      )}
                    </TableCell>
                  ))}
                  <TableCell>
                    <div className="flex h-8 w-8 flex-row space-x-2 p-0 text-right">
                      <EditEmployeeDialog employeeData={row.original} />
                      <button
                        type="button"
                        onClick={() => onDelete(row.original.id.toString())}
                      >
                        <Trash2
                          style={{ color: "red" }}
                          className="h-5 w-5 p-0"
                        />
                      </button>
                    </div>
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>

      <div className="flex items-center justify-end space-x-2 py-4">
        <Button
          variant="outline"
          size="sm"
          onClick={() => table.previousPage()}
          disabled={!table.getCanPreviousPage()}
          className="dark:border-zinc-800"
        >
          Previous
        </Button>
        <Button
          variant="outline"
          size="sm"
          onClick={() => table.nextPage()}
          disabled={!table.getCanNextPage()}
          className="dark:border-zinc-800"
        >
          Next
        </Button>
      </div>
    </div>
  );
}
