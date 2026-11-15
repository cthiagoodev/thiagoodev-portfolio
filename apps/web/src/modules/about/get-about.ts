import { fetchApi } from "@/api/client";
import type { About } from "@/modules/about/about.types";

interface AboutRow {
  uuid: string;
  name: string;
  phone: string | null;
  description: string | null;
  text: string | null;
}

export async function getAbout(): Promise<About | null> {
  const data = await fetchApi<AboutRow>("/about");

  if (!data || !data.name.trim()) {
    return null;
  }

  return {
    uuid: data.uuid,
    name: data.name,
    phone: data.phone,
    description: data.description,
    text: data.text,
  };
}
