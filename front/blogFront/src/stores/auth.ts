import { writable } from "svelte/store";
import { fetchLogin } from "../api/auth";

const authStore = () => {
  const state = {
    loginNick: "",
  };
  const { update, set, subscribe } = writable(state);
  const methods = {
    async postLogin(id: string, pw: string): Promise<string> {
      return await fetchLogin(id, pw);
    },
    setNick(id: string) {
      update((val) => {
        val.loginNick = id;
        return val;
      });
    },
    getCookie(key: string) {
      let cookieValue = null;
      if (document.cookie) {
        let array = document.cookie.split(escape(key) + "=");
        if (array.length >= 2) {
          let arraySub = array[1].split(";");
          cookieValue = unescape(arraySub[0]);
        }
      }
      return cookieValue;
    },
  };

  return {
    update,
    set,
    subscribe,
    ...methods,
  };
};

export default authStore();
