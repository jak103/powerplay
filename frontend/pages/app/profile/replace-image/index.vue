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
        <div class="top-profile">Replace Image</div>
        <div class="top-username">{{ user }}</div>
      </div>
      <div></div>
    </div>

  <div class="contianer">
      <div class="replace-image-container">
          <img src="/_nuxt/assets/ProfilePic.webp" width="150" height="150">
          <div class="circle-overlay"></div>
        </div>
        
        <button class="delete-button">Delete Image</button>
        <div>Select Image</div>
        <input type="file" accept="image/*" >

  <div class="bottom-buttons">
    <div class="vstack d-inline-flex gap-2" >
      <NuxtLink class="btn btn-primary" to="/app/profile">Cancel</NuxtLink>
      <NuxtLink class="btn btn-primary" to="/app/profile">Save Image</NuxtLink>
    </div>
  </div>
</div>
</template>

<style>

.delete-button{
    background-color: red;
    border-style: none;
    border-radius: 5px;
    padding: 5px 10px 5px 10px;
    color: white;
    margin-bottom: 20px;
}

.circle-overlay {
  position: absolute;
  top: 95px;
  width: 150px;
  height: 149px;
  border-radius: 50%;
  border: 1px solid rgba(255, 255, 255, 0.8);
  pointer-events: none;
  box-shadow: 0 0 0 45px rgba(50, 50, 50, 0.7);
  overflow: hidden;
}

.replace-image-container{
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(0, 0, 0, 0.8);
    width: 100%;
    margin-top: 25px;
    padding: 45px 25px 45px 25px;
    margin-bottom: 25px;
    overflow: hidden;
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