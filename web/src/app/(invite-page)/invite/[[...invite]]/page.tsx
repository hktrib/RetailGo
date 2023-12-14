import { RedirectToSignIn, SignIn, SignedIn, SignedOut, useUser } from "@clerk/nextjs";
import { GetStoreByUUID } from "./queries";
import Header from "@/components/landing-page/Header";
import Footer from "@/components/landing-page/footer";
import InviteButtons from "@/components/app-page/InviteButtons";

export default async function InvitePage({
  searchParams,
}: {
  searchParams: { code: string };
}) {
  let store = await GetStoreByUUID({ uuid: searchParams.code });
  return (
    <div className="flex flex-col justify-between h-screen"> {/* Added flex container */}
      <Header />
      <div className="flex-grow flex flex-col justify-center items-center px-6 pt-14 lg:px-8"> {/* Adjusted for centering */}
        <SignedIn>
          <h1 className="text-center text-2xl">
            <span className="font-bold">&quot;{store.store.store_name}&quot;</span> has
            invited you to their business
          </h1>
          <InviteButtons store_uuid={searchParams.code} />
        </SignedIn>
        <SignedOut>
          <RedirectToSignIn />
        </SignedOut>
      </div>
      <Footer />
    </div>
  );
}
