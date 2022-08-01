import { writable } from "svelte/store";
import {
  fetchCrateCareer,
  fetchDeleteCareer,
  fetchIntroduce,
  fetchUpdateCareer,
} from "../api/about";
import type { CareerListType } from "../api/types/about";
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
            state.info = resp.profile;
            state.careers = resp.careers;
            state.skills = resp.skills;
            state.project = resp.project;
            state.keywords = resp.keywords;
            Object.assign(state.careers, {
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
