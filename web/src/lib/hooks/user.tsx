import { Item, ItemWithoutId } from "@/models/item";
import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from './config';
import { PostJoinStoreModel } from "@/models/user";
import router from "next/router";
import { toast } from "react-toastify";


const storeURL = config.serverURL + "/user/";



export function HasStore() {
  // const queryClient = useQueryClient()
  const authFetch = useFetch();
  return useQuery({
    queryKey: ["hasStore"],
    queryFn: () => authFetch(storeURL + "store", {}, {}, true),
  });
}

export function PostJoinStore(storeId: string) {
  const authFetch = useFetch();
  return useMutation({
    mutationFn: (clerkId: string) =>
    authFetch(
      `${storeURL}join`,
      {
        method: "POST",
        body: JSON.stringify({"ClerkUserID": clerkId, "StoreId": storeId}),
      },
      {
        "Content-Type": "application/json",
      }
      
    ),
    onError: (err, email) => {
      console.log("Error while sending invite to", email, ":", err);
    },
    onSuccess: (email) => {
      toast.success("Successfully! joined store", {
        position: toast.POSITION.TOP_RIGHT,
        autoClose: 10000,
      });
      console.log("Invite sent successfully to", email);
      router.push("/store");
      // You can add any additional logic you need on successful invite here
    },
  });
}

export function PutUser(user_id: string) {
  const authFetch = useFetch();
  return useMutation({
    mutationFn: (Employee) =>
    authFetch(
      `${storeURL}${user_id}`,
      {
        method: "PUT",
        body: JSON.stringify(Employee),
      },
      {
        "Content-Type": "application/json",
      }
    ),
    onError: (err, email, context) => {
      console.log("Error updating user", ":", err);
    },
    onSuccess: (email) => {
      console.log("Updated User", email);
      // You can add any additional logic you need on successful invite here
    },
  });
}

export function DeleteUser() {
  const authFetch = useFetch();
  return useMutation({
    mutationFn: (user_id: string) =>
    authFetch(
      `${storeURL}${user_id}`,
      {
        method: "DELETE",
      },
      {
        "Content-Type": "application/json",
      }
    ),
    onError: (err, email, context) => {
      console.log("Error deleting user", ":", err);
    },
    onSuccess: (email) => {
      console.log("Deleted User", email);
      // You can add any additional logic you need on successful invite here
    },
  });
}


export { config };

