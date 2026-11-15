import { fetchApi } from "@/api/client";
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
  const data = await fetchApi<WorkExperienceRow[]>("/work-experiences");

  return groupByCompany(data ?? []);
}
