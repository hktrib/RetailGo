"use client";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils";

import { Client } from "@clerk/nextjs/server";
import { GetStaffByStore } from "@/lib/hooks/staff";

// Dummy inventory data
const employees = [
  {
    firstName: "John",
    lastName: "Doe",
    email: "john.doe@example.com",
    name: "John Doe",
    employeeId: "E1001",
    position: "Software Engineer",
    department: "IT",
    hireDate: "2021-01-15",
    location: "New York",
  },
  {
    firstName: "Jane",
    lastName: "Smith",
    email: "jane.smith@example.com",
    name: "Jane Smith",
    employeeId: "E1002",
    position: "Graphic Designer",
    department: "Marketing",
    hireDate: "2019-08-23",
    location: "Los Angeles",
  },
  {
    firstName: "Emily",
    lastName: "Johnson",
    email: "emily.johnson@example.com",
    name: "Emily Johnson",
    employeeId: "E1003",
    position: "Sales Manager",
    department: "Sales",
    hireDate: "2020-05-12",
    location: "Chicago",
  },
  {
    firstName: "Michael",
    lastName: "Brown",
    email: "michael.brown@example.com",
    name: "Michael Brown",
    employeeId: "E1004",
    position: "HR Coordinator",
    department: "Human Resources",
    hireDate: "2018-11-30",
    location: "Miami",
  },
  {
    firstName: "David",
    lastName: "Wilson",
    email: "david.wilson@example.com",
    name: "David Wilson",
    employeeId: "E1005",
    position: "Product Manager",
    department: "Product",
    hireDate: "2022-02-14",
    location: "Seattle",
  },
  {
    firstName: "Linda",
    lastName: "Garcia",
    email: "linda.garcia@example.com",
    name: "Linda Garcia",
    employeeId: "E1006",
    position: "Accountant",
    department: "Finance",
    hireDate: "2017-07-09",
    location: "Boston",
  },
  {
    firstName: "Robert",
    lastName: "Martinez",
    email: "robert.martinez@example.com",
    name: "Robert Martinez",
    employeeId: "E1007",
    position: "Network Administrator",
    department: "IT",
    hireDate: "2021-06-20",
    location: "San Francisco",
  },
  {
    firstName: "Sarah",
    lastName: "Robinson",
    email: "sarah.robinson@example.com",
    name: "Sarah Robinson",
    employeeId: "E1008",
    position: "Customer Service Rep",
    department: "Support",
    hireDate: "2019-03-18",
    location: "Austin",
  },
  {
    firstName: "Thomas",
    lastName: "Clark",
    email: "thomas.clark@example.com",
    name: "Thomas Clark",
    employeeId: "E1009",
    position: "Quality Assurance",
    department: "Production",
    hireDate: "2022-07-22",
    location: "Denver",
  },
  {
    firstName: "Nancy",
    lastName: "Rodriguez",
    email: "nancy.rodriguez@example.com",
    name: "Nancy Rodriguez",
    employeeId: "E1010",
    position: "Research Analyst",
    department: "Research",
    hireDate: "2018-12-10",
    location: "Atlanta",
  },
];

export default function EmployeeTable() {
  let data = employees;
  const itemQuery = GetStaffByStore("1381");
  console.log("ItemQuery is : " + itemQuery);
  // let data = initialInventory;

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
