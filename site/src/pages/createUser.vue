<script lang="ts" setup>

import {ref} from 'vue'

const options = ref([
      'Youth League', 'High School', 'Adult League Only', 'College or Higher'
      ])

const isPwd = ref(true)
const isConPwd = ref(true)

// user inputs
const levelPlayed = ref('')
const firstName =ref('')
const lastName =ref('')
const password = ref('')
const confirmPassword = ref('')
const birthDate = ref('')
const phoneNumberInput = ref('')
const phoneNumber = ref('')
const email = ref('')
const experience = ref('')

//const yearsOfExp = 0

// error messages if needed
const first_errorMessage = ref('')
const last_errorMessage = ref('')
const pass_errorMessage = ref('')
const birthday_errorMessage = ref('')
const phone_errorMessage = ref('')
const email_errorMessage = ref('')
const option_errorMessage = ref('')
const experience_errorMessage = ref('')

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
  const isExpValid = validateYearsOfExp()
  const isOptionValid = validateOption()

  console.log('Called?')
  if (isPasswordValid && isPhoneValid && isBirthdayValid && isEmailValid && isLastValid && isFirstValid && isExpValid && isOptionValid) {
    // All validations passed
    console.log()
    //TODO: figure out backend post requests and reroute user to login page
    // useRouter().push('/app')
  }
}
function validateFirst(){
if(firstName.value == ''){
  first_errorMessage.value = 'Please enter your first name!'
  return false
}
return true
  
}
function validateLast(){
  if(lastName.value == ''){
  last_errorMessage.value = 'Please enter your last name!'
  return false
}
return true

}

function validatePassword() {
  if (password.value !== confirmPassword.value) {
    pass_errorMessage.value = 'Your password does not match!'
    return false
  }
  if (password.value.length < 8) {
    pass_errorMessage.value = 'Your password must be at least 8 characters long!'
    return false
  }

  const specialCharacters = ['!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '{', '}', '[', ']', '|', '\\', ':', ';', '"', "'", '<', '>', ',', '.', '/']
  const specialChar = password.value.split('').some(char => specialCharacters.includes(char))
  
  if (!specialChar) {
    pass_errorMessage.value = 'Your password must contain a special character!'
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
    phone_errorMessage.value = 'The phone number you entered is invalid!'
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

  console.log(age+' '+monthDiff+' '+dayDiff)
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    return age - 1 < 18
  }

  return false
}

function validateEmail() {
  email_errorMessage.value = ''
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailPattern.test(email.value)|| !email.value.endsWith('.com')) {
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
  experience_errorMessage.value = ''
  if(experience.value == '') {
    experience_errorMessage.value = 'Please enter years of experience!'
    return false
  }
  return true
}

function validateOption(){
  option_errorMessage.value = ''
  if(levelPlayed.value == '') {
    option_errorMessage.value = 'Please select level of experience!'
    return false
  }
  return true
}
</script>

<template>
  <q-card class="my-card white text-black" style="margin-left:20%; margin-right:20%; margin-top:5%; margin-bottom:5%;">
    <q-card-section align="middle">
      <div class="text-h6"><h3>Create an Account</h3></div>

      <!-- First Name -->  
      <div class="q-mt-md">
        <q-input filled v-model="firstName" label="First Name" @keyup="validateIsAlphaOnly" />
        <q-banner v-if="first_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p align="left">{{ first_errorMessage }}</p>
        </q-banner>
      </div>
      
      <!-- Last Name -->  
      <div class="q-mt-md">
        <q-input filled v-model="lastName" label="Last Name" @keyup="validateIsAlphaOnly" />
        <q-banner v-if="last_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p>{{ last_errorMessage }}</p>
        </q-banner>
      </div>
      
      <!-- Email -->   
      <div class="q-mt-md">
        <q-input filled v-model="email" label="Email" @input="validateIsAlphaOnly" />
        <q-banner v-if="email_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p>{{ email_errorMessage }}</p>
        </q-banner>
      </div>
      
      <!-- Password -->
      <div class="q-mt-md">
        <q-input v-model="password" filled :type="isPwd ? 'password' : 'text'" label="Password">
          <template v-slot:append>
            <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer" @click="isPwd = !isPwd" />
          </template>
        </q-input>
      </div>
      
      <div class="q-mt-md">
        <q-input v-model="confirmPassword" filled :type="isConPwd ? 'password' : 'text'" label="Confirm password">
          <template v-slot:append>
            <q-icon :name="isConPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer" @click="isConPwd = !isConPwd" />
          </template>
        </q-input>
        <q-banner v-if="pass_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p>{{ pass_errorMessage }}</p>
        </q-banner>
      </div>

      <!-- Phone Number -->
      <div class="q-mt-md">
        <q-input filled v-model="phoneNumberInput" label="Phone Number" />
        <q-banner v-if="phone_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p>{{ phone_errorMessage }}</p>
        </q-banner>
      </div>

      <!-- Birthday -->
      <div class="q-mt-md">
        <q-input filled type="date" v-model="birthDate" label="Birth Date" />
        <q-banner v-if="birthday_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p>{{ birthday_errorMessage }}</p>
        </q-banner>
      </div>

      <!-- Experience -->
      <div class="q-mt-md">
        <q-input
          filled
          v-model="experience"
          type="number"
          label="Years of experience"
          :rules="[val => val >= 0 || 'The number cannot be negative']"
          min="0"
        />
        <q-banner v-if="experience_errorMessage" class="text-white bg-red" style="border-radius: 4px 4px 4px 4px;">
          <p>{{ experience_errorMessage }}</p>
        </q-banner>
      </div>
      
      <!-- Level of Play -->
      <div class="q-mt-md">
        <q-select filled v-model="levelPlayed" :options="options" label="Highest Level of Play" />
        <q-banner v-if="option_errorMessage" class="text-white bg-red" style="border-radius: 0px 0px 4px 4px;">
          <p>{{ option_errorMessage }}</p>
        </q-banner>
      </div>
    </q-card-section>

    <q-card-actions align="center">
      <q-btn label="Create Account" type="submit" color="primary" @click="createAccount" />
    </q-card-actions>
  </q-card>
</template>

