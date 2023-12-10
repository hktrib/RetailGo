"use client";
import { columns } from "./columns";
import { DataTable } from "./data-table";

export default function EmployeeTable({
  employees,
}: {
  employees: Employee[];
})  {

 


  return (
    <div>
      <DataTable columns={columns} data={employees} />
    </div>
  );
}
