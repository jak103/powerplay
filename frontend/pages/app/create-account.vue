<script lang="ts" setup>
definePageMeta({
  layout: "auth-layout",
});

// user inputs
const password = ref('')
const confirmPassword = ref('')
const birthDate = ref('')
const phoneNumberInput = ref('')
const phoneNumber = ref('')
const email = ref('')

// error messages if needed
const pass_errorMessage = ref('')
const birthday_errorMessage = ref('')
const phone_errorMessage = ref('')
const email_errorMessage = ref('')

const createAccount = () => {
  pass_errorMessage.value = ''
  birthday_errorMessage.value = ''
  phone_errorMessage.value = ''
  email_errorMessage.value = ''

  const isPasswordValid = validatePassword()
  const isPhoneValid = validatePhone()
  const isBirthdayValid = validateBirthday()
  const isEmailValid = validateEmail()

  if (isPasswordValid && isPhoneValid && isBirthdayValid && isEmailValid) {
    // All validations passed
    // useRouter().push('/app')
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

  return true
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
</script>

<template>
  <Title>Create an Account</Title>
  <h1>Create an Account</h1>
  <div class="vstack gap-3">
    <div>
      <label for="first-name" class="form-label">First Name</label>
      <input id="first-name" class="form-control" />
    </div>
    <div>
      <label for="last-name" class="form-label">Last Name</label>
      <input id="last-name" class="form-control" />
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