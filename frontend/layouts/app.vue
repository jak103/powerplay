<script lang="ts" setup>

</script>

<template>
  <div>&lt; Power Play ...</div>
  <slot />

  <ul>
    <li><NuxtLink to="/appl">Home</NuxtLink></li>
    <li><NuxtLink to="/appl/chat">Chat</NuxtLink></li>
    <li><NuxtLink to="/appl/profile">Profile</NuxtLink></li>
    <li><NuxtLink to="/">Website Home</NuxtLink></li>
  </ul>
</template>
