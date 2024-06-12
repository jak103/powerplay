<script lang="ts" setup>
  import { useRouter } from 'vue-router';
  import { useUserStore } from '../../../app/src/stores/user-store';

  const signIn = async () => {
    const router = useRouter();
    const userStore = useUserStore(); 
    const emailElement = document.getElementById('email') as HTMLInputElement;
    const passwordElement = document.getElementById('password') as HTMLInputElement;

    if (!emailElement || !passwordElement) {
      console.error('Email or password field is missing');
      return;
    }

    const email = emailElement.value;
    const password = passwordElement.value;

    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email: email,
          password: password
        })
      });

      if (response.ok) {
        const data = await response.json();
        userStore.setUser(data.user);
        router.push('/app');
      } else {
        console.error('Login failed');
      }
    } catch (error) {
      console.error('An error occurred:', error);
    }
  };
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
      />
    </div>
    <div>
      <label for="password" class="form-label">Password</label>
      <input type="password" class="form-control" id="password" />
    </div>
    <div class="hstack gap-3">
      <button class="btn btn-primary" @click="signIn">Sign In</button>
      <NuxtLink to="/app/create-account" class="btn btn-link">Create an Account</NuxtLink>
    </div>
  </div>
</template>