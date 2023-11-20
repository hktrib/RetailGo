import { NextRequest, NextResponse } from "next/server";
import { headers } from "next/headers";
import { clerkClient } from "@clerk/nextjs";
import { Webhook } from "svix";
import type { User } from "@clerk/nextjs/api";
import { useAuth } from "@clerk/nextjs";

type UnwantedKeys = "emailAddresses" | "primaryEmailAddressId" | "id";

interface UserInterface extends Omit<User, UnwantedKeys> {
  email_addresses: {
    email_address: string;
    id: string;
  }[];
  primary_email_address_id: string;
  id: string;
}

type Event = {
  data: UserInterface;
  object: "event";
  type: EventType;
};
type EventType = "user.created" | "user.updated" | "*";

const webhookSecret: string = process.env.CLERK_WEBHOOK_SECRET || "";

export async function POST(req: NextRequest) {
  const payload = await req.json();
  const payloadStr = JSON.stringify(payload);

  // get headers
  const headerPayload = headers();
  const svixId = headerPayload.get("svix-id");
  const svixIdTimeStamp = headerPayload.get("svix-timestamp");
  const svixSignature = headerPayload.get("svix-signature");

  // err if no headers
  if (!svixId || !svixIdTimeStamp || !svixSignature) {
    return new Response("Error occured", {
      status: 400,
    });
  }

  const svixHeaders = {
    "svix-id": svixId,
    "svix-timestamp": svixIdTimeStamp,
    "svix-signature": svixSignature,
  };
  const wh = new Webhook(webhookSecret);
  let evt: Event | null = null;
  try {
    evt = wh.verify(payloadStr, svixHeaders) as Event;
  } catch (_) {
    console.error("Error verifying webhook");
    return new Response("Error verifying webhook", {
      status: 400,
    });
  }

  const eventType: EventType = evt.type;
  if (eventType === "user.created") {
    const { email_addresses, primary_email_address_id, id } = evt.data;
    // parse email
    const emailObject = email_addresses?.find(
      (email) => email.id === primary_email_address_id
    );

    if (!emailObject)
      return new Response("Error locating user", { status: 400 });

    const email = emailObject.email_address;

    // add user to db:
    fetch("https://retailgo-production.up.railway.app/create/user", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    })
    
  }

  return NextResponse.json({ status: 200 });
}
