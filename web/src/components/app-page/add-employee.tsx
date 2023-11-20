"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, SubmitHandler } from "react-hook-form";
import { useFetch } from "../../lib/utils"

import * as z from "zod";

import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const formSchema = z.object({
  firstName: z.string().min(1, "First name is required"),
  lastName: z.string().min(1, "Last name is required"),
  email: z.string().email("Invalid email address")
});

export default function AddEmployee() {

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  let authFetch = useFetch()

  const onNewEmployee: SubmitHandler<z.infer<typeof formSchema>> = async (values) => {
    try {
      const response = await authFetch("http://localhost:8080/employees/create", 
        {
          method: 'POST',
          body: JSON.stringify(values)
        },
        {
          'Content-Type': 'application/json'
        }
      );

      if (!response.id) {
        throw new Error('Failed to add the employee.');
      }

    } catch (error) {
      console.error("There was an error:", error);
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <button className="bg-amber-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
          Add Employee
        </button>
      </DialogTrigger>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onNewEmployee)}>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>
                Add Employee
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

            <DialogFooter className="gap-x-4">
              <button type="submit">Save</button>
            </DialogFooter>
          </DialogContent>
        </form>
      </Form>
    </Dialog>
  );
}
