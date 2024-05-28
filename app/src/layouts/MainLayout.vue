<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated style="background-color: #343333;">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
          v-if="!$q.screen.lt.md"
        />

        <q-toolbar-title> Power Play </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <div class="column">
        <q-btn
          v-for="item in navItems"
          :key="item.label"
          v-bind="item"
          class="q-mt-md q-pt-md q-pb-md"
          style="display: flex; align-items: start; width: 90%;"
          color="black"
        />
      </div>
    </q-drawer>

    <q-footer class="q-pa-md bg-white" elevated v-if="$q.screen.lt.md">
        <div class="row justify-evenly">
          <q-btn v-for="item in navItems" :key="item.label" v-bind="item" color="black" label=""/>
        </div>
    </q-footer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<style scoped>
  .column {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>

<script setup lang="ts">
import { ref } from 'vue';

defineOptions({
  name: 'MainLayout',
});

const navItems = [
  {
    label: 'Home',
    icon: 'home',
    to: '/',
    flat: true,
    dense: true,
  },
  {
    label: 'Schedule',
    icon: 'event',
    to: '/schedule',
    flat: true,
    dense: true,
  },
  {
    label: 'Chat',
    icon: 'chat_bubble_outline',
    to: '/chat',
    flat: true,
    dense: true,
  },
  {
    label: 'Profile',
    icon: 'person',
    to: '/profile',
    flat: true,
    dense: true,
  },
  {
    label: 'Schedule List',
    caption: 'Example Schedule List Component',
    icon: 'view_list',
    to: '/schedule-list'
    flat: true,
    dense: true,
  },
];

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
