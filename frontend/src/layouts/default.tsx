import { Link } from "@nextui-org/link";

import { Navbar } from "@/components/navbar";
import { CircularProgress } from "@nextui-org/react";

export default function DefaultLayout({
  children,
  loading,
}: {
  children: React.ReactNode;
  loading?: boolean;
}) {
  return (
    <div className="relative flex flex-col h-screen">
      <Navbar />
      {/* <main className="container mx-auto max-w-7xl px-6 flex-grow pt-14 overflow-auto"> */}
      <main className="container mx-auto max-w-7xl px-6 flex-grow pt-5 overflow-auto relative">
        {loading && (
          <div className="absolute inset-0 flex justify-center items-center bg-white bg-opacity-80 z-50">
            <CircularProgress aria-label="Loading..." />
          </div>
        )}
        {children}
      </main>
      <footer className="w-full flex items-center justify-center py-3">
        <Link
          isExternal
          className="flex items-center gap-1 text-current"
          href="https://github.com/manhrev"
          title="nextui.org homepage"
        >
          <span className="text-default-600">Made by</span>
          <p className="text-primary">manhrev</p>
        </Link>
      </footer>
    </div>
  );
}
