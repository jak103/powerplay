<script setup lang="ts">
import { QBtn } from 'quasar'
import { useCounterStore } from '../stores/counter'

const counter = useCounterStore()
</script>

<template>
  <div class="card">
    <QBtn fab icon="add" color="accent" @click="counter.increment()"/>
    <p>Count is {{ counter.count }}</p>
    <p>Count times 2 is {{ counter.doubleCount }}</p>
  </div>
</template>
