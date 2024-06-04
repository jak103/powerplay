<template>
    <q-card @click="handleClick" class="q-mb-md q-pa-md q-mr-sm q-ml-sm">
      <q-card-section class="row no-wrap items-center">
        <div v-if="image" class="styled-image-container">
          <q-img :src="image" class="styled-image" />
        </div>
        <div v-else-if="type !== 'channel'" :style="{ backgroundColor: color() }" class="box">{{ name.charAt(0).toUpperCase() }}</div>
        <div v-else class="channel">#</div>
        <div class="title q-ml-md">{{ name }}</div>
        <q-badge v-if="unread >= 1" color="negative" class="unread">{{ unread }}</q-badge>
      </q-card-section>
    </q-card>
  </template>
  
  <script setup lang="ts">
  import { defineProps } from 'vue';
  import { useRouter } from 'vue-router';
  
  const props = defineProps<{
    name: string;
    link: string;
    unread: number;
    type: string;
    image: string | null;
  }>();
  
  const router = useRouter();
  
  const handleClick = () => {
    router.push({ path: props.link, query: { chatName: props.name } });
  };
  
  const color = () => {
    const colors = ['#187eb3', '#F86C55', '#DF2F56', '#EF5C48', '#F57A45', '#D74E9C', '#21B2EC', '#d1a2ff'];
    const firstLetter = props.name.charAt(0).toUpperCase();
    const asciiValue = firstLetter.charCodeAt(0);
    const colorIndex = asciiValue % colors.length;
    return colors[colorIndex];
  };
  </script>
  
  <style scoped>
  .styled-image-container {
    margin-left: 0.5rem;
  }
  
  .styled-image {
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 10%;
  }
  
  .channel {
    font-size: x-large;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-left: 0.5rem;
    width: 30px;
    height: 30px;
  }
  
  .unread {
    font-size: 12px;
    background-color: #dc3545;
    display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    padding-left: 0.4rem;
    padding-right: 0.4rem;
    height: 1rem;
    border-radius: 1rem;
    color: white;
    float: right;
    margin-left: auto;
  }
  
  .box {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-left: 0.5rem;
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 10%;
  }
  
  .title {
    margin-left: 1rem;
  }
  
  .q-card {
    display: flex;
    align-items: center;
    width: 100%;
    height: 3rem;
    padding: 0;
  }
  
  .q-card-section {
    display: flex;
    align-items: center;
    width: 100%;
    height: 100%;
    padding: 0;
  }
  </style>
  