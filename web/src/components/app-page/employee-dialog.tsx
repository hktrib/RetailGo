"use client";

import { useState } from "react";
import router from "next/router";
import { useParams } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { updateUser } from "@/app/(app-page)/store/[store_id]/employees/actions";

import { toast } from "react-toastify";
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
import { PencilIcon } from "lucide-react";

import type { EmployeeData } from "./employee-table/columns";

// Schema for form validation using Zod
const formSchema = z.object({
  first_name: z.string().min(1, "First name is required"),
  last_name: z.string(),
  email: z.string().regex(/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/),
});

export default function EmployeeDialog({
  employeeData,
}: {
  employeeData: EmployeeData;
}) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      first_name: employeeData.first_name,
      last_name: employeeData.last_name,
      email: employeeData.email,
    },
  });

  const [isDialogOpen, setDialogOpen] = useState(false);
  const params = useParams();

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (!params.store_id) return;

    let response = await updateUser({
      user: values,
      user_id: employeeData.id.toString(),
    });

    if (response) {
      setDialogOpen(false);
      toast.success("Employee updated successfully!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000,
      });
    } else {
      toast.error("Error updating employee!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000,
      });
    }
  };

  return (
    <Dialog open={isDialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <button className="icon-button">
          <PencilIcon className="h-5 w-5 p-0 text-amber-500" />
        </button>
      </DialogTrigger>

      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Employee</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-1">
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

            <DialogFooter>
              <button type="submit">Update</button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
