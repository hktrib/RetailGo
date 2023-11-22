"use client";

import {useState} from "react"
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, SubmitHandler } from "react-hook-form";
import { useFetch } from "../../lib/utils"
import { useCreateItem, useEditItem } from "@/app/(app-page)/app/hooks/items";

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
import { Item, ItemWithoutId } from "@/models/item";
import { useEffect } from "react";
import { PencilIcon } from "lucide-react";

const formSchema = z.object({
  name: z.string(),
  price: z.coerce.number(),
  quantity: z.coerce.number(),
  category: z.string(),
});

export default function ItemDialog({ item, mode = 'add' }: { item: Item, mode?: string }) {

  const form = useForm<z.infer <typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });;

  const createItemMutation = useCreateItem("1")
  const editItemMutation = useEditItem("1")

  const [inputName, setInputName] = useState(item.name)
  const [inputCategory, setInputCategory] = useState(item.category)
  const [inputPrice, setInputPrice] = useState(item.price)
  const [inputQuantity, setInputQuantity] = useState(item.quantity)

  return (
    <Dialog>
      <DialogTrigger asChild>
        {mode === "edit" ? (
          <button className="icon-button">
            <PencilIcon style={{ color: "orange" }} className="h-5 w-5 p-0"></PencilIcon>
          </button>
        ) : (
          <button className="bg-amber-500 text-sm px-3 py-1.5 text-white font-medium rounded-md">
            Add item
          </button>
        )}
      </DialogTrigger>

      <Form {...form}>
        <form>
          <DialogContent>
            <form onSubmit = {form.handleSubmit((data) => {
              console.log("Data:", data)
              return mode === "edit" ? editItemMutation.mutate({...item, ...data}) : createItemMutation.mutate(data)
            }
              , (data) => console.log("Error:", data, "InputName:", inputName, "inputCategory:", inputCategory, "InputQuantity:", inputQuantity))}>
            <DialogHeader>
              <DialogTitle>
              {mode === "edit" ? 'Edit' : 'Add'} Item
              </DialogTitle>
            </DialogHeader>

            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field} value = {inputName} onChange={(e) => setInputName(e.currentTarget.value)}/>
                  </FormControl>
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="category"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Category</FormLabel>
                  <FormControl>
                    <Input {...field} value = {inputCategory} onChange={(e) => setInputCategory(e.currentTarget.value)}/>
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
                    <Input {...field} type = "number" value = {inputPrice} onChange={(e) => setInputPrice(parseFloat(e.currentTarget.value))}/>
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
                    <Input {...field} type = "number" value = {inputQuantity} onChange={(e) => setInputQuantity(parseInt(e.currentTarget.value))}/>
                  </FormControl>
                </FormItem>
              )}
            />

            <DialogFooter className="gap-x-4">
              <button type="submit">{mode === "edit" ? 'Save' : 'Add'}</button>
            </DialogFooter>
            </form>
          </DialogContent>
        </form>
      </Form>
    </Dialog>
  );
}
