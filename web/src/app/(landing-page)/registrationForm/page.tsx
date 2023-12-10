"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { useFetch } from "@/lib/utils";
import { useUser } from "@clerk/nextjs";
import { useRouter } from 'next/navigation'
import {config} from "@/lib/hooks/config";
import {createStore} from "./actions"

// type Member = {
//   firstName: string;
//   lastName: string;
//   email: string;
//   role: string;
// }

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Router } from "next/router";
import { toast } from "react-toastify";

const formSchema = z.object({
  storeName: z.string().min(0),
  phoneNumber: z.string().min(7, { message: "Invalid phone number" }),
  address1: z.string().min(0),
  address2: z.string().optional(),
  businessType: z.string().min(0),
});

const businessTypes = [
  "Clothing",
  "Grocery",
  "Convenience",
  "Departmnet",
  "Restaurant",
  "Other",
];

export default function RegistrationForm() {
  const { user } = useUser();
  const router = useRouter()

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  console.log(config.serverURL)

  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    // not doing anything with `address2` yet

    const { storeName, phoneNumber, address1, address2, businessType } = values;

    // data to be sent in POST request body
    const postData = {
      store_name: storeName || "",
      store_phone: phoneNumber || "",
      store_address: address1 || "",
      store_type: businessType || "",
      owner_email: user?.emailAddresses[0]?.emailAddress || "",
    };

    try {
      let response : Boolean = await createStore({postData});

      if (response === true) {
        // router.refresh()
        toast.success("Store created successfully!", {
          position: toast.POSITION.TOP_RIGHT,
          autoClose: 10000,
        });
        router.push("/store?refresh='refresh'") 
      } else {
        toast.error("Error creating store!", {
          position: toast.POSITION.TOP_RIGHT,
          autoClose: 10000,
        });
        throw "Failed to create store"
      }

    } catch (error) {
      toast.error("Error creating store!", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000,
      });
      console.error("Error making create store request:", error);
    }
  };

  return (
    <div className="relative isolate px-6 pt-14 lg:px-8 flex-1 flex flex-col justify-center items-center">
      <div className="mx-auto max-w-2xl w-full bg-gray-50 py-16 px-12 rounded-xl">
        <h1 className="text-center font-bold text-3xl">
          Register your business
        </h1>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="space-y-8 mt-12"
          >
            <FormField
              control={form.control}
              name="storeName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Store name</FormLabel>
                  <FormControl>
                    <Input placeholder="Store name" {...field} />
                  </FormControl>
                  {/* <FormDescription>
                This is your public display name.
              </FormDescription> */}
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="phoneNumber"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Phone number</FormLabel>
                  <FormControl>
                    <Input placeholder="Store phone number" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="address1"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Address line 1</FormLabel>
                  <FormControl>
                    <Input placeholder="Address line 1" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="address2"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Address line 2</FormLabel>
                  <FormControl>
                    <Input placeholder="Address line 2" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="businessType"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Business type</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a business type" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {businessTypes.map((businessType) => (
                        <SelectItem key={businessType} value={businessType}>
                          {businessType}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />

            <div className="text-right">
              <Button type="submit">Submit</Button>
            </div>
          </form>
        </Form>
      </div>
    </div>
  );
}
