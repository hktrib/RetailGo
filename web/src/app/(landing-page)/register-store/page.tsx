"use client";

import { useRouter } from "next/navigation";
import { useUser } from "@clerk/nextjs";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { createStore } from "./actions";

import { toast } from "react-toastify";
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
  const router = useRouter();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  // not doing anything with `address2` yet
  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    const { storeName, phoneNumber, address1, address2, businessType } = values;

    const postData = {
      store_name: storeName,
      store_phone: phoneNumber,
      store_address: address1,
      store_type: businessType,
      owner_email: user?.emailAddresses[0]?.emailAddress || "",
    };

    try {
      let response: Boolean = await createStore({ postData });

      if (response === true) {
        toast.success("Store created successfully!", {
          position: toast.POSITION.TOP_RIGHT,
          autoClose: 10000,
        });
        router.push("/store");
      } else {
        toast.error("Error creating store!", {
          position: toast.POSITION.TOP_RIGHT,
          autoClose: 10000,
        });
        throw "Failed to create store";
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
    <div className="relative isolate flex flex-1 flex-col items-center justify-center px-6 pt-14 lg:px-8">
      <div className="mx-auto w-full max-w-2xl rounded-xl bg-gray-50 px-12 py-16">
        <h1 className="text-center text-3xl font-bold">
          Register your business
        </h1>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="mt-12 space-y-8"
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
