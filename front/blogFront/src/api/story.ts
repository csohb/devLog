import reqresApi from "./common";
import type { StoryResp } from "./types/story";

export function fetchStoryList(
  page: number,
  count: number
): Promise<StoryResp> {
  return reqresApi.get("/story/list?page=1&count=10");
}
