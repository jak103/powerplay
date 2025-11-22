<template>
  <q-card class="bg-secondary" square>
    <q-card-section>
      <q-table square
        :title="league"
        :rows="standings"
        :columns="standingsColumns"
        row-key="team"
        hide-bottom
        :loading="!isLoaded"
      />
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { QTableProps } from 'quasar';

type Standing = {
  team: string,
  pts: number,
  w: number,
  l: number,
  otw: number,
  otl: number,
  gf: number,
  ga: number,
  gd: number,
  gp: number
}

const league = ref('');
const standings = ref<Standing[]>([]);

const isLoaded = ref(false); // Sets the tables to loading when data is being fetched

// Delays code for a given time period
// Used to simulate delay for testing the loading on data
function delay(milliseconds: number) {
  return new Promise(resolve => setTimeout(resolve, milliseconds));
}

// Used to customize columns in the table
// TODO: create tooltips for the stats columns in the headers
const standingsColumns: QTableProps['columns'] = [
  { name: 'team', required: true, label: 'Team', align: 'left', field: 'team' },
  { name: 'pts', required: true, label: 'PTS', align: 'center', field: 'pts' },
  { name: 'w', required: true, label: 'W', align: 'center', field: 'w' },
  { name: 'l', required: true, label: 'L', align: 'center', field: 'l' },
  { name: 'otw', required: true, label: 'OTW', align: 'center', field: 'otw' },
  { name: 'otl', required: true, label: 'OTL', align: 'center', field: 'otl' },
  { name: 'gf', required: true, label: 'GF', align: 'center', field: 'gf' },
  { name: 'ga', required: true, label: 'GA', align: 'center', field: 'ga' },
  { name: 'gd', required: true, label: 'GD', align: 'center', field: 'gd' },
  { name: 'gp', required: true, label: 'GP', align: 'center', field: 'gp' },
]

async function fetchTeamStandings() {
  // Simulate delay to test loading bar
  // TODO: remove artificial delay, along with the delay function
  await delay(1000);
  isLoaded.value = true;

  // Filler data
  // TODO: replace filler data with API call
  const response = {
    league: 'A League',
    standings: [
      {
        team: 'Chiefs',
        pts: 15, w: 5, l: 0, otw: 0, otl: 0, gf: 30, ga: 17, gd: 13, gp: 5
      },
      {
        team: 'Chiefs',
        pts: 15, w: 5, l: 0, otw: 0, otl: 0, gf: 30, ga: 17, gd: 13, gp: 5
      },
      {
        team: 'Chiefs',
        pts: 15, w: 5, l: 0, otw: 0, otl: 0, gf: 30, ga: 17, gd: 13, gp: 5
      },
      {
        team: 'Chiefs',
        pts: 15, w: 5, l: 0, otw: 0, otl: 0, gf: 30, ga: 17, gd: 13, gp: 5
      },
      {
        team: 'Chiefs',
        pts: 15, w: 5, l: 0, otw: 0, otl: 0, gf: 30, ga: 17, gd: 13, gp: 5
      }
    ]
  }

  league.value = response.league;
  standings.value = response.standings;
}

fetchTeamStandings();
</script>
