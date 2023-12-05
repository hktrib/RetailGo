"use client";

import { useRouter } from "next/navigation";
import { PostJoinStore } from "@/lib/hooks/user";
import { SignedIn, SignedOut, useUser } from "@clerk/nextjs";
import { useSearchParams } from "next/navigation";

const InvitePage = () => {
  const router = useRouter();
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
    router.push("/app");
  };

  return (
    <div className="relative isolate px-6 pt-14 lg:px-8 flex-1 flex flex-col justify-center items-center">
      <SignedIn>
        <h1 className="text-center text-2xl">
          <span className="font-bold">&quot;Stores are us&quot;</span> has
          invited you to their business
        </h1>

        <div className="py-10 flex items-center justify-center space-x-5">
          <button
            onClick={() => onSubmit()}
            className="bg-green-500 hover:bg-green-600 active:bg-green-700 focus:outline-none focus:ring focus:ring-green-300 text-white font-medium px-3 py-3 rounded-md text-sm"
          >
            Accept Invitation
          </button>
          <button
            onClick={OnDecline}
            className="bg-red-500 hover:bg-red-600 active:bg-red-700 focus:outline-none focus:ring focus:ring-red-300 text-white font-medium px-3 py-3 rounded-md text-sm"
          >
            Decline Invitation
          </button>
        </div>
      </SignedIn>
      <SignedOut>
        <div>You are signed Out</div>
      </SignedOut>
    </div>
  );
};

export default InvitePage;
