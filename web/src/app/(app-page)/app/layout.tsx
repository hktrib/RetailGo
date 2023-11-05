import MobileNav from "@/components/app-page/mobile-nav";
import Sidebar from "@/components/app-page/sidebar";

export default function AppLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen h-full flex flex-col">
      <MobileNav />
      <Sidebar />
      <div className="xl:pl-64 h-full flex-grow flex flex-col">{children}</div>
    </div>
  );
}
