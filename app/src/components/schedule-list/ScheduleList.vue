<script setup lang="ts">
import {
  QIcon, QImg, QTable
} from 'quasar'
import {ref} from 'vue';
/* eslint-disable */
const props = defineProps({color: { type: String, default: 'bg-teal-5 text-black'}})

interface Column {
  name: string;
  label: string;
  field: string | ((row: any) => any);
  required?: boolean;
  align?: "left" | "center" | "right";
  sortable?: boolean;
  sort?: (a: any, b: any, rowA: any, rowB: any) => number;
  format?: (val: any, row: any) => any;
  style?: string | ((row: any) => any);
  classes?: string | ((row: any) => any);
  headerStyle?: string;
  headerClasses?: string;
}
const columns: Column[] = [
  {
    name: 'opponent',
    required: true,
    label: 'opponent',
    align: 'left',
    field: 'opponent',
    sortable: false
  },
  { name: 'outcome', align: 'center', label: 'Win / Loss', field: 'outcome', sortable: false },
  { name: 'score', label: 'score', field: 'score', sortable: false },
  { name: 'game_status', label: 'game_status', field: 'game_status' },
  { name: 'game_start', label: 'game_start', field: 'game_start' },
  { name: 'venue', label: 'venue', field: 'venue' },
]

const rows = ref([
  {
    outcome: 'L',
    score: '6-4',
    game_status: 'Final',
    game_start: '8:30PM MDT',
    venue: 'George S. Eccles Ice Center - Sheet 1',
    image: 'https://se-team-service-production.s3.amazonaws.com/uploads/team/logo/11ef1251-5574-992a-91e3-7aea041d6990/medium_logo.png',
    teamName: 'Esqueletos'
  },
  {
    teamName: 'P-Wings',
    game_start: '8:45PM MDT',
    venue: 'George S. Eccles Ice Center - Sheet 1',
  },
  {
    teamName: 'Jr Yetis',
    game_start: '9:00PM MDT',
    venue: 'George S. Eccles Ice Center - Sheet 1',
  },
  {
    teamName: 'Puck Bandits',
    game_start: '9:00PM MDT',
    venue: 'George S. Eccles Ice Center - Sheet 1',
  },
  {
    teamName: 'Rusty Blades',
    game_start: '10:15PM MDT',
    venue: 'George S. Eccles Ice Center - Sheet 1',
  }
]);

</script>

<template>
  <q-table
    :card-class="color"
    :rows="rows"
    :columns="columns"
    row-key="name"
    hide-header
    hide-bottom
    binary-state-sort
  >
    <template v-slot:body-cell-opponent="props">
      <q-td :props="props">
        <span>VS</span>
          <q-img width="30px" height="30px" v-if="props.row.image" :src="props.row.image"></q-img>
          <q-icon size="lg" v-else name="sports_hockey"></q-icon>
        <span>{{props.row.teamName}}</span>
      </q-td>
    </template>
  </q-table>
</template>

<style scoped>

</style>
