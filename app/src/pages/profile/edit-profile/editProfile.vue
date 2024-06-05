<template>
    <q-page class="column items-center">
        <div class="container" >
            <div class="info-title">Personal Information</div>
            <div class="input-title">First Name</div>
            <q-input outlined v-model="firstName" label="First Name" dense style="width: 70vw;" class="input"/>
            <div class="input-title">Preferred Name</div>
            <q-input outlined v-model="preferredName" label="Preferred Name" dense style="width: 70vw;" class="input"/>
            <div class="input-title">Last Name</div>
            <q-input outlined v-model="lastName" label="Last Name" dense style="width: 70vw;" class="input"/>
            <div class="input-title">Birth Date</div>
            <div class="q-pa-md" style="max-width: 300px">
    <q-input filled v-model="birthDate" mask="date" :rules="['date']" dense style="width: 70vw;margin-left: -15px; margin-top: -10px;">
      <template v-slot:append>
        <q-icon name="event" class="cursor-pointer">
          <q-popup-proxy cover transition-show="scale" transition-hide="scale">
            <q-date v-model="birthDate">
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
            <q-input outlined v-model="email" label="Email" type="email" dense style="width: 70vw;" class="input"/>
            <div class="input-title">Phone Number</div>
            <q-input outlined v-model="number" label="Phone Number" dense style="width: 70vw;" class="input"/>
            <div class="info-title">Password</div>
            <div class="input-title">Current Password</div>
            <q-input outlined v-model="oldPassword" label="Old Password" dense type="password" style="width: 70vw;" class="input"/>
            <div class="input-title">New Password</div>
            <q-input outlined v-model="newPassword" label="New Password" dense type="password" style="width: 70vw;" class="input"/>
            <div class="input-title">Confirm New Password</div>
            <q-input v-model="confirmPassword" label="Confirm Password" type="password" outlined dense style="width: 70vw;" class="input"/>
            <q-btn flat color="primary" label="Forgot your password?" />
        </div>
        <q-item class="button-container">
      <q-btn 
      color="primary"
       label="Cancel"
        to="/profile" 
        style="width: 40vw;" 
        outline
        />
      <q-btn unelevated color="primary" 
      label="Save Info"
      to="/profile"
       class="q-pa-sm q-ma-sm rounded-edges" 
       style="width: 40vw;"/>
    </q-item>

    </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const firstName = ref<string>('');
const preferredName = ref<string>('');
const lastName = ref<string>('');
const birthDate = ref<string>('');
const email = ref<string>('');
const oldPassword = ref<string>('');
const newPassword = ref<string>('');
const confirmPassword = ref<string>('');
const number = ref<number | null>(null);
</script>

<style>
.input{
    margin-bottom: 15px;
}
.info-title {
    font-size: 25px;
    font-weight: bold;
    margin-top: 25px;
    margin-bottom: 10px;
}
.input-title{
    margin-bottom: 5px;
}
.container{
    display: flex;
    justify-content: baseline;
    flex-direction: column;
}
</style>