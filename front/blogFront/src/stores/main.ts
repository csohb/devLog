import { writable } from "svelte/store";
import { fetchProfileList } from "../api/main";

import type { MainStore, Profiles } from "./types/main";

const mainStore = () => {
  const state: MainStore = {
    profiles: [],
  };

  const { subscribe, set, update } = writable(state);

  const methods = {
    async getProfiles() {
      let temp: Profiles[] = [];
      await fetchProfileList().then((resp: MainStore) => {
        if (resp !== null) {
          temp = resp.profiles;
        }
      });

      update((state) => {
        state.profiles = temp;
        return state;
      });
    },
    imgError(idx: number) {
      update((state) => {
        state.profiles[idx].image = "	https://picsum.photos/seed/picsum/536/536";
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

export default mainStore();
