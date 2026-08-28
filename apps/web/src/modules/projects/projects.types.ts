export interface Project {
  uuid: string;
  name: string;
  description: string | null;
  url: string | null;
  startDate: string;
  endDate: string | null;
}
