"use client"
// Importing necessary dependencies
import { useState } from "react";
import { useParams } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { SendInvite } from "@/lib/hooks/staff";

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
import { Button } from "@/components/ui/button";

// Function component to invite an employee
export default function InviteEmployee() {
  // State to control the dialog
  const [isDialogOpen, setDialogOpen] = useState(false);

  // Get the parameters from the URL
  const params = useParams();

  // Define the form validation schema using zod
  const formSchema = z.object({
    email: z.string().regex(/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/),
    name: z.string(),
  });

  // Initialize the form using react-hook-form and zodResolver
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  // Get the store ID from the URL parameters
  const id = params.store_id;

  // Create a mutation function to send the invite
  const inviteMutation = SendInvite(id.toString());

  // Function to handle form submission
  const onSubmit = form.handleSubmit((data: any) => {
    console.log(JSON.stringify(data));

    // Call the inviteMutation function to send the invite
    inviteMutation.mutate(data);

    // Close the dialog
    setDialogOpen(false);
  });

  // Render the dialog component with the form
  return (
    <Dialog open={isDialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <button
          className="rounded-md bg-sky-500 px-4 py-1.5 text-sm font-medium text-white dark:bg-gradient-to-r dark:from-blue-600 dark:to-indigo-600"
          onClick={() => setDialogOpen(true)}
        >
          Invite
        </button>
      </DialogTrigger>

      <DialogContent className="dark:border-zinc-900 dark:bg-zinc-950">
        <DialogHeader>
          <DialogTitle>Invite employee</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={onSubmit} {...form} className="space-y-2">
            {/* Form field for employee name */}
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Bob Jones"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            {/* Form field for employee email */}
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="person@google.com"
                      className="dark:border-zinc-800 dark:focus:ring-zinc-700"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            {/* Form submission button */}
            <DialogFooter className="pt-2">
              <Button type="submit">Invite</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
