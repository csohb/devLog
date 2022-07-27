import type { BlogList, BlogResp } from "../stores/types/blog";
import reqresApi from "./common";
import type { BlogSaveRequestBody, BlogUpdateRequestBody } from "./types/blog";

// 리턴값 오류일 경우 확인 필요.....

export function fetchBlogSave(req: BlogSaveRequestBody): Promise<any> {
  return reqresApi.post("/blog/save", req);
}

export function fetchBlogList(
  page: number,
  count: number
): Promise<BlogResp | any> {
  let url = "/blog/list?";
  if (page >= 1) {
    url += `page=${page}`;
  }
  if (count >= 5) {
    url += `&count=${count}`;
  }
  return reqresApi.get(url);
}

export function fetchBlogDetail(id: string): Promise<BlogList | any> {
  let url = "/blog/get";
  if (!!id) {
    url += `/${id}`;
  }
  return reqresApi.get(url);
}

export function fetchBlogUpdate(req: BlogUpdateRequestBody): Promise<any> {
  return reqresApi.put("/blog/update", req);
}

export function fetchBlogDelete(id: string): Promise<any> {
  return reqresApi.delete(`/blog/delete/${id}`);
}

export function fetchBlogTag(
  tag: string,
  page: number,
  count: number
): Promise<any> {
  let url = "/blog/tags?";
  if (page >= 1) {
    url += `page=${page}`;
  }
  if (count >= 5) {
    url += `&count=${count}`;
  }
  if (tag !== "") {
    url += `&tag=${tag}`;
  }
  console.log("tag url:", url);
  return reqresApi.get(url);
}

export function fetchViewCount(id: string): Promise<any> {
  return reqresApi.put("/blog/view", {
    id,
    count: 1,
  });
}
