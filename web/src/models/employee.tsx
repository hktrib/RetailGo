export class Employee {
  id!: number;
  first_name!: string;
  last_name!: string;
  email!: string;
  name!: string;
  employeeId!: string;
  position!: string;
  department!: string;
  hireDate!: string;
  location!: string;

  constructor(employeeData?: Partial<Employee>) {
    // Initialize the class properties based on the employeeData object, with default values
    this.first_name = employeeData?.first_name || "";
    this.last_name = employeeData?.last_name || "";
    this.email = employeeData?.email || "";
    this.name = employeeData?.name || "";
    this.employeeId = employeeData?.employeeId || "";
    this.position = employeeData?.position || "";
    this.department = employeeData?.department || "";
    this.hireDate = employeeData?.hireDate || "";
    this.location = employeeData?.location || "";
  }
}

// Create a function to convert an Employee instance to a plain data object
export function convertEmployeeToPlainObject(employee: Employee): Record<string, any> {
  return {
    first_name: employee.first_name,
    last_name: employee.last_name,
    email: employee.email,
    name: employee.name,
    employeeId: employee.employeeId,
    position: employee.position,
    department: employee.department,
    hireDate: employee.hireDate,
    location: employee.location,
  };
}