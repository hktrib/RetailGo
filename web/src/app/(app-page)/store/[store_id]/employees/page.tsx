import EmployeeTable from "./employee-table";
import InviteEmployeeDialog from "./invite-employee-dialog";
import { GetStaffByStore } from "../../queries";

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
    <main className="h-full flex-grow bg-gray-50 dark:bg-zinc-900">
      <div className="mx-auto max-w-6xl px-6 py-6 md:px-8 lg:ml-0">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-semibold tracking-wide dark:text-white">
            Employees
          </h1>

          <>
            <InviteEmployeeDialog />
          </>
        </div>

        <hr className="my-4 dark:border-zinc-800" />

        <div className="mt-6">
          <EmployeeTable employees={employees} />
        </div>
      </div>
    </main>
  );
}
