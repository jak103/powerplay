<template>
  <q-card class="match-info">
    <q-card-section class="q-pa-xs">
      <q-item class="team-wrapper">
        <q-item-section avatar>
          <q-img :src="homeTeamLogo" alt="Logo" class="team-logo" />
        </q-item-section>
        <q-item-section class="team-details">
          <q-item-label
            :class="{ 'text-bold': homeScore > awayScore }"
            class="team-name"
          >
            {{ homeTeam }}
          </q-item-label>
          <q-item-label caption class="home-away-label">Home</q-item-label>
        </q-item-section>
        <q-item-section side class="score-section">
          <q-item-label
            :class="{ 'text-bold': homeScore > awayScore }"
          >
            {{ homeScore !== null ? homeScore : '' }}
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-card-section>
    <q-card-section class="q-pa-xs">
      <q-item class="team-wrapper">
        <q-item-section avatar>
          <q-img :src="awayTeamLogo" alt="Logo" class="team-logo" />
        </q-item-section>
        <q-item-section class="team-details">
          <q-item-label
            :class="{ 'text-bold': awayScore > homeScore }"
            class="team-name"
          >
            {{ awayTeam }}
          </q-item-label>
          <q-item-label caption class="home-away-label">Away</q-item-label>
        </q-item-section>
        <q-item-section side class="score-section">
          <q-item-label
            :class="{ 'text-bold': awayScore > homeScore }"
            class="score"
          >
            {{ awayScore !== null ? awayScore : '' }}
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-card-section>
    <q-card-section class="q-pa-xs">
      <q-item class="datetime-wrapper">
        <q-item-section class="datetime-details">
          <q-item-label>{{ date }}</q-item-label>
          <q-item-label>{{ time }}</q-item-label>
        </q-item-section>
        <q-item-section v-if="showRsvpButton" side class="rsvp-section">
          <q-btn v-if="!hasRsvped" size="sm" color="primary" @click="goToGameDetails">RSVP</q-btn>
          <q-icon
            v-else
            name="check_circle"
            color="green"
            size="md"
            class="checkmark"
          />
        </q-item-section>
      </q-item>
    </q-card-section>
  </q-card>
</template>

<script>
export default {
  props: {
    homeTeam: {
      type: String,
      default: 'The Homeys',
    },
    awayTeam: {
      type: String,
      default: 'A Way Good Team',
    },
    date: {
      type: String,
      default: 'Wed, Jan 25, 2024',
    },
    time: {
      type: String,
      default: '9:00 - 10:15 PM',
    },
    hasRsvped: {
      type: Boolean,
      default: false,
    },
    homeScore: {
      type: Number,
      default: 0,
    },
    awayScore: {
      type: Number,
      default: 0,
    },
    homeTeamLogo: {
      type: String,
      default: 'src/assets/test/homeys.png',
    },
    awayTeamLogo: {
      type: String,
      default: 'src/assets/test/awaygoodteam.png',
    },
    showRsvpButton: {
      type: Boolean,
      default: true,
    },
  },
  methods: {
    goToGameDetails() {
      this.$router.push({ path: '/game-details' });
    },
  },
};
</script>

<style scoped>
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

@media (min-width: 768px) {
  .team-logo {
    width: 4vw;
    height: 4vw;
  }
}
</style>
