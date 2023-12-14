// Importing necessary dependencies
"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { updateUser } from "@/app/(app-page)/store/[store_id]/employees/actions";

import toast from "react-hot-toast";
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
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { PencilIcon } from "lucide-react";

import type { EmployeeData } from "./employee-table/columns";

// Schema for form validation using Zod
const formSchema = z.object({
  first_name: z.string().min(1, "First name is required"),
  last_name: z.string(),
  email: z.string().regex(/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/),
});

// Function component for editing an employee
export default function EditEmployeeDialog({
  employeeData,
}: {
  employeeData: EmployeeData;
}) {
  // Initializing form using react-hook-form
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      first_name: employeeData.first_name,
      last_name: employeeData.last_name,
      email: employeeData.email,
    },
  });

  // State for dialog open/close
  const [isDialogOpen, setDialogOpen] = useState(false);
  const params = useParams();

  // Function to handle form submission
  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (!params.store_id) return;

    // Calling the updateUser action to update the employee
    let response = await updateUser({
      user: values,
      user_id: employeeData.id.toString(),
    });

    // Checking the response and displaying appropriate toast message
    if (response) {
      setDialogOpen(false);
      toast.success("Employee updated successfully!");
    } else {
      toast.error("Error updating employee!");
    }
  };

  return (
    <Dialog open={isDialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <button className="icon-button">
          <PencilIcon className="h-5 w-5 p-0 text-amber-500" />
        </button>
      </DialogTrigger>

      <DialogContent className="dark:border-zinc-900 dark:bg-zinc-950">
        <DialogHeader>
          <DialogTitle>Edit Employee</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-2">
            <FormField
              control={form.control}
              name="first_name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>First Name</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="First name"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="last_name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Last Name</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Last name"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
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
                    <Input
                      {...field}
                      type="email"
                      disabled={true}
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            <DialogFooter className="pt-2">
              <Button type="submit">Update</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
