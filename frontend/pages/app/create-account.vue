<script lang="ts" setup>
const user = ref({
  firstName: '',
  lastName: '',
  email: '',
  password: '',
  confirmPassword: '',
  phone: '',
  birthDate: '',
  experience: '',
  highestLevel: 'Adult League Only',
});

const levels = [
  'Youth Hockey',
  'High School',
  'College',
  'Adult League Only'
]

const signIn = () => {
  alert(`Creating user ${JSON.stringify(user.value)}`);
};
</script>

<template>
  <div>
    <Title>Create an Account</Title>
    <h1>Create an Account</h1>
    <q-form @submit="signIn">
      <q-input v-model="user.firstName" label="First Name" />
      <q-input v-model="user.lastName" label="Last Name" />
      <q-input v-model="user.email" label="Email" type="email" hint="This will be your username" />
      <q-input v-model="user.password" label="Password" type="password" />
      <q-input v-model="user.confirmPassword" label="Confirm Password" type="password" />
      <q-input v-model="user.phone" label="Phone Number" />
      <q-input v-model="user.birthDate" label="Birth Date" type="date" />
      <q-input v-model="user.experience" label="Years of Experience" type="number" step=".01" />
      <q-select v-model="user.highestLevel" label="Highest Level of Experience" :options="levels" />
      <q-btn type="submit" @click="">Submit</q-btn>
      <q-btn to="/app/create-account">Create an Account</q-btn>
    </q-form>
  </div>
</template>