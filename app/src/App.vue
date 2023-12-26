<template>

  <!-- TODO add a prop that passes the current logged in user-->
  <mobile-layout v-if="$q.screen.lt.md" />
  <web-layout v-if="$q.screen.gt.sm" />

</template>

<script setup lang="ts">

import MobileLayout from './layouts/MobileLayout.vue';
import WebLayout from './layouts/WebLayout.vue';

</script>
