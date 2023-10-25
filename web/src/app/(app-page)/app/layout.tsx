import Sidebar from "@/components/app-page/sidebar";

export default function AppLayout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <Sidebar />
      <div className="xl:pl-72">{children}</div>
    </div>
  );
}
