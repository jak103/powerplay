<template>
  <div>
    <GameInfo
      v-for="(game, index) in games"
      :key="index"
      :game="game"
      :showRsvpButton="true"
      :hasRsvped="index % 2 === 0 ? true : false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useScheduleStore } from 'src/stores/scheduleStore';
import GameInfo from 'components/GameInfo.vue';
import { Game } from 'src/models/Game';

const store = useScheduleStore();
const games = ref<Game[]>([]);

// load the example data on component mount
onMounted(() => {
  store.loadExampleData();
  games.value = store.games;
});

</script>
