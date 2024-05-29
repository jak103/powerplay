<template>
  <q-page class="q-pa-md">

    <q-card-section>
          <q-item>
            <q-avatar size="70px" class="q-mr-md">
              <img src="league.logo" alt="Logo">
            </q-avatar>
            <q-item-section>
              <q-item-label>
                <div class="text-h3">{{ league[0].name }}</div>
                <div class="text-subtitle2">Team Manager:</div>
                <div class="text-subtitle2"> {{ league[0].manager }}</div>
              </q-item-label>
            </q-item-section>
          </q-item>
    </q-card-section>

    <q-card-section>
      <div class="row justify-between">
        <q-card v-for="item in items" :key="item.label" class="q-ma-sm card-square">
          <q-card-section class="text-center">
            <q-icon :name="item.icon" size="40px" class="q-mb-sm" />
            <div class="text-subtitle2">{{ item.label }}</div>
          </q-card-section>
        </q-card>
      </div>
    </q-card-section>
  </q-page>
</template>


<script setup>
import { ref } from 'vue';
//import { useRouter } from 'vue-router';

//const router = useRouter();
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
  { icon: 'list_alt', label: 'Sub List' },
  { icon: 'assessment', label: 'Standings' },
  { icon: 'insights', label: 'Statistics' },
]);
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
