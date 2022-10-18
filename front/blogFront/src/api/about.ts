import reqresApi from "./common";
import type {
  CrateCareerRequest,
  CareerListType,
  CrateSkillRequest,
  UpdateProfile,
  CrateProjectRequest,
  CrateProjectsReq,
  UpdateSkillRequest,
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
export function fetchUpdateSkill(request: UpdateSkillRequest): Promise<any> {
  return reqresApi.put("introduce/tech/update", request);
}
export function fetchDeleteSkill(id: string): Promise<any> {
  return reqresApi.delete(`introduce/tech/delete/${id}`);
}

export function fetchUpdate(request: UpdateProfile): Promise<any> {
  return reqresApi.put("/introduce/update", request);
}

export function fetchCrateProject(request: CrateProjectRequest): Promise<any> {
  return reqresApi.post("/introduce/project/create", request);
}

export function fetchDeleteProject(id: string): Promise<any> {
  return reqresApi.delete(`/introduce/project/delete/${Number(id)}`);
}

export function fetchUpdateProject(
  id: string,
  req: CrateProjectsReq
): Promise<any> {
  let request = req;

  Object.assign(request, {
    id: String(id),
  });

  return reqresApi.put("/introduce/project/update", request);
}
