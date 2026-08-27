import { createSupabaseClient } from "@/api/supabase/client";
import type {
  WorkExperienceGroup,
  WorkExperienceRole,
} from "@/modules/work-experience/work-experience.types";

interface WorkExperienceRow {
  uuid: string;
  role: string;
  description: string;
  company: string;
  start_date: string;
  end_date: string | null;
}

function groupByCompany(rows: WorkExperienceRow[]): WorkExperienceGroup[] {
  const orderedRows = [...rows].sort((first, second) =>
    second.start_date.localeCompare(first.start_date),
  );
  const groups = new Map<string, WorkExperienceRole[]>();

  for (const row of orderedRows) {
    const roles = groups.get(row.company) ?? [];

    roles.push({
      uuid: row.uuid,
      role: row.role,
      description: row.description,
      startDate: row.start_date,
      endDate: row.end_date,
    });
    groups.set(row.company, roles);
  }

  return Array.from(groups, ([company, roles]) => ({ company, roles })).sort(
    (first, second) => {
      const firstIsCurrent = first.roles.some((role) => role.endDate === null);
      const secondIsCurrent = second.roles.some((role) => role.endDate === null);

      if (firstIsCurrent !== secondIsCurrent) {
        return firstIsCurrent ? -1 : 1;
      }

      return second.roles[0].startDate.localeCompare(first.roles[0].startDate);
    },
  );
}

export async function getWorkExperience(): Promise<WorkExperienceGroup[]> {
  const supabase = createSupabaseClient();

  if (!supabase) {
    console.warn(
      "Supabase não configurado; experiências profissionais não serão carregadas.",
    );
    return [];
  }

  try {
    const { data, error } = await supabase
      .from("work_experience")
      .select("uuid, role, description, company, start_date, end_date")
      .order("start_date", { ascending: false })
      .overrideTypes<WorkExperienceRow[], { merge: false }>();

    if (error) {
      console.error(
        "Não foi possível carregar as experiências profissionais do Supabase.",
      );
      return [];
    }

    return groupByCompany(data ?? []);
  } catch {
    console.error(
      "Não foi possível conectar ao Supabase para carregar as experiências profissionais.",
    );
    return [];
  }
}
