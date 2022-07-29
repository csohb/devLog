type InfoType = {
  addr: string;
  birthday: string;
  developer: string;
  email: string;
  img: string;
  name: string;
  nick_name: string;
};

type CareerType = {
  company: string;
  end_date: string;
  job_detail: string;
  job_title: string;
  start_date: string;
  isEditMode: boolean;
};
type KeywordType = {
  id: string;
  name: string;
};
type SkillType = {
  id: string;
  name: string;
  percentage: number;
};

export type AboutType = {
  description: string;
  info: InfoType;
  careers: CareerType[];
  keywords: KeywordType[];
  skills: SkillType[];
  project: any[];
};
