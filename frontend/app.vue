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
