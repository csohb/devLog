type StoryRespListItem = {
  content: string;
  created_at: string;
  description: string;
  id: string;
  title: string;
  view: number;
  writer: string;
  imgUrl?: string;
};
export type StoryResp = {
  list: StoryRespListItem[];
  total: number;
};

export type StoryRespDetail = {
  content: string;
  created_at: string;
  description: string;
  id: string;
  type: string;
  title: string;
  view: number;
  writer: string;
};

export type StoryUpdateResp = {};
export type StorySaveResp = {};

// request

export type StorySaveRequest = {
  title: string;
  content: string;
  description: string;
};

export type StoryUpdateRequest = {
  id: string;
  title: string;
  content: string;
  description: string;
};
