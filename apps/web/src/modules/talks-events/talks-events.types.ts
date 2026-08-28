export interface TalkEventLink {
  label: string;
  url: string;
}

export interface TalkEvent {
  uuid: string;
  title: string;
  description: string | null;
  location: string | null;
  startDate: string;
  endDate: string | null;
  imageUrl: string | null;
  links: readonly TalkEventLink[];
}
