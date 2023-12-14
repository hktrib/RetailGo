import Link from "next/link";
import { UserButton, auth } from "@clerk/nextjs";
import Image from "next/image";

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
            <Image
              src="/retailgo-black.svg"
              width={100}
              height={100}
              alt="RetailGo"
            />
          </Link>
        </div>

        <div className="hidden items-center space-x-4 lg:flex lg:flex-1 lg:justify-end">
          {userId ? (
            <>
              <Link
                href="/store"
                className="rounded-md bg-sky-500 px-3 py-1.5 text-sm font-medium text-white"
              >
                My stores
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
                className="rounded-md bg-sky-500 px-3 py-1.5 text-sm font-medium text-white"
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
