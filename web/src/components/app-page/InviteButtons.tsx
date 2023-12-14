"use client"
import { GetStoreByUUID } from '@/app/(invite-page)/invite/[[...invite]]/queries';
import { PostJoinStore } from '@/lib/hooks/user';
import { RedirectToSignIn, SignedIn, SignedOut, useUser } from '@clerk/nextjs';
import React from 'react';

export default function InviteButtons({store_uuid}: {store_uuid: string}){
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
    );


}