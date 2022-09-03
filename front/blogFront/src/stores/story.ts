import { writable } from "svelte/store";
import { fetchStoryList } from "../api/story";

const storyStore = () => {
  //   const state: StoryStore = {
  //     storyList: [],
  //   };
  const state = {};

  const { subscribe, set, update } = writable(state);

  const methods = {
    async setStoryList(page: number, count: number) {
      await fetchStoryList(page, count).then((resp) => {
        console.log("story resp :", resp);
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

export default storyStore();
