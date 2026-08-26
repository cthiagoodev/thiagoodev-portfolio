import { ThemeProvider } from "@/common/theme/ThemeProvider";
import { TooltipProvider } from "@/common/ui/tooltip";
import Navbar, { type NavbarProps } from "@/modules/navigation/Navbar";

export default function NavbarIsland(props: NavbarProps) {
  return (
    <ThemeProvider attribute="class" defaultTheme="light" enableSystem={false}>
      <TooltipProvider delayDuration={0}>
        <Navbar {...props} />
      </TooltipProvider>
    </ThemeProvider>
  );
}
