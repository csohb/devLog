import reqresApi from "./common";
import type { ContactRequestBody } from "./types/contact";

export function fetchEmailSend(req: ContactRequestBody): Promise<any> {
  return reqresApi.post("/email", req);
}

export function imgUpload(data: any) {
  console.log("data:", data);
  return reqresApi.post("/upload", data);
}
