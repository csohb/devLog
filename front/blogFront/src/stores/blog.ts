import { writable } from "svelte/store";
import {
  fetchBlogList,
  fetchBlogDetail,
  fetchBlogTag,
  fetchViewCount,
  fetchHeartCount,
} from "../api/blog";
import type { BlogStore, BlogList, BlogResp } from "./types/blog";

const blogStore = () => {
  const state: BlogStore = {
    blogList: [],
    listTotal: 0,
    blogDetail: {
      id: "",
      heart: 0,
      view: 0,
      tags: [],
      title: "",
      content: "",
      date: "",
      description: "",
      writer: "",
    },
    tagList: [],
  };

  const { subscribe, set, update } = writable(state);

  const methods = {
    async setBlogList(page: number, count: number) {
      let temp: BlogList[] = [];
      let total: number = 0;
      await fetchBlogList(page, count)
        .then((resp: BlogResp) => {
          if (resp !== null) {
            temp = resp.list;
            total = resp.total;

            // temp.forEach((val) => {
            //   Object.assign(val, {
            //     description: "",
            //     date: "2022년 01월 22일",
            //   });
            // });
          }
        })
        .catch((err) => {
          console.log("블로그 리스트 불러오기 err : ", err);
        });

      update((state) => {
        state.blogList = temp;
        state.listTotal = total;
        state.tagList = temp.reduce((acc, cur) => {
          cur.tags.map((val) => {
            const idx = acc.indexOf(val);
            if (idx === -1) {
              acc.push(val);
            }
          });
          return acc;
        }, []);

        return state;
      });
    },
    resetBlogList() {
      update((state) => {
        state.blogList = [];
        return state;
      });
    },
    async setTagList(tag: string, page: number, count: number) {
      let temp: BlogList[] = [];
      let total: number = 0;
      await fetchBlogTag(tag, page, count).then((resp) => {
        if (resp.list != null) {
          temp = resp.list;
          total = resp.total;
        }
      });
      update((state) => {
        state.blogList = temp;
        state.listTotal = total;
        return state;
      });
    },
    async setBlogDetail(id: string) {
      let temp: any = {};
      await fetchBlogDetail(id)
        .then((resp: BlogList) => {
          if (resp !== null) {
            temp = resp;
          }
        })
        .catch((err) => {
          console.log("블로그 세부페이지 불러오기 err : ", err);
        });

      update((state) => {
        // Object.assign(temp, {
        //   description: "",
        // });
        state.blogDetail = temp;
        return state;
      });
    },
    resetBlogDetail() {
      update((state) => {
        state.blogDetail = {
          id: "",
          heart: 0,
          view: 0,
          tags: [],
          title: "",
          content: "",
          date: "",
          description: "",
          writer: "",
        };
        return state;
      });
    },
    async viewCount(id: string) {
      await fetchViewCount(id)
        .then(() => {
          update((state) => {
            state.blogDetail.view += 1;
            return state;
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
    async heartCount(id: string, isAdd: boolean) {
      if (!isAdd && state.blogDetail.heart <= 0) {
        return;
      }
      await fetchHeartCount(id, isAdd)
        .then(() => {
          update((state) => {
            isAdd
              ? (state.blogDetail.heart += 1)
              : (state.blogDetail.heart -= 1);
            return state;
          });
        })
        .catch((err) => {
          console.log(err);
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

export default blogStore();
