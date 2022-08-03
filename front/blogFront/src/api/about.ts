import reqresApi from "./common";
import type {
  CrateCareerRequest,
  CareerListType,
  CrateSkillRequest,
} from "./types/about";

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
    id: Number(id),
  });
  return reqresApi.put("/career/update", request);
}

export function fetchDeleteCareer(id: string): Promise<any> {
  return reqresApi.delete(`/career/delete/${Number(id)}`);
}

export function fetchCrateSkill(request: CrateSkillRequest): Promise<any> {
  return reqresApi.post("/introduce/tech/create", request);
}
