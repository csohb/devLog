export type BlogSaveRequestBody = {
  title: string;
  content: string;
  writer: string;
  tags: string[];
};

export type BlogUpdateRequestBody = {
  id: string;
  title?: string;
  content?: string;
  tags?: string[];
};
