import { fetchApi } from "@/api/client";
import type { Project } from "@/modules/projects/projects.types";

interface ProjectSkillRow {
  label: string;
}

interface ProjectRow {
  uuid: string;
  name: string;
  description: string | null;
  url: string | null;
  repo_url: string | null;
  skills?: ProjectSkillRow[] | null;
  pushed_at: string | null;
  created_at: string;
}

export async function getProjects(): Promise<Project[]> {
  const data = await fetchApi<ProjectRow[]>("/projects");

  return (data ?? [])
    .map((row) => ({
      uuid: row.uuid,
      name: row.name,
      description: row.description,
      url: row.url?.trim() || row.repo_url?.trim() || null,
      startDate: row.created_at.slice(0, 10),
      endDate: row.pushed_at ? row.pushed_at.slice(0, 10) : null,
      skills: (row.skills ?? []).map((skill) => skill.label),
    }))
    .sort((first, second) => second.startDate.localeCompare(first.startDate));
}
