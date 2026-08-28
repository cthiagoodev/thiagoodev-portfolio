import { Icons, type IconProps } from "@/common/icons/icons";
import { DockIcon } from "@/common/magic-ui/Dock";
import {
  Tooltip,
  TooltipArrow,
  TooltipContent,
  TooltipTrigger,
} from "@/common/ui/tooltip";
import type {
  SocialContact,
  SocialIcon,
} from "@/modules/social-contacts/social-contacts.types";
import type { ComponentType } from "react";

const socialIcons = {
  github: Icons.github,
  linkedin: Icons.linkedin,
  x: Icons.x,
  youtube: Icons.youtube,
  email: Icons.email,
} satisfies Record<SocialIcon, ComponentType<IconProps>>;

export default function SocialLinks({
  contacts,
}: {
  contacts: readonly SocialContact[];
}) {
  return contacts
    .filter((social) => social.navbar)
    .map((social) => {
      const isExternal = social.url.startsWith("http");
      const IconComponent = socialIcons[social.icon];

      return (
        <Tooltip key={social.uuid}>
          <TooltipTrigger asChild>
            <a
              href={social.url}
              target={isExternal ? "_blank" : undefined}
              rel={isExternal ? "noopener noreferrer" : undefined}
            >
              <DockIcon className="rounded-3xl cursor-pointer size-full bg-background p-0 text-foreground hover:text-foreground hover:bg-muted backdrop-blur-3xl border border-border transition-colors">
                <IconComponent className="size-full rounded-sm overflow-hidden object-contain" />
              </DockIcon>
            </a>
          </TooltipTrigger>
          <TooltipContent
            side="top"
            sideOffset={8}
            className="rounded-xl bg-primary text-primary-foreground px-4 py-2 text-sm shadow-[0_10px_40px_-10px_rgba(0,0,0,0.3)] dark:shadow-[0_10px_40px_-10px_rgba(0,0,0,0.5)]"
          >
            <p>{social.name}</p>
            <TooltipArrow className="fill-primary" />
          </TooltipContent>
        </Tooltip>
      );
    });
}
