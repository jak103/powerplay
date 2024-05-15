<script lang="ts" setup>
  import FloatingButton from '~/components/FloatingButton.vue';
  import AppHeader from '~/components/AppHeader.vue';
  import Card from '~/components/Card.vue';


  // TODO: pull data in this form from the backend
  const data = {
    channels : [
      {
        "channel_name": "Toronto Maple Leafs",
        "channel_id": "mapleleafs123",
        "channel_image": "mapleleafs_image.jpg",
        "missed_chats": 5,
        "number_of_members": 100
      },
      {
        "channel_name": "Montreal Canadiens",
        "channel_id": "canadiens456",
        "channel_image": "canadiens_image.jpg",
        "missed_chats": 2,
        "number_of_members": 50
      },
      {
        "channel_name": "Edmonton Oilers",
        "channel_id": "oilers789",
        "channel_image": "oilers_image.jpg",
        "missed_chats": 10,
        "number_of_members": 200
      },
      {
        "channel_name": "Vancouver Canucks",
        "channel_id": "canucks101",
        "channel_image": "canucks_image.jpg",
        "missed_chats": 1,
        "number_of_members": 30
      }
    ]
  }

  const createChannel = "/app/chat/create-channel"
  const chatRoom = "/app/chat/chat-room"
  const name = "example"
  const chat = "Chat"

</script>


<template>
  <AppHeader :name="chat"/>
  <div id="container">
    <div>
      <h2>Teams</h2>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <div class="line"></div>

      <br>
      
      <h2>Leagues</h2>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <div class="line"></div>

      <br>

      <h2>Channels</h2>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <div class="line"></div>

      <br>

      <h2>Direct Messages</h2>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <Card :name="name" :link="chatRoom"/>
      <div class="line"></div>
    </div>
    <FloatingButton :link="createChannel"/>
  </div>
</template>

<style scoped>
    .line{
        border-bottom: 1px solid #f1f0f0; /* Adjust thickness and color as needed */
        margin: 0; /* Adjust margin as needed */
    }
</style>