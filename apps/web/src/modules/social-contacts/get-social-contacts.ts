import { fetchApi } from "@/api/client";
import type {
  SocialContact,
  SocialIcon,
} from "@/modules/social-contacts/social-contacts.types";

interface SocialContactRow {
  uuid: string;
  label: string;
  url: string;
}

const SOCIAL_ICONS: Readonly<Record<string, SocialIcon>> = {
  github: "github",
  linkedin: "linkedin",
  x: "x",
  youtube: "youtube",
  email: "email",
};

const SOCIAL_ORDER: readonly SocialIcon[] = [
  "github",
  "linkedin",
  "x",
  "youtube",
  "email",
];

const orderByIcon = new Map(SOCIAL_ORDER.map((icon, index) => [icon, index]));

function resolveSocialIcon(label: string): SocialIcon | undefined {
  const normalized = label.toLocaleLowerCase();

  if (SOCIAL_ICONS[normalized]) {
    return SOCIAL_ICONS[normalized];
  }

  const matchedKey = Object.keys(SOCIAL_ICONS).find((key) =>
    normalized.includes(key),
  );

  return matchedKey ? SOCIAL_ICONS[matchedKey] : undefined;
}

function mapSocialContact(row: SocialContactRow): SocialContact | null {
  const label = row.label.trim();
  const icon = resolveSocialIcon(label);
  const url = row.url.trim();

  if (!icon) {
    console.warn(`Contato social ignorado por label desconhecida: ${label}`);
    return null;
  }

  if (!url) {
    console.warn(`Contato social ignorado por URL vazia: ${label}`);
    return null;
  }

  return {
    uuid: row.uuid,
    name: label,
    url,
    icon,
    navbar: true,
  };
}

export async function getSocialContacts(): Promise<SocialContact[]> {
  const data = await fetchApi<SocialContactRow[]>("/social-contacts");

  return (data ?? [])
    .map(mapSocialContact)
    .filter((contact): contact is SocialContact => contact !== null)
    .sort(
      (first, second) =>
        (orderByIcon.get(first.icon) ?? Number.MAX_SAFE_INTEGER) -
        (orderByIcon.get(second.icon) ?? Number.MAX_SAFE_INTEGER),
    );
}
