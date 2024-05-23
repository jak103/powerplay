<script lang="ts" setup>
  import { useCounterStore } from '../store'


  const store = useCounterStore();
  const test = () => { store.increment() }
  const dec = () => { store.decrement() }
  const addChannel = "Add Channel"

</script>


<template>
  <AppHeader :name="addChannel"/>
  <div>
    <h1>Create Channel</h1>
    <p>This is where channels can be created</p>

    <p>This is a test of the store</p>
    <div>
      <h1>{{ store.count }}</h1>
      <button @click="test">increment</button>
      <button @click="dec">decrement</button>
    </div>
  </div>
</template>