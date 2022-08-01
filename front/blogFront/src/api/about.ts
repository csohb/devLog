import reqresApi from "./common";
import type { CrateCareerRequest, CareerListType } from "./types/about";

export function fetchIntroduce(id: string): Promise<any> {
  return reqresApi.get(`/introduce/${id}`);
}

export function fetchCrateCareer(req: CrateCareerRequest): Promise<any> {
  return reqresApi.post("/career/save", req);
}

export function fetchUpdateCareer(
  id: string,
  req: CareerListType
): Promise<any> {
  let request = req;

  Object.assign(request, {
    id: Number(id), // 현재 아이디를 get 할때 가져오지 못하는 것으로 알고 있음...
  });
  return reqresApi.put("/career/update", request);
}

export function fetchDeleteCareer(id: string): Promise<any> {
  return reqresApi.delete(`/career/delete/${Number(id)}`);
}
