<script lang="ts" setup>
  import FloatingButton from '~/components/FloatingButton.vue';
  import AppHeader from '~/components/AppHeader.vue';
  import Card from '~/components/Card.vue';

  
  // TODO: pull data in this form from the backend
  const data = {
    channels : [
      {
        "channel_name": "Trash Pandas",
        "channel_id": "example1",
        "channel_image": "mapleleafs_image.jpg",
        "missed_chats": 1,
        "type" : "team"
        
      },
      {
        "channel_name": "District 5",
        "channel_id": "example2",
        "channel_image": "canadiens_image.jpg",
        "missed_chats": 16,
        "type" : "team"
      },
      {
        "channel_name": "C-League",
        "channel_id": "example3",
        "channel_image": null,
        "missed_chats": 0,
        "type" : "league"
      },
      {
        "channel_name": "D-League",
        "channel_id": "example4",
        "channel_image": null,
        "missed_chats": 0,
        "type" : "league"
      },
      {
        "channel_name": "C-League Subs",
        "channel_id": "example5",
        "channel_image": null,
        "missed_chats": 3,
        "type" : "channel"
      },
      {
        "channel_name": "D-League Subs",
        "channel_id": "example6",
        "channel_image": null,
        "missed_chats": 0,
        "type" : "channel"
      },
      {
        "channel_name": "News & Announcements",
        "channel_id": "example7",
        "channel_image": null,
        "missed_chats": 1,
        "type" : "channel"
      },
      {
        "channel_name": "Corey Truetel",
        "channel_id": "example8",
        "channel_image": "oilers_image.jpg",
        "missed_chats": 0,
        "type" : "dm"
      },
      {
        "channel_name": "D-League Subs",
        "channel_id": "example9",
        "channel_image": "canucks_image.jpg",
        "missed_chats": 3,
        "type" : "dm"
      },
      {
        "channel_name": "Scott Pococha",
        "channel_id": "example10",
        "channel_image": "canucks_image.jpg",
        "missed_chats": 0,
        "type" : "dm"
      }
    ]
  }

  const createChannel = "/app/chat/create-channel"
  const chatRoom = "/app/chat/chat-room"
  const chat = "Chat"

  const teams = data.channels.filter(channel => channel.type === "team");
  const leagues = data.channels.filter(channel => channel.type === "league");
  const channels = data.channels.filter(channel => channel.type === "channel");
  const dms = data.channels.filter(channel => channel.type === "dm");

</script>


<template>
  <AppHeader :name="chat"/>
  <div id="container">
    <div>
      <h2>Teams</h2>
      <Card v-for="channel in teams" :key="channel.channel_id" :type="channel.type"
      :name="channel.channel_name" :link="chatRoom" :unread="channel.missed_chats"/>
      <div class="line"></div>
      <br>
      
      <h2>Leagues</h2>
      <Card v-for="channel in leagues" :key="channel.channel_id" :type="channel.type"
      :name="channel.channel_name" :link="chatRoom" :unread="channel.missed_chats"/>
      <div class="line"></div>
      <br>

      <h2>Channels</h2>
      <Card v-for="channel in channels" :key="channel.channel_id" :type="channel.type"
      :name="channel.channel_name" :link="chatRoom" :unread="channel.missed_chats"/>
      <div class="line"></div>
      <br>

      <h2>Direct Messages</h2>
      <Card v-for="channel in dms" :key="channel.channel_id" :type="channel.type"
      :name="channel.channel_name" :link="chatRoom" :unread="channel.missed_chats"/>
      <div class="line"></div>
      <br>
      <br>
      <br>
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