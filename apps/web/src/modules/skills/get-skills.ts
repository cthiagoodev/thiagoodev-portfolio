import { createSupabaseClient } from "@/api/supabase/client";
import type { Skill } from "@/modules/skills/skills.types";

interface SkillRow {
  uuid: string;
  label: string;
  url: string | null;
  image_path: string | null;
}

const SKILL_PRIORITY = [
  "Flutter",
  "Dart",
  "Go",
  "Kotlin",
  "PostgreSQL",
  "Docker",
  "Kubernetes",
  "Linux",
  "AWS",
  "GCP",
  "Firebase",
  "Rust",
] as const;

const priorityByLabel = new Map(
  SKILL_PRIORITY.map((label, index) => [label.toLocaleLowerCase(), index]),
);

function sortSkills(skills: Skill[]) {
  return skills.sort((first, second) => {
    const firstPriority = priorityByLabel.get(first.label.toLocaleLowerCase());
    const secondPriority = priorityByLabel.get(second.label.toLocaleLowerCase());

    if (firstPriority !== undefined || secondPriority !== undefined) {
      if (firstPriority === undefined) return 1;
      if (secondPriority === undefined) return -1;
      return firstPriority - secondPriority;
    }

    return first.label.localeCompare(second.label, "pt-BR", {
      sensitivity: "base",
    });
  });
}

export async function getSkills(): Promise<Skill[]> {
  const supabase = createSupabaseClient();

  if (!supabase) {
    console.warn("Supabase não configurado; habilidades não serão carregadas.");
    return [];
  }

  try {
    const { data, error } = await supabase
      .from("skills")
      .select("uuid, label, url, image_path")
      .overrideTypes<SkillRow[], { merge: false }>();

    if (error) {
      console.error("Não foi possível carregar as habilidades do Supabase.");
      return [];
    }

    const skills = (data ?? []).map((row) => ({
      uuid: row.uuid,
      label: row.label,
      url: row.url?.trim() || null,
    }));

    return sortSkills(skills);
  } catch {
    console.error("Não foi possível conectar ao Supabase para carregar habilidades.");
    return [];
  }
}
