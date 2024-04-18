<script lang="ts" setup>
useHead({
  titleTemplate: (title) => title ? `${title} | Power Play` : `Power Play`,
});
</script>


<template>
  <div class="d-flex align-items-center justify-content-center py-5" style="min-height: 100vh;">
    <div class="container pb-5" style="width: 400px;">
      <slot />
    </div>
  </div>
</template>
