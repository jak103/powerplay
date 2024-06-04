<template>
  <q-page padding>
    <ColorButtonComponent />

    <q-input class="example-input" filled v-model:model-value="textInput" label="Type Here" />
  </q-page>
</template>

<script setup lang="ts">
import ColorButtonComponent from 'components/ColorButtonComponent.vue';
import { ref } from 'vue';

defineOptions({
  name: 'ExamplePage'
});

const textInput = ref('');
</script>

<style lang="scss" scoped>
.example-input {
  margin: 1px;
}
</style>
