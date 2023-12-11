"use client";

import { useSearchParams } from "next/navigation";
import {
  RedirectToSignIn,
  SignIn,
  SignedIn,
  SignedOut,
  useUser,
} from "@clerk/nextjs";
import { PostJoinStore } from "@/lib/hooks/user";

const InvitePage = () => {
  const searchParams = useSearchParams();
  const store_uuid = searchParams.get("code") || "";
  const { user } = useUser();

  function OnDecline() {
    console.log("Declined");
    // Handle decline action
  }
  const joinMutation = PostJoinStore(store_uuid);
  const onSubmit = () => {
    if (!user) return;
    joinMutation.mutate(user.id);
    user.reload();
  };

  return (
    <div className="relative isolate flex flex-1 flex-col items-center justify-center px-6 pt-14 lg:px-8">
      <SignedIn>
        <h1 className="text-center text-2xl">
          <span className="font-bold">&quot;Stores are us&quot;</span> has
          invited you to their business
        </h1>

        <div className="flex items-center justify-center space-x-5 py-10">
          <button
            onClick={() => onSubmit()}
            className="rounded-md bg-green-500 px-3 py-3 text-sm font-medium text-white hover:bg-green-600 focus:outline-none focus:ring focus:ring-green-300 active:bg-green-700"
          >
            Accept Invitation
          </button>
          <button
            onClick={OnDecline}
            className="rounded-md bg-red-500 px-3 py-3 text-sm font-medium text-white hover:bg-red-600 focus:outline-none focus:ring focus:ring-red-300 active:bg-red-700"
          >
            Decline Invitation
          </button>
        </div>
      </SignedIn>

      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </div>
  );
};

export default InvitePage;
