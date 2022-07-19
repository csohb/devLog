import reqresApi from "./common";
import type { BlogSaveRequestBody, BlogUpdateRequestBody } from "./types/blog";

export function fetchBlogSave(req: BlogSaveRequestBody): Promise<any> {
  return reqresApi.post("/blog/save", req);
}

export function fetchBlogList(page: number, count: number): Promise<any> {
  let url = "/blog/list?";
  if (page >= 1) {
    url += `page=${page}`;
  }
  if (count >= 10) {
    url += `&count=${count}`;
  }
  return reqresApi.get(url);
}

export function fetchBlogDetail(id: string): Promise<any> {
  let url = "/blog/get";
  if (!!id) {
    url += `/${id}`;
  }
  return reqresApi.get(url);
}

export function fetchBlogUpdate(req: BlogUpdateRequestBody): Promise<any> {
  return reqresApi.put("/blog/update", req);
}
