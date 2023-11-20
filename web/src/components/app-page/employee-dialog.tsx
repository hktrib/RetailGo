// Import necessary libraries and components
import { useEffect } from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';

import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { useFetch } from '@/lib/utils';
import { PencilIcon } from 'lucide-react';
import { Employee } from '@/models/employee';


// Schema for form validation using Zod
const formSchema = z.object({
  firstName: z.string().min(1, 'First name is required'),
  lastName: z.string().min(1, 'Last name is required'),
  email: z.string().email('Invalid email address'),
});

// The component definition
export default function EmployeeDialog({ employeeData, mode = 'add' }: { employeeData: Employee, mode?: string }) {
  const form = useForm({
    resolver: zodResolver(formSchema),
  });

  // Set up your fetch utility
  let authFetch = useFetch();

  // Pre-fill the form when editing an employee
  useEffect(() => {

    if (employeeData != null) {
      console.log('Employee Data:', employeeData);
      form.reset(employeeData);
    }
  }, [employeeData, form]);

  // Handler for form submission
  const onEmployeeSubmit: SubmitHandler<z.infer<typeof formSchema>> = async (values) => {
    const url = employeeData
      ? `http://localhost:8080/employees/${employeeData}/update`
      : 'http://localhost:8080/employees/create';
    const method = employeeData ? 'PUT' : 'POST';

    try {
      const response = await authFetch(url, {
        method: method,
        body: JSON.stringify(values),
      }, {
        'Content-Type': 'application/json',
      });

      if (!response.id) {
        throw new Error(`Failed to ${employeeData ? 'update' : 'add'} the employee.`);
      }

    } catch (error) {
      console.error('There was an error:', error);
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        {mode === "edit" ? (
          <button className="icon-button">
            <PencilIcon style={{ color: "orange" }} className="h-5 w-5 p-0"></PencilIcon>
          </button>
        ) : (
          <button className="bg-amber-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
            Add item
          </button>
        )}
      </DialogTrigger>

      <Form {...form}>
        <form>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>
                {mode === "edit" ? 'Edit' : 'Add'} Employee
              </DialogTitle>
            </DialogHeader>

            <FormField
              control={form.control}
              name="firstName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>First Name</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="lastName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Last Name</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input {...field} type="email" />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="position"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Position</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="hireDate"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Hire Date</FormLabel>
                  <FormControl>
                  <Input {...field} type='date' defaultValue={new Date().toJSON()} />

                  </FormControl>
                </FormItem>
              )}
            />

            <DialogFooter className="gap-x-4">
              <button type="submit">
                {mode === "edit" ? 'Update' : 'Save'}
              </button>
            </DialogFooter>
          </DialogContent>
        </form>
      </Form>
    </Dialog>
  );
}
