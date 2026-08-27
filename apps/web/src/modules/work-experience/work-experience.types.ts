export interface WorkExperienceRole {
  uuid: string;
  role: string;
  description: string;
  startDate: string;
  endDate: string | null;
}

export interface WorkExperienceGroup {
  company: string;
  roles: WorkExperienceRole[];
}
