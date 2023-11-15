'use client'
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { useFetch } from "../../../lib/utils"

import { Client } from "@clerk/nextjs/server";

// Dummy inventory data
const employees = [
    {
      name: "John Doe",
      employeeId: "E1001",
      position: "Software Engineer",
      department: "IT",
      hireDate: "2021-01-15",
      location: "New York"
    },
    {
      name: "Jane Smith",
      employeeId: "E1002",
      position: "Graphic Designer",
      department: "Marketing",
      hireDate: "2019-08-23",
      location: "Los Angeles"
    },
    {
      name: "Emily Johnson",
      employeeId: "E1003",
      position: "Sales Manager",
      department: "Sales",
      hireDate: "2020-05-12",
      location: "Chicago"
    },
    {
      name: "Michael Brown",
      employeeId: "E1004",
      position: "HR Coordinator",
      department: "Human Resources",
      hireDate: "2018-11-30",
      location: "Miami"
    },
    {
      name: "David Wilson",
      employeeId: "E1005",
      position: "Product Manager",
      department: "Product",
      hireDate: "2022-02-14",
      location: "Seattle"
    },
    {
      name: "Linda Garcia",
      employeeId: "E1006",
      position: "Accountant",
      department: "Finance",
      hireDate: "2017-07-09",
      location: "Boston"
    },
    {
      name: "Robert Martinez",
      employeeId: "E1007",
      position: "Network Administrator",
      department: "IT",
      hireDate: "2021-06-20",
      location: "San Francisco"
    },
    {
      name: "Sarah Robinson",
      employeeId: "E1008",
      position: "Customer Service Rep",
      department: "Support",
      hireDate: "2019-03-18",
      location: "Austin"
    },
    {
      name: "Thomas Clark",
      employeeId: "E1009",
      position: "Quality Assurance",
      department: "Production",
      hireDate: "2022-07-22",
      location: "Denver"
    },
    {
      name: "Nancy Rodriguez",
      employeeId: "E1010",
      position: "Research Analyst",
      department: "Research",
      hireDate: "2018-12-10",
      location: "Atlanta"
    }
  ];
  

export default  function EmployeeTable(){
  let data = employees;
  let authFetch = useFetch()
  try{
    //data = await authFetch("http://localhost:8080/store/1391/inventory/");
  }
  catch{

  }

  return(
    <div>
      <DataTable columns = {columns} data = {data} />
    </div>
  );
}