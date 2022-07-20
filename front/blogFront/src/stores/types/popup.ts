export type PopupItem = {
  title?: string;
  text: string;
  btn?: string;
  type: "alert" | "confirm";
  isShow?: boolean;
  action?: Function;
};

export type PopupType = {
  popupTask: PopupItem[];
};
