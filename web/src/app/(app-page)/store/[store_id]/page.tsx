import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { getItemRecommendations } from "../queries";

async function RemoteImage({imageURL}: {imageURL: string}){
  
  // https://www.freecodecamp.org/news/check-if-a-javascript-string-is-a-url/
  const isValidUrl = (urlString: string) => {
    try { 
      return Boolean(new URL(urlString)); 
    }
    catch(e){ 
      return false; 
    }
  }

  if (isValidUrl(imageURL)){
    return (<div className="w-full h-40 bg-gray-100">
        <img src = {imageURL} />
      </div>)
  }
  return (
    <div className="w-full h-40 bg-gray-100 text-center">
      No Image Available.
    </div>
  ) 
}

async function ItemGrid({storeId}: {storeId: string}){
  
  const recommendedItems = await getItemRecommendations({store_id: storeId})

  if (!recommendedItems.success){
    return (<div>
      Nothing to recommend for now!
    </div>)
  }

  return (
    recommendedItems.items.map(
    (item: any) => (
    <div className="bg-gray-200 rounded-md col-span-2 row-span-2" key = {item.id}>
      <article className="bg-white rounded-md shadow-sm w-52">
        <RemoteImage imageURL = {item.imageURL} key = {item.id}/>
        <div className="p-4">
          <span className="font-medium">{item.name} ({item.categoryName})</span>
        </div>
      </article>
    </div>)
  )
  )
}

function DashboardPage({params}: {params: {store_id: string}}) {
  return (
    <main className="bg-gray-50 h-full flex-grow flex">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0 flex-grow">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold">Dashboard</h1>
        </div>

        <Tabs defaultValue="overview" className="w-full pt-6 -mt-2 h-full">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="activity">Activity</TabsTrigger>
          </TabsList>
          <TabsContent value="overview" className="pt-14 h-full -mt-10 pb-4">
            <ItemGrid storeId={String(params.store_id)}/>
          </TabsContent>
          <TabsContent value="activity">Activity</TabsContent>
        </Tabs>
      </div>
    </main>
  );
}

export default DashboardPage;