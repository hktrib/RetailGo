import Header from "@/components/landing-page/Header";
import Footer from "@/components/landing-page/footer";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div>
      <Header />
      {children}
      <Footer />
    </div>
  );
}
