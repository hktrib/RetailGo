"use server";

import { revalidatePath } from "next/cache";
import { config } from "@/lib/hooks/config";
export const updateUser = async ({
  user,
  user_id,
}: {
  user: {
    first_name: string;
    last_name: string;
    email: string;
  };
  user_id: string;
}) => {

  let serverUrl = `${config.serverURL}/user/${user_id}/`;

  console.log(`attempting to update user ${user_id}`);

  try {
    let response = await fetch(serverUrl, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    console.log(`response: ${response.status} + ${response.statusText}`);
    if (response.status === 200 || response.status === 201 || response.status === 202) {
      revalidatePath("/user/[user]", "page");
      return true;
    }
  } catch (err) {
    console.error(`error updating user ${user_id}: ${err}`);
    return false;
  }
  return false;
};
