import Sidebar from "@/components/app-page/sidebar";

export default function AppLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen h-full flex flex-col">
      <Sidebar />
      <div className="xl:pl-72 h-full flex-grow flex flex-col">{children}</div>
    </div>
  );
}
