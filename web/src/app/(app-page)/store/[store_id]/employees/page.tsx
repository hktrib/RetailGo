import { GetStaffByStore } from "../../queries";

import EmployeeTable from "./employee-table";
import InviteEmployeeDialog from "./invite-employee-dialog";
import { Users2 } from "lucide-react";

export default async function Employees({
  params,
}: {
  params: { store_id: string };
}) {
  const itemQuery = await GetStaffByStore({ store_id: params.store_id });
  const employees = itemQuery.employees;
  if (!itemQuery.success) {
    return <div>Failed to fetch Employees</div>;
  }

  return (
    <main className="h-full flex-grow">
      <div className="flex h-16 items-center justify-between px-4 py-5 md:px-6 xl:px-8">
        <div className="mt-1 flex items-center gap-x-3">
          <Users2 className="h-5 w-5 text-gray-800 dark:text-zinc-200" />
          <h1 className="text-xl font-medium tracking-wide">Employees</h1>
        </div>

        <InviteEmployeeDialog />
      </div>

      <hr className="mb-4 border-gray-100 dark:border-zinc-800" />

      <div className="px-4 md:px-6 xl:px-8">
        <div className="mt-6">
          <EmployeeTable employees={employees} />
        </div>
      </div>
    </main>
  );
}
