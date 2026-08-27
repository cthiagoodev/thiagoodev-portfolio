import BlurFade from "@/common/magic-ui/BlurFade";
import BlurFadeText from "@/common/magic-ui/BlurFadeText";
import { Avatar, AvatarFallback, AvatarImage } from "@/common/ui/avatar";

interface AboutHeroProps {
  name: string;
  description: string | null;
  delay: number;
}

export default function AboutHero({
  name,
  description,
  delay,
}: AboutHeroProps) {
  const initials = name
    .split(/\s+/)
    .filter(Boolean)
    .slice(0, 2)
    .map((part) => part[0])
    .join("")
    .toUpperCase();

  return (
    <section id="hero">
      <div className="mx-auto w-full max-w-2xl space-y-8">
        <div className="gap-2 gap-y-6 flex flex-col md:flex-row justify-between">
          <div className="gap-2 flex flex-col order-2 md:order-1">
            <BlurFadeText
              delay={delay}
              className="text-3xl font-semibold tracking-tighter sm:text-4xl lg:text-5xl"
              yOffset={8}
              text={`Olá, eu sou ${name.split(" ")[0]}`}
            />
            {description && (
              <BlurFadeText
                className="text-muted-foreground max-w-[600px] md:text-lg lg:text-xl"
                delay={delay}
                text={description}
              />
            )}
          </div>
          <BlurFade delay={delay} className="order-1 md:order-2">
            <Avatar className="size-24 md:size-32 border rounded-full shadow-lg ring-4 ring-muted">
              <AvatarImage alt={name} />
              <AvatarFallback>{initials}</AvatarFallback>
            </Avatar>
          </BlurFade>
        </div>
      </div>
    </section>
  );
}
