import { useFetch } from "../utils";
import { useMutation } from "@tanstack/react-query";
import { config } from "./config";
import { PostEmailInviteModel } from "@/models/staff";
import toast from "react-hot-toast";

const storeURL = config.serverURL + "/store/";

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
        },
      ),
    onError: (err, email, context) => {
      console.error("Error while sending invite to", email, ":", err);
      toast.error("Error sending invite!");
    },
    onSuccess: (email) => {
      console.log("Invite sent successfully to", email);
      toast.success("Successfully sent invite!");
      // You can add any additional logic you need on successful invite here
    },
  });
}
