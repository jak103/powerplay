<script lang="ts" setup>

</script>

<template>
  <div class="app-wrapper">
    <div class="app-menu text-bg-primary">
      <NuxtLink active-class="" exact-active-class="active" to="/app" class="btn btn-primary btn-lg rounded-0 text-start hstack gap-3">
        <Icon name="home" />
        Home
      </NuxtLink>
      <NuxtLink to="/app/schedule" class="btn btn-primary btn-lg rounded-0 text-start hstack gap-3">
        <Icon name="calendar-month" />
        Schedule
      </NuxtLink>
      <NuxtLink to="/app/chat" class="btn btn-primary btn-lg rounded-0 text-start hstack gap-3">
        <Icon name="message-outline" />
        Chat
      </NuxtLink>
      <NuxtLink to="/app/profile" class="btn btn-primary btn-lg rounded-0 text-start hstack gap-3">
        <Icon name="account-circle" />
        Profile
      </NuxtLink>
    </div>
    <div class="container">
      <NuxtPage />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.app-wrapper{
  display: flex;
  flex: 1 1 0;
  height: 100%;
  .app-menu{
    min-width: 250px;
  }
}
</style>
