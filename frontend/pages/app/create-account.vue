<script lang="ts" setup>
definePageMeta({
  layout: "auth-layout",
});

// user inputs
const firstName =ref('')
const lastName =ref('')
const password = ref('')
const confirmPassword = ref('')
const birthDate = ref('')
const phoneNumberInput = ref('')
const phoneNumber = ref('')
const email = ref('')
const yearsOfExp = 0

// error messages if needed
const first_errorMessage = ref('')
const last_errorMessage = ref('')
const pass_errorMessage = ref('')
const birthday_errorMessage = ref('')
const phone_errorMessage = ref('')
const email_errorMessage = ref('')

const createAccount = () => {
  pass_errorMessage.value = ''
  birthday_errorMessage.value = ''
  phone_errorMessage.value = ''
  email_errorMessage.value = ''

  const isFirstValid = validateFirst()
  const isLastValid = validateLast()
  const isPasswordValid = validatePassword()
  const isPhoneValid = validatePhone()
  const isBirthdayValid = validateBirthday()
  const isEmailValid = validateEmail()

  if (isPasswordValid && isPhoneValid && isBirthdayValid && isEmailValid) {
    // All validations passed

    //TODO: figure out backend post requests and reroute user to login page
    // useRouter().push('/app')
  }
}
function validateFirst(){
if(firstName.value == ''){
  first_errorMessage.value = 'Please enter your first name!'
}
  
}
function validateLast(){
  if(lastName.value == ''){
  first_errorMessage.value = 'Please enter your last name!'
}
}

function validatePassword() {
  if (password.value !== confirmPassword.value) {
    pass_errorMessage.value = "Your password doesn't match!"
    return false
  }
  if (password.value.length < 8) {
    pass_errorMessage.value = "Your password must be at least 8 characters long!"
    return false
  }

  const specialCharacters = ['!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '{', '}', '[', ']', '|', '\\', ':', ';', '"', "'", '<', '>', ',', '.', '/']
  const specialChar = password.value.split('').some(char => specialCharacters.includes(char))
  
  if (!specialChar) {
    pass_errorMessage.value = "Your password must contain a special character!"
    return false
  }

  return true
}

function validatePhone() {
  phoneNumber.value = ''
  phone_errorMessage.value = ''

  const charCodeZero = '0'.charCodeAt(0)
  const charCodeNine = '9'.charCodeAt(0)

  for (let i = 0; i < phoneNumberInput.value.length; i++) {
    const char = phoneNumberInput.value[i]
    if (char.charCodeAt(0) >= charCodeZero && char.charCodeAt(0) <= charCodeNine) {
      phoneNumber.value += char
    }
  }
  
  if (phoneNumber.value.length !== 10) {
    phone_errorMessage.value = "The phone number you entered is invalid!"
    return false
  }

  return true
}

function validateBirthday() {
  birthday_errorMessage.value = ''
  if (birthDate.value === '') {
    birthday_errorMessage.value = 'Please enter your birthday!'
    return false
  } 

  const [year, month, day] = birthDate.value.split('-').map(part => parseInt(part, 10))
  const currentDate = new Date()
  const currentYear = currentDate.getFullYear()
  const currentMonth = currentDate.getMonth() + 1
  const currentDay = currentDate.getDate()
  
  if (year > currentYear || (year === currentYear && (month > currentMonth || (month === currentMonth && day > currentDay)))) {
    birthday_errorMessage.value = 'Invalid birthday entered!'
    return false
  }

  if(isUnderage()){
    birthday_errorMessage.value = 'Sorry! You must be 18 years or older to play!'
  }

  return true
}

function isUnderage() {
  const [year, month, day] = birthDate.value.split('-').map(part => parseInt(part, 10))
  const currentDate = new Date()
  const birthDateObj = new Date(year, month - 1, day)
  const age = currentDate.getFullYear() - birthDateObj.getFullYear()
  const monthDiff = currentDate.getMonth() - birthDateObj.getMonth()
  const dayDiff = currentDate.getDate() - birthDateObj.getDate()

  if(age > 18){ 
    return false
  }

  console.log(age+" "+monthDiff+" "+dayDiff)
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    return age - 1 < 18
  }

  return false
}

function validateEmail() {
  email_errorMessage.value = ''
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailPattern.test(email.value)|| !email.value.endsWith(".com")) {
    email_errorMessage.value = 'Invalid email!'
    return false
  }
  return true
}

function validateIsAlphaOnly() {
    first_errorMessage.value = ''
    last_errorMessage.value = ''
    
      const regex = /^[A-Za-z]+$/
      if (!regex.test(firstName.value)&& firstName.value != '') {
        first_errorMessage.value= 'Name can only contain letters!'
      }
      if (!regex.test(lastName.value) && lastName.value != '') {
        last_errorMessage.value= 'Name can only contain letters!'
      }
    }

function validateYearsOfExp(){
  //TODO: Validate Years of experience ei) not negative or decimal numbers

}
</script>

<template>
  <Title>Create an Account</Title>
  <h1>Create an Account</h1>
  <div class="vstack gap-3">
    <div>
      <label for="first-name" class="form-label">First Name</label>
      <input id="first-name" class="form-control" @input="validateIsAlphaOnly" v-model="firstName" />
    </div>
    <div v-if="first_errorMessage" class="alert alert-danger">
      <p>{{ first_errorMessage }}</p>
    </div>
    <div>
      <label for="last-name" class="form-label">Last Name</label>
      <input id="last-name" class="form-control" @input="validateIsAlphaOnly" v-model="lastName"/>
    </div>
    <div v-if="last_errorMessage" class="alert alert-danger">
      <p>{{ last_errorMessage }}</p>
    </div>
    <div>
      <label for="email" class="form-label">Email</label>
      <input id="email" class="form-control" type="email" v-model="email"/>
    </div>
    <div v-if="email_errorMessage" class="alert alert-danger">
      <p>{{ email_errorMessage }}</p>
    </div>
    <div>
      <label for="password" class="form-label">Password</label>
      <input id="password" class="form-control" type="password" v-model="password"/>
    </div>
    <div>
      <label for="confirm-password" class="form-label">Confirm Password</label>
      <input id="confirm-password" class="form-control" type="password" v-model="confirmPassword"/>
    </div>
    <div v-if="pass_errorMessage" class="alert alert-danger">
      <p>{{ pass_errorMessage }}</p>
    </div>
    <div>
      <label for="phone-number" class="form-label">Phone Number</label>
      <input id="phone-number" class="form-control" type="text" v-model="phoneNumberInput"/>
    </div>
    <div v-if="phone_errorMessage" class="alert alert-danger">
      <p>{{ phone_errorMessage }}</p>
    </div>
    <div>
      <label for="birth-date" class="form-label">Birth Date</label>
      <input id="birth-date" class="form-control" type="date" v-model="birthDate" />
    </div>
    <div v-if="birthday_errorMessage" class="alert alert-danger">
      <p>{{ birthday_errorMessage }}</p>
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