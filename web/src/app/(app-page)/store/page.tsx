import { auth } from "@clerk/nextjs";
import Link from "next/link";
import { notFound, redirect } from "next/navigation";

type StoreMetadata = {
  Permission_level: number;
  id: number;
  storename: string;
};

export default function StoreViewPage() {
  const { sessionClaims } = auth();

  if (!sessionClaims || !sessionClaims.publicMetadata)
    return redirect("/registrationForm");

  const publicMetadata = sessionClaims.publicMetadata as {
    stores?: StoreMetadata[];
  };
  if (!publicMetadata.stores) redirect("/registrationForm");

  return (
    <main className="bg-gray-50 h-full flex-grow flex">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0 flex-grow">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold">My stores</h1>
        </div>

        <hr className="my-4" />

        <section className="flex flex-row flex-wrap gap-4">
          {publicMetadata.stores.map((store: StoreMetadata) => (
            <StoreCard
              key={store.id}
              id={store.id}
              storeName={store.storename}
            />
          ))}
        </section>
      </div>
    </main>
  );
}

const StoreCard = ({ id, storeName }: { id: number; storeName: string }) => {
  const createUrl = (id: number) => {
    const baseUrl = `/store/${id}`;
    return `${baseUrl}`;
  };

  return (
    <Link href={createUrl(id)}>
      <article className="bg-white rounded-md shadow-sm w-52">
        <div className="w-full h-40 bg-gray-100" />
        <div className="p-4">
          <span className="font-medium">{storeName}</span>
        </div>
      </article>
    </Link>
  );
};
