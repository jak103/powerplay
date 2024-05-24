<script>
  const userInfo = {
    "name": "Ricky Bobby",
    "phone-number": "901-123-4567",
    "email": "RickyBobby@gmail.com",
    "birthday": "12/12/1998",
    "teams": [{
      name: "Trash Pandas",
      logo: "/_nuxt/assets/trash-panda.webp"}, 
      {
      name: "District 5",
      logo: "/_nuxt/assets/District-5.webp"}],
      "leagues": ["C", "D"]
}
export default {
  data() {
    return {
      user: userInfo["name"],
      phoneNumber: userInfo["phone-number"],
      email: userInfo["email"],
      birthday: userInfo["birthday"],
      teams: userInfo["teams"],
      leagues: userInfo["leagues"]
    };
  }
};
</script>

<template>
  <link rel="stylesheet" 
href="https://fonts.googleapis.com/icon?family=Material+Icons">
    <div class="top-banner">
      <div @click="() => this.$router.go(-1)" >Back</div>
      <div class="heading-container">
        <div class="top-profile">Profile</div>
        <div class="top-username">{{ user }}</div>
      </div>
      <div>Menu</div>
    </div>
  <img src="../../../assets/ProfilePic.webp" width="100" height="100" 
class="profile-image-container">
  <NuxtLink class="btn btn-primary" to="/app/sign-in" id="btn-image" 
style="font-size: 0px;">
    <i class="material-icons icon">photo_camera</i>
  </NuxtLink>
  <div class="contianer">
  <div class="profile-container">
    <h1>{{user}}</h1>
    <div>{{phoneNumber}}</div>
    <div>{{email}}</div>
    <div>Birthday: {{birthday}}</div>
  </div>
  <div class="teams-container">
    <h4 class="teams-heading">Teams</h4>
    <ul>
      <li v-for="team in teams">
      <img :src="team.logo" width="25" height="25" class="team-logo"> {{ 
team.name }}
    </li>
    </ul>
  </div>
  <div class="teams-container">
    <h4 class="teams-heading">Leagues</h4>
    <ul>
      <li v-for="league in leagues">
      <img :src="'/_nuxt/assets/' + league + '.jpeg'" width="25" 
height="25" class="team-logo"> {{ league }}-League
    </li>
    </ul>
    <div class="emergency-contact">
  <h5>Emergency Contact Info</h5>
  <div>Contact Name: Reese Bobby</div>
  <div>Contact Number: 801-123-1234</div>
  <NuxtLink class="btn btn-primary" to="/app/sign-in">Edit</NuxtLink>
    </div>
  </div>
  <div class="bottom-buttons">
    <div class="vstack d-inline-flex gap-2" >
      <NuxtLink class="btn btn-primary" to="/app/sign-in">Sign 
Out</NuxtLink>
      <NuxtLink class="btn btn-primary" to="/app/sign-in">Edit 
Profile</NuxtLink>
    </div>
  </div>
</div>
</template>

<style>
#btn-image{
  position: absolute;
  border-radius: 100px;
  top: 150px;
  width: 10px;
  height: 25px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: 50px;
}
.material-icons.icon{
  font-size: 13px;
}
.emergency-contact{
  margin-top: 50px;
  font-size: 12px;
  margin-bottom: 50px;
}
.emergency-contact div{
  margin-bottom: 10px;
}
.team-logo{
  border-radius: 5px;
  margin-right: 3px;
}
.profile-image-container{
  border-radius: 100px;
  margin-top: 45px;
  margin-bottom: 5px;
}
.teams-heading{
  align-self: baseline;
}
.top-profile{
  font-size: 20px;
}
.top-username{
  font-size: 10px;
  margin-bottom: 5px;
}
.bottom-buttons div{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: space-around;
  width: 100%;
}
.bottom-buttons div a{
  width: 45%;
}
.profile-container{
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  font-size: 13px;
}
.profile-container div{
  margin-top: 10px;
}
.teams-container{
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  width: 80vw;
}
ul{
  width: 80vw;
}
li {
  list-style-type: none;
  width: 90%;
  padding: 10px 10px 10px 10px;
  border-bottom-style: groove;
}
li:first-child{
  border-top-style: groove;
}
.container{
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}
.top-banner{
  display: none;
}
@media (max-width:991px)
{.top-banner{
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  background-color: #303030;
  color: white;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  outline:none;
}
.heading-container{
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
}
</style>