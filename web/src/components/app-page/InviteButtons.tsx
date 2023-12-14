// Importing necessary dependencies and modules
"use client"
import { GetStoreByUUID } from '@/app/(invite-page)/invite/[[...invite]]/queries';
import { PostJoinStore } from '@/lib/hooks/user';
import { RedirectToSignIn, SignedIn, SignedOut, useUser } from '@clerk/nextjs';
import React from 'react';

// Defining the functional component InviteButtons
export default function InviteButtons({ store_uuid }: { store_uuid: string }) {
    // Accessing the user object from the useUser hook
    const { user } = useUser();

    // Declaring the OnDecline function
    function OnDecline() {
        console.log("Declined");
        // Handle decline action
    }

    // Creating a joinMutation object using the PostJoinStore hook
    const joinMutation = PostJoinStore(store_uuid);

    // Defining the onSubmit function
    const onSubmit = () => {
        // Checking if the user object exists
        if (!user) return;

        // Invoking the joinMutation function with the user's id as the argument
        joinMutation.mutate(user.id);

        // Reloading the user object
        user.reload();
    };

    // Rendering the component
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