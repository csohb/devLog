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
  };

  return {
    update,
    set,
    subscribe,
    ...methods,
  };
};

export default authStore();
