import { createSupabaseClient } from "@/api/supabase/client";
import type { Education } from "@/modules/education/education.types";

interface EducationRow {
  uuid: string;
  course: string;
  description: string;
  educational_institution: string;
  start_date: string;
  end_date: string | null;
}

export async function getEducation(): Promise<Education[]> {
  const supabase = createSupabaseClient();

  if (!supabase) {
    console.warn("Supabase não configurado; formações não serão carregadas.");
    return [];
  }

  try {
    const { data, error } = await supabase
      .from("education")
      .select(
        "uuid, course, description, educational_institution, start_date, end_date",
      )
      .order("start_date", { ascending: false })
      .overrideTypes<EducationRow[], { merge: false }>();

    if (error) {
      console.error("Não foi possível carregar as formações do Supabase.");
      return [];
    }

    return (data ?? []).map((row) => ({
      uuid: row.uuid,
      course: row.course,
      description: row.description,
      educationalInstitution: row.educational_institution,
      startDate: row.start_date,
      endDate: row.end_date,
    }));
  } catch {
    console.error(
      "Não foi possível conectar ao Supabase para carregar formações.",
    );
    return [];
  }
}
