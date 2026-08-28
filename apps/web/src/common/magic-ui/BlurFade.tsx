import type { CSSProperties, ReactNode } from "react";

interface BlurFadeProps {
  children: ReactNode;
  className?: string;
  duration?: number;
  delay?: number;
  yOffset?: number;
  blur?: string;
}

interface BlurFadeStyle extends CSSProperties {
  "--blur-fade-y": string;
  "--blur-fade-blur": string;
}

const BlurFade = ({
  children,
  className,
  duration = 0.4,
  delay = 0,
  yOffset = 6,
  blur = "6px",
}: BlurFadeProps) => {
  const style: BlurFadeStyle = {
    animationDelay: `${0.04 + delay}s`,
    animationDuration: `${duration}s`,
    "--blur-fade-y": `${yOffset}px`,
    "--blur-fade-blur": blur,
  };

  return (
    <div className={`blur-fade ${className ?? ""}`} style={style}>
      {children}
    </div>
  );
};

export default BlurFade;
