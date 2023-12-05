"use client"
import { useState } from "react"; // Import useState
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

  const [isDialogOpen, setDialogOpen] = useState(false); // State to control the dialog

  // Assuming storeId is known or retrieved from somewhere
  const storeId = "1";
  const inviteMutation = SendInvite(storeId);

  const onSubmit = form.handleSubmit((data) => {
    console.log(JSON.stringify(data));
    inviteMutation.mutate(data);

    // Close the dialog on successful submission
    if (inviteMutation.isSuccess) {
      setDialogOpen(false);
    }
  });

  return (
    <Dialog>
      <DialogTrigger asChild>
        <button
          className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md"
          onClick={() => setDialogOpen(true)} // Open the dialog on button click
        >
          Invite
        </button>
      </DialogTrigger>

      {isDialogOpen && (
        <Form {...form}>
          <form onSubmit={onSubmit} {...form}>
            <DialogContent>
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
            </DialogContent>
            <DialogFooter className="gap-x-4">
              <button
                type="submit"
                className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md mt-5"
              >
                Invite
              </button>
            </DialogFooter>
          </form>
        </Form>
      )}
    </Dialog>
  );
}
