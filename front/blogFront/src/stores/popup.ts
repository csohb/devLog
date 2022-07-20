import { writable, Writable } from "svelte/store";
import type { PopupType, PopupItem } from "./types/popup";

const popupStore = () => {
  const state: PopupType = {
    popupTask: [],
  };

  const { update, set, subscribe } = writable(state);

  const methods = {
    open(val: any) {
      const item: PopupItem = {
        title: val.title,
        text: val.text,
        type: val.type,
        btn: val.btn,
        isShow: val.isShow,
        action: val.action,
      };
      if (state.popupTask.length === 0) {
        item.isShow = true;
      }
      update((state) => {
        state.popupTask.push(item);
        return state;
      });
    },
    close() {
      if (state.popupTask[0].action != null) {
        state.popupTask[0].action();
      }
      this.cancel();
    },
    cancel() {
      update((state) => {
        state.popupTask.shift();
        this.visible();
        return state;
      });
    },
    visible() {
      if (state.popupTask.length === 0) {
        return;
      }
      console.log("visible");
      update((state) => {
        state.popupTask[0].isShow = true;
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

export default popupStore();
