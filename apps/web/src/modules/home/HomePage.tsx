import BlurFade from "@/common/magic-ui/BlurFade";
import AboutHero from "@/modules/about/AboutHero";
import AboutSection from "@/modules/about/AboutSection";
import type { About } from "@/modules/about/about.types";
import ContactSection from "@/modules/contact/ContactSection";
import EducationSection from "@/modules/education/EducationSection";
import HackathonsSection from "@/modules/hackathons/HackathonsSection";
import PhotosSection from "@/modules/photos/PhotosSection";
import ProjectsSection from "@/modules/projects/ProjectsSection";
import SkillsSection from "@/modules/skills/SkillsSection";
import type { Skill } from "@/modules/skills/skills.types";
import WorkExperienceSection from "@/modules/work-experience/WorkExperienceSection";
import type { WorkExperienceGroup } from "@/modules/work-experience/work-experience.types";

const BLUR_FADE_DELAY = 0.04;

interface HomePageProps {
  about: About | null;
  workExperience: WorkExperienceGroup[];
  skills: Skill[];
}

export default function HomePage({
  about,
  workExperience,
  skills,
}: HomePageProps) {
  const name = about?.name.trim();
  const description = about?.description?.trim() || null;
  const summary = about?.text?.trim();

  return (
    <main className="min-h-dvh flex flex-col gap-14 relative">
      {name && (
        <AboutHero
          name={name}
          description={description}
          delay={BLUR_FADE_DELAY}
        />
      )}

      {summary && (
        <AboutSection
          heading="Sobre"
          summary={summary}
          delay={BLUR_FADE_DELAY}
        />
      )}

      <section id="work">
        <div className="flex min-h-0 flex-col gap-y-6">
          <BlurFade delay={BLUR_FADE_DELAY * 5}>
            <h2 className="text-xl font-bold">Experiência Profissional</h2>
          </BlurFade>
          <BlurFade delay={BLUR_FADE_DELAY * 6}>
            <WorkExperienceSection
              items={workExperience}
              presentLabel="Atual"
            />
          </BlurFade>
        </div>
      </section>

      <EducationSection
        heading="Formação"
        education={[]}
        delay={BLUR_FADE_DELAY}
      />

      <SkillsSection
        heading="Habilidades"
        skills={skills}
        delay={BLUR_FADE_DELAY}
      />

      <section id="projects">
        <BlurFade delay={BLUR_FADE_DELAY * 11}>
          <ProjectsSection
            label="Meus Projetos"
            heading="Projetos em destaque"
            text="Conheça alguns dos projetos que desenvolvi."
            projects={[]}
          />
        </BlurFade>
      </section>

      <PhotosSection heading="Viagens recentes" photos={[]} />

      <section id="hackathons">
        <BlurFade delay={BLUR_FADE_DELAY * 13}>
          <HackathonsSection
            label="Hackathons"
            heading="Projetos e experiências em hackathons"
            text="Confira minha participação em hackathons."
            hackathons={[]}
          />
        </BlurFade>
      </section>

      <section id="contact">
        <BlurFade delay={BLUR_FADE_DELAY * 16}>
          <ContactSection
            label="Contato"
            heading="Entre em contato"
            text="Não há informações de contato disponíveis no momento."
          />
        </BlurFade>
      </section>
    </main>
  );
}
