import reqresApi from "./common";
import type { ContactRequestBody } from "./types/contact";

export function fetchEmailSend(req: ContactRequestBody): Promise<any> {
  return reqresApi.post("/email", req);
}
