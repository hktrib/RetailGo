import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from "./config";
import { useRouter } from "next/navigation";
import router from "next/router";
import toast from "react-hot-toast";

const storeURL = config.serverURL + "/user/";

/**
 * Custom hook to check if the user has a store.
 * 
 * @returns {Object} - The result of the query.
 *   - isLoading: boolean - Indicates if the query is in progress.
 *   - isError: boolean - Indicates if an error occurred during the query.
 *   - data: any - The data returned from the query.
 *   - error: any - The error object if an error occurred.
 */
export function HasStore() {
  const authFetch = useFetch();
  return useQuery({
    queryKey: ["hasStore"],
    queryFn: () => authFetch(storeURL + "store", {}, {}, true),
  });
}

/**
 * Mutation function to join a store.
 * 
 * @param {string} storeId - The ID of the store to join.
 * 
 * @returns {Object} - The result of the mutation.
 *   - mutate: function - Function to trigger the mutation.
 *   - isLoading: boolean - Indicates if the mutation is in progress.
 *   - isError: boolean - Indicates if an error occurred during the mutation.
 *   - error: any - The error object if an error occurred.
 *   - reset: function - Function to reset the mutation state.
 */
export function PostJoinStore(storeId: string) {
  const router = useRouter();
  const authFetch = useFetch();
  return useMutation({
    mutationFn: (clerkId: string) =>
      authFetch(
        `${storeURL}join`,
        {
          method: "POST",
          body: JSON.stringify({ ClerkUserID: clerkId, StoreId: storeId }),
        },
        {
          "Content-Type": "application/json",
        },
      ),
    onError: (err, email) => {
      console.log("Error while sending invite to", email, ":", err);
    },
    onSuccess: (email) => {
      toast.success("Successfully joined store!");
      console.log("Invite sent successfully to", email);
      router.push("/store");
      // You can add any additional logic you need on successful invite here
    },
  });
}

/**
 * Mutation function to update a user.
 * 
 * @param {string} user_id - The ID of the user to update.
 * 
 * @returns {Object} - The result of the mutation.
 *   - mutate: function - Function to trigger the mutation.
 *   - isLoading: boolean - Indicates if the mutation is in progress.
 *   - isError: boolean - Indicates if an error occurred during the mutation.
 *   - error: any - The error object if an error occurred.
 *   - reset: function - Function to reset the mutation state.
 */
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
        },
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

/**
 * Mutation function to delete a user.
 * 
 * @returns {Object} - The result of the mutation.
 *   - mutate: function - Function to trigger the mutation.
 *   - isLoading: boolean - Indicates if the mutation is in progress.
 *   - isError: boolean - Indicates if an error occurred during the mutation.
 *   - error: any - The error object if an error occurred.
 *   - reset: function - Function to reset the mutation state.
 */
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
        },
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
