import { defineStore } from 'pinia';

interface User {
  id: number;
  name: string;
  // Add other properties as needed
}

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as User | null,
  }),
  actions: {
    setUser(user: User | null) {
      this.user = user;
    },
    clearUser() {
      this.user = null;
    },
  },
});