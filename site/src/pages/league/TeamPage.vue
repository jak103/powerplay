<template>
  <q-page style="margin: 50px">
    <q-card class="bg-primary" style="padding: 20px; display: flex" square>
      <q-img
        src="https://cdn.quasar.dev/img/image-src.png"
        spinner-color="white"
        style="height: 100px; max-width: 100px"
      />
      <p class="title">{{ team }}</p>
      <p class="record">(1 - 3)</p>
      <q-btn flat bordered @click="table = 'stats'" style="color: white">Stats</q-btn>
      <q-btn flat bordered @click="table = 'standings'" style="color: white">Standings</q-btn>
    </q-card>
    <TeamStats v-if="table === 'stats'"/>
    <TeamStandings v-if="table == 'standings'"/>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import TeamStats from './TeamStats.vue';
import TeamStandings from './TeamStandings.vue';

const route = useRoute();
const team = ref(route.params.team);
const table = ref('stats');
</script>

<style scoped>
  .title {
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    font-size: 50px;
    color: white;
    text-align: center;
    padding: 15px;
  }

  .record {
    font-family: Arial, Helvetica, sans-serif;
    font-size: 20px;
    color: white;
    text-align: center;
    padding: 15px;
  }
</style>
