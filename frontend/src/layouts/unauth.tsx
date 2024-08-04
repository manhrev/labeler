// import { ThemeSwitch } from "@/components/theme-switch";
import { Link } from "@nextui-org/link";
import {
  Navbar as NextUINavbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
} from "@nextui-org/navbar";

export default function UnauthenticatedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="relative flex flex-col h-screen">
      <NextUINavbar maxWidth="xl" position="sticky">
        <NavbarContent className="basis-1/5 sm:basis-full" justify="start">
          <NavbarBrand className="gap-3 max-w-fit">
            <Link
              className="flex justify-start items-center gap-1"
              color="foreground"
              href="/"
            >
              <p className="font-bold text-inherit">LABELER</p>
            </Link>
          </NavbarBrand>
        </NavbarContent>
        <NavbarContent
          className="hidden sm:flex basis-1/5 sm:basis-full"
          justify="end"
        >
          <NavbarItem className="hidden sm:flex gap-2">
            {/* <ThemeSwitch /> */}
          </NavbarItem>
        </NavbarContent>
      </NextUINavbar>
      <main className="container mx-auto max-w-7xl px-6 flex-grow pt-32">
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
