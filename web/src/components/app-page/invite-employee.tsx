"use client"

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



export default function InviteEmployee() {
  const formSchema = z.object({
    email: z.string()
  });

const form = useForm<z.infer <typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });;

  const onSubmit: SubmitHandler<z.infer<typeof formSchema>> = async (data) => {
    try {
      const response = await fetch('/ai/send-invite', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
  
      if (response.ok) {
        // Handle success - maybe show a message to the user
      } else {
        // Handle error - maybe show an error message
      }
    } catch (error) {
      // Handle error - maybe show an error message
    }
  };

  return (    <Dialog>
    <DialogTrigger asChild>
      <button className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
        Invite
      </button>
    </DialogTrigger>

    <Form {...form} >
      <form>
        <DialogContent>
          <form onSubmit={form.handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>
              Invite employee
            </DialogTitle>
          </DialogHeader>

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input {...field} placeholder="person@google.com"/>
                </FormControl>
              </FormItem>
            )}
          />
          <DialogFooter className="gap-x-4">
            <button type="submit" className="bg-blue-500 text-sm px-3 py-1.5 text-white font-medium rounded-md mt-5">Invite</button>
          </DialogFooter>
          </form>
        </DialogContent>
      </form>
    </Form>
  </Dialog>)

}