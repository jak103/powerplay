<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-toolbar-title>
          PowerPlay
        </q-toolbar-title>

        <q-btn stretch flat label="Home" to="/" />
        <q-separator dark vertical inset />
        <q-btn stretch flat label="League" to="/league" />
        <q-separator dark vertical inset />
        <q-btn stretch flat label="*Calendar" to="/" />
        <q-separator dark vertical inset />
        <q-btn stretch flat label="*Registration" to="/" />
        <q-separator dark vertical inset />

        <q-btn-dropdown stretch flat label="Information">
          <q-list>
            <q-item clickable to="/information/how-to-join">
              <q-item-section>How to Join</q-item-section>
            </q-item>
            <q-item clickable to="/information/players">
              <q-item-section>For Players</q-item-section>
            </q-item>
            <q-item clickable to="/information/managers">
              <q-item-section>For Managers</q-item-section>
            </q-item>
            <q-item clickable to="/information/staff">
              <q-item-section>Our Staff</q-item-section>
            </q-item>
            <q-item clickable to="/information/rink">
              <q-item-section>Our Rink</q-item-section>
            </q-item>
            <q-item clickable to="/information/substitution">
              <q-item-section>Substitutions</q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>

        <q-separator dark vertical inset />
        <q-btn stretch flat label="*Documents" to="/" />
        <q-separator dark vertical inset />
        <q-btn stretch flat label="*Sign In" to="/sign-in" />
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
defineOptions({
  name: 'MainLayout'
});
</script>
