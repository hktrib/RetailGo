import Image from "next/image";
import { getItemRecommendations } from "../queries";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { HomeIcon } from "lucide-react";

function DashboardPage({ params }: { params: { store_id: string } }) {
  return (
    <main className="h-full flex-grow">
      <div className="flex h-16 items-center justify-between px-4 py-5 md:px-6 xl:px-8">
        <div className="mt-1 flex items-center gap-x-3">
          <HomeIcon className="h-5 w-5 text-gray-800 dark:text-zinc-200" />
          <h1 className="text-xl font-medium tracking-wide">Dashboard</h1>
        </div>
      </div>

      <hr className="mb-4 border-gray-100 dark:border-zinc-800" />

      <div className="px-4 md:px-6 xl:px-8">
        <Tabs defaultValue="overview" className="-mt-2 h-full w-full pt-4">
          <TabsList className="dark:border-zinc-800">
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="activity">Activity</TabsTrigger>
          </TabsList>
          <TabsContent value="overview" className="mt-6">
            <h2 className="text-lg font-medium dark:text-zinc-200">
              Recommended Items
            </h2>
            <ItemRecommendations storeId={String(params.store_id)} />
          </TabsContent>

          <TabsContent value="activity">Activity</TabsContent>
        </Tabs>
      </div>
    </main>
  );
}

async function ItemRecommendations({ storeId }: { storeId: string }) {
  const recommendedItems = await getItemRecommendations({ store_id: storeId });

  if (!recommendedItems.success || !recommendedItems.items) {
    return (
      <p className="text-sm leading-6 text-gray-700 dark:text-zinc-300">
        Nothing to recommend for now!
      </p>
    );
  }

  return (
    <section className="mt-2 flex flex-row flex-wrap gap-4">
      {recommendedItems.items.map((item: ItemRecommendation) => (
        <article
          className="rounded-md bg-gray-200 shadow-sm dark:bg-zinc-800"
          key={item.name}
        >
          <div className="w-52 rounded-md bg-white shadow-sm dark:bg-zinc-800">
            <RemoteImage imageURL={item.imageURL} />

            <div className="flex flex-col p-4">
              <p className="text-sm font-medium">{item.name}</p>
              <p className="mt-0.5 text-xs dark:text-zinc-300">
                {item.categoryName}
              </p>
            </div>
          </div>
        </article>
      ))}
    </section>
  );
}

async function RemoteImage({ imageURL }: { imageURL: string }) {
  // https://www.freecodecamp.org/news/check-if-a-javascript-string-is-a-url/
  const isValidUrl = (urlString: string) => {
    try {
      return Boolean(new URL(urlString));
    } catch (e) {
      return false;
    }
  };

  return (
    <div className="flex h-40 w-full items-center justify-center rounded-t-md bg-gray-100 dark:bg-zinc-700">
      {isValidUrl(imageURL) ? (
        <Image src={imageURL} alt={imageURL} />
      ) : (
        <span className="text-xs dark:text-zinc-300">No image available.</span>
      )}
    </div>
  );
}

export default DashboardPage;
