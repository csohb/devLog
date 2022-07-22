export type CustomResponseFormat<T = any> = {
  data: T;
  code: number;
  error?: { message?: string };
};
