<script lang="ts" setup>
definePageMeta({
  layout: "auth-layout",
});
const signIn = () => {
  useRouter().push('/app')
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
