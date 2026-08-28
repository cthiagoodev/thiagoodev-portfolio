import { cn } from "@/common/utils/cn";
import type { CSSProperties } from "react";

interface BlurFadeTextProps {
  text: string;
  as?: "div" | "h1" | "p";
  className?: string;
  duration?: number;
  characterDelay?: number;
  delay?: number;
  yOffset?: number;
  animateByCharacter?: boolean;
}

interface BlurFadeTextStyle extends CSSProperties {
  "--blur-fade-y": string;
  "--blur-fade-blur": string;
}

const BlurFadeText = ({
  text,
  as: Root = "div",
  className,
  duration = 0.4,
  characterDelay = 0.03,
  delay = 0,
  yOffset = 8,
  animateByCharacter = false,
}: BlurFadeTextProps) => {
  const characters = Array.from(text);

  const getStyle = (animationDelay: number): BlurFadeTextStyle => ({
    animationDelay: `${animationDelay}s`,
    animationDuration: `${duration}s`,
    "--blur-fade-y": `${yOffset}px`,
    "--blur-fade-blur": "8px",
  });

  if (animateByCharacter) {
    return (
      <Root className="flex">
        {characters.map((char, i) => {
          return (
            <span
              key={i}
              className={cn("blur-fade inline-block", className)}
              style={{
                ...getStyle(delay + i * characterDelay),
                width: char.trim() === "" ? "0.2em" : "auto",
              }}
            >
              {char}
            </span>
          );
        })}
      </Root>
    );
  }

  return (
    <Root className="flex">
      <span
        className={cn("blur-fade inline-block", className)}
        style={getStyle(delay)}
      >
        {text}
      </span>
    </Root>
  );
};

export default BlurFadeText;
