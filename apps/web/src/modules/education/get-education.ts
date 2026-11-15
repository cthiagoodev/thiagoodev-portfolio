import { fetchApi } from "@/api/client";
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
  const data = await fetchApi<EducationRow[]>("/educations");

  return (data ?? [])
    .sort((first, second) => second.start_date.localeCompare(first.start_date))
    .map((row) => ({
      uuid: row.uuid,
      course: row.course,
      description: row.description,
      educationalInstitution: row.educational_institution,
      startDate: row.start_date,
      endDate: row.end_date,
    }));
}
