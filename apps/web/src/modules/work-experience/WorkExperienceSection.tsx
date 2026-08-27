import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/common/ui/accordion";
import { cn } from "@/common/utils/cn";
import type { WorkExperienceGroup } from "@/modules/work-experience/work-experience.types";
import { ChevronDown, ChevronRight } from "lucide-react";

interface WorkExperienceSectionProps {
  items: readonly WorkExperienceGroup[];
  presentLabel: string;
}

const dateFormatter = new Intl.DateTimeFormat("pt-BR", {
  month: "short",
  year: "numeric",
  timeZone: "UTC",
});

function formatDate(date: string) {
  const formattedDate = dateFormatter.format(new Date(`${date}T00:00:00Z`));
  return formattedDate.charAt(0).toUpperCase() + formattedDate.slice(1);
}

function formatPeriod(
  startDate: string,
  endDate: string | null,
  presentLabel: string,
) {
  return `${formatDate(startDate)} – ${endDate ? formatDate(endDate) : presentLabel}`;
}

function CompanyInitials({ company }: { company: string }) {
  const initials = company
    .split(/\s+/)
    .filter(Boolean)
    .slice(0, 2)
    .map((part) => part[0])
    .join("")
    .toUpperCase();

  return (
    <div
      aria-hidden
      className="size-8 md:size-10 border rounded-full shadow ring-2 ring-border bg-muted flex items-center justify-center flex-none text-xs font-semibold text-muted-foreground"
    >
      {initials}
    </div>
  );
}

export default function WorkExperienceSection({
  items,
  presentLabel,
}: WorkExperienceSectionProps) {
  if (items.length === 0) {
    return (
      <p className="text-sm text-muted-foreground">
        Não há experiências profissionais disponíveis no momento.
      </p>
    );
  }

  return (
    <Accordion type="single" collapsible className="w-full grid gap-6">
      {items.map((experience) => {
        const latestRole = experience.roles[0];

        return (
          <AccordionItem
            key={experience.company}
            value={experience.company}
            className="w-full border-b-0 grid gap-2"
          >
            <AccordionTrigger className="hover:no-underline p-0 cursor-pointer transition-colors rounded-none group [&>svg]:hidden">
              <div className="flex items-center gap-x-3 justify-between w-full text-left">
                <div className="flex items-center gap-x-3 flex-1 min-w-0">
                  <CompanyInitials company={experience.company} />
                  <div className="flex-1 min-w-0 gap-0.5 flex flex-col">
                    <div className="font-semibold leading-none flex items-center gap-2">
                      {experience.company}
                      <span className="relative inline-flex items-center w-3.5 h-3.5">
                        <ChevronRight
                          className={cn(
                            "absolute h-3.5 w-3.5 shrink-0 text-muted-foreground stroke-2 transition-all duration-300 ease-out",
                            "translate-x-0 opacity-0",
                            "group-hover:translate-x-1 group-hover:opacity-100",
                            "group-data-[state=open]:opacity-0 group-data-[state=open]:translate-x-0",
                          )}
                        />
                        <ChevronDown
                          className={cn(
                            "absolute h-3.5 w-3.5 shrink-0 text-muted-foreground stroke-2 transition-all duration-200",
                            "opacity-0 rotate-0",
                            "group-data-[state=open]:opacity-100 group-data-[state=open]:rotate-180",
                          )}
                        />
                      </span>
                    </div>
                    <div className="font-sans text-sm text-muted-foreground">
                      {latestRole.role}
                    </div>
                  </div>
                </div>
                <div className="flex items-center gap-1 text-xs tabular-nums text-muted-foreground text-right flex-none">
                  <span>
                    {formatPeriod(
                      latestRole.startDate,
                      latestRole.endDate,
                      presentLabel,
                    )}
                  </span>
                </div>
              </div>
            </AccordionTrigger>
            <AccordionContent className="p-0 ml-11 md:ml-13 text-xs sm:text-sm text-muted-foreground">
              <div className="grid gap-5">
                {experience.roles.map((role) => (
                  <article
                    key={role.uuid}
                    className={cn(
                      "relative grid gap-2",
                      experience.roles.length > 1 &&
                        "pl-4 before:absolute before:inset-y-1 before:left-0 before:w-px before:bg-border",
                    )}
                  >
                    {experience.roles.length > 1 && (
                      <span
                        aria-hidden
                        className="absolute left-0 top-1.5 size-1.5 -translate-x-[2.5px] rounded-full bg-muted-foreground"
                      />
                    )}
                    <div className="grid gap-0.5">
                      <h3 className="font-medium text-foreground">{role.role}</h3>
                      <time className="text-xs tabular-nums text-muted-foreground">
                        {formatPeriod(
                          role.startDate,
                          role.endDate,
                          presentLabel,
                        )}
                      </time>
                    </div>
                    <div className="grid gap-3 leading-relaxed">
                      {role.description
                        .split(/\n{2,}/)
                        .filter(Boolean)
                        .map((paragraph, paragraphIndex) => (
                          <p
                            key={`${role.uuid}-paragraph-${paragraphIndex}`}
                            className="whitespace-pre-line"
                          >
                            {paragraph}
                          </p>
                        ))}
                    </div>
                  </article>
                ))}
              </div>
            </AccordionContent>
          </AccordionItem>
        );
      })}
    </Accordion>
  );
}
