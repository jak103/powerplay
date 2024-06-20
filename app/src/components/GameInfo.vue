<template>
  <q-card bordered class="match-info" @click="goToGameDetails">
    <q-card-section class="q-pa-xs">
      <q-item class="team-wrapper">
        <q-item-section avatar>
          <q-img :src="getLogoSrc(game.home_team.logo_id)" alt="Home Team Logo" class="team-logo" />
        </q-item-section>
        <q-item-section class="team-details">
          <q-item-label :class="{ 'text-bold': game.home_team_score > game.away_team_score }" class="team-name">
            {{ game.home_team.name }}
          </q-item-label>
          <q-item-label caption class="home-away-label">Home</q-item-label>
        </q-item-section>
        <q-item-section side class="score-section">
          <q-item-label :class="{ 'text-bold': game.home_team_score > game.away_team_score }">
            {{ game.home_team_score !== null ? game.home_team_score : '' }}
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-card-section>
    <q-card-section class="q-pa-xs">
      <q-item class="team-wrapper">
        <q-item-section avatar>
          <q-img :src="getLogoSrc(game.away_team.logo_id)" alt="Away Team Logo" class="team-logo" />
        </q-item-section>
        <q-item-section class="team-details">
          <q-item-label :class="{ 'text-bold': game.away_team_score > game.home_team_score }" class="team-name">
            {{ game.away_team.name }}
          </q-item-label>
          <q-item-label caption class="home-away-label">Away</q-item-label>
        </q-item-section>
        <q-item-section side class="score-section">
          <q-item-label :class="{ 'text-bold': game.away_team_score > game.home_team_score }">
            {{ game.away_team_score !== null ? game.away_team_score : '' }}
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-card-section>
    <q-card-section class="q-pa-xs">
      <q-item class="datetime-wrapper">
        <q-item-section class="datetime-details">
          <q-item-label>{{ formatTimeInterval(game.start) }}</q-item-label>
          <q-item-label>{{ formatDate(game.start) }}</q-item-label>
        </q-item-section>
        <q-item-section v-if="showRsvpButton" side class="rsvp-section">
          <q-btn v-if="!props.hasRsvped" size="sm" color="primary" class="active-button no-cursor-change no-pointer-events">
            RSVP
          </q-btn>
          <q-icon v-else name="check_circle" color="green" size="md" class="checkmark" />
        </q-item-section>
      </q-item>
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';
import { Game } from 'src/models/Game';
import { useRouter } from 'vue-router';
import { useScheduleStore } from 'src/stores/scheduleStore';
import { onMounted, ref } from 'vue';
import type { LogoId } from 'src/models/ids';

const router = useRouter();
const store = useScheduleStore();
const logos = ref<{ id: LogoId, src: string }[]>([]);

onMounted(() => {
  store.loadExampleData();
  logos.value = store.logos;
});

interface Props {
  game: Game;
  showRsvpButton: boolean;
  hasRsvped: boolean;
}

const props = defineProps<Props>();

const getLogoSrc = (logoId: LogoId): string => {
  const logo = store.logos.find(logo => logo.id === logoId);
  return logo ? logo.src : '';
};

const formatTime = (timeString: string): string => {
  const date = new Date(timeString);
  let hours = date.getHours();
  const minutes = date.getMinutes();
  const ampm = hours >= 12 ? 'PM' : 'AM';
  hours = hours % 12;
  hours = hours ? hours : 12; // the hour '0' should be '12'
  const minutesStr = minutes < 10 ? '0' + minutes : minutes;
  return `${hours}:${minutesStr} ${ampm}`;
};

const formatTimeInterval = (startTime: string): string => {
  const startDate = new Date(startTime);
  const endDate = new Date(startDate.getTime() + 75 * 60 * 1000); // Add 1 hour and 15 minutes
  return `${formatTime(startDate.toISOString())} - ${formatTime(endDate.toISOString())}`;
};

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  const options: Intl.DateTimeFormatOptions = {
    weekday: 'short',
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  };
  return date.toLocaleDateString('en-US', options);
};
function goToGameDetails() {
  if (props.showRsvpButton) {
    router.push({ name: 'GameDetailsPage', params: { gameId: props.game.id, teamId: props.game.away_team.id } });
  }
}
</script>
<style scoped>
.match-info {
  cursor: pointer;
}
.team-logo {
  width: 8vw;
  height: 8vw;
  border-radius: 10%;
}
.text-bold {
  font-weight: 900;
  color: black;
}
.checkmark {
  font-size: 2.5em;
}
.active-button {
  background-color: #1976d2;
  color: white;
  opacity: 1;
}

/* prevents the cursor from changing to a pointer when hovering over the button */
.no-cursor-change {
  cursor: default;
}
.no-pointer-events {
  pointer-events: none;
}
@media (min-width: 768px) {
  .team-logo {
    width: 4vw;
    height: 4vw;
  }
}
</style>
