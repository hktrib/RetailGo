import Link from "next/link";

export default function Header() {
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
          <Link
            href="/sign-in"
            className="text-sm font-semibold leading-6 text-gray-900"
          >
            Sign in
          </Link>
          <Link
            href="/inventory"
            className="bg-sky-500 text-white font-semibold p-2 rounded-md text-sm"
          >
            Inventory <span aria-hidden="true">&rarr;</span>
          </Link>
          <Link
            href="/sign-up"
            className="bg-amber-500 text-white font-semibold p-2 rounded-md text-sm"
          >
            Get started <span aria-hidden="true">&rarr;</span>
          </Link>
        </div>
      </nav>
    </header>
  );
}
