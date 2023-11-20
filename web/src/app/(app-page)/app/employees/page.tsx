import EmployeeTable from '@/components/app-page/employee-table';
import React from 'react';

export default function Employees() {
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
      location: "New York"
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
      location: "Los Angeles"
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
      location: "Chicago"
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
      location: "Miami"
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
      location: "Seattle"
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
      location: "Boston"
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
      location: "San Francisco"
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
      location: "Austin"
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
      location: "Denver"
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
      location: "Atlanta"
    }
  ];
  
      
    return (
        <main className="bg-gray-50 h-full flex-grow">
            <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0">
                <div className="flex items-center justify-between ">
                    <h1 className="text-2xl font-bold">Employees</h1>
                </div>
                <hr className="my-4" />
                <div className="mt-6">
                        <EmployeeTable />
                    </div>
            </div>
        </main>

    );
}