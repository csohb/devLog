import { writable } from "svelte/store";
import {
  fetchStoryDetail,
  fetchStoryList,
  fetchStoryViewCount,
} from "../api/story";
import type { StoryRespDetail, StoryRespListItem } from "../api/types/story";

const storyStore = () => {
  //   const state: StoryStore = {
  //     storyList: [],
  //   };
  const state = {
    storyList: [],
    listTotal: 0,
    storyDetail: {
      id: "",
      type: "",
      view: 0,
      content: "",
      created_at: "",
      description: "",
      title: "",
      writer: "",
      images: [],
    },
  };

  const { subscribe, set, update } = writable(state);

  const methods = {
    async setStoryList(page: number, count: number) {
      let temp: StoryRespListItem[] = [];
      let total: number = 0;
      await fetchStoryList(page, count)
        .then((resp) => {
          if (resp !== null) {
            temp = resp.list;
            total = resp.total;
          }
        })
        .catch((err) => {
          console.log(err);
        });

      update((state) => {
        state.listTotal = total;
        state.storyList = temp;

        // state.storyList.map((val) => {
        //   Object.assign(val, {
        //     image:
        //       "https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg",
        //   });
        // });

        return state;
      });
    },
    async setStoryDetail(id: string) {
      let temp: any = {};
      await fetchStoryDetail(id)
        .then((resp: StoryRespDetail) => {
          if (resp !== null) {
            temp = resp;
          }
        })
        .catch((err) => {
          console.log(err);
        });

      update((state) => {
        state.storyDetail = temp;
        if (state.storyDetail.type === "F") {
          state.storyDetail.type = "FrontEnd";
        }
        if (state.storyDetail.type === "B") {
          state.storyDetail.type = "BackEnd";
        }
        if (!state.storyDetail.type) {
          state.storyDetail.type = "Etc";
        }
        return state;
      });
    },
    async viewCount(id: string) {
      await fetchStoryViewCount(id)
        .then(() => {
          update((state) => {
            state.storyDetail.view += 1;
            return state;
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
    resetStoryList() {
      update((state) => {
        state.storyList = [];
        state.listTotal = 0;
        return state;
      });
    },
    resetStoryDetail() {
      update((state) => {
        state.storyDetail = {
          id: "",
          type: "",
          view: 0,
          content: "",
          created_at: "",
          description: "",
          title: "",
          writer: "",
          images: [],
        };
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

export default storyStore();
