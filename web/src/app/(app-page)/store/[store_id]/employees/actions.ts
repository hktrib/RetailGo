"use server";

import { useAuth } from "@clerk/nextjs";
import { revalidatePath } from "next/cache";
import { headers } from "next/headers";

let serverUrl = "https://retailgo-production.up.railway.app"
const env: string = process.env.NODE_ENV;
if (env === "development") {
  serverUrl = "http://localhost:8080"
}

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
  const env: string = process.env.NODE_ENV;

  let serverUrl = `https://retailgo-production.up.railway.app/user/${user_id}/`;
  if (env === "development") {
    serverUrl = `http://localhost:8080/user/${user_id}/`;
  }
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
    if (response.status === 200) {
      revalidatePath("/user/[user]", "page");
      return true;
    }
  } catch (err) {
    console.error(`error updating user ${user_id}: ${err}`);
    return false;
  }
  return false;
};

export const deleteUser = async ({
  user_id,
}: {
  user_id: string;
}) => {
  const env: string = process.env.NODE_ENV;

  let serverUrl = `https://retailgo-production.up.railway.app/user/${user_id}/`;
  if (env === "development") {
    serverUrl = `http://localhost:8080/user/${user_id}/`;
  }
  console.log(`attempting to delete user ${user_id}`);
  try {
    let response = await fetch(serverUrl, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.log(`response: ${response.status} + ${response.statusText}`);
    if (response.status === 200) {
      return true;
    }
  } catch (err) {
    console.error(`error updating user ${user_id}: ${err}`);
    return false;
  }
  return false;
};

/*export const PostInviteUser = async ({
  invite,
  store_id,
}: {
  invite: {
    Email: string;
    Name: string;
  };
  store_id: string;
}) => {
  const { getToken, isLoaded, isSignedIn } = useAuth();
  const fetchUrl = `${serverUrl}/store/${store_id}/staff/invite`;
  console.log(`attempting to invite user ${invite.Email}`);
  try {
    let response = await fetch(fetchUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${await getToken()}`,
      },
      body: JSON.stringify(invite),
    });
    console.log(`response: ${response.status} + ${response.statusText}`);
    if (response.status === 200 || response.status === 201) {
      return true;
    }
  } catch (err) {
    console.error(`error inviting user ${invite}`);
    return false;
  }
  return false;
};*/
