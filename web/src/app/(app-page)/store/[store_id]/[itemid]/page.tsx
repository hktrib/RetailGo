export default function Page({ params }: { params: { itemid: string } }) {
  const id = params.itemid;

  return (
    <main className="h-full flex-grow bg-gray-50">
      <div className="mx-auto max-w-6xl px-6 py-6 md:px-8 lg:ml-0">
        <div className="flex items-center justify-between ">
          <h1 className="text-2xl font-bold">{params.itemid}</h1>
        </div>
        <hr className="my-4" />
        <div className="mt-6"></div>
      </div>
    </main>
  );
}
