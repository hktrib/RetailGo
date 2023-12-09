"use client";

import { useParams } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useForm } from "react-hook-form";
import {
  createItem,
  updateItem,
} from "@/app/(app-page)/store/[store_id]/inventory/actions";
import { cn } from "@/lib/utils";
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
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ChevronsUpDown, PencilIcon } from "lucide-react";

const formSchema = z.object({
  name: z.string(),
  price: z.coerce.number(),
  quantity: z.coerce.number(),
  category_name: z.string(),
});

export default function ItemDialog({
  item,
  mode = "add",
  categories,
}: {
  item?: Item;
  mode?: string;
  categories: Category[];
}) {
  const params = useParams();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: item ? item.name : undefined,
      price: item ? item.price : undefined,
      quantity: item ? item.quantity : undefined,
      category_name: item ? item.category_name : undefined,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (!params.store_id) return;

    if (mode === "edit") {
      if (!item) return;

      await updateItem({
        item: { id: item.id, ...values },
        store_id: params.store_id as string,
      });

      return;
    }

    await createItem({ item: values, store_id: params.store_id as string });
  };

  const displayCategory = (value: string) => {
    if (!value) return "Select category";
    if (!categories || !categories.length) return value;

    const categoryName = categories.find((category) => category.name === value)
      ?.name;

    if (categoryName) return categoryName;

    return value;
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        {mode === "edit" ? (
          <button className="icon-button">
            <PencilIcon className="text-amber-500 h-5 w-5 p-0" />
          </button>
        ) : (
          <button className="bg-amber-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
            Add item
          </button>
        )}
      </DialogTrigger>

      <DialogContent>
        <DialogHeader>
          <DialogTitle>{mode === "edit" ? "Edit" : "Add"} Item</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-2">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder="Item name" />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="price"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Price</FormLabel>
                  <FormControl>
                    <Input {...field} type="number" placeholder="Item price" />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="quantity"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Quantity</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      type="number"
                      placeholder="Item quantity"
                    />
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="category_name"
              render={({ field }) => (
                <FormItem className="group pt-2">
                  <Popover>
                    <PopoverTrigger asChild>
                      <FormControl>
                        <Button
                          variant="outline"
                          role="combobox"
                          className={cn(
                            "justify-between w-full",
                            !field.value && "text-muted-foreground"
                          )}
                        >
                          {displayCategory(field.value)}
                          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>

                    <PopoverContent className="w-72 p-2">
                      <Input {...field} placeholder="Enter a category" />

                      {categories.length ? (
                        <div className="mt-4 flex flex-col space-y-0.5">
                          {categories.map((category) => (
                            <Button
                              key={category.id}
                              variant="default"
                              value={category.name}
                              // @ts-ignore
                              onClick={(e) => field.onChange(e.target.value)}
                              className="bg-white text-black shadow-none hover:text-black hover:bg-gray-100 text-left justify-start"
                            >
                              {category.name}
                            </Button>
                          ))}
                        </div>
                      ) : (
                        <div className="pt-2 text-center">
                          <span className="text-sm text-muted-foreground">
                            You don{"'"}t have any categories created yet.
                            Create a new one now!
                          </span>
                        </div>
                      )}
                    </PopoverContent>
                  </Popover>
                </FormItem>
              )}
            />

            <DialogFooter className="gap-x-4 pt-2">
              <Button type="submit">
                {mode === "edit" ? "Update" : "Add"}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
