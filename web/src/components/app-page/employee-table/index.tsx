"use client";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils";

import { Client } from "@clerk/nextjs/server";
import { GetStaffByStore } from "@/lib/hooks/staff";
import { useSelectedStore } from "@/components/storeprovider";
import { useParams } from "next/navigation";


export default function EmployeeTable() {
  const params = useParams()
  const id = params.store_id;

  if (!id) {
    // Handle the scenario where the id is not available
    // This could be rendering a placeholder, an error message, or returning null
    return <div>No store selected</div>;
  }
  
  const itemQuery = GetStaffByStore(id.toString());
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
