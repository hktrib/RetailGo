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
import { useState } from "react";
import { ToastContainer, toast } from "react-toastify";
import { useSelectedStore } from "@/components/storeprovider";
import { useParams  } from "next/navigation";


export default function InviteEmployee() {
  const [isDialogOpen, setDialogOpen] = useState(false); // State to control the dialog
  const { selectedStore, selectStore } = useSelectedStore();
  const params = useParams()

  const formSchema = z.object({
    email: z.string().regex(/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/),
    name: z.string(),
  });

  const notify = () => {
    toast.success("Invite sent successfully!", {
      position: toast.POSITION.TOP_RIGHT,
      autoClose: 10000
    });
  }

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  const id = params.store_id;
  const inviteMutation = SendInvite(id.toString());

  const onSubmit = form.handleSubmit((data: any) => {
    console.log(JSON.stringify(data));
    inviteMutation.mutate(data);
    if (inviteMutation.isSuccess) {
      setDialogOpen(false);
      toast.success("Invite sent successfully!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000
      });
    }else{
      toast.error("Error sending invite!", {
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
        <button className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md"
          onClick={() => setDialogOpen(true)}
        >
          Invite
        </button>
      </DialogTrigger>
      <Form {...form}>
        <form>
          <DialogContent>
            <form onSubmit={onSubmit} {...form}>
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
