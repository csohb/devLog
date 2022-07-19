import { writable } from "svelte/store";
import { fetchBlogList, fetchBlogDetail } from "../api/blog";
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
  };

  const { subscribe, set, update } = writable(state);

  const methods = {
    async setBlogList(page: number, count: number) {
      let temp: BlogList[] = [];
      await fetchBlogList(page, count)
        .then((resp: BlogResp) => {
          if (resp !== null) {
            temp = resp.list;

            temp.forEach((val) => {
              Object.assign(val, {
                description: "",
                date: "2022년 01월 22일",
              });
            });
          }
        })
        .catch((err) => {
          console.log("블로그 리스트 불러오기 err : ", err);
        });

      update((state) => {
        state.blogList = temp;
        return state;
      });
    },
    async setBlogDetail(id: string) {
      let temp: any = {};
      await fetchBlogDetail(id)
        .then((resp: BlogList) => {
          console.log(resp);
          if (resp !== null) {
            temp = resp;
          }
        })
        .catch((err) => {
          console.log("블로그 세부페이지 불러오기 err : ", err);
        });

      update((state) => {
        Object.assign(temp, {
          description: "",
          date: "2022년 01월 22일",
        });
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
  };

  return {
    subscribe,
    set,
    update,
    ...methods,
  };
};

export default blogStore();
