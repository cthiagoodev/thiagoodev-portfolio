import { createSupabaseClient } from "@/api/supabase/client";
import type { About } from "@/modules/about/about.types";

interface AboutRow {
  uuid: string;
  name: string;
  phone: string | null;
  description: string | null;
  text: string | null;
}

export async function getAbout(): Promise<About | null> {
  const supabase = createSupabaseClient();

  if (!supabase) {
    console.warn(
      "Supabase não configurado; usando os dados locais como fallback.",
    );
    return null;
  }

  try {
    const { data, error } = await supabase
      .from("about")
      .select("uuid, name, phone, description, text")
      .limit(1)
      .maybeSingle<AboutRow>();

    if (error) {
      console.error("Não foi possível carregar os dados de about do Supabase.");
      return null;
    }

    if (!data) {
      return null;
    }

    return {
      uuid: data.uuid,
      name: data.name,
      phone: data.phone,
      description: data.description,
      text: data.text,
    };
  } catch {
    console.error("Não foi possível conectar ao Supabase para carregar about.");
    return null;
  }
}
