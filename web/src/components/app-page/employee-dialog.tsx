"use client"
// Import necessary libraries and components
import { useEffect, useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
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
import { useFetch } from "@/lib/utils";
import { PencilIcon } from "lucide-react";
import { Employee } from "@/models/employee";
import { config } from "@/lib/hooks/config";
import { useParams } from "next/navigation";
import { toast } from "react-toastify";
import { PutUser } from "@/lib/hooks/user";

// Schema for form validation using Zod
const formSchema = z.object({
  first_name: z.string().min(1, "First name is required"),
  last_name: z.string(),
  email: z.string().regex(/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/),
});

// The component definition
export default function EmployeeDialog({
  employeeData = new Employee(),
  mode = "add",
}: {
  employeeData: Employee;
  mode?: string;
}) {
  const form = useForm({
    resolver: zodResolver(formSchema),
  });

  // Set up your fetch utility
  const [isDialogOpen, setDialogOpen] = useState(false); // State to control the dialog
  const params = useParams()
  const id = params.store_id;

  // Pre-fill the form when editing an employee
  useEffect(() => {
    if (employeeData != null) {
      console.log("Employee Data:", employeeData);
      form.reset(employeeData);
    }
  }, [employeeData, form]);

  const inviteMutation = PutUser(employeeData.id.toString());
  const onSubmit = form.handleSubmit((data: any) => {
    console.log(JSON.stringify(data));
    inviteMutation.mutate(data);
    if (inviteMutation.isSuccess) {
      setDialogOpen(false);
      toast.success("User updated successfully!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000
      });
    }else{
      toast.error("Failed to update user!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000
      }
      );
      console.log(inviteMutation.error)
    }

  });

  return (
    <Dialog open={isDialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
      <button className="icon-button">
            <PencilIcon
              style={{ color: "orange" }}
              className="h-5 w-5 p-0"
            ></PencilIcon>
          </button>
      </DialogTrigger>

      <Form {...form}>
        <form>
          <DialogContent>
          <form onSubmit={onSubmit} {...form}>

            <DialogHeader>
              <DialogTitle>
                {mode === "edit" ? "Edit" : "Add"} Employee
              </DialogTitle>
            </DialogHeader>

            <FormField
              control={form.control}
              name="first_name"
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
              name="last_name"
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
              <button type="submit">
                {mode === "edit" ? "Update" : "Save"}
              </button>
            </DialogFooter>
            </form>
          </DialogContent>
        </form>
      </Form>

    </Dialog>
  );
}
