"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, SubmitHandler } from "react-hook-form";

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
import { SendInvite } from "@/lib/hooks/staff";

export default function InviteEmployee() {
  const formSchema = z.object({
    email: z.string(),
    name: z.string(),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  const inviteMutation = SendInvite("13");

  return (
    <Dialog>
      <DialogTrigger asChild>
        <button className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
          Invite
        </button>
      </DialogTrigger>

      <Form {...form}>
        <form>
          <DialogContent>
            <form onSubmit={form.handleSubmit(
                (data) => {
                  console.log("Data:", data);
                  return inviteMutation;
                })}>
              <DialogHeader>
                <DialogTitle>Invite employee</DialogTitle>
              </DialogHeader>

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
              <DialogFooter className="gap-x-4">
                <button
                  type="submit"
                  className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md mt-5"
                >
                  Invite
                </button>
              </DialogFooter>
            </form>
          </DialogContent>
        </form>
      </Form>
    </Dialog>
  );
}
