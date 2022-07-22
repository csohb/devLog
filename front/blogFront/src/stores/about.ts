import { writable } from "svelte/store";
import { fetchIntroduce } from "../api/about";
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
          console.log("about resp:", resp);
          update((state) => {
            state.info = resp.profile;
            state.careers = resp.careers;
            state.skills = resp.skills;
            state.project = resp.project;
            state.keywords = resp.keywords;
            return state;
          });
        })
        .catch((err) => {
          console.log("about err:", err);
          this.resetAbout();
        });
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
