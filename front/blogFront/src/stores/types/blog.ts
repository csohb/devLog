export type BlogList = {
  id: string;
  heart: number;
  view: number;
  tags: string[];
  title: string;
  content: string; //description
  date?: string;
  writer: string;
  description?: string;
  images: string[];
};

export type BlogStore = {
  blogList: BlogList[];
  listTotal: number;
  blogDetail: BlogList;
  tagList: string[];
};

export type BlogResp = {
  list: BlogList[];
  total: number;
};
