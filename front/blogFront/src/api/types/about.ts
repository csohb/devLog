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
