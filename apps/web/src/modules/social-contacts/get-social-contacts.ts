import { createSupabaseClient } from "@/api/supabase/client";
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

function mapSocialContact(row: SocialContactRow): SocialContact | null {
  const label = row.label.trim();
  const icon = SOCIAL_ICONS[label.toLocaleLowerCase()];
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
  const supabase = createSupabaseClient();

  if (!supabase) {
    console.warn(
      "Supabase não configurado; contatos sociais não serão carregados.",
    );
    return [];
  }

  try {
    const { data, error } = await supabase
      .from("social_contacts")
      .select("uuid, label, url")
      .overrideTypes<SocialContactRow[], { merge: false }>();

    if (error) {
      console.error(
        "Não foi possível carregar os contatos sociais do Supabase.",
      );
      return [];
    }

    return (data ?? [])
      .map(mapSocialContact)
      .filter((contact): contact is SocialContact => contact !== null)
      .sort(
        (first, second) =>
          (orderByIcon.get(first.icon) ?? Number.MAX_SAFE_INTEGER) -
          (orderByIcon.get(second.icon) ?? Number.MAX_SAFE_INTEGER),
      );
  } catch {
    console.error(
      "Não foi possível conectar ao Supabase para carregar contatos sociais.",
    );
    return [];
  }
}
