"use client";

import { useRouter } from "next/navigation";
import { useUser } from "@clerk/nextjs";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { createStore } from "./actions";

import toast from "react-hot-toast";
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
import { Store } from "lucide-react";

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
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  const router = useRouter();
  const { user } = useUser();
  if (!user) return null;

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

      if (response) {
        toast.success(`Successfully created ${storeName}!`);
        user.reload();
        router.push("/store");
      } else {
        toast.error("Error creating store!");
        throw "Failed to create store";
      }
    } catch (error) {
      toast.error("Error creating store!");
      console.error("Create store request error:", error);
    }
  };

  return (
    <div className="relative isolate flex flex-1 flex-col items-center justify-center overflow-hidden bg-white px-6 pt-14 lg:px-8">
      <div className="absolute left-1/2 top-4 -z-10 h-[1026px] w-[1026px] -translate-x-1/2 stroke-gray-300/70 [mask-image:linear-gradient(to_bottom,white_20%,transparent_75%)] sm:top-16 lg:-top-16 xl:top-8 xl:ml-0">
        <svg
          viewBox="0 0 1026 1026"
          fill="none"
          aria-hidden="true"
          className="animate-spin-slow absolute inset-0 h-full w-full"
        >
          <path
            d="M1025 513c0 282.77-229.23 512-512 512S1 795.77 1 513 230.23 1 513 1s512 229.23 512 512Z"
            stroke="#D4D4D4"
            stroke-opacity="0.7"
          />
          <path
            d="M513 1025C230.23 1025 1 795.77 1 513"
            stroke="url(#:S2:-gradient-1)"
            stroke-linecap="round"
          />
          <defs>
            <linearGradient
              id=":S2:-gradient-1"
              x1="1"
              y1="513"
              x2="1"
              y2="1025"
              gradientUnits="userSpaceOnUse"
            >
              <stop stop-color="#06b6d4"></stop>
              <stop offset="1" stop-color="#06b6d4" stop-opacity="0"></stop>
            </linearGradient>
          </defs>
        </svg>
        <svg
          viewBox="0 0 1026 1026"
          fill="none"
          aria-hidden="true"
          className="animate-spin-reverse-slower absolute inset-0 h-full w-full"
        >
          <path
            d="M913 513c0 220.914-179.086 400-400 400S113 733.914 113 513s179.086-400 400-400 400 179.086 400 400Z"
            stroke="#D4D4D4"
            stroke-opacity="0.7"
          />
          <path
            d="M913 513c0 220.914-179.086 400-400 400"
            stroke="url(#:S2:-gradient-2)"
            stroke-linecap="round"
          />
          <defs>
            <linearGradient
              id=":S2:-gradient-2"
              x1="913"
              y1="513"
              x2="913"
              y2="913"
              gradientUnits="userSpaceOnUse"
            >
              <stop stop-color="#06b6d4"></stop>
              <stop offset="1" stop-color="#06b6d4" stop-opacity="0"></stop>
            </linearGradient>
          </defs>
        </svg>
      </div>

      <div className="mx-auto w-full max-w-xl rounded-3xl bg-gray-50 px-6 py-16 shadow md:px-16">
        <div className="flex flex-col items-center">
          <Store className="h-8 w-8 text-gray-600" />

          <div className="mt-4 text-center">
            <h1 className="text-2xl font-semibold text-gray-900">
              Register your business
            </h1>
            <p className="text-sm leading-6 text-gray-600">
              Enter your business details below
            </p>
          </div>
        </div>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="mt-4 space-y-6"
          >
            <FormField
              control={form.control}
              name="storeName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-gray-600">Store name</FormLabel>
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
                  <FormLabel className="text-gray-600">Phone number</FormLabel>
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
                  <FormLabel className="text-gray-600">
                    Address line 1
                  </FormLabel>
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
                  <FormLabel className="text-gray-600">
                    Address line 2
                  </FormLabel>
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
                  <FormLabel className="text-gray-600">Business type</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger className="bg-white text-black">
                        <SelectValue placeholder="Select a business type" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent className="bg-white text-black">
                      {businessTypes.map((businessType) => (
                        <SelectItem
                          key={businessType}
                          value={businessType}
                          className="focus:bg-[#f3f4f6] focus:text-[#111827]"
                        >
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
    </main>
  );
}
