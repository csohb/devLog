import { writable } from 'svelte/store';
import { fetchLogin } from '../api/auth';

const authStore = () => {
  const methods = {
    async postLogin(id: string, pw: string): Promise<string> {
      return await fetchLogin(id, pw);
    },
  };

  return {
    ...methods,
  };
};

export default authStore();
