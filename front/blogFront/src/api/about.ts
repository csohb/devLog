import reqresApi from "./common";

export function fetchIntroduce(id: string): Promise<any> {
  return reqresApi.get(`/introduce/${id}`);
}
