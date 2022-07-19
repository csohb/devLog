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
};

export type BlogStore = {
  blogList: BlogList[];
  listTotal: number;
  blogDetail: BlogList;
};

export type BlogResp = {
  list: BlogList[];
  total: number;
};
