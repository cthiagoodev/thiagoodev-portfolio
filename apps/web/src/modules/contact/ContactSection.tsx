import { FlickeringGrid } from "@/common/magic-ui/FlickeringGrid";

interface ContactSectionProps {
  label: string;
  heading: string;
  text: string;
}

export default function ContactSection({
  label,
  heading,
  text,
}: ContactSectionProps) {
  return (
    <div className="border rounded-xl p-10 relative">
      <div className="absolute -top-4 border bg-primary z-10 rounded-xl px-4 py-1 left-1/2 -translate-x-1/2">
        <span className="text-background text-sm font-medium">{label}</span>
      </div>
      <div className="absolute inset-0 top-0 left-0 right-0 h-1/2 rounded-xl overflow-hidden">
        <FlickeringGrid
          className="h-full w-full"
          squareSize={2}
          gridGap={2}
          style={{
            maskImage: "linear-gradient(to bottom, black, transparent)",
            WebkitMaskImage: "linear-gradient(to bottom, black, transparent)",
          }}
        />
      </div>
      <div className="relative flex flex-col items-center gap-4 text-center">
        <h2 className="text-3xl font-bold tracking-tighter sm:text-5xl">
          {heading}
        </h2>
        <p className="mx-auto max-w-lg text-muted-foreground text-balance">
          {text}
        </p>
      </div>
    </div>
  );
}
