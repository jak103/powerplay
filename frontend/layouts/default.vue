<template>
  <div class="app-wrapper">
    <NavBar />
    <div class="flex-fill vstack">
      <NuxtPage />
    </div>
    <Footer />
  </div>
</template>

<style lang="scss" scoped>
.app-wrapper {
  min-height: 100vh;
  height: 100%;
  display: flex;
  flex-direction: column;
}
</style>
