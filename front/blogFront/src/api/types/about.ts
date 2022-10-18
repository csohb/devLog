export type CareerListType = {
  company: string;
  start_date: string;
  end_date: string;
  job_title: string;
  job_detail: string;
};

export type CrateCareerRequest = {
  id: string;
  career: CareerListType[];
};

export type SkillListType = {
  name: string;
  percentage?: number;
};

export type CrateSkillRequest = {
  user_id: string;
  tech: SkillListType[];
};
export type UpdateSkillRequest = {
  id: string;
  name: string;
  percentage?: number;
};

export type ProfileType = {
  intro: string;
  email: string;
  addr: string;
};

export type UpdateProfile = {
  user_id: string;
  profile: ProfileType;
};

export type CrateProjectsReq = {
  name: string;
  is_personal: boolean;
  start_date: string;
  end_date: string;
  description: string;
  stack: string[];
};

export type CrateProjectRequest = {
  id: string;
  projects: CrateProjectsReq[];
};

type IntroduceProfileRes = {
  name: string;
  nick_name: string;
  developer: string;
  description: string;
  img: string;
  addr: string;
  email: string;
  birthday: string;
};
type IntroduceCareersRes = {
  id: string;
  company: string;
  start_date: string;
  end_date: string;
  job_title: string;
  job_detail: string;
};
type IntroduceProjectRes = {
  id: string;
  name: string;
  is_personal: boolean;
  start_date: string;
  end_date: string;
  description: string;
  stack: string[];
};
type IntroduceKeywordsRes = { id: string; name: string };
type IntroduceSkillsRes = {
  id: string;
  name: string;
  percentage: number;
};
export type IntroduceResponse = {
  profile: IntroduceProfileRes;
  careers: IntroduceCareersRes[];
  project: IntroduceProjectRes[];
  keywords: IntroduceKeywordsRes[];
  skills: IntroduceSkillsRes[];
};
