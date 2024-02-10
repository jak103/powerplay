<script lang="ts" setup>
const username = ref("");
const password = ref("");

const signIn = () => {
  alert("Sign In here...");
};
</script>

<template>
  <div>
    <h1>Sign In</h1>
    <q-form @submit="signIn">
      <q-input v-model="username" label="Username" />
      <q-input v-model="password" label="Password" type="password" />
      <q-btn type="submit" @click="">Submit</q-btn>
      <q-btn to="/app/create-account">Create an Account</q-btn>
    </q-form>
  </div>
</template>
