import BlurFade from "@/common/magic-ui/BlurFade";
import Markdown from "react-markdown";

interface AboutSectionProps {
  heading: string;
  summary: string;
  delay: number;
}

export default function AboutSection({
  heading,
  summary,
  delay,
}: AboutSectionProps) {
  return (
    <section id="about">
      <div className="flex min-h-0 flex-col gap-y-4">
        <BlurFade delay={delay * 3}>
          <h2 className="text-xl font-bold">{heading}</h2>
        </BlurFade>
        <BlurFade delay={delay * 4}>
          <div className="prose max-w-full text-pretty font-sans leading-relaxed text-muted-foreground dark:prose-invert">
            <Markdown>{summary}</Markdown>
          </div>
        </BlurFade>
      </div>
    </section>
  );
}
