<template>
  <q-page class="q-pa-md">
    
    <div class="text-h6">Upcoming Events</div>

    <div class="q-pa-md">
      <q-card class="q-mb-md">
        <q-card-section>
          <q-item>
            <q-avatar size="50px" class="q-mr-md">
              <img src="path/to/logo.png" alt="Logo">
            </q-avatar>
            <q-item-section>
              <q-item-label>
                <div class="text-h5">{{ game[0].hometeam }}</div>
                <div class="text-subtitle2">Home</div>
              </q-item-label>
            </q-item-section>
          </q-item>
          <q-item>
            <q-avatar size="50px" class="q-mr-md">
              <img src="path/to/logo.png" alt="Logo">
            </q-avatar>
            <q-item-section>
              <q-item-label>
                <div class="text-h5">{{ game[0].awayteam }}</div>
                <div class="text-subtitle2">Locker Room 2</div>
              </q-item-label>
            </q-item-section>
          </q-item>
          <div class="text-subtitle2 q-ml-md">{{ game[0].date }}</div>
          <div class="text-subtitle2 q-ml-md">{{ game[0].time }}</div>
        </q-card-section>
      </q-card>
    </div>
    <div class="text-h6">My Teams</div>

    <q-item v-for="team in teams" :key="team.id" class="q-mb-sm border-bottom">
        <q-item clickable v-ripple @click="goToTeamInfo(team.name)">
            <q-avatar size="30px" class="q-mr-md">
                <img src="team.logo" alt="Logo">
            </q-avatar>
            <q-item-section>
                <div class="text-h6 text-weight-regular">{{ team.name }}</div>
            </q-item-section>
        </q-item>
    </q-item>

    <div class="text-h6">My Leagues</div>

    <q-item v-for="team in teams" :key="team.id" class="q-mb-sm border-bottom">
        <q-item clickable v-ripple @click="goToLeagueInfo(team.league)">
            <q-avatar size="30px" class="q-mr-md">
                <img src="team.logo" alt="Logo">
            </q-avatar>
            <q-item-section>
                <div class="text-h6 text-weight-regular">{{ team.league }} League</div>
            </q-item-section>
        </q-item>
    </q-item>

  </q-page>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const teams = ref([
  {
    id: 1,
    name: 'District 5',
    logo: 'path/to/district5_logo.png',
    league: 'B',
    manager: 'Captain Hook',
  },
  {
    id: 2,
    name: 'Trash Pandas',
    logo: 'path/to/trashpandas_logo.png',
    league: 'C',
    manager: 'Jacob Christensen',
  }
]) 

const game = ref([
  {
    id: 1,
    hometeam: 'Red Shift',
    homelogo: 'path/to/redshift_logo.png',
    awayteam: 'District 5',
    awaylogo: 'path/to/redshift_logo.png',
    date: 'Wed, January 31, 2024',
    time: '9:00 – 10:15 PM',
  },
]) 

const goToTeamInfo = (teamName) => {
  const encodedTeamName = encodeURIComponent(teamName);
  router.push({ name: 'TeamInfo', params: { teamName: encodedTeamName } });
};

const goToLeagueInfo = (league) => {
  router.push({ name: 'LeagueInfo', params: { id: league } });
};

</script>

<style scoped>
.border-bottom {
  border-bottom: 1px solid #ccc; /* Adjust the color and width as needed */
}
</style>
