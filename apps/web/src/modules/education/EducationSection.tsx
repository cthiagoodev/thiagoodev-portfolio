import BlurFade from "@/common/magic-ui/BlurFade";
import { ArrowUpRight } from "lucide-react";

interface EducationItem {
  school: string;
  href: string;
  degree: string;
  logoUrl?: string;
  start: string;
  end: string;
}

interface EducationSectionProps {
  heading: string;
  education: readonly EducationItem[];
  delay: number;
}

export default function EducationSection({
  heading,
  education,
  delay,
}: EducationSectionProps) {
  return (
    <section id="education">
      <div className="flex min-h-0 flex-col gap-y-6">
        <BlurFade delay={delay * 7}>
          <h2 className="text-xl font-bold">{heading}</h2>
        </BlurFade>
        <div className="flex flex-col gap-8">
          {education.map((item, index) => (
            <BlurFade
              key={item.school}
              delay={delay * 8 + index * 0.05}
            >
              <a
                href={item.href}
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center gap-x-3 justify-between group"
              >
                <div className="flex items-center gap-x-3 flex-1 min-w-0">
                  {item.logoUrl ? (
                    <img
                      src={item.logoUrl}
                      alt={item.school}
                      className="size-8 md:size-10 p-1 border rounded-full shadow ring-2 ring-border overflow-hidden object-contain flex-none"
                    />
                  ) : (
                    <div className="size-8 md:size-10 p-1 border rounded-full shadow ring-2 ring-border bg-muted flex-none" />
                  )}
                  <div className="flex-1 min-w-0 flex flex-col gap-0.5">
                    <div className="font-semibold leading-none flex items-center gap-2">
                      {item.school}
                      <ArrowUpRight
                        className="h-3.5 w-3.5 text-muted-foreground opacity-0 -translate-x-2 group-hover:opacity-100 group-hover:translate-x-0 transition-all duration-200"
                        aria-hidden
                      />
                    </div>
                    <div className="font-sans text-sm text-muted-foreground">
                      {item.degree}
                    </div>
                  </div>
                </div>
                <div className="flex items-center gap-1 text-xs tabular-nums text-muted-foreground text-right flex-none">
                  <span>{item.start} - {item.end}</span>
                </div>
              </a>
            </BlurFade>
          ))}
        </div>
      </div>
    </section>
  );
}
