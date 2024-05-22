

<template>
  <div>

    <AppHeader :name="Home"/>

    <!-- League Info -->
    <div class="section">
      <div class="team-info">
        <div class="team-logo">
          <img :src="league.logo">
        </div>
        <div class="team-text">
          <h2 class="team-name">{{ league[0].name }}</h2>
          <p class="team-subtitle">Manager: {{ league[0].manager }}</p>
        </div>
      </div>
    </div>

    <!-- Box container for schedule, chat, roster, standings, and statistics -->
    <div class="box-container">
      <NuxtLink :to='"/app/schedule"' class="box">
        <i class="fas fa-calendar-alt"></i> <!-- Font Awesome icon for Schedule -->
        Schedule
      </NuxtLink>
      <NuxtLink :to='"/app/chat"' class="box">
        <i class="fas fa-comments"></i> <!-- Font Awesome icon for Chat -->
        Chat
      </NuxtLink>
      <NuxtLink :to='"/app/schedule"' class="box">
        <i class="fas fa-users"></i> <!-- Font Awesome icon for Roster -->
        Roster
      </NuxtLink>
      <NuxtLink :to='"/app/schedule"' class="box">
        <i class="fas fa-user-plus"></i> <!-- Font Awesome icon for Standings -->
        Sub List
      </NuxtLink>
      <NuxtLink :to='"/app/schedule"' class="box">
        <i class="fas fa-trophy"></i> <!-- Font Awesome icon for Standings -->
        Standings
      </NuxtLink>
      <NuxtLink :to='"/app/schedule"' class="box">
        <i class="fas fa-chart-bar"></i> <!-- Font Awesome icon for Statistics -->
        Statistics
      </NuxtLink>
    </div>
  </div>
</template>

<script>
import '@fortawesome/fontawesome-free/css/all.css'; // Import Font Awesome CSS
export default {
  data() {
    return {
      league: [
        {
          id: 1,
          name: 'A League',
          logo: 'path/to/logo.png',
          manager: 'Jacob Christensen',
        },
      ]
    };
  },
};
</script>

<style>
/* Add your CSS styles here */
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  position: fixed;
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
  align-items: left;
}

.box-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  flex-direction: row;
  gap: 20px;
  margin: 20px;
}

.box {
  width: 120px;
  height: 120px;
  border: 1px solid #ccc;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  font-size: 1.2rem;
  border-radius: 10px;
  text-decoration: none;
  color: inherit;
}
</style>

