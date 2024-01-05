<template>
  <template v-for="button in buttons" :key="button.text">
    <q-btn round flat color="grey-8" stack no-caps size="26px" class="GPL__side-btn" @click="$router.push(button.link)">
      <q-icon :name="button.icon" />
      <div class="GPL__side-btn__label">{{ button.text }}</div>
      <q-badge v-if="button.notifications > 0" color="red" text-color="white" floating style="top: 8px; right: 16px">
        {{ button.notifications }}
      </q-badge>
    </q-btn>
  </template>
</template>

<script setup lang="ts">
const buttons = [
  {
    text: 'Teams',
    icon: 'sports_hockey',
    notifications: 0,
    link: '/teams',
  },
  {
    text: 'Chat',
    icon: 'chat',
    notifications: 2,
    link: '/chat',
  },
  {
    text: 'Schedule',
    icon: 'calendar_month',
    notifications: 1,
    link: '/schedule',
  },
  {
    text: 'Finance',
    icon: 'paid',
    notifications: 0,
    link: '/finance',
  },
  {
    text: 'Account',
    icon: 'account_circle',
    notifications: 0,
    link: '/account',
  },
];
</script>