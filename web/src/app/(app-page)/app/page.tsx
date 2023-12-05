"use client"

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import isAuth from "@/components/isAuth";

function DashboardPage() {
  return (
    <main className="bg-gray-50 h-full flex-grow flex">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0 flex-grow">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold">Dashboard</h1>
        </div>

        <Tabs defaultValue="overview" className="w-full pt-6 -mt-2 h-full">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="activity">Activity</TabsTrigger>
          </TabsList>
          <TabsContent value="overview" className="pt-14 h-full -mt-10 pb-4">
            <div className="grid grid-cols-4 grid-rows-4 grid-flow-row gap-4 h-full">
              <div className="bg-gray-200 rounded-md col-span-2 row-span-2" />
              <div className="bg-gray-200 rounded-md col-span-2 row-span-2" />

              <div className="bg-gray-200 rounded-md col-span-1 row-span-1" />
              <div className="bg-gray-200 rounded-md col-span-1 row-span-1" />
              <div className="bg-gray-200 rounded-md col-span-1 row-span-1" />
              <div className="bg-gray-200 rounded-md col-span-1 row-span-1" />

              <div className="bg-gray-200 rounded-md col-span-2 row-span-2" />
              <div className="bg-gray-200 rounded-md col-span-2 row-span-2" />
            </div>
          </TabsContent>
          <TabsContent value="activity">Activity</TabsContent>
        </Tabs>
      </div>
    </main>
  );
}

export default isAuth(DashboardPage);
