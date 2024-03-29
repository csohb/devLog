import { writable } from "svelte/store";
import {
  fetchCrateCareer,
  fetchCrateProject,
  fetchCrateSkill,
  fetchDeleteCareer,
  fetchDeleteProject,
  fetchDeleteSkill,
  fetchIntroduce,
  fetchUpdate,
  fetchUpdateCareer,
  fetchUpdateProject,
} from "../api/about";
import type {
  CareerListType,
  CrateProjectsReq,
  CrateSkillRequest,
} from "../api/types/about";
import type { AboutType } from "./types/about";

const aboutStore = () => {
  const state: AboutType = {
    description: "",
    info: {
      addr: "",
      birthday: "",
      developer: "",
      email: "",
      img: "",
      name: "",
      nick_name: "",
    },
    careers: [],
    keywords: [],
    skills: [],
    project: [],
  };

  const { subscribe, set, update } = writable(state);

  const methods = {
    async setIntroduce(name: string) {
      await fetchIntroduce(name)
        .then((resp) => {
          update((state) => {
            state.description = resp.profile.description;
            state.info = resp.profile;
            state.careers = resp.careers;
            state.skills = resp.skills;
            state.project = resp.project;
            state.keywords = resp.keywords;
            Object.assign(state.careers, {
              isEditMode: false,
            });
            Object.assign(state.project, {
              isEditMode: false,
            });
            Object.assign(state.keywords, {
              isEditMode: false,
            });
            return state;
          });
        })
        .catch((err) => {
          console.log("about err:", err);
          this.resetAbout();
        });
    },
    async updateIntroduce(
      name: string,
      intro: string,
      email: string,
      addr: string,
      img: string
    ): Promise<any> {
      let requestData = {
        user_id: name,
        profile: {
          intro,
          email,
          addr,
          img,
        },
      };
      return await fetchUpdate(requestData);
    },
    careerEditMode(idx: number) {
      update((state) => {
        if (!state.careers[idx].isEditMode) {
          state.careers.map((val) => {
            val.isEditMode = false;
          });
        }
        state.careers[idx].isEditMode = !state.careers[idx].isEditMode;
        return state;
      });
    },
    async crateCareer(id: string, list: any[]): Promise<any> {
      let req = {
        id,
        career: list,
      };
      return await fetchCrateCareer(req);
    },
    async updateCareer(id: string, req: CareerListType): Promise<any> {
      return await fetchUpdateCareer(id, req);
    },
    async deleteCareer(id: string): Promise<any> {
      return await fetchDeleteCareer(id);
    },
    projectEditMode(idx: number) {
      update((state) => {
        if (!state.project[idx].isEditMode) {
          state.project.map((val) => {
            val.isEditMode = false;
          });
        }
        state.project[idx].isEditMode = !state.project[idx].isEditMode;
        return state;
      });
    },
    async crateProject(id: string, list: CrateProjectsReq[]) {
      // 프로젝트 추가
      let request = {
        id,
        projects: list,
      };
      return await fetchCrateProject(request);
    },
    async updateProject(id: string, req: CrateProjectsReq) {
      return await fetchUpdateProject(id, req);
    },
    async deleteProject(id: string) {
      return await fetchDeleteProject(id);
    },
    async crateSkill(req: CrateSkillRequest): Promise<any> {
      return await fetchCrateSkill(req);
    },
    async deleteSkill(id: string): Promise<any> {
      return await fetchDeleteSkill(id);
    },
    skillEditMode(idx: number) {
      update((state) => {
        if (!state.keywords[idx].isEditMode) {
          state.keywords.map((val) => {
            val.isEditMode = false;
          });
        }
        state.keywords[idx].isEditMode = !state.keywords[idx].isEditMode;
        return state;
      });
    },
    getSkillPercentage(id: string) {
      let percentage = 0;
      state.skills.map((val) => {
        if (val.id === id) {
          percentage = val.percentage;
        }
      });
      return percentage;
    },
    resetAbout() {
      update((state) => {
        state.description = "";
        state.info = {
          addr: "",
          birthday: "",
          developer: "",
          email: "",
          img: "",
          name: "",
          nick_name: "",
        };
        state.careers = [];
        state.keywords = [];
        state.project = [];
        state.skills = [];
        return state;
      });
    },
  };

  return {
    subscribe,
    set,
    update,
    ...methods,
  };
};

export default aboutStore();
