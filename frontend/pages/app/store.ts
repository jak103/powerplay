import { defineStore } from "pinia"

export const useCounterStore = defineStore('counter', () => {
    const count = ref(0); 
    function increment() { count.value++ }
    function decrement() { count.value-- }

    return { count, increment, decrement}
})