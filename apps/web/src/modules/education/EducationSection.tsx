import BlurFade from "@/common/magic-ui/BlurFade";
import type { Education } from "./education.types";

interface EducationSectionProps {
  heading: string;
  education: readonly Education[];
  delay: number;
}

const dateFormatter = new Intl.DateTimeFormat("pt-BR", {
  month: "short",
  year: "numeric",
  timeZone: "UTC",
});

function formatDate(date: string) {
  return dateFormatter.format(new Date(`${date}T00:00:00Z`));
}

function getInitials(name: string) {
  return name
    .split(/\s+/)
    .filter(Boolean)
    .slice(0, 2)
    .map((part) => part[0])
    .join("")
    .toUpperCase();
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
        {education.length === 0 && (
          <BlurFade delay={delay * 8}>
            <p className="text-sm text-muted-foreground">
              Não há informações de formação disponíveis no momento.
            </p>
          </BlurFade>
        )}
        <div className="flex flex-col gap-8">
          {education.map((item, index) => (
            <BlurFade key={item.uuid} delay={delay * 8 + index * 0.05}>
              <article className="flex items-center gap-x-3 justify-between">
                <div className="flex items-center gap-x-3 flex-1 min-w-0">
                  <div
                    aria-hidden="true"
                    className="size-8 md:size-10 border rounded-full shadow ring-2 ring-border bg-muted flex items-center justify-center text-xs font-semibold flex-none"
                  >
                    {getInitials(item.educationalInstitution)}
                  </div>
                  <div className="flex-1 min-w-0 flex flex-col gap-0.5">
                    <h3 className="font-semibold leading-none">
                      {item.educationalInstitution}
                    </h3>
                    <p className="font-sans text-sm text-muted-foreground">
                      {item.course}
                    </p>
                    {item.description && (
                      <p className="font-sans text-sm text-muted-foreground">
                        {item.description}
                      </p>
                    )}
                  </div>
                </div>
                <div className="text-xs tabular-nums text-muted-foreground text-right flex-none">
                  <span>
                    {formatDate(item.startDate)} -{" "}
                    {item.endDate ? formatDate(item.endDate) : "Atual"}
                  </span>
                </div>
              </article>
            </BlurFade>
          ))}
        </div>
      </div>
    </section>
  );
}
