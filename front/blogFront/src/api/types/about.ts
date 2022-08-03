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
