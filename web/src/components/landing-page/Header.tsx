import Link from "next/link";
import { UserButton, auth } from "@clerk/nextjs";

export default function Header() {
  const { userId } = auth();

  return (
    <header className="absolute inset-x-0 top-0 z-50">
      <nav
        className="flex items-center justify-between p-6 lg:px-8"
        aria-label="Global"
      >
        <div className="flex lg:flex-1">
          <Link href="/" className="-m-1.5 p-1.5">
            <span className="sr-only">RetailGo</span>
            <span className="font-bold tracking-wide">RetailGo</span>
          </Link>
        </div>

        <div className="hidden lg:flex lg:flex-1 lg:justify-end items-center space-x-4">
          {userId ? (
            <>
              <Link
                href="/store"
                className="bg-amber-500 text-white font-medium px-3 py-1.5 rounded-md text-sm"
              >
                My stores
              </Link>
              <Link
                href="/registrationForm"
                className="bg-amber-600 text-white font-medium px-3 py-1.5 rounded-md text-sm"
              >
                Add Store
              </Link>
              <UserButton afterSignOutUrl="/" />

            </>
          ) : (
            <>
              <Link
                href="/sign-in"
                className="text-sm font-semibold leading-6 text-gray-900"
              >
                Sign in
              </Link>
              <Link
                href="/sign-up"
                className="bg-amber-500 text-white font-medium px-3 py-1.5 rounded-md text-sm"
              >
                Get started <span aria-hidden="true">&rarr;</span>
              </Link>
            </>
          )}
        </div>
      </nav>
    </header>
  );
}
