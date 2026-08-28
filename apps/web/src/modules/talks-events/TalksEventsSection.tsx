/* eslint-disable @next/next/no-img-element */
import { Badge } from "@/common/ui/badge";
import {
  Timeline,
  TimelineConnectItem,
  TimelineItem,
} from "@/modules/talks-events/Timeline";
import type { TalkEvent } from "@/modules/talks-events/talks-events.types";
import { ArrowUpRight } from "lucide-react";

interface TalksEventsSectionProps {
  label: string;
  heading: string;
  text: string;
  items: readonly TalkEvent[];
}

const dateFormatter = new Intl.DateTimeFormat("pt-BR", {
  day: "2-digit",
  month: "short",
  year: "numeric",
  timeZone: "UTC",
});

function formatDate(date: string) {
  return dateFormatter.format(new Date(`${date}T00:00:00Z`));
}

function formatPeriod(startDate: string, endDate: string | null) {
  return endDate
    ? `${formatDate(startDate)} – ${formatDate(endDate)}`
    : formatDate(startDate);
}

export default function TalksEventsSection({
  label,
  heading,
  text,
  items,
}: TalksEventsSectionProps) {
  return (
    <section id="talks-events" className="overflow-hidden">
      <div className="flex min-h-0 flex-col gap-y-8 w-full">
        <div className="flex flex-col gap-y-4 items-center justify-center">
          <div className="flex items-center w-full">
            <div className="flex-1 h-px bg-linear-to-r from-transparent from-5% via-border via-95% to-transparent" />
            <div className="border bg-primary z-10 rounded-xl px-4 py-1">
              <span className="text-background text-sm font-medium">
                {label}
              </span>
            </div>
            <div className="flex-1 h-px bg-linear-to-l from-transparent from-5% via-border via-95% to-transparent" />
          </div>
          <div className="flex flex-col gap-y-3 items-center justify-center">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl">
              {heading}
            </h2>
            <p className="text-muted-foreground md:text-lg/relaxed lg:text-base/relaxed xl:text-lg/relaxed text-balance text-center">
              {text.replace("{count}", String(items.length))}
            </p>
          </div>
        </div>
        {items.length === 0 && (
          <p className="text-sm text-muted-foreground text-center">
            Não há palestras ou eventos disponíveis no momento.
          </p>
        )}
        <Timeline>
          {items.map((item) => (
            <TimelineItem
              key={item.uuid}
              className="w-full flex items-start justify-between gap-10"
            >
              <TimelineConnectItem className="flex items-start justify-center">
                {item.imageUrl ? (
                  <img
                    src={item.imageUrl}
                    alt={item.title}
                    className="size-10 bg-card z-10 shrink-0 overflow-hidden p-1 border rounded-full shadow ring-2 ring-border object-contain flex-none"
                  />
                ) : (
                  <div className="size-10 bg-card z-10 shrink-0 overflow-hidden p-1 border rounded-full shadow ring-2 ring-border flex-none" />
                )}
              </TimelineConnectItem>
              <div className="flex flex-1 flex-col justify-start gap-2 min-w-0">
                <time className="text-xs text-muted-foreground">
                  {formatPeriod(item.startDate, item.endDate)}
                </time>
                {item.title && (
                  <h3 className="font-semibold leading-none">{item.title}</h3>
                )}
                {item.location && (
                  <p className="text-sm text-muted-foreground">
                    {item.location}
                  </p>
                )}
                {item.description && (
                  <p className="text-sm text-muted-foreground leading-relaxed wrap-break-word">
                    {item.description}
                  </p>
                )}
                {item.links.length > 0 && (
                  <div className="mt-1 flex flex-row flex-wrap items-start gap-2">
                    {item.links.map((link) => {
                      const isExternal = link.url.startsWith("http");

                      return (
                        <a
                          href={link.url}
                          key={`${link.label}-${link.url}`}
                          target={isExternal ? "_blank" : undefined}
                          rel={isExternal ? "noopener noreferrer" : undefined}
                        >
                          <Badge className="flex items-center gap-1.5 text-xs bg-primary text-primary-foreground">
                            <ArrowUpRight className="size-3" aria-hidden />
                            {link.label}
                          </Badge>
                        </a>
                      );
                    })}
                  </div>
                )}
              </div>
            </TimelineItem>
          ))}
        </Timeline>
      </div>
    </section>
  );
}
