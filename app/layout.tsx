import "./globals.css";
import type { Metadata } from "next";
import { Nunito } from "next/font/google";
import ClientOnly from "./components/ClientOnly";
import Siderbar from "./components/sidebar/Sidebar";

const font = Nunito({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Business Manager",
  description: "Custom business manager for shop",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={font.className}>
        <ClientOnly>
          <Siderbar children={children} />
        </ClientOnly>
      </body>
    </html>
  );
}
