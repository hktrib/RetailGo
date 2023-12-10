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
import { DeleteUser, PutUser } from "@/lib/hooks/user";
import { updateUser } from "@/app/(app-page)/store/[store_id]/employees/action";
import router from "next/router";

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
  employeeData?: Employee;
  mode?: string;
}) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      first_name: employeeData ? employeeData.first_name : undefined,
      last_name: employeeData ? employeeData.last_name : undefined,
      email: employeeData ? employeeData.email : undefined,
    },
  });

  // Set up your fetch utility
  const [isDialogOpen, setDialogOpen] = useState(false); // State to control the dialog
  const params = useParams()
  // Pre-fill the form when editing an employee
  useEffect(() => {
    if (employeeData != null) {
      console.log("Employee Data:", employeeData);
      form.reset(employeeData);
    }
  }, [employeeData, form]);

  const inviteMutation = PutUser(employeeData.id.toString());

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (!params.store_id) return;
    let response = await updateUser({ user: values, user_id: employeeData.id.toString() });
    if (response) {
      setDialogOpen(false);
      toast.success("Employee updated successfully!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000
      });
      router.reload();

    }else{
      toast.error("Error updating employee!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000
      });
    }
    
  }

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
        <DialogContent>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-1">

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
                    <Input {...field} type="email" disabled={true} />
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
      </Form>

    </Dialog>
  );
}
