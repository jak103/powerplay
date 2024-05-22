

<template>
  <div>
    <AppHeader :name="Home"/>

    <!-- Upcoming Events Section -->
    <div class="section">
      <h2>Upcoming Events</h2>
      <!-- Box for each game -->
      <div v-for="game in upcomingGames" :key="game.id" class="game-box" @click="handleGameClick(game)">
        <div class="team-info">
          <div class="team-logo">
            <img :src="game.homeLogo">
          </div>
          <div class="team">
            <div class="team-name">{{ game.hometeam }}</div>
            <div class="team-subtitle">Home</div>
          </div>
        </div>
        <div class="team-info">
          <div class="team-logo">
            <img :src="game.awayLogo">
          </div>
          <div class="team">
            <div class="team-name">{{ game.awayteam }}</div>
            <div class="team-subtitle">{{ game.locker }}</div>
          </div>
        </div>
        <div class="date-time">
          <div>{{ game.date }}</div>
          <div>{{ game.time }}</div>
        </div>
      </div>
    </div>

    <!-- My Teams Section -->
    <div class="section">
      <h2 class="section-header">My Teams</h2>
      <!-- List of teams -->
      <NuxtLink v-for="team in myTeams" :key="team.id" :to='"/app/team-info"' class="team-box" @click="handleTeamClick(league)">
        <div class="line"></div> <!-- Line above team box -->
        <div class="logo-name-container">
          <img :src="team.logo" class="team-logo">
          <div class="league-name">{{ team.name }}</div>
        </div>
      </NuxtLink>
    </div>

    <!-- My Leagues Section -->
    <div class="section">
      <h2 class="section-header">My Leagues</h2>
      <!-- List of leagues -->
      <NuxtLink v-for="team in myTeams" :key="team.id" :to='"/app/league-info"' class="team-box" @click="handleTeamClick(league)">
        <div class="line"></div> <!-- Line above league box -->
        <div class="logo-name-container">
          <img :src="team.logo"  class="team-logo">
          <div class="league-name">{{ team.division }}</div>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>

<script>
import AppHeader from '~/components/AppHeader.vue';
export default {
  data() {
    return {
      upcomingGames: [
        {
          id: 1,
          hometeam: 'Red Shift',
          awayteam: 'District 5',
          locker: 'Locker Room 2',
          date: 'Wed, January 31, 2024',
          time: '9:00 – 10:15 PM'
        },
        // Add more game objects as needed
      ],
      myTeams: [
        {
          id: 1,
          name: 'District 5',
          logo: 'path/to/team_a_logo.png',
          division: 'A League',
          // Add more team objects as needed
        },
        {
          id: 2,
          name: 'Trash Pandas',
          logo: 'path/to/logo.png',
          division: 'B League',
        }
        // Add more team objects as needed
      ],
      myLeagues: [
        {
          id: 1,
          name: 'A League',
          logo: 'path/to/logo.png',
        },
        {
          id: 2,
          name: 'B League',
          logo: 'path/to/logo.png',
        }
      ]
    };
  },
  methods: {
    handleGameClick(game) {
      // Handle tap/click on a game
      console.log('Game tapped:', game.title);
      // Add your logic here, such as navigating to a game details page
    },
    handleTeamClick(team) {
      // Handle tap/click on a team
      console.log('Team tapped:', team.name);
      // Add your logic here, such as navigating to team details page
    },
  },
};
</script>

<style>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
}

.back-arrow {
  cursor: pointer;
}

.title {
  font-size: 1.5rem;
}

.options {
  font-size: 1.2rem;
}

.section {
  margin-bottom: 20px;
  margin-top: 60px;
}

.team-box {
  display: flex;
  flex-direction: column;
  align-items: left;
  position: relative; /* For positioning the lines */
  text-decoration: none;
  color: inherit;
}

.line {
  width: 100%;
  height: 1px;
  background-color: #ccc; /* Line color */
}

.logo-name-container {
  display: flex;
  align-items: center;
}

.team-logo {
  width: 50px; /* Adjust logo size as needed */
  height: 50px; /* Adjust logo size as needed */
  margin-right: 10px; /* Margin between logo and name */
}

.league-name {
  font-size: 1.2rem;
  text-align: left;
}

.game-box {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: left;
  padding: 15px;
  background-color: #f0f0f0;
  border: 1px solid #ccc;
  border-radius: 5px;
  margin-bottom: 10px;
}

.team-info {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.team-details {
  text-align: center;
}

.team-logo img {
  width: 50px; /* Adjust logo size as needed */
  height: 50px; /* Adjust logo size as needed */
  margin-bottom: 10px; /* Adjust margin between logo and details */
}

.team-name {
  font-weight: bold;
}

.team-subtitle {
  color: #666;
  font-size: 0.8rem;
}

.date-time {
  text-align: left;
}

.team {
  display: flex;
  flex-direction: column;
}

</style>

