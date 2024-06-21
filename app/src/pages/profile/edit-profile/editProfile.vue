<template>
  <q-page class="column items-center">
    <div class="container">
      <div class="info-title">Personal Information</div>
      <div class="input-title">First Name</div>
      <q-input
        outlined
        v-model="firstName"
        label="First Name"
        :ref="setRef"
        dense
        style="width: 70vw"
        class="input"
      />
      <div class="input-title">Preferred Name</div>
      <q-input
        outlined
        v-model="preferredName"
        label="Preferred Name"
        :ref="setRef"
        dense
        style="width: 70vw"
        class="input"
      />
      <div class="input-title">Last Name</div>
      <q-input
        outlined
        v-model="lastName"
        label="Last Name"
        :ref="setRef"
        dense
        style="width: 70vw"
        class="input"
      />
      <div class="input-title">Birth Date</div>
      <div class="q-pa-md" style="max-width: 300px">
        <q-input
          filled
          v-model="birthDate"
          :ref="setRef"
          mask="##/##/####"
          :rules="[
            (val) =>
              /^-?[0-1]\d\/[0-3]\d\/[\d]+$/.test(val) ||
              'Please enter a valid date!',
          ]"
          dense
          style="width: 70vw; margin-left: -15px; margin-top: -10px"
        >
          <template v-slot:append>
            <q-icon name="event" class="cursor-pointer">
              <q-popup-proxy
                cover
                transition-show="scale"
                transition-hide="scale"
              >
                <q-date v-model="birthDate" mask="MM/DD/YYYY">
                  <div class="row items-center justify-end">
                    <q-btn v-close-popup label="Close" color="primary" flat />
                  </div>
                </q-date>
              </q-popup-proxy>
            </q-icon>
          </template>
        </q-input>
      </div>
      <div class="info-title">Contact Information</div>
      <div class="input-title">Email</div>
      <q-input
        outlined
        v-model="email"
        type="email"
        :ref="setRef"
        dense
        style="width: 70vw"
        class="input"
        :rules="[val => validateEmail(val) || 'Please enter a valid email!']"
      />
      <div class="input-title">Phone Number</div>
      <q-input
        outlined
        v-model="number"
        label="Phone Number"
        :ref="setRef"
        dense
        style="width: 70vw"
        class="input"
        mask="phone"
        :rules="[val => /^\([\d]{3}\)\s[\d]{3}\s-\s[\d]{4}$/.test(val) || 'Please enter a valid phone number!']"
      />
      <div class="info-title">Password</div>
      <div class="input-title">Current Password</div>
      <q-input
        outlined
        v-model="oldPassword"
        label="Old Password"
        :ref="setRef"
        dense
        type="password"
        style="width: 70vw"
        class="input"
      />
      <div class="input-title">New Password</div>
      <q-input
        outlined
        v-model="newPassword"
        label="New Password"
        :ref="setRef"
        dense
        type="password"
        style="width: 70vw"
        class="input"
      />
      <div class="input-title">Confirm New Password</div>
      <q-input
        v-model="confirmPassword"
        label="Confirm Password"
        :ref="setRef"
        type="password"
        outlined
        dense
        style="width: 70vw"
        class="input"
        :rules="[val => validatePassword(val) || 'Passwords do not match!']"
        />
        <a href='/user/profile'>I forgot my password</a>
    </div>
    <q-item class="button-container">
      <q-btn
        color="primary"
        label="Cancel"
        to="/profile"
        style="width: 35vw"
        class="q-pa-sm q-ma-sm rounded-edges"
        outline
      />
      <q-btn
      unelevated
        color="primary"
        label="Save Info"
        type="submit"
        class="q-pa-sm q-ma-sm rounded-edges"
        style="width: 35vw"
        @click="saveUserInfo()"
      />
    </q-item>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter()

// TODO: replace this canned user info with info from backend/stores
const userInfo = {
  firstName: 'Ricky',
  preferredName: 'Ricky Bobby',
  lastName: 'Bobby',
  'phone-number': '9011234567',
  email: 'RickyBobby@gmail.com',
  birthday: '12/12/1998',
  password: 'hunter1',
};

const inputRefs = ref([])

const setRef = (element) => {
  inputRefs.value.push(element)
}

const validateFields = () => {
  inputRefs.value.forEach(field => {
    if (!field.validate()) {
      return false;
    }
  })
  return true;
}

const firstName = ref<string>(userInfo['firstName']);
const preferredName = ref<string>(userInfo['preferredName']);
const lastName = ref<string>(userInfo['lastName']);
const birthDate = ref<string>(userInfo['birthday']);
const email = ref<string>(userInfo['email']);
const oldPassword = ref<string>(userInfo['password']);
const newPassword = ref<string>('');
const confirmPassword = ref<string>('');
const number = ref<number | null>(userInfo['phone-number']);

// Regex pulled directly from Quasar's default patterns and used here to allow an error message to be displayed.
const validateEmail = (v) => {
  return /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(v)
}

const validatePassword = (v) => {
  return v == newPassword.value
}

const saveUserInfo = () => {
  if (!validateFields()) {
    console.log('Validation failed :(')
    return;
  }

  // Add changed items to request body
  const apiBody = {}
  const fields = [
    { key: 'firstName', value: firstName.value },
    { key: 'preferredName', value: preferredName.value },
    { key: 'lastName', value: lastName.value },
    { key: 'birthday', value: birthDate.value },
    { key: 'email', value: email.value },
    { key: 'phone-number', value: number.value }
  ]

  fields.forEach(field => {
    if (field.value != userInfo[field.key]) {
      apiBody[field.key] = field.value
    }
  })

  if (newPassword.value) {
    apiBody['password'] = newPassword.value
  }

  if (Object.keys(apiBody).length === 0) {
    router.push('/profile')
    return;
  }

  // TODO: Replace the following log to console with a call to the backend, using the object just created
  console.log('See API request body below:')
  console.log(apiBody)
  router.push('/profile')
}
</script>

<style>
.input {
  margin-bottom: 15px;
}
.info-title {
  font-size: 25px;
  font-weight: bold;
  margin-top: 25px;
  margin-bottom: 10px;
}
.input-title {
  margin-bottom: 5px;
}
.container {
  display: flex;
  justify-content: baseline;
  flex-direction: column;
}
</style>
