<template>
  <q-btn class="example" :color="color" text-color="white" label="Press" @click="change" />
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'ColorButtonComponent',
  data: () => ({
    color: 'blue'
  }),
  methods: {
    change() {
      if (this.color === 'blue') {
        this.color = 'yellow';
      } else {
        this.color = 'blue';
      }
      return this.color;
    }
  }
})
</script>

<style lang="scss" scoped>
.example {
  margin: 1px;
}
</style>
