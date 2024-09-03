import type { Metadata } from "next";
import "./globals.scss";

export const metadata: Metadata = {
  title: "Click the planet",
  description: "Click the planet v3 - Simple clicker game",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
