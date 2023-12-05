"use client";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils";

import { Client } from "@clerk/nextjs/server";
import { GetStaffByStore } from "@/lib/hooks/staff";
import { useSelectedStore } from "@/components/storeprovider";


export default function EmployeeTable() {
  const { selectedStore, selectStore } = useSelectedStore();

  const itemQuery = GetStaffByStore(selectedStore?.id ?? 0);
  if (itemQuery.isLoading) {
    return <div>LOADING...</div>;
  }
  if (!itemQuery.isSuccess) {
    console.error("Error loading items:", itemQuery.error);
    return <div>There was an error loading your items. Please try again!</div>;
  }

  return (
    <div>
      <DataTable columns={columns} data={itemQuery.data} />
    </div>
  );
}
