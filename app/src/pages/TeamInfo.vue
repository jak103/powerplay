<template>
  <q-layout>
  <q-page-container>
  <q-page padding>
    <q-card-section>
          <q-item>
            <q-avatar size="65px" class="q-mr-md">
              <img src="path/to/logo.png" alt="Logo">
            </q-avatar>
            <q-item-section>
              <q-item-label>
                <div class="text-h4">{{ team.name }}</div>
                <div class="text-subtitle2">Team Manager:</div>
                <div class="text-subtitle2 text-weight-regular"> {{ team.manager }}</div>
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
  </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const teamId = route.params.id;

const team = ref({
  id: teamId,
  name: 'District 5',
  logo: 'path/to/district5_logo.png',
  manager: 'Captain Hook',
})

const items = ref([
  { icon: 'event', label: 'Schedule' },
  { icon: 'chat', label: 'Chat' },
  { icon: 'people', label: 'Roster' },
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
