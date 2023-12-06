import { redirect } from "next/navigation";

export default function AppLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { store_id: string };
}) {
  const { store_id } = params;
  if (!store_id) redirect("/store");

  return children;
}
