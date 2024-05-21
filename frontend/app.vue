<script lang="ts" setup>
useHead({
  script: [
    {
      src: "https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js",
    },
    {
      src: '/initTheme.js',
      tagPosition: 'bodyOpen'
    }
  ],
  link: [
    {
      rel: "manifest",
      href: "../manifest.json"
    }
  ],
  titleTemplate: (title) =>
    title ? `${title} | Eccles Ice Hockey` : "Eccles Ice Hockey",
});

// Register service worker
onMounted(async () => {
        if ('serviceWorker' in navigator) {
          try {
            const registration = await navigator.serviceWorker.register('/sw.js');
            console.log('ServiceWorker registration successful with scope: ', registration.scope);
          } catch (err) {
            console.log('ServiceWorker registration failed: ', err);
          }
        }
      });
</script>

<template>
  <DevOnly><DevBar /></DevOnly>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<style>
  /* Disable text selection on touchscreen devices */
  @media (hover: none) { 
    body {
      /* Prevent text selection */
      -webkit-user-select: none; /* Safari */
      -moz-user-select: none; /* Firefox */
      -ms-user-select: none; /* Internet Explorer/Edge */
      user-select: none; /* Non-prefixed version, currently supported by Chrome, Opera, and Edge */
    }
  }
</style>