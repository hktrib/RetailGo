import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { getItemRecommendations } from "../queries";
import Image from "next/image";

async function RemoteImage({ imageURL }: { imageURL: string }) {
  // https://www.freecodecamp.org/news/check-if-a-javascript-string-is-a-url/
  const isValidUrl = (urlString: string) => {
    try {
      return Boolean(new URL(urlString));
    } catch (e) {
      return false;
    }
  };

  if (isValidUrl(imageURL)) {
    return (
      <div className="h-40 w-full bg-gray-100 dark:bg-zinc-700">
        <Image src={imageURL} alt={imageURL} />
      </div>
    );
  }

  return (
    <div className="h-40 w-full bg-gray-100 text-center dark:bg-zinc-700">
      No Image Available.
    </div>
  );
}

async function ItemGrid({ storeId }: { storeId: string }) {
  const recommendedItems = await getItemRecommendations({ store_id: storeId });

  if (!recommendedItems.success) {
    return <div>Nothing to recommend for now!</div>;
  }

  return await recommendedItems.items.map((item: any) => (
    <div
      className="col-span-2 row-span-2 rounded-md bg-gray-200 dark:bg-zinc-800"
      key={item.id}
    >
      <article className="w-52 rounded-md bg-white shadow-sm dark:bg-zinc-600">
        <RemoteImage imageURL={item.imageURL} key={item.id} />
        <div className="p-4">
          <span className="font-medium">
            {item.name} ({item.categoryName})
          </span>
        </div>
      </article>
    </div>
  ));
}

function DashboardPage({ params }: { params: { store_id: string } }) {
  return (
    <main className="flex h-full flex-grow bg-gray-50 dark:bg-zinc-900">
      <div className="mx-auto max-w-6xl flex-grow px-6 py-6 md:px-8 lg:ml-0">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-semibold tracking-wide">Dashboard</h1>
        </div>

        <Tabs defaultValue="overview" className="-mt-2 h-full w-full pt-6">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="activity">Activity</TabsTrigger>
          </TabsList>
          <TabsContent value="overview" className="-mt-10 h-full pb-4 pt-14">
            @ts-expect-error Async Server Component
            <ItemGrid storeId={String(params.store_id)} />
          </TabsContent>
          <TabsContent value="activity">Activity</TabsContent>
        </Tabs>
      </div>
    </main>
  );
}

export default DashboardPage;
