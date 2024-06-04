<template>
  <div>
    <GameInfo
      v-for="(game, index) in games"
      :key="index"
      :homeTeam="game.homeTeam"
      :awayTeam="game.awayTeam"
      :date="game.date"
      :time="game.time"
      :hasRsvped="game.hasRsvped"
      :homeScore="game.homeScore"
      :awayScore="game.awayScore"
      :homeLogo="game.homeLogo"
      :awayLogo="game.awayLogo"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import GameInfo from 'components/GameInfo.vue';

interface Game {
  homeTeam: string;
  awayTeam: string;
  date: string;
  time: string;
  hasRsvped: boolean;
  homeScore: number;
  awayScore: number;
  homeLogo: string;
  awayLogo: string;
}

const games = ref<Game[]>([
  {
    homeTeam: 'The Homeys',
    awayTeam: 'A Way Good Team',
    date: 'Wed, Jan 25, 2024',
    time: '9:00 - 10:15 PM',
    hasRsvped: false,
    homeScore: 0,
    awayScore: 0,
    homeLogo: 'assets/homeys.png',
    awayLogo: 'assets/awaygoodteam.png',
  },
  {
    homeTeam: 'Team Alpha',
    awayTeam: 'Team Beta',
    date: 'Thu, Jan 26, 2024',
    time: '8:00 - 9:15 PM',
    hasRsvped: true,
    homeScore: 5,
    awayScore: 4,
    homeLogo: 'assets/homeys.png',
    awayLogo: 'assets/awaygoodteam.png',
  },
  {
    homeTeam: 'Squad Goals',
    awayTeam: 'Dream Team',
    date: 'Fri, Jan 27, 2024',
    time: '7:00 - 8:15 PM',
    hasRsvped: true,
    homeScore: 5,
    awayScore: 6,
    homeLogo: 'assets/homeys.png',
    awayLogo: 'assets/awaygoodteam.png',
  },
  {
    homeTeam: 'Champions',
    awayTeam: 'Contenders',
    date: 'Sat, Jan 28, 2024',
    time: '6:00 - 7:15 PM',
    hasRsvped: false,
    homeScore: 7,
    awayScore: 8,
    homeLogo: 'assets/homeys.png',
    awayLogo: 'assets/awaygoodteam.png',
  },
]);
</script>
