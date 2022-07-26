import type { MainStore } from "../stores/types/main";
import reqresApi from "./common";

export function fetchProfileList(): Promise<MainStore> {
  return reqresApi.get("/main/introduce");
}
