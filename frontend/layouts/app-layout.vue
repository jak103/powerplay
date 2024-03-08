<script lang="ts" setup></script>

<template>
  <div class="app-wrapper">
    <div class="app-menu p-3 vstack gap-1 text-bg-light shadow">
      <h2>Power Play</h2>
      <NuxtLink
        active-class=""
        exact-active-class="active"
        to="/app"
        class="btn btn-light btn-lg hstack gap-3"
      >
        <Icon name="home" />
        Home
      </NuxtLink>
      <NuxtLink to="/app/schedule" class="btn btn-light btn-lg hstack gap-3">
        <Icon name="calendar-month" />
        Schedule
      </NuxtLink>
      <NuxtLink to="/app/chat" class="btn btn-light btn-lg hstack gap-3">
        <Icon name="message-outline" />
        Chat
      </NuxtLink>
      <NuxtLink to="/app/profile" class="btn btn-light btn-lg hstack gap-3">
        <Icon name="account-circle" />
        Profile
      </NuxtLink>
      <NuxtLink class="btn btn-link mt-auto" to="/"
        >Website Home</NuxtLink
      >
      <NuxtLink class="btn btn-primary" to="/app/sign-in"
        >Sign Out</NuxtLink
      >
    </div>
    <div class="container py-4">
      <slot />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.app-wrapper {
  min-height: 100vh;
  display: flex;
  flex: 1 1 0;
  height: 100%;
  .app-menu {
    height: 100vh;
    min-width: 300px;
    position: sticky;
    top: 0;
  }
}
</style>
