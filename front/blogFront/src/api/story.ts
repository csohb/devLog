import reqresApi from "./common";
import type {
  StoryResp,
  StoryRespDetail,
  StoryUpdateResp,
  StorySaveResp,
  StorySaveRequest,
  StoryUpdateRequest,
} from "./types/story";

export function fetchStoryList(
  page: number,
  count: number
): Promise<StoryResp | any> {
  let url = "/story/list?";
  if (page >= 1) {
    url += `page=${page}`;
  }
  if (count >= 5) {
    url += `&count=${count}`;
  }
  return reqresApi.get(url);
}

export function fetchStoryDetail(id: string): Promise<StoryRespDetail | any> {
  let url = "/story/get";
  if (!!id) {
    url += `/${id}`;
  }
  return reqresApi.get(url);
}

export function fetchStoryUpdate(
  req: StoryUpdateRequest
): Promise<StoryUpdateResp | any> {
  return reqresApi.put("/update", req);
}

export function fetchStoryDelete(id: string): Promise<any> {
  return reqresApi.delete(`/story/delete/${id}`);
}

export function fetchStorySave(
  req: StorySaveRequest
): Promise<StorySaveResp | any> {
  return reqresApi.post("/story/save", req);
}

export function fetchStoryViewCount(id: string): Promise<null> {
  return reqresApi.put(`/story/view/${id}`);
}
