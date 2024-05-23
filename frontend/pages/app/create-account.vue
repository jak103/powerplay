<script lang="ts" setup>
definePageMeta({
  layout: "auth-layout",
});
const password = ref('')
const confirmPassword = ref('')

const pass_errorMessage = ref('')

const createAccount = () => {
  //useRouter().push('/app')
  console.log('called')
  pass_errorMessage.value = ''
  validatePassword()
  console.log(pass_errorMessage.value)
}
function validatePassword(){
if(password.value != confirmPassword.value){
  pass_errorMessage.value = "Your password doesn't match!"
  return false
}
if(password.value.length < 8){
  pass_errorMessage.value = "Your password must be at least 8 characters long!"
  return false
}

var specialCharacters = ['!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '{', '}', '[', ']', '|', '\\', ':', ';', '"', "'", '<', '>', ',', '.', '/'];
var specialChar = false 
for (let i = 0; i < password.value.length; i++) {
  if (specialCharacters.includes(password.value[i])) {
    specialChar = true
    break;
  }
}
if(!specialChar){pass_errorMessage.value = "Your password must contain a special character!"
  return false}
}


</script>

<template>
  <Title>Create an Account</Title>
  <h1>Create an Account</h1>
  <div class="vstack gap-3">
    <div>
      <label for="first-name" class="form-label">First Name</label>
      <input if="first-name" class="form-control" />
    </div>
    <div>
      <label for="last-name" class="form-label">Last Name</label>
      <input id="last-name" class="form-control" />
    </div>
    <div>
      <label for="email" class="form-label">Email</label>
      <input id="email" class="form-control" type="email"/>
    </div>
    <div>
      <label for="password" class="form-label">Password</label>
      <input id="password" class="form-control" type="password" v-model="password">
    </div>
    <div>
      <label for="confirm-password" class="form-label">Confirm Password</label>
      <input id="confirm-password" class="form-control" type="password" v-model="confirmPassword"/>
    </div>
    <div v-if="pass_errorMessage != ''"  style='background-color: #ff4444 ; border-radius: 8px; line-height:center;outline-color: #CC0000;padding-top:8px; padding-left:6px;'>
      <p>{{pass_errorMessage}}</p>
    </div>
    <div>
      <label for="phone-number" class="form-label">Phone Number</label>
      <input id="phone-number" class="form-control" type="text" />
    </div>
    <div>
      <label for="birth-date" class="form-label">Birth Date</label>
      <input id="birth-date" class="form-control" type="date" />
    </div>
    <div>
      <label for="experience" class="form-label">Years of Experience</label>
      <input id="experience" class="form-control" type="number" />
    </div>
    <div>
      <label for="level" class="form-label">Highest Level of Play</label>
      <select id="level" class="form-select">
        <option>Adult League Only</option>
        <option>Youth League</option>
        <option>High School</option>
        <option>College or Higher</option>
      </select>
    </div>
    <div class="hstack gap-3">
      <button class="btn btn-primary" @click="createAccount">Create an Account</button>
      <NuxtLink to="/app/sign-in" class="btn btn-link">Sign In</NuxtLink>
    </div>
  </div>
</template>
