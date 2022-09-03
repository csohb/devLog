import reqresApi from "./common";

export function fetchLogin(id: string, pw: string): Promise<string> {
  return reqresApi.post("/login", {
    id,
    password: pw,
  });
}

export function fetchLogout(): Promise<string> {
  return reqresApi.post("/logout");
}
