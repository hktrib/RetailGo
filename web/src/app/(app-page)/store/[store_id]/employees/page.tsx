import React from "react";
import EmployeeTable from "@/components/app-page/employee-table";
import { Employee } from "@/models/employee";
import InviteEmployee from "@/components/app-page/invite-employee";
import AddEmployee from "@/components/app-page/employee-dialog";
import { GetStaffByStore} from "../../queries";

export default async function Employees({
  params,
}: {
  params: { store_id: string };
}) {
  const itemQuery = await GetStaffByStore({store_id: params.store_id});
  const employees = itemQuery.employees;
  if (!itemQuery.success) {
    return <div>Failed to fetch Employees</div>;
  }

  return (
    <main className="bg-gray-50 h-full flex-grow">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold">Employees</h1>
          <div className="flex items-center space-x-4">
            <InviteEmployee />
          </div>
        </div>

        <hr className="my-4" />
        <div className="mt-6">
          <EmployeeTable employees={employees} />
        </div>
      </div>
    </main>
  );
}
