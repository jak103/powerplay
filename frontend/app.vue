<template>
  <Head>
    <Script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></Script>
    <Title>EcclesIceHockey</Title>
  </Head>

  <div class="app-wrapper">
    <DevOnly><DevBar /></DevOnly>
    <NavBar />
    <div class="flex-fill vstack">
      <NuxtPage />
    </div>
    <Footer />
  </div>
</template>

<style lang="scss" scoped>
.app-wrapper{
  min-height: 100vh;
  height: 100%;
display: flex;
flex-direction: column;
}
</style>
