<script lang="ts" setup>
  import FloatingButtonGeneric from '~/components/FloatingButtonGeneric.vue';

  const addChannel = "Add Channel"
  const chatRoom = "/app/chat/chat-room"
<<<<<<< main
=======
  const channelName = ref("");
  const createChannel = () => {

  }
>>>>>>> main

</script>


<template>
  <AppHeader :name="addChannel"/>
  <div id="container">
    <h1>Create a channel</h1>
    <p class="maintext">Name</p>
<<<<<<< main
    <input id="channelName" type="text">
=======
    <input id="channelName" type="text" v-model="channelName">
>>>>>>> main
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
<<<<<<< main
    <FloatingButtonGeneric text="Create Channel" :link="chatRoom"/>
=======
    <FloatingButtonGeneric onclick="" text="Create Channel" :link="chatRoom"/>
>>>>>>> main
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
<<<<<<< main
</style>
=======
</style>
>>>>>>> main
