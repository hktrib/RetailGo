"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { SendInvite } from "@/lib/hooks/staff";

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

export default function InviteEmployee() {
  const [isDialogOpen, setDialogOpen] = useState(false); // State to control the dialog
  const params = useParams();

  const formSchema = z.object({
    email: z.string().regex(/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/),
    name: z.string(),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  const id = params.store_id;
  const inviteMutation = SendInvite(id.toString());

  const onSubmit = form.handleSubmit((data: any) => {
    console.log(JSON.stringify(data));

    inviteMutation.mutate(data);
    setDialogOpen(false);
  });

  return (
    <Dialog open={isDialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <button
          className="rounded-md bg-blue-500 px-4 py-1.5 text-sm font-medium text-white dark:bg-gradient-to-r dark:from-blue-600 dark:to-indigo-600"
          onClick={() => setDialogOpen(true)}
        >
          Invite
        </button>
      </DialogTrigger>

      <DialogContent>
        <DialogHeader>
          <DialogTitle>Invite employee</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={onSubmit} {...form}>
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder="Bob Jones" />
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
                    <Input {...field} placeholder="person@google.com" />
                  </FormControl>
                </FormItem>
              )}
            />

            <DialogFooter>
              <button
                type="submit"
                className="mt-5 rounded-md bg-blue-500 px-3 py-1.5 text-sm font-medium text-white"
              >
                Invite
              </button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
