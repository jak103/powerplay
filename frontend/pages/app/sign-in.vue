<script lang="ts" setup>
import { ref } from 'vue';

const email = ref('');
const password = ref('');
definePageMeta({
  layout: "auth-layout",
});
const signIn = async () => {
  try {
    const response = await fetch('http://localhost:9001/api/v1/auth', {
      method: 'POST',
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: email.value,
        password: password.value
      })
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    localStorage.setItem('jwt', data.response_data.jwt);
    useRouter().push('/app');
  } catch (error) {
    console.error('Login failed:', error);
  }
}
</script>

<template>
  <Title>Sign In</Title>
  <h1>Sign In</h1>
  <div class="vstack gap-3">
    <div>
      <label for="email" class="form-label">Email</label>
      <input
        type="email"
        class="form-control"
        id="email"
        v-model="email"
      />
    </div>
    <div>
      <label for="password" class="form-label">Password</label>
      <input type="password" class="form-control" id="password" v-model="password" />
    </div>
    <div class="hstack gap-3">
      <button class="btn btn-primary" @click="signIn">Sign In</button>
      <NuxtLink to="/app/create-account" class="btn btn-link">Create an Account</NuxtLink>
    </div>
  </div>
</template>
