<script lang="ts" setup>
<<<<<<< HEAD
  import { useCounterStore } from '../store'


  const store = useCounterStore();
  const test = () => { store.increment() }
  const dec = () => { store.decrement() }
=======
  import FloatingButtonGeneric from '~/components/FloatingButtonGeneric.vue';

>>>>>>> d1e57528dc4a1af6e507bc111af6da07ac0d6411
  const addChannel = "Add Channel"
  const chatRoom = "/app/chat/chat-room"

</script>


<template>
  <AppHeader :name="addChannel"/>
<<<<<<< HEAD
  <div>
    <h1>Create Channel</h1>
    <p>This is where channels can be created</p>

    <p>This is a test of the store</p>
    <div>
      <h1>{{ store.count }}</h1>
      <button @click="test">increment</button>
      <button @click="dec">decrement</button>
    </div>
=======
  <div id="container">
    <h1>Create a channel</h1>
    <p class="maintext">Name</p>
    <input id="channelName" type="text">
    <p class="subtext">Channels are where conversations happen around a topic. Use a name that is easy to find and understand.</p>
    <p class="maintext">Visibility</p>
    <div class="channelTypeSelectorContainer">
      <input type="radio"/>
      <p class="subtext">Public - anyone in Power Play</p>
    </div>
    <div class="channelTypeSelectorContainer">
      <input type="radio"/>
      <p class="subtext">Private - Only specific people</p>
    </div>
    <FloatingButtonGeneric text="Create Channel" :link="chatRoom"/>
>>>>>>> d1e57528dc4a1af6e507bc111af6da07ac0d6411
  </div>
</template>

<style scoped>

input[type="text"] {
  width: 100%;
  box-sizing: border-box;
  padding: 8px;
  font-size: 16px;
}
.maintext {
  margin-top: 16px;
  margin-bottom: 16px;
}
.subtext {
  font-size: x-small;
  margin: 0;
}
.channelTypeSelectorContainer{
  display: flex;
  align-items: center;
}

.channelTypeSelectorContainer input[type="radio"] {
  margin-right: 8px;
  vertical-align: middle;
}
</style>