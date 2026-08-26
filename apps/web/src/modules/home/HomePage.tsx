import React from "react";
import BlurFade from "@/common/magic-ui/BlurFade";
import { DATA } from "@/data/resume";
import AboutHero from "@/modules/about/AboutHero";
import AboutSection from "@/modules/about/AboutSection";
import type { About } from "@/modules/about/about.types";
import ContactSection from "@/modules/contact/ContactSection";
import EducationSection from "@/modules/education/EducationSection";
import HackathonsSection from "@/modules/hackathons/HackathonsSection";
import PhotosSection from "@/modules/photos/PhotosSection";
import ProjectsSection from "@/modules/projects/ProjectsSection";
import SkillsSection from "@/modules/skills/SkillsSection";
import WorkExperienceSection from "@/modules/work-experience/WorkExperienceSection";

const BLUR_FADE_DELAY = 0.04;

function getSectionComponents(
  about: About | null,
): Record<string, React.ReactNode> {
  const summary = about?.text?.trim() || DATA.summary;

  return {
    about: (
      <AboutSection
        heading={DATA.sections.about.heading}
        summary={summary}
        delay={BLUR_FADE_DELAY}
      />
    ),
    work: (
      <section id="work">
        <div className="flex min-h-0 flex-col gap-y-6">
          <BlurFade delay={BLUR_FADE_DELAY * 5}>
            <h2 className="text-xl font-bold">
              {DATA.sections.work.heading}
            </h2>
          </BlurFade>
          <BlurFade delay={BLUR_FADE_DELAY * 6}>
            <WorkExperienceSection
              items={DATA.work}
              presentLabel={DATA.sections.work.presentLabel}
            />
          </BlurFade>
        </div>
      </section>
    ),
    education: (
      <EducationSection
        heading={DATA.sections.education.heading}
        education={DATA.education}
        delay={BLUR_FADE_DELAY}
      />
    ),
    skills: (
      <SkillsSection
        heading={DATA.sections.skills.heading}
        skills={DATA.skills}
        delay={BLUR_FADE_DELAY}
      />
    ),
    projects: (
      <section id="projects">
        <BlurFade delay={BLUR_FADE_DELAY * 11}>
          <ProjectsSection
            label={DATA.sections.projects.label}
            heading={DATA.sections.projects.heading}
            text={DATA.sections.projects.text}
            projects={DATA.projects}
          />
        </BlurFade>
      </section>
    ),
    hackathons: (
      <section id="hackathons">
        <BlurFade delay={BLUR_FADE_DELAY * 13}>
          <HackathonsSection
            label={DATA.sections.hackathons.label}
            heading={DATA.sections.hackathons.heading}
            text={DATA.sections.hackathons.text}
            hackathons={DATA.hackathons}
          />
        </BlurFade>
      </section>
    ),
    photos: (
      <PhotosSection
        heading={DATA.sections.photos.heading}
        photos={DATA.photos}
      />
    ),
    contact: (
      <section id="contact">
        <BlurFade delay={BLUR_FADE_DELAY * 16}>
          <ContactSection
            label={DATA.sections.contact.label}
            heading={DATA.sections.contact.heading}
            text={DATA.sections.contact.text}
          />
        </BlurFade>
      </section>
    ),
  };
}

interface HomePageProps {
  about: About | null;
}

export default function HomePage({ about }: HomePageProps) {
  const name = about?.name?.trim() || DATA.name;
  const description = about?.description?.trim() || DATA.description;
  const sectionComponents = getSectionComponents(about);
  const orderedSections = Object.entries(DATA.sections)
    .filter(([, section]) => section.enabled)
    .sort(([, first], [, second]) => first.order - second.order)
    .map(([key]) => key);

  return (
    <main className="min-h-dvh flex flex-col gap-14 relative">
      <AboutHero
        name={name}
        description={description}
        avatarUrl={DATA.avatarUrl}
        initials={DATA.initials}
        delay={BLUR_FADE_DELAY}
      />
      {orderedSections.map((key) => (
        <React.Fragment key={key}>{sectionComponents[key]}</React.Fragment>
      ))}
    </main>
  );
}
