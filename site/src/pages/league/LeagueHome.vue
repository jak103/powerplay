<template>
  <q-page style="margin: 50px">
    <q-card v-for="(league, i) in leagues" :key="i" class="bg-primary" style="color: white; margin: 20px; padding: 10px; text-align: center">
      <h1>{{ league.name }}</h1>
      <q-btn v-for="(team, j) in league.teams" :key="j" :to="'league/' + team" class="
      bg-secondary" style="margin: 10px; color: black">{{ team }}</q-btn>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';

type League = {
  name: string,
  teams: string[]
}

const leagues = ref<League[]>([]);

async function fetchLeagues() {
  // Filler data for now that can be replaced with appropriate API calls
  // TODO: Replace data with API call
  const response = [
    {
      name: 'A League',
      teams: ['Chiefs', 'Mountain Men', 'Pine Hill', 'Reapers', 'Sharknado', 'Yetis']
    },
    {
      name: 'B League',
      teams: ['Chicken Tendies', 'Esqueletos', 'Jr Yetis', 'P-Wings', 'Puck Bandits', 'Rusty Blades']
    },
    {
      name: 'C League',
      teams: ['Big Willies', 'Marauders', 'Rebels', 'The Crew', 'The Neighborhood', 'Trash Pandas']
    },
    {
      name: 'D League',
      teams: ['District 5', 'Iceholes', 'Pistol Shrimp', 'Redshift', 'Sentinels', 'Trashers']
    }
  ];

  leagues.value = response;
}

fetchLeagues();
</script>
