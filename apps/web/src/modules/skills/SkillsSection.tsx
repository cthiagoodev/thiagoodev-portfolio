import BlurFade from "@/common/magic-ui/BlurFade";
import type { ElementType } from "react";

interface SkillItem {
  name: string;
  icon?: ElementType<{ className?: string }>;
}

interface SkillsSectionProps {
  heading: string;
  skills: readonly SkillItem[];
  delay: number;
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
        <div className="flex flex-wrap gap-2">
          {skills.map((skill, id) => (
            <BlurFade key={skill.name} delay={delay * 10 + id * 0.05}>
              <div className="border bg-background border-border ring-2 ring-border/20 rounded-xl h-8 w-fit px-4 flex items-center gap-2">
                {skill.icon && (
                  <skill.icon className="size-4 rounded overflow-hidden object-contain" />
                )}
                <span className="text-foreground text-sm font-medium">
                  {skill.name}
                </span>
              </div>
            </BlurFade>
          ))}
        </div>
      </div>
    </section>
  );
}
