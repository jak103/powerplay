<template>
    <div class="app-wrapper">
      <div class="app-menu">
        <h2 class="d-none d-lg-block">Power Play</h2>
        <div class="nav-icon-btn-group">
          <NuxtLink
            active-class=""
            exact-active-class="active"
            to="/app"
            class="btn btn-lg nav-icon-btn"
          >
            <Icon name="home" />
            <div class="nav-text">Home</div>
          </NuxtLink>
          <NuxtLink
            to="/app/schedule"
            class="btn btn-lg nav-icon-btn"
          >
            <Icon name="calendar-month" />
            <div class="nav-text">Schedule</div>
          </NuxtLink>
          <NuxtLink to="/app/chat" class="btn btn-lg nav-icon-btn">
            <Icon name="message-outline" />
            <div class="nav-text">Chat</div>
          </NuxtLink>
          <NuxtLink to="/app/profile" class="btn btn-lg nav-icon-btn">
            <Icon name="account-circle" />
            <div class="nav-text">Profile</div>
          </NuxtLink>
        </div>
        <NuxtLink class="btn btn-link mt-auto d-none d-lg-inline" to="/"
        >Website Home</NuxtLink
        >
        <NuxtLink class="btn btn-primary d-none d-lg-inline" to="/app/sign-in"
        >Sign Out</NuxtLink
        >
      </div>
      <div class="app-container container pt-4">
        <slot />
      </div>
    </div>
</template>

<style lang="scss" scoped>
.app-menu{
  background: var(--bs-tertiary-bg);
}
.nav-icon-btn.active{
  background: var(--bs-secondary-color) !important;
  color: var(--bs-body-bg) !important;
  // color: white !important;
  border: 1px solid transparent;
}
@media (min-width: 992px) {
  // lg and bigger
  .app-wrapper {
    display: flex;
    flex-direction: row;
    flex: 1 1 0;
    height: 100%;
    .app-container {
      margin-left: 300px;
    }
    .app-menu {
      order: 0;
      width: 300px;
      position: fixed;
      top: 0;
      bottom: 0;
      left: 0;
      padding: 16px;
      display: flex;
      flex-direction: column;
    }
    .nav-icon-btn-group {
      display: flex;
      flex-direction: column;
    }
    .nav-icon-btn {
      display: flex;
      align-items: center;
      gap: 12px;
    }
  }
}
@media (max-width: 991.98px) {
  // md and smaller
  .app-wrapper {
    display: flex;
    flex-direction: column;
    .app-menu {
      order: 1;
      position: fixed;
      bottom: 0;
      left: 0;
      right: 0;
      flex: 0 0 0;
    }
    .app-container {
      flex: 1 1 0;
      margin-bottom: 64px;
    }
    .nav-icon-btn-group {
      display: flex;
      gap: 4px;
      height: 56px;
      align-items: center;
      padding: 0 12px;
    }
    .nav-icon-btn {
      display: inline-flex;
      flex-direction: column;
      align-items: center;
      flex: 1 1 0;
      width: 0;
      padding: 4px;
    }
    .nav-text {
      font-size: 0.6em;
    }
  }
}
</style>
