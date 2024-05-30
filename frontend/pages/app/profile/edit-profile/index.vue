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
      <div @click="() => this.$router.go(-1)" ><</div>
      <div class="heading-container">
        <div class="top-profile">Edit Profile</div>
        <div class="top-username">{{ user }}</div>
      </div>
      <div></div>
    </div>

  <div class="contianer">
  <div class="edit-info-container">
    <div class="info-title">Personal Information</div>
    <div class="input-title">First Name</div>
    <input type="text">
    <div class="input-title">Preferred Name</div>
    <input type="text">
    <div class="input-title">Last Name</div>
    <input type="text">
    <div class="input-title">Date of Birth</div>
    <input type="text">
  </div>
  <div class="edit-info-container">
    <div class="info-title">Contact Information</div>
    <div class="input-title">Email & Username</div>
    <input type="text">
    <div class="input-title">Phone Number</div>
    <input type="text">
  </div>
  
  <div class="edit-info-container">
    <div class="info-title">Password</div>
    <div class="input-title">Current Password</div>
    <input type="password">
    <div class="input-title">New Password</div>
    <input type="password">
    <div class="input-title">Confirm New Password</div>
    <input type="password">
  </div>
<div class="forgot-password">
  <NuxtLink to="/app/profile">I forgot my password</NuxtLink>
</div>

  <div class="bottom-buttons">
    <div class="vstack d-inline-flex gap-2" >
      <NuxtLink class="btn btn-primary" to="/app/profile">Cancel</NuxtLink>
      <NuxtLink class="btn btn-primary" to="/app/profile" @click="saveInfo">Save Info</NuxtLink>
    </div>
  </div>
</div>
</template>

<style>

.forgot-password{
  margin-top: 10px;
  margin-bottom: 20px;
}

.info-title {
    font-size: 25px;
    font-weight: bold;
    margin-top: 25px;
}

input{
  width: 65vw;
  border: 1px solid;
  border-color: lightgray;
  border-radius: 5px;
  margin-bottom: 10px;
}

.input-title{
  margin-bottom: 5px;
}

.edit-info-container{
  width:100%;
}

.container{
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  width: 100vw;
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