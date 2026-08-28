import { createSupabaseClient } from "@/api/supabase/client";
import type { Project } from "@/modules/projects/projects.types";

interface ProjectRow {
  uuid: string;
  name: string;
  description: string | null;
  url: string | null;
  start_date: string;
  end_date: string | null;
}

export async function getProjects(): Promise<Project[]> {
  const supabase = createSupabaseClient();

  if (!supabase) {
    console.warn("Supabase não configurado; projetos não serão carregados.");
    return [];
  }

  try {
    const { data, error } = await supabase
      .from("projects")
      .select("uuid, name, description, url, start_date, end_date")
      .order("start_date", { ascending: false })
      .overrideTypes<ProjectRow[], { merge: false }>();

    if (error) {
      console.error("Não foi possível carregar os projetos do Supabase.");
      return [];
    }

    return (data ?? []).map((row) => ({
      uuid: row.uuid,
      name: row.name,
      description: row.description,
      url: row.url?.trim() || null,
      startDate: row.start_date,
      endDate: row.end_date,
    }));
  } catch {
    console.error(
      "Não foi possível conectar ao Supabase para carregar projetos.",
    );
    return [];
  }
}
