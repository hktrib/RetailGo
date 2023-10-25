import Sidebar from "@/components/app-page/sidebar";
import Header from "@/components/landing-page/header";

export default function AppLayout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <Sidebar />
      <div className="xl:pl-72">{children}</div>
    </div>
  );
}
