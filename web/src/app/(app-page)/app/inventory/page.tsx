import AddItemDialog from "@/components/app-page/add-item-dialog";
import InventoryTable from "@/components/app-page/inventory-table";

export default function Inventory() {
  const stats = [
    {
      name: "Total items",
      value: "5",
      change: "+0%",
      changeType: "positive",
    },
    {
      name: "Coming soon",
      value: "0",
      change: "+0%",
      changeType: "positive",
    },
    {
      name: "Coming soon",
      value: "0",
      change: "+0%",
      changeType: "positive",
    },
    {
      name: "Coming soon",
      value: "0",
      change: "+0%",
      changeType: "positive",
    },
  ];

  return (
    <main className="bg-gray-50 h-full flex-grow">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0">
        <div className="flex items-center justify-between ">
          <h1 className="text-2xl font-bold">Inventory</h1>

          <div>
            <AddItemDialog />
          </div>
        </div>
        <hr className="my-4" />

        <div className="mt-6">
          <dl className="border rounded-md mx-auto grid grid-cols-1 gap-px bg-gray-900/5 sm:grid-cols-2 lg:grid-cols-4">
            {stats.map((stat, idx) => (
              <div
                key={stat.name}
                className={`${idx === 0 ? "rounded-tl-md rounded-bl-md" : ""} ${
                  idx === stats.length - 1 ? "rounded-tr-md rounded-br-md" : ""
                } flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 bg-white px-4 py-10 sm:px-6 xl:px-8`}
              >
                <dt className="text-sm font-medium leading-6 text-gray-500">
                  {stat.name}
                </dt>
                {/* <dd className="text-gray-700 text-sm font-medium">
                  {stat.change}
                </dd> */}
                <dd className="w-full flex-none text-3xl font-medium leading-10 tracking-tight text-gray-900">
                  {stat.value}
                </dd>
              </div>
            ))}
          </dl>
        </div>

        <div className="mt-6">
          <InventoryTable />
        </div>
      </div>
    </main>
  );
}
