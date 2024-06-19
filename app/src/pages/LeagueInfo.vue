<template>
  <q-layout>
  <q-page-container>
  <q-page class="q-pa-md">

    <q-card-section>
          <q-item>
            <q-avatar size="65px" class="q-mr-md">
              <img src="path/to/logo.png" alt="Logo">
            </q-avatar>
            <q-item-section>
              <q-item-label>
                <div class="text-h4">{{ league.name }} League</div>
                <div class="text-subtitle2">League Manager:</div>
                <div class="text-subtitle2 text-weight-regular"> {{ league.manager }}</div>
              </q-item-label>
            </q-item-section>
          </q-item>
    </q-card-section>

    <q-card-section>
      <div class="row justify-between">
        <router-link
              v-for="item in items"
              :key="item.label"
              :id="'label-' + item.label"
              :to="{ name: routesMapping[item.label] }"
              class="q-ma-sm card-square rounded-borders shadow-2"
              style="text-decoration: none; color: inherit;"
            >
          <q-card-section class="text-center">
            <q-icon :name="item.icon" size="40px" class="q-mb-sm" />
            <div class="text-subtitle2">{{ item.label }}</div>
          </q-card-section>
        </router-link>
      </div>
    </q-card-section>
  </q-page>
  </q-page-container>
  </q-layout>
</template>


<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const leagueName = route.params.id;

// Replace this with actual league data retrieval logic
const league = ref({
  name: leagueName,
  logo: 'path/to/league_logo.png',
  manager: 'Jacob Christensen',
})

const items = ref([
  { icon: 'event', label: 'Schedule' },
  { icon: 'chat', label: 'Chat' },
  { icon: 'people', label: 'Roster' },
  { icon: 'list_alt', label: 'Sub' },
  { icon: 'assessment', label: 'Standings' },
  { icon: 'insights', label: 'Statistics' },
]);

const routesMapping = {
  'Schedule': 'SchedulePage',
  'Chat': 'ChatPage',
  'Roster': 'RosterPage', /* change to RosterPage when we have */
  'Sub': 'SubPage', /* change to SubPage when we have */
  'Standings': 'StandingsPage', /* change to StandingsPage when we have */
  'Statistics': 'StatsPage', /* change to StatisticsPage when we have */
};
</script>

<style scoped>
.card-square {
  width: 120px; /* Adjust size as needed */
  height: 120px;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
