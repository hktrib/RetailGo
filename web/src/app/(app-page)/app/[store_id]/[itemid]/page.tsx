export default function Page({ params }: { params: { itemid: string } }) {
  const id = params.itemid;

  return (
    <main className="bg-gray-50 h-full flex-grow">
      <div className="py-6 px-6 md:px-8 max-w-6xl mx-auto lg:ml-0">
        <div className="flex items-center justify-between ">
          <h1 className="text-2xl font-bold">{params.itemid}</h1>
        </div>
        <hr className="my-4" />
        <div className="mt-6"></div>
      </div>
    </main>
  );
}
