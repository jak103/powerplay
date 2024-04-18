<script lang="ts" setup>
useHead({
  script: [
    {
      src: "https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js",
    },
    {
      src: '/theme.js',
      tagPosition: 'bodyOpen'
    }
  ],
  titleTemplate: (title) =>
    title ? `${title} | Eccles Ice Hockey` : "Eccles Ice Hockey",
});
</script>

<template>
  <DevOnly><DevBar /></DevOnly>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
