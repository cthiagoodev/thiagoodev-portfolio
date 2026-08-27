import { Aws } from "@/common/icons/technologies/aws";
import { Dart } from "@/common/icons/technologies/dart";
import { Docker } from "@/common/icons/technologies/docker";
import { Firebase } from "@/common/icons/technologies/firebase";
import { Flutter } from "@/common/icons/technologies/flutter";
import { Golang } from "@/common/icons/technologies/golang";
import { GoogleCloud } from "@/common/icons/technologies/googleCloud";
import { Kubernetes } from "@/common/icons/technologies/kubernetes";
import { Kotlin } from "@/common/icons/technologies/kotlin";
import { Linux } from "@/common/icons/technologies/linux";
import { Postgresql } from "@/common/icons/technologies/postgresql";
import { Rust } from "@/common/icons/technologies/rust";
import BlurFade from "@/common/magic-ui/BlurFade";
import type { Skill } from "@/modules/skills/skills.types";
import { Code2 } from "lucide-react";
import type { ComponentType, SVGProps } from "react";

interface SkillsSectionProps {
  heading: string;
  skills: readonly Skill[];
  delay: number;
}

type SkillIcon = ComponentType<SVGProps<SVGSVGElement>>;

const skillIcons: Record<string, SkillIcon> = {
  flutter: Flutter,
  dart: Dart,
  go: Golang,
  kotlin: Kotlin,
  postgresql: Postgresql,
  docker: Docker,
  kubernetes: Kubernetes,
  linux: Linux,
  aws: Aws,
  gcp: GoogleCloud,
  firebase: Firebase,
  rust: Rust,
};

function resolveSkillIcon(label: string): SkillIcon {
  return skillIcons[label.trim().toLocaleLowerCase()] ?? Code2;
}

function SkillBadge({ skill }: { skill: Skill }) {
  const Icon = resolveSkillIcon(skill.label);
  const className =
    "border bg-background border-border ring-2 ring-border/20 rounded-xl h-8 w-fit px-4 flex items-center gap-2";
  const content = (
    <>
      <Icon
        aria-hidden
        className="size-4 rounded overflow-hidden object-contain"
      />
      <span className="text-foreground text-sm font-medium">{skill.label}</span>
    </>
  );

  if (!skill.url) {
    return <div className={className}>{content}</div>;
  }

  const isExternal = skill.url.startsWith("http");

  return (
    <a
      href={skill.url}
      target={isExternal ? "_blank" : undefined}
      rel={isExternal ? "noopener noreferrer" : undefined}
      className={className}
    >
      {content}
    </a>
  );
}

export default function SkillsSection({
  heading,
  skills,
  delay,
}: SkillsSectionProps) {
  return (
    <section id="skills">
      <div className="flex min-h-0 flex-col gap-y-4">
        <BlurFade delay={delay * 9}>
          <h2 className="text-xl font-bold">{heading}</h2>
        </BlurFade>
        {skills.length === 0 && (
          <BlurFade delay={delay * 10}>
            <p className="text-sm text-muted-foreground">
              Não há habilidades disponíveis no momento.
            </p>
          </BlurFade>
        )}
        <div className="flex flex-wrap justify-center gap-2">
          {skills.map((skill, id) => (
            <BlurFade key={skill.uuid} delay={delay * 10 + id * 0.05}>
              <SkillBadge skill={skill} />
            </BlurFade>
          ))}
        </div>
      </div>
    </section>
  );
}
