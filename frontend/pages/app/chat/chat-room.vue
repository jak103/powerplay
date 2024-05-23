<script lang="ts" setup>
  import { useCounterStore } from '../store'


  const store = useCounterStore();
  const test = () => { store.increment() }
  const dec = () => { store.decrement() }
  const chatName = "Chat Name"

</script>

<template>
  <AppHeader :name="chatName"/>
  <div>
    <h1>Chat Room</h1>
    <p>This is the actual chat room</p>

    <p>This is a test of the store</p>
    <div>
      <h1>{{ store.count }}</h1>
      <button @click="test">increment</button>
      <button @click="dec">decrement</button>
    </div>
  </div>
</template>