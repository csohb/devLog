import axios, {
  AxiosInstance,
  AxiosInterceptorManager,
  AxiosRequestConfig,
  AxiosResponse,
} from "axios";

import type { CustomResponseFormat } from "./types/common";
import popupStore from "../stores/popup";

interface CustomInstance extends AxiosInstance {
  interceptors: {
    request: AxiosInterceptorManager<AxiosRequestConfig>;
    response: AxiosInterceptorManager<AxiosResponse<CustomResponseFormat>>;
  };
  getUri(config?: AxiosRequestConfig): string;
  request<T>(config: AxiosRequestConfig): Promise<T>;
  get<T>(url: string, config?: AxiosRequestConfig): Promise<T>;
  delete<T>(url: string, config?: AxiosRequestConfig): Promise<T>;
  head<T>(url: string, config?: AxiosRequestConfig): Promise<T>;
  options<T>(url: string, config?: AxiosRequestConfig): Promise<T>;
  post<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>;
  put<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>;
  patch<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T>;
}

//const url: any = __devlog.env.baseUrl;
const url = process.env.BASE_URL;
// console.log("env:", url);

// axios 객체 하나만들어두고 재사용하기
const reqresApi: CustomInstance = axios.create({
  baseURL: `${url}/api/v1`, // Url 나중에 env 로 빼기
  timeout: 15000, // timeout 15초
});

// Axios 에는 interceptors 라는 기능이 있다. 이를 통해서 request / response 에 선행,후행 처리를 커스텀
reqresApi.interceptors.response.use((res) => {
  if (res.data.code === 200) {
    return res.data.data;
  } else {
    if (res.data.error != null && res.data.error.message != null) {
      popupStore.open({
        title: "ERROR",
        text: res.data.error.message,
        type: "alert",
        isShow: false,
      });
    }
    return Promise.reject(res.data);
  }
});

reqresApi.interceptors.request.use((req) => {
  // req.headers.authorization = TokenManager.accessToken; // jwt token
  if (req.url === "/upload") {
    req.headers.Accept = "multipart/form-data";
  }
  return req;
});

export default reqresApi;
