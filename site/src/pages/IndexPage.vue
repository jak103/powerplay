<template>
  <q-page class="row items-center justify-evenly">
    <LoginComponent />
  </q-page>
</template>

<script setup lang="ts">
// import { ref } from 'vue';
import LoginComponent from 'components/LoginComponent.vue';

defineOptions({
  name: 'IndexPage'
});


</script>
