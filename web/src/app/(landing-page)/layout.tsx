import Header from "@/components/landing-page/Header";
import Footer from "@/components/landing-page/footer";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="relative flex h-full min-h-screen flex-col">
      <Header />
      <div className="flex flex-1 flex-col">{children}</div>
      <Footer />
    </div>
  );
}
