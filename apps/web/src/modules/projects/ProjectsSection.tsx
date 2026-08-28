import BlurFade from "@/common/magic-ui/BlurFade";
import { ProjectCard } from "@/modules/projects/ProjectCard";
import type { Project } from "@/modules/projects/projects.types";

const BLUR_FADE_DELAY = 0.04;

interface ProjectsSectionProps {
  label: string;
  heading: string;
  text: string;
  projects: readonly Project[];
}

const dateFormatter = new Intl.DateTimeFormat("pt-BR", {
  month: "short",
  year: "numeric",
  timeZone: "UTC",
});

function formatDate(date: string) {
  return dateFormatter.format(new Date(`${date}T00:00:00Z`));
}

function formatPeriod(startDate: string, endDate: string | null) {
  return `${formatDate(startDate)} – ${endDate ? formatDate(endDate) : "Atual"}`;
}

export default function ProjectsSection({
  label,
  heading,
  text,
  projects,
}: ProjectsSectionProps) {
  return (
    <section id="projects">
      <div className="flex min-h-0 flex-col gap-y-8">
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
              {text}
            </p>
          </div>
        </div>
        {projects.length === 0 && (
          <p className="text-sm text-muted-foreground text-center">
            Não há projetos disponíveis no momento.
          </p>
        )}
        <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 max-w-[800px] mx-auto auto-rows-fr">
          {projects.map((project, id) => (
            <BlurFade
              key={project.uuid}
              delay={BLUR_FADE_DELAY * 12 + id * 0.05}
              className="h-full"
            >
              <ProjectCard
                href={project.url ?? undefined}
                title={project.name}
                description={project.description ?? ""}
                dates={formatPeriod(project.startDate, project.endDate)}
                tags={[]}
              />
            </BlurFade>
          ))}
        </div>
      </div>
    </section>
  );
}
