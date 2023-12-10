import { Item, ItemWithoutId } from "@/models/item";
import { useFetch } from "../utils";
import { auth } from "@clerk/nextjs";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { config } from './config';
import { PostEmailInviteModel } from "@/models/staff";


const storeURL = config.serverURL + "store/";



export function SendInvite(storeId: string) {
  const authFetch = useFetch();
  return useMutation({
    mutationFn: (body: PostEmailInviteModel) =>
      authFetch(
        `${storeURL}${storeId}/staff/invite`,
        {
          method: "POST",
          body: JSON.stringify(body),
        },
        {
          "Content-Type": "application/json",
        }
        
      ),
    onError: (err, email, context) => {
      console.log("Error while sending invite to", email, ":", err);
    },
    onSuccess: (email) => {
      console.log("Invite sent successfully to", email);
      // You can add any additional logic you need on successful invite here
    },
  });
}
