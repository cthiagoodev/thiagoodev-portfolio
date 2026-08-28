export type SocialIcon = "github" | "linkedin" | "x" | "youtube" | "email";

export interface SocialContact {
  uuid: string;
  name: string;
  url: string;
  icon: SocialIcon;
  navbar: boolean;
}
