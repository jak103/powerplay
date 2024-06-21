<template>
  <q-card class="bg-secondary" square>
    <q-card-section>
      <q-table square
        title="Skaters"
        :rows="skaters"
        :columns="skaterColumns"
        row-key="name"
        hide-bottom
        :loading="!isLoaded"
      />
    </q-card-section>
    <q-card-section>
      <q-table square
        title="Goalies"
        :rows="goalies"
        :columns="goalieColumns"
        row-key="name"
        hide-bottom
        :loading="!isLoaded"
      />
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { QTableProps } from 'quasar';

type Skater = {
  name: string,
  number?: number,
  position?: number,
  gp: number,
  g: number,
  a: number,
  pts: number,
  pim: number,
  pen: number
}

type Goalie = {
  name: string,
  number?: number,
  gp: number,
  w: number,
  l: number,
  min: number
}

const skaters = ref<Skater[]>([]);
const goalies = ref<Goalie[]>([]);

const isLoaded = ref(false); // Sets the tables to loading when data is being fetched

// Used to customize columns in the table
// TODO: create tooltips for the stats columns in the headers
const skaterColumns: QTableProps['columns'] = [
  { name: 'name', required: true, label: 'Name', align: 'left', field: 'name' },
  { name: 'number', label: 'Number', align: 'left', field: 'number'},
  { name: 'position', label: 'Position', align: 'left', field: 'number'},
  { name: 'gp', required: true, label: 'GP', align: 'center', field: 'gp' },
  { name: 'g', required: true, label: 'G', align: 'center', field: 'g' },
  { name: 'a', required: true, label: 'A', align: 'center', field: 'a' },
  { name: 'pts', required: true, label: 'PTS', align: 'center', field: 'pts' },
  { name: 'pim', required: true, label: 'PIM', align: 'center', field: 'pim' },
  { name: 'pen', required: true, label: 'PEN', align: 'center', field: 'pen' },
];

const goalieColumns: QTableProps['columns'] = [
  { name: 'name', required: true, label: 'Name', align: 'left', field: 'name' },
  { name: 'number', label: 'Number', align: 'left', field: 'number'},
  { name: 'gp', required: true, label: 'GP', align: 'center', field: 'gp' },
  { name: 'w', required: true, label: 'W', align: 'center', field: 'w' },
  { name: 'l', required: true, label: 'L', align: 'center', field: 'l' },
  { name: 'min', required: true, label: 'MIN', align: 'center', field: 'min' },
];

// Delays code for a given time period
// Used to simulate delay for testing the loading on data
function delay(milliseconds: number) {
  return new Promise(resolve => setTimeout(resolve, milliseconds));
}

async function fetchTeamStats() {
  // Simulate delay to test loading bar
  // TODO: remove artificial delay, along with the delay function
  await delay(1000);
  isLoaded.value = true;

  // Filler data
  // TODO: replace filler data with API call
  const response = {
    skaters: [
      {
        name: 'Skater 1', number: 1,
        gp: 1, g: 1, a: 0, pts: 1, pim: 0, pen: 0
      },
      {
        name: 'Skater 2', number: 53,
        gp: 3, g: 0, a: 0, pts: 0, pim: 0, pen: 0
      }
    ],
    goalies: [
      {
        name: 'Goalie 1', number: 40,
        gp: 1, w: 0, l: 1, min: 3600
      }
    ],
  }

  skaters.value = response.skaters;
  goalies.value = response.goalies;
};

fetchTeamStats();
</script>
